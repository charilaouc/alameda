package vsphere

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/google/uuid"
	machinev1 "github.com/openshift/machine-api-operator/pkg/apis/machine/v1beta1"
	vspherev1 "github.com/openshift/machine-api-operator/pkg/apis/vsphereprovider/v1alpha1"
	machinecontroller "github.com/openshift/machine-api-operator/pkg/controller/machine"
	"github.com/openshift/machine-api-operator/pkg/controller/vsphere/session"
	"github.com/pkg/errors"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vapi/tags"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog"
)

const (
	minMemMB              = 2048
	minCPU                = 2
	fullCloneDiskMoveType = string(types.VirtualMachineRelocateDiskMoveOptionsMoveAllDiskBackingsAndConsolidate)
	linkCloneDiskMoveType = string(types.VirtualMachineRelocateDiskMoveOptionsCreateNewChildDiskBacking)
	ethCardType           = "vmxnet3"
	providerIDPrefix      = "vsphere://"
	regionKey             = "region"
	zoneKey               = "zone"
)

// These are the guestinfo variables used by Ignition.
// https://access.redhat.com/documentation/en-us/openshift_container_platform/4.1/html/installing/installing-on-vsphere
const (
	GuestInfoIgnitionData     = "guestinfo.ignition.config.data"
	GuestInfoIgnitionEncoding = "guestinfo.ignition.config.data.encoding"
	GuestInfoHostname         = "guestinfo.hostname"
)

// Reconciler runs the logic to reconciles a machine resource towards its desired state
type Reconciler struct {
	*machineScope
}

func newReconciler(scope *machineScope) *Reconciler {
	return &Reconciler{
		machineScope: scope,
	}
}

// create creates machine if it does not exists.
func (r *Reconciler) create() error {
	if err := validateMachine(*r.machine); err != nil {
		return fmt.Errorf("%v: failed validating machine provider spec: %v", r.machine.GetName(), err)
	}

	moTask, err := r.session.GetTask(r.Context, r.providerStatus.TaskRef)
	if err != nil {
		if !isRetrieveMONotFound(r.providerStatus.TaskRef, err) {
			return err
		}
	}
	if taskIsFinished, err := taskIsFinished(moTask); err != nil || !taskIsFinished {
		if !taskIsFinished {
			return fmt.Errorf("task %v has not finished", moTask.Reference().Value)
		}
		return err
	}

	if _, err := findVM(r.machineScope); err != nil {
		if !isNotFound(err) {
			return err
		}
		if r.machineScope.session.IsVC() {
			klog.Infof("%v: cloning", r.machine.GetName())
			task, err := clone(r.machineScope)
			if err != nil {
				conditionFailed := conditionFailed()
				conditionFailed.Message = err.Error()
				statusError := setProviderStatus(task, conditionFailed, r.machineScope, nil)
				if statusError != nil {
					return errors.Wrap(err, "Failed to set provider status")
				}
				return err
			}

			return setProviderStatus(task, conditionSuccess(), r.machineScope, nil)
		}
		return fmt.Errorf("%v: not connected to a vCenter", r.machine.GetName())
	}

	return nil
}

