package nodeinfo

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/containers-ai/alameda/datahub/pkg/entities"
	operatorutils "github.com/containers-ai/alameda/operator/pkg/utils"
	"github.com/containers-ai/alameda/pkg/provider"
	mahcinev1beta1 "github.com/openshift/machine-api-operator/pkg/apis/machine/v1beta1"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type role = string

const (
	masterRole role = "master"
	workerRole role = "worker"

	defaultNodeStorageSize = "100Gi"
)

var (
	roleMap = map[string]role{
		"node-role.kubernetes.io/master": masterRole,
	}
)

// NodeInfo flats the k8s node information from labels, spec and status
type NodeInfo struct {
	Name                string
	CreatedTime         int64
	Namespace           string
	Kind                string
	Role                string
	Region              string
	Zone                string
	Size                string
	InstanceType        string
	OS                  string
	Provider            string
	InstanceID          string
	StorageSize         int64
	CPUCores            int64
	MemoryBytes         int64
	MachineSetNamespace string
	MachineSetName      string
	MachineCreateTime   int64
}

// NewNodeInfo creates node from k8s node
func NewNodeInfo(k8sNode corev1.Node, k8sClient client.Client) (NodeInfo, error) {
	machimeList := mahcinev1beta1.MachineList{}
	err := k8sClient.List(context.Background(), &machimeList, &client.ListOptions{})
	if err != nil {
	}

	node := NodeInfo{Name: k8sNode.Name, Namespace: k8sNode.Namespace, Kind: k8sNode.Kind}
	for _, ms := range machimeList.Items {
		if ms.GetName() == k8sNode.GetName() {
			for _, or := range ms.GetOwnerReferences() {
				if or.Kind == "MachineSet" {
					node.MachineSetNamespace = ms.GetNamespace()
					node.MachineSetName = or.Name
					node.MachineCreateTime = ms.GetCreationTimestamp().Unix()
					break
				}
			}
			break
		}
	}

	rf := reflect.TypeOf(node)
	rv := reflect.ValueOf(&node).Elem()
	for i := 0; i < rf.NumField(); i++ {
		key := rf.Field(i).Name
		// parse node label information
		for labelKey, labelV := range k8sNode.Labels {
			if strings.Contains(labelKey, "stackpoint.") && strings.Contains(labelKey, "stackpoint.io/role") == false {
				continue
			}
			value := parseKeyValue(labelKey, key, labelV)
			if len(value) > 0 {
				rValue := rv.FieldByName(strings.Title(key))
				rValue.SetString(string(labelV))
				break
			}
		}
		switch key {
		case "StorageSize":
			node.StorageSize = k8sNode.Status.Capacity.StorageEphemeral().Value()
		}
	}

	if node.Role == "" {
		node.patchRoleByK8SLabels(k8sNode.Labels)
	}

	if len(k8sNode.Spec.ProviderID) > 0 {
		provider, _, instanceID := parseProviderID(k8sNode.Spec.ProviderID)
		node.Provider = provider
		node.InstanceID = instanceID
	}

	// Below ard original convert logic
	node.CreatedTime = k8sNode.ObjectMeta.GetCreationTimestamp().Unix()

	cpuCores, ok := k8sNode.Status.Capacity.Cpu().AsInt64()
	if !ok {
		return NodeInfo{}, errors.Errorf("cannot convert cpu capacity from k8s Node")
	}
	node.CPUCores = cpuCores

	memoryBytes, ok := k8sNode.Status.Capacity.Memory().AsInt64()
	if !ok {
		return NodeInfo{}, errors.Errorf("cannot convert memory capacity from k8s Node")
	}
	node.MemoryBytes = memoryBytes

	if regionMap, exist := provider.ProviderRegionMap[node.Provider]; exist {
		if region, exist := regionMap[node.Region]; exist {
			node.Region = region
		}
	}

	node.setDefaultValue()

	return node, nil
}

// DatahubNode converts nodeInfo to Datahub Node
func (n NodeInfo) DatahubNode(clusterUID string) entities.ResourceClusterStatusNode {

	node := entities.ResourceClusterStatusNode{
		Name:                n.Name,
		ClusterName:         clusterUID,
		NodeCPUCores:        n.CPUCores,
		NodeMemoryBytes:     n.MemoryBytes,
		CreateTime:          n.CreatedTime,
		MachinesetNamespace: n.MachineSetNamespace,
		MachinesetName:      n.MachineSetName,
		IOProvider:          n.Provider,
		IOInstanceType:      n.InstanceType,
		IORegion:            n.Region,
		IOZone:              n.Zone,
		IOOs:                n.OS,
		IORole:              n.Role,
		IOInstanceId:        n.InstanceID,
		IOStorageSize:       n.StorageSize,
	}

	return node
}

func (n *NodeInfo) patchRoleByK8SLabels(labels map[string]string) {
	found := false
	for key, role := range roleMap {
		if _, exist := labels[key]; exist {
			found = true
			n.Role = role
			break
		}
	}
	if !found {
		n.Role = workerRole
	}
}

func (n *NodeInfo) setDefaultValue() {

	storageSize := operatorutils.GetNodeInfoDefaultStorageSizeBytes()
	if storageSize == "" {
		storageSize = defaultNodeStorageSize
	}
	defaultNodeStorageQuantity := resource.MustParse(storageSize)
	if n.StorageSize == 0 {
		n.StorageSize = defaultNodeStorageQuantity.Value()
	}
}

func parseKeyValue(strParse string, key string, value string) string {
	pattern, err := regexp.Compile(strings.ToLower(fmt.Sprintf("/%s$", key)))
	if err != nil {
		return ""
	}
	if len(pattern.FindString(strings.Replace(strParse, "-", "", -1))) > 0 {
		return value
	}
	return ""
}

func parseProviderID(providerID string) (string, string, string) {
	var provider string
	var region string
	var instanceID string
	rex, err := regexp.Compile("([^\\:/]+)")
	if err != nil {
		fmt.Println(err)
		return "", "", ""
	}
	res := rex.FindAllString(providerID, -1)
	if res == nil || len(res) == 0 {
		return "", "", ""
	}
	for i := 0; i < len(res) && i < 3; i++ {
		switch i {
		case 0:
			provider = res[i]
		case 1:
			region = res[i]
		case 2:
			instanceID = res[i]
		}
	}
	return provider, region, instanceID
}