// update finds a vm and reconciles the machine resource status against it.
func (r *Reconciler) update() error {
	if err := validateMachine(*r.machine); err != nil {
		return fmt.Errorf("%v: failed validating machine provider spec: %v", r.machine.GetName(), err)
	}

	motask, err := r.session.GetTask(r.Context, r.providerStatus.TaskRef)
	if err != nil {
		if !isRetrieveMONotFound(r.providerStatus.TaskRef, err) {
			return err
		}
	}
	if taskIsFinished, err := taskIsFinished(motask); err != nil || !taskIsFinished {
		if !taskIsFinished {
			return fmt.Errorf("task %v has not finished", motask.Reference().Value)
		}
		return err
	}

	vmRef, err := findVM(r.machineScope)
	if err != nil {
		if !isNotFound(err) {
			return err
		}
		return errors.Wrap(err, "vm not found on update")
	}

	vm := &virtualMachine{
		Context: r.machineScope.Context,
		Obj:     object.NewVirtualMachine(r.machineScope.session.Client.Client, vmRef),
		Ref:     vmRef,
	}

	if err := vm.reconcileTags(r.Context, r.session, r.machine); err != nil {
		return errors.Wrapf(err, "failed to reconcile tags")
	}

	// TODO: we won't always want to reconcile power state
	//  but as per comment in clone() function, powering on right on creation might be problematic
	ok, task, err := vm.reconcilePowerState()
	if err != nil || !ok {
		return err
	}

	return r.reconcileMachineWithCloudState(vm, task)
}

// exists returns true if machine exists.
func (r *Reconciler) exists() (bool, error) {
	if err := validateMachine(*r.machine); err != nil {
		return false, fmt.Errorf("%v: failed validating machine provider spec: %v", r.machine.GetName(), err)
	}

	if _, err := findVM(r.machineScope); err != nil {
		if !isNotFound(err) {
			return false, err
		}
		klog.Infof("%v: does not exist", r.machine.GetName())
		return false, nil
	}
	klog.Infof("%v: already exists", r.machine.GetName())
	return true, nil
}

func (r *Reconciler) delete() error {
	moTask, err := r.session.GetTask(r.Context, r.providerStatus.TaskRef)
	if err != nil {
		if !isRetrieveMONotFound(r.providerStatus.TaskRef, err) {
			return err
		}
	}
	if taskIsFinished, err := taskIsFinished(moTask); err != nil || !taskIsFinished {
		if !taskIsFinished {
			return fmt.Errorf("task %v has not finished", moTask.Reference().Value)
		}
		return err
	}

	vmRef, err := findVM(r.machineScope)
	if err != nil {
		if !isNotFound(err) {
			return err
		}
		klog.Infof("%v: vm does not exist", r.machine.GetName())
		return nil
	}

	vm := &virtualMachine{
		Context: r.Context,
		Obj:     object.NewVirtualMachine(r.machineScope.session.Client.Client, vmRef),
		Ref:     vmRef,
	}

	if _, err := vm.powerOffVM(); err != nil {
		return err
	}

	task, err := vm.Obj.Destroy(r.Context)
	if err != nil {
		return fmt.Errorf("%v: failed to destroy vm: %v", r.machine.GetName(), err)
	}

	if err := setProviderStatus(task.Reference().Value, conditionSuccess(), r.machineScope, vm); err != nil {
		return errors.Wrap(err, "Failed to set provider status")
	}

	return fmt.Errorf("destroying vm in progress, reconciling")
}

// reconcileMachineWithCloudState reconcile machineSpec and status with the latest cloud state
func (r *Reconciler) reconcileMachineWithCloudState(vm *virtualMachine, taskRef string) error {
	klog.V(3).Infof("%v: reconciling machine with cloud state", r.machine.GetName())
	// TODO: reconcile task

	if err := r.reconcileRegionAndZoneLabels(vm); err != nil {
		// Not treating this is as a fatal error for now.
		klog.Errorf("Failed to reconcile region and zone labels: %v", err)
	}

	klog.V(3).Infof("%v: reconciling providerID", r.machine.GetName())
	if err := r.reconcileProviderID(vm); err != nil {
		return err
	}

	klog.V(3).Infof("%v: reconciling network", r.machine.GetName())
	if err := r.reconcileNetwork(vm); err != nil {
		return err
	}

	return setProviderStatus(taskRef, conditionSuccess(), r.machineScope, vm)
}

// reconcileRegionAndZoneLabels reconciles the labels on the Machine containing
// region and zone information -- provided the vSphere cloud provider has been
// configured with the labels that identify region and zone, and the configured
// tags are found somewhere in the ancestry of the given virtual machine.
func (r *Reconciler) reconcileRegionAndZoneLabels(vm *virtualMachine) error {
	if r.vSphereConfig == nil {
		klog.Warning("No vSphere cloud provider config. " +
			"Will not set region and zone labels.")
		return nil
	}

	regionLabel := r.vSphereConfig.Labels.Region
	zoneLabel := r.vSphereConfig.Labels.Zone

	var res map[string]string

	err := r.session.WithRestClient(vm.Context, func(c *rest.Client) error {
		var err error
		res, err = vm.getRegionAndZone(c, regionLabel, zoneLabel)

		return err
	})

	if err != nil {
		return err
	}

	if r.machine.Labels == nil {
		r.machine.Labels = make(map[string]string)
	}

	r.machine.Labels[machinecontroller.MachineRegionLabelName] = res[regionKey]
	r.machine.Labels[machinecontroller.MachineAZLabelName] = res[zoneKey]

	return nil
}

func (r *Reconciler) reconcileProviderID(vm *virtualMachine) error {
	providerID, err := convertUUIDToProviderID(vm.Obj.UUID(vm.Context))
	if err != nil {
		return err
	}
	r.machine.Spec.ProviderID = &providerID
	return nil
}

// convertUUIDToProviderID transforms a UUID string into a provider ID.
func convertUUIDToProviderID(UUID string) (string, error) {
	parsedUUID, err := uuid.Parse(UUID)
	if err != nil {
		return "", err
	}
	return providerIDPrefix + parsedUUID.String(), nil
}

func (r *Reconciler) reconcileNetwork(vm *virtualMachine) error {
	currentNetworkStatusList, err := vm.getNetworkStatusList(r.session.Client.Client)
	if err != nil {
		return fmt.Errorf("error getting network status: %v", err)
	}

	//If the VM is powered on then issue requeues until all of the VM's
	//networks have IP addresses.
	expectNetworkLen, currentNetworkLen := len(r.providerSpec.Network.Devices), len(currentNetworkStatusList)
	if expectNetworkLen != currentNetworkLen {
		return errors.Errorf("invalid network count: expected=%d current=%d", expectNetworkLen, currentNetworkLen)
	}

	var ipAddrs []corev1.NodeAddress
	for _, netStatus := range currentNetworkStatusList {
		for _, ip := range netStatus.IPAddrs {
			ipAddrs = append(ipAddrs, corev1.NodeAddress{
				Type:    corev1.NodeInternalIP,
				Address: ip,
			})
		}
	}

	// Using Name() if InventoryPath is empty will return empty name
	// see: https://github.com/vmware/govmomi/blob/master/object/common.go#L66-L75
	// Using ObjectName() as it will query from VirtualMachine properties

	vmName, err := vm.Obj.ObjectName(vm.Context)
	if err != nil {
		return fmt.Errorf("error getting virtual machine name: %v", err)
	}

	ipAddrs = append(ipAddrs, corev1.NodeAddress{
		Type:    corev1.NodeInternalDNS,
		Address: vmName,
	})

	klog.V(3).Infof("%v: reconciling network: IP addresses: %v", r.machine.GetName(), ipAddrs)
	r.machine.Status.Addresses = ipAddrs
	return nil
}

func validateMachine(machine machinev1.Machine) error {
	if machine.Labels[machinev1.MachineClusterIDLabel] == "" {
		return machinecontroller.InvalidMachineConfiguration("%v: missing %q label", machine.GetName(), machinev1.MachineClusterIDLabel)
	}

	return nil
}

func findVM(s *machineScope) (types.ManagedObjectReference, error) {
	uuid := string(s.machine.UID)
	objRef, err := s.GetSession().FindRefByInstanceUUID(s, uuid)
	if err != nil {
		return types.ManagedObjectReference{}, err
	}
	if objRef == nil {
		return types.ManagedObjectReference{}, errNotFound{instanceUUID: true, uuid: uuid}
	}
	return objRef.Reference(), nil
}

// errNotFound is returned by the findVM function when a VM is not found.
type errNotFound struct {
	instanceUUID bool
	uuid         string
}

func (e errNotFound) Error() string {
	if e.instanceUUID {
		return fmt.Sprintf("vm with instance uuid %s not found", e.uuid)
	}
	return fmt.Sprintf("vm with bios uuid %s not found", e.uuid)
}

func isNotFound(err error) bool {
	switch err.(type) {
	case errNotFound, *errNotFound:
		return true
	default:
		return false
	}
}

func isRetrieveMONotFound(taskRef string, err error) bool {
	return err.Error() == fmt.Sprintf("ServerFaultCode: The object 'vim.Task:%v' has already been deleted or has not been completely created", taskRef)
}

func clone(s *machineScope) (string, error) {
	userData, err := s.GetUserData()
	if err != nil {
		return "", err
	}

	vmTemplate, err := s.GetSession().FindVM(*s, s.providerSpec.Template)
	if err != nil {
		return "", err
	}

	var snapshotRef *types.ManagedObjectReference

	// If a linked clone is requested then a MoRef for a snapshot must be
	// found with which to perform the linked clone.
	// Empty clone mode is linked clone
	if s.providerSpec.CloneMode == "" || s.providerSpec.CloneMode == vspherev1.LinkedClone {
		if s.providerSpec.Snapshot == "" {
			klog.V(3).Infof("%v: no snapshot name provided, getting snapshot using template", s.machine.GetName())
			var vm mo.VirtualMachine
			if err := vmTemplate.Properties(s.Context, vmTemplate.Reference(), []string{"snapshot"}, &vm); err != nil {
				return "", errors.Wrapf(err, "error getting snapshot information for template %s", vmTemplate.Name())
			}

			if vm.Snapshot != nil {
				snapshotRef = vm.Snapshot.CurrentSnapshot
			}
		} else {
			klog.V(3).Infof("%v: searching for snapshot by name %s", s.machine.GetName(), s.providerSpec.Snapshot)
			var err error
			snapshotRef, err = vmTemplate.FindSnapshot(s.Context, s.providerSpec.Snapshot)
			if err != nil {
				klog.V(3).Infof("%v: failed to find snapshot %s", s.machine.GetName(), s.providerSpec.Snapshot)
			}
		}
	}

	// The type of clone operation depends on whether or not there is a snapshot
	// from which to do a linked clone.
	diskMoveType := fullCloneDiskMoveType
	if snapshotRef != nil {
		// TODO: write clone mode to status
		diskMoveType = linkCloneDiskMoveType
	}

	var folderPath, datastorePath, resourcepoolPath string
	if s.providerSpec.Workspace != nil {
		folderPath = s.providerSpec.Workspace.Folder
		datastorePath = s.providerSpec.Workspace.Datastore
		resourcepoolPath = s.providerSpec.Workspace.ResourcePool
	}

	folder, err := s.GetSession().Finder.FolderOrDefault(s, folderPath)
	if err != nil {
		return "", errors.Wrapf(err, "unable to get folder for %q", folderPath)
	}

	datastore, err := s.GetSession().Finder.DatastoreOrDefault(s, datastorePath)
	if err != nil {
		return "", errors.Wrapf(err, "unable to get datastore for %q", datastorePath)
	}

	resourcepool, err := s.GetSession().Finder.ResourcePoolOrDefault(s, resourcepoolPath)
	if err != nil {
		return "", errors.Wrapf(err, "unable to get resource pool for %q", resourcepool)
	}

	numCPUs := s.providerSpec.NumCPUs
	if numCPUs < minCPU {
		numCPUs = minCPU
	}
	numCoresPerSocket := s.providerSpec.NumCoresPerSocket
	if numCoresPerSocket == 0 {
		numCoresPerSocket = numCPUs
	}
	memMiB := s.providerSpec.MemoryMiB
	if memMiB == 0 {
		memMiB = minMemMB
	}

	devices, err := vmTemplate.Device(s.Context)
	if err != nil {
		return "", fmt.Errorf("error getting devices %v", err)
	}

	// Create a new list of device specs for cloning the VM.
	deviceSpecs := []types.BaseVirtualDeviceConfigSpec{}

	// Only non-linked clones may expand the size of the template's disk.
	if snapshotRef == nil {
		diskSpec, err := getDiskSpec(s, devices)
		if err != nil {
			return "", errors.Wrapf(err, "error getting disk spec for %q", s.providerSpec.Snapshot)
		}
		deviceSpecs = append(deviceSpecs, diskSpec)
	}

	klog.V(3).Infof("Getting network devices")
	networkDevices, err := getNetworkDevices(s, devices)
	if err != nil {
		return "", fmt.Errorf("error getting network specs: %v", err)
	}

	deviceSpecs = append(deviceSpecs, networkDevices...)

	extraConfig := []types.BaseOptionValue{}

	extraConfig = append(extraConfig, IgnitionConfig(userData)...)
	extraConfig = append(extraConfig, &types.OptionValue{
		Key:   GuestInfoHostname,
		Value: s.machine.GetName(),
	})

	spec := types.VirtualMachineCloneSpec{
		Config: &types.VirtualMachineConfigSpec{
			Annotation: s.machine.GetName(),
			// Assign the clone's InstanceUUID the value of the Kubernetes Machine
			// object's UID. This allows lookup of the cloned VM prior to knowing
			// the VM's UUID.
			InstanceUuid:      string(s.machine.UID),
			Flags:             newVMFlagInfo(),
			ExtraConfig:       extraConfig,
			DeviceChange:      deviceSpecs,
			NumCPUs:           numCPUs,
			NumCoresPerSocket: numCoresPerSocket,
			MemoryMB:          memMiB,
		},
		Location: types.VirtualMachineRelocateSpec{
			Datastore:    types.NewReference(datastore.Reference()),
			Folder:       types.NewReference(folder.Reference()),
			Pool:         types.NewReference(resourcepool.Reference()),
			DiskMoveType: diskMoveType,
		},
		// This is implicit, but making it explicit as it is important to not
		// power the VM on before its virtual hardware is created and the MAC
		// address(es) used to build and inject the VM with cloud-init metadata
		// are generated.
		PowerOn:  false,
		Snapshot: snapshotRef,
	}

	task, err := vmTemplate.Clone(s, folder, s.machine.GetName(), spec)
	if err != nil {
		return "", errors.Wrapf(err, "error triggering clone op for machine %v", s)
	}

	klog.V(3).Infof("%v: running task: %+v", s.machine.GetName(), s.providerStatus.TaskRef)
	return task.Reference().Value, nil
}

func getDiskSpec(s *machineScope, devices object.VirtualDeviceList) (types.BaseVirtualDeviceConfigSpec, error) {
	disks := devices.SelectByType((*types.VirtualDisk)(nil))
	if len(disks) != 1 {
		return nil, errors.Errorf("invalid disk count: %d", len(disks))
	}

	disk := disks[0].(*types.VirtualDisk)
	disk.CapacityInKB = int64(s.providerSpec.DiskGiB) * 1024 * 1024

	return &types.VirtualDeviceConfigSpec{
		Operation: types.VirtualDeviceConfigSpecOperationEdit,
		Device:    disk,
	}, nil
}

func getNetworkDevices(s *machineScope, devices object.VirtualDeviceList) ([]types.BaseVirtualDeviceConfigSpec, error) {
	var networkDevices []types.BaseVirtualDeviceConfigSpec
	// Remove any existing NICs
	for _, dev := range devices.SelectByType((*types.VirtualEthernetCard)(nil)) {
		networkDevices = append(networkDevices, &types.VirtualDeviceConfigSpec{
			Device:    dev,
			Operation: types.VirtualDeviceConfigSpecOperationRemove,
		})
	}

	// Add new NICs based on the machine config.
	for i := range s.providerSpec.Network.Devices {
		netSpec := &s.providerSpec.Network.Devices[i]
		klog.V(3).Infof("Adding device: %v", netSpec.NetworkName)

		ref, err := s.GetSession().Finder.Network(s.Context, netSpec.NetworkName)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to find network %q", netSpec.NetworkName)
		}

		backing, err := ref.EthernetCardBackingInfo(s.Context)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to create new ethernet card backing info for network %q", netSpec.NetworkName)
		}

		dev, err := object.EthernetCardTypes().CreateEthernetCard(ethCardType, backing)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to create new ethernet card %q for network %q", ethCardType, netSpec.NetworkName)
		}

		// Get the actual NIC object. This is safe to assert without a check
		// because "object.EthernetCardTypes().CreateEthernetCard" returns a
		// "types.BaseVirtualEthernetCard" as a "types.BaseVirtualDevice".
		nic := dev.(types.BaseVirtualEthernetCard).GetVirtualEthernetCard()
		// Assign a temporary device key to ensure that a unique one will be
		// generated when the device is created.
		nic.Key = int32(i)

		networkDevices = append(networkDevices, &types.VirtualDeviceConfigSpec{
			Device:    dev,
			Operation: types.VirtualDeviceConfigSpecOperationAdd,
		})
		klog.V(3).Infof("Adding device: eth card type: %v, network spec: %+v, device info: %+v",
			ethCardType, netSpec, dev.GetVirtualDevice().Backing)
	}

	return networkDevices, nil
}

func newVMFlagInfo() *types.VirtualMachineFlagInfo {
	diskUUIDEnabled := true
	return &types.VirtualMachineFlagInfo{
		DiskUuidEnabled: &diskUUIDEnabled,
	}
}

func taskIsFinished(task *mo.Task) (bool, error) {
	if task == nil {
		return true, nil
	}

	// Otherwise the course of action is determined by the state of the task.
	klog.V(3).Infof("task: %v, state: %v, description-id: %v", task.Reference().Value, task.Info.State, task.Info.DescriptionId)
	switch task.Info.State {
	case types.TaskInfoStateQueued:
		return false, nil
	case types.TaskInfoStateRunning:
		return false, nil
	case types.TaskInfoStateSuccess:
		return true, nil
	case types.TaskInfoStateError:
		return true, nil
	default:
		return false, errors.Errorf("task: %v, unknown state %v", task.Reference().Value, task.Info.State)
	}
}

func setProviderStatus(taskRef string, condition vspherev1.VSphereMachineProviderCondition, scope *machineScope, vm *virtualMachine) error {
	klog.Infof("%s: Updating provider status", scope.machine.Name)

	if vm != nil {
		id := vm.Obj.UUID(scope.Context)
		scope.providerStatus.InstanceID = &id

		// This can return an error if machine is being deleted
		powerState, err := vm.getPowerState()
		if err != nil {
			klog.V(3).Infof("%s: Failed to get power state during provider status update: %v", scope.machine.Name, err)
		} else {
			powerStateString := string(powerState)
			scope.providerStatus.InstanceState = &powerStateString
		}
	}

	if taskRef != "" {
		scope.providerStatus.TaskRef = taskRef
	}

	scope.providerStatus.Conditions = setVSphereMachineProviderConditions(condition, scope.providerStatus.Conditions)

	return nil
}

type virtualMachine struct {
	context.Context
	Ref types.ManagedObjectReference
	Obj *object.VirtualMachine
}

func (vm *virtualMachine) getAncestors() ([]mo.ManagedEntity, error) {
	client := vm.Obj.Client()
	pc := client.ServiceContent.PropertyCollector

	return mo.Ancestors(vm.Context, client, pc, vm.Ref)
}

// getRegionAndZone checks the virtual machine and each of its ancestors for the
// given region and zone labels and returns their values if found.
func (vm *virtualMachine) getRegionAndZone(c *rest.Client, regionLabel, zoneLabel string) (map[string]string, error) {
	result := make(map[string]string)
	tagsMgr := tags.NewManager(c)

	objects, err := vm.getAncestors()
	if err != nil {
		klog.Errorf("Failed to get ancestors for %s: %v", vm.Ref, err)
		return nil, err
	}

	for i := range objects {
		obj := objects[len(objects)-1-i] // Reverse order.
		klog.V(4).Infof("getRegionAndZone: Name: %s, Type: %s",
			obj.Self.Value, obj.Self.Type)

		tags, err := tagsMgr.ListAttachedTags(vm.Context, obj)
		if err != nil {
			klog.Warningf("Failed to list attached tags: %v", err)
			return nil, err
		}

		for _, value := range tags {
			tag, err := tagsMgr.GetTag(vm.Context, value)
			if err != nil {
				klog.Errorf("Failed to get tag: %v", err)
				return nil, err
			}

			category, err := tagsMgr.GetCategory(vm.Context, tag.CategoryID)
			if err != nil {
				klog.Errorf("Failed to get tag category: %v", err)
				return nil, err
			}

			switch {
			case regionLabel != "" && category.Name == regionLabel:
				result[regionKey] = tag.Name
				klog.V(2).Infof("%s has region tag (%s) with value %s",
					vm.Ref, category.Name, tag.Name)

			case zoneLabel != "" && category.Name == zoneLabel:
				result[zoneKey] = tag.Name
				klog.V(2).Infof("%s has zone tag (%s) with value %s",
					vm.Ref, category.Name, tag.Name)
			}

			// We've found both tags, return early.
			if result[regionKey] != "" && result[zoneKey] != "" {
				return result, nil
			}
		}
	}

	return result, nil
}

func (vm *virtualMachine) reconcilePowerState() (bool, string, error) {
	powerState, err := vm.getPowerState()
	if err != nil {
		return false, "", err
	}
	switch powerState {
	case types.VirtualMachinePowerStatePoweredOff:
		klog.Infof("%v: powering on", vm.Obj.Reference().Value)
		task, err := vm.powerOnVM()
		if err != nil {
			return false, "", errors.Wrapf(err, "failed to trigger power on op for vm %q", vm)
		}

		klog.Infof("%v: requeue to wait for power on state", vm.Obj.Reference().Value)
		return false, task, nil
	case types.VirtualMachinePowerStatePoweredOn:
		klog.Infof("%v: powered on", vm.Obj.Reference().Value)
		return true, "", nil
	default:
		return false, "", errors.Errorf("unexpected power state %q for vm %q", powerState, vm)
	}
}

func (vm *virtualMachine) powerOnVM() (string, error) {
	task, err := vm.Obj.PowerOn(vm.Context)
	if err != nil {
		return "", err
	}
	return task.Reference().Value, nil
}

func (vm *virtualMachine) powerOffVM() (string, error) {
	task, err := vm.Obj.PowerOff(vm.Context)
	if err != nil {
		return "", err
	}
	return task.Reference().Value, nil
}

func (vm *virtualMachine) getPowerState() (types.VirtualMachinePowerState, error) {
	powerState, err := vm.Obj.PowerState(vm.Context)
	if err != nil {
		return "", err
	}

	switch powerState {
	case types.VirtualMachinePowerStatePoweredOn:
		return types.VirtualMachinePowerStatePoweredOn, nil
	case types.VirtualMachinePowerStatePoweredOff:
		return types.VirtualMachinePowerStatePoweredOff, nil
	case types.VirtualMachinePowerStateSuspended:
		return types.VirtualMachinePowerStateSuspended, nil
	default:
		return "", errors.Errorf("unexpected power state %q for vm %v", powerState, vm)
	}
}

// reconcileTags ensures that the required tags are present on the virtual machine, eg the Cluster ID
// that is used by the installer on cluster deletion to ensure ther are no leaked resources.
func (vm *virtualMachine) reconcileTags(ctx context.Context, session *session.Session, machine *machinev1.Machine) error {
	if err := session.WithRestClient(vm.Context, func(c *rest.Client) error {
		klog.Infof("%v: Reconciling attached tags", machine.GetName())

		m := tags.NewManager(c)

		clusterID := machine.Labels[machinev1.MachineClusterIDLabel]

		// the tag should already be created by installer
		if err := m.AttachTag(ctx, clusterID, vm.Ref); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

type NetworkStatus struct {
	// Connected is a flag that indicates whether this network is currently
	// connected to the VM.
	Connected bool

	// IPAddrs is one or more IP addresses reported by vm-tools.
	IPAddrs []string

	// MACAddr is the MAC address of the network device.
	MACAddr string

	// NetworkName is the name of the network.
	NetworkName string
}

func (vm *virtualMachine) getNetworkStatusList(client *vim25.Client) ([]NetworkStatus, error) {
	var obj mo.VirtualMachine
	var pc = property.DefaultCollector(client)
	var props = []string{
		"config.hardware.device",
		"guest.net",
	}

	if err := pc.RetrieveOne(vm.Context, vm.Ref, props, &obj); err != nil {
		return nil, errors.Wrapf(err, "unable to fetch props %v for vm %v", props, vm.Ref)
	}
	klog.V(3).Infof("Getting network status: object reference: %v", obj.Reference().Value)
	if obj.Config == nil {
		return nil, errors.New("config.hardware.device is nil")
	}

	var networkStatusList []NetworkStatus
	for _, device := range obj.Config.Hardware.Device {
		if dev, ok := device.(types.BaseVirtualEthernetCard); ok {
			nic := dev.GetVirtualEthernetCard()
			klog.V(3).Infof("Getting network status: device: %v, macAddress: %v", nic.DeviceInfo.GetDescription().Summary, nic.MacAddress)
			netStatus := NetworkStatus{
				MACAddr: nic.MacAddress,
			}
			if obj.Guest != nil {
				klog.V(3).Infof("Getting network status: getting guest info")
				for _, i := range obj.Guest.Net {
					klog.V(3).Infof("Getting network status: getting guest info: network: %+v", i)
					if strings.EqualFold(nic.MacAddress, i.MacAddress) {
						//TODO: sanitizeIPAddrs
						netStatus.IPAddrs = i.IpAddress
						netStatus.NetworkName = i.Network
						netStatus.Connected = i.Connected
					}
				}
			}
			networkStatusList = append(networkStatusList, netStatus)
		}
	}

	return networkStatusList, nil
}

// IgnitionConfig returns a slice of option values that set the given data as
// the guest's ignition config.
func IgnitionConfig(data []byte) []types.BaseOptionValue {
	config := EncodeIgnitionConfig(data)

	if config == "" {
		return nil
	}

	return []types.BaseOptionValue{
		&types.OptionValue{
			Key:   GuestInfoIgnitionData,
			Value: config,
		},
		&types.OptionValue{
			Key:   GuestInfoIgnitionEncoding,
			Value: "base64",
		},
	}
}

// EncodeIgnitionConfig attempts to decode the given data until it looks to be
// plain-text, then returns a base64 encoded version of that plain-text.
func EncodeIgnitionConfig(data []byte) string {
	if len(data) == 0 {
		return ""
	}

	for {
		decoded, err := base64.StdEncoding.DecodeString(string(data))
		if err != nil {
			break
		}

		data = decoded
	}

	return base64.StdEncoding.EncodeToString(data)
}
