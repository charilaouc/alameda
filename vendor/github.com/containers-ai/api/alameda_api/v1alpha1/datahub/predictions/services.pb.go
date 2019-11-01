// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/predictions/services.proto

package predictions

import (
	fmt "fmt"
	common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	proto "github.com/golang/protobuf/proto"
	status "google.golang.org/genproto/googleapis/rpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Represents a request for creating predictions of containers' metric data belonging to a pod
type CreatePodPredictionsRequest struct {
	PodPredictions       []*PodPrediction `protobuf:"bytes,1,rep,name=pod_predictions,json=podPredictions,proto3" json:"pod_predictions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CreatePodPredictionsRequest) Reset()         { *m = CreatePodPredictionsRequest{} }
func (m *CreatePodPredictionsRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePodPredictionsRequest) ProtoMessage()    {}
func (*CreatePodPredictionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f99ee8b129e06c96, []int{0}
}

func (m *CreatePodPredictionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePodPredictionsRequest.Unmarshal(m, b)
}
func (m *CreatePodPredictionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePodPredictionsRequest.Marshal(b, m, deterministic)
}
func (m *CreatePodPredictionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePodPredictionsRequest.Merge(m, src)
}
func (m *CreatePodPredictionsRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePodPredictionsRequest.Size(m)
}
func (m *CreatePodPredictionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePodPredictionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePodPredictionsRequest proto.InternalMessageInfo

func (m *CreatePodPredictionsRequest) GetPodPredictions() []*PodPrediction {
	if m != nil {
		return m.PodPredictions
	}
	return nil
}

// Represents a request for creating predictions of a node metric data
type CreateNodePredictionsRequest struct {
	NodePredictions      []*NodePrediction `protobuf:"bytes,1,rep,name=node_predictions,json=nodePredictions,proto3" json:"node_predictions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateNodePredictionsRequest) Reset()         { *m = CreateNodePredictionsRequest{} }
func (m *CreateNodePredictionsRequest) String() string { return proto.CompactTextString(m) }
func (*CreateNodePredictionsRequest) ProtoMessage()    {}
func (*CreateNodePredictionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f99ee8b129e06c96, []int{1}
}

func (m *CreateNodePredictionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNodePredictionsRequest.Unmarshal(m, b)
}
func (m *CreateNodePredictionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNodePredictionsRequest.Marshal(b, m, deterministic)
}
func (m *CreateNodePredictionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNodePredictionsRequest.Merge(m, src)
}
func (m *CreateNodePredictionsRequest) XXX_Size() int {
	return xxx_messageInfo_CreateNodePredictionsRequest.Size(m)
}
func (m *CreateNodePredictionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNodePredictionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNodePredictionsRequest proto.InternalMessageInfo

func (m *CreateNodePredictionsRequest) GetNodePredictions() []*NodePrediction {
	if m != nil {
		return m.NodePredictions
	}
	return nil
}

// Represents a request for listing predictions of pods
type ListPodPredictionsRequest struct {
	QueryCondition       *common.QueryCondition    `protobuf:"bytes,1,opt,name=query_condition,json=queryCondition,proto3" json:"query_condition,omitempty"`
	NamespacedName       *resources.NamespacedName `protobuf:"bytes,2,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	Granularity          int64                     `protobuf:"varint,3,opt,name=granularity,proto3" json:"granularity,omitempty"`
	FillDays             int64                     `protobuf:"varint,4,opt,name=fill_days,json=fillDays,proto3" json:"fill_days,omitempty"`
	ModelId              string                    `protobuf:"bytes,5,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	PredictionId         string                    `protobuf:"bytes,6,opt,name=prediction_id,json=predictionId,proto3" json:"prediction_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ListPodPredictionsRequest) Reset()         { *m = ListPodPredictionsRequest{} }
func (m *ListPodPredictionsRequest) String() string { return proto.CompactTextString(m) }
func (*ListPodPredictionsRequest) ProtoMessage()    {}
func (*ListPodPredictionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f99ee8b129e06c96, []int{2}
}

func (m *ListPodPredictionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPodPredictionsRequest.Unmarshal(m, b)
}
func (m *ListPodPredictionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPodPredictionsRequest.Marshal(b, m, deterministic)
}
func (m *ListPodPredictionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPodPredictionsRequest.Merge(m, src)
}
func (m *ListPodPredictionsRequest) XXX_Size() int {
	return xxx_messageInfo_ListPodPredictionsRequest.Size(m)
}
func (m *ListPodPredictionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPodPredictionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPodPredictionsRequest proto.InternalMessageInfo

func (m *ListPodPredictionsRequest) GetQueryCondition() *common.QueryCondition {
	if m != nil {
		return m.QueryCondition
	}
	return nil
}

func (m *ListPodPredictionsRequest) GetNamespacedName() *resources.NamespacedName {
	if m != nil {
		return m.NamespacedName
	}
	return nil
}

func (m *ListPodPredictionsRequest) GetGranularity() int64 {
	if m != nil {
		return m.Granularity
	}
	return 0
}

func (m *ListPodPredictionsRequest) GetFillDays() int64 {
	if m != nil {
		return m.FillDays
	}
	return 0
}

func (m *ListPodPredictionsRequest) GetModelId() string {
	if m != nil {
		return m.ModelId
	}
	return ""
}

func (m *ListPodPredictionsRequest) GetPredictionId() string {
	if m != nil {
		return m.PredictionId
	}
	return ""
}

// Represents a request for listing predictions of nodes
type ListNodePredictionsRequest struct {
	QueryCondition       *common.QueryCondition `protobuf:"bytes,1,opt,name=query_condition,json=queryCondition,proto3" json:"query_condition,omitempty"`
	NodeNames            []string               `protobuf:"bytes,2,rep,name=node_names,json=nodeNames,proto3" json:"node_names,omitempty"`
	Granularity          int64                  `protobuf:"varint,3,opt,name=granularity,proto3" json:"granularity,omitempty"`
	ModelId              string                 `protobuf:"bytes,4,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	PredictionId         string                 `protobuf:"bytes,5,opt,name=prediction_id,json=predictionId,proto3" json:"prediction_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ListNodePredictionsRequest) Reset()         { *m = ListNodePredictionsRequest{} }
func (m *ListNodePredictionsRequest) String() string { return proto.CompactTextString(m) }
func (*ListNodePredictionsRequest) ProtoMessage()    {}
func (*ListNodePredictionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f99ee8b129e06c96, []int{3}
}

func (m *ListNodePredictionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNodePredictionsRequest.Unmarshal(m, b)
}
func (m *ListNodePredictionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNodePredictionsRequest.Marshal(b, m, deterministic)
}
func (m *ListNodePredictionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNodePredictionsRequest.Merge(m, src)
}
func (m *ListNodePredictionsRequest) XXX_Size() int {
	return xxx_messageInfo_ListNodePredictionsRequest.Size(m)
}
func (m *ListNodePredictionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNodePredictionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListNodePredictionsRequest proto.InternalMessageInfo

func (m *ListNodePredictionsRequest) GetQueryCondition() *common.QueryCondition {
	if m != nil {
		return m.QueryCondition
	}
	return nil
}

func (m *ListNodePredictionsRequest) GetNodeNames() []string {
	if m != nil {
		return m.NodeNames
	}
	return nil
}

func (m *ListNodePredictionsRequest) GetGranularity() int64 {
	if m != nil {
		return m.Granularity
	}
	return 0
}

func (m *ListNodePredictionsRequest) GetModelId() string {
	if m != nil {
		return m.ModelId
	}
	return ""
}

func (m *ListNodePredictionsRequest) GetPredictionId() string {
	if m != nil {
		return m.PredictionId
	}
	return ""
}

// Represents a response for a listing predictions of pods request
type ListPodPredictionsResponse struct {
	Status               *status.Status   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	PodPredictions       []*PodPrediction `protobuf:"bytes,2,rep,name=pod_predictions,json=podPredictions,proto3" json:"pod_predictions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListPodPredictionsResponse) Reset()         { *m = ListPodPredictionsResponse{} }
func (m *ListPodPredictionsResponse) String() string { return proto.CompactTextString(m) }
func (*ListPodPredictionsResponse) ProtoMessage()    {}
func (*ListPodPredictionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f99ee8b129e06c96, []int{4}
}

func (m *ListPodPredictionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPodPredictionsResponse.Unmarshal(m, b)
}
func (m *ListPodPredictionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPodPredictionsResponse.Marshal(b, m, deterministic)
}
func (m *ListPodPredictionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPodPredictionsResponse.Merge(m, src)
}
func (m *ListPodPredictionsResponse) XXX_Size() int {
	return xxx_messageInfo_ListPodPredictionsResponse.Size(m)
}
func (m *ListPodPredictionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPodPredictionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPodPredictionsResponse proto.InternalMessageInfo

func (m *ListPodPredictionsResponse) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListPodPredictionsResponse) GetPodPredictions() []*PodPrediction {
	if m != nil {
		return m.PodPredictions
	}
	return nil
}

// Represents a response for a listing predictions of nodes request
type ListNodePredictionsResponse struct {
	Status               *status.Status    `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	NodePredictions      []*NodePrediction `protobuf:"bytes,2,rep,name=node_predictions,json=nodePredictions,proto3" json:"node_predictions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListNodePredictionsResponse) Reset()         { *m = ListNodePredictionsResponse{} }
func (m *ListNodePredictionsResponse) String() string { return proto.CompactTextString(m) }
func (*ListNodePredictionsResponse) ProtoMessage()    {}
func (*ListNodePredictionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f99ee8b129e06c96, []int{5}
}

func (m *ListNodePredictionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNodePredictionsResponse.Unmarshal(m, b)
}
func (m *ListNodePredictionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNodePredictionsResponse.Marshal(b, m, deterministic)
}
func (m *ListNodePredictionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNodePredictionsResponse.Merge(m, src)
}
func (m *ListNodePredictionsResponse) XXX_Size() int {
	return xxx_messageInfo_ListNodePredictionsResponse.Size(m)
}
func (m *ListNodePredictionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNodePredictionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListNodePredictionsResponse proto.InternalMessageInfo

func (m *ListNodePredictionsResponse) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListNodePredictionsResponse) GetNodePredictions() []*NodePrediction {
	if m != nil {
		return m.NodePredictions
	}
	return nil
}

func init() {
	proto.RegisterType((*CreatePodPredictionsRequest)(nil), "containersai.alameda.v1alpha1.datahub.predictions.CreatePodPredictionsRequest")
	proto.RegisterType((*CreateNodePredictionsRequest)(nil), "containersai.alameda.v1alpha1.datahub.predictions.CreateNodePredictionsRequest")
	proto.RegisterType((*ListPodPredictionsRequest)(nil), "containersai.alameda.v1alpha1.datahub.predictions.ListPodPredictionsRequest")
	proto.RegisterType((*ListNodePredictionsRequest)(nil), "containersai.alameda.v1alpha1.datahub.predictions.ListNodePredictionsRequest")
	proto.RegisterType((*ListPodPredictionsResponse)(nil), "containersai.alameda.v1alpha1.datahub.predictions.ListPodPredictionsResponse")
	proto.RegisterType((*ListNodePredictionsResponse)(nil), "containersai.alameda.v1alpha1.datahub.predictions.ListNodePredictionsResponse")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/predictions/services.proto", fileDescriptor_f99ee8b129e06c96)
}

var fileDescriptor_f99ee8b129e06c96 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x5f, 0x8b, 0xd3, 0x4e,
	0x14, 0x25, 0xe9, 0x6e, 0x7f, 0xdb, 0xe9, 0xcf, 0x56, 0xf2, 0x62, 0xb6, 0x55, 0x08, 0xf5, 0xa5,
	0x08, 0x4e, 0xe8, 0x8a, 0x08, 0x22, 0xf8, 0x67, 0x15, 0x59, 0x90, 0x65, 0x8d, 0x6f, 0xbe, 0x84,
	0xbb, 0x99, 0x6b, 0x3b, 0x90, 0xcc, 0xa4, 0x33, 0x93, 0x85, 0xbe, 0xfb, 0xe0, 0x83, 0xdf, 0xc6,
	0x07, 0x3f, 0x94, 0x5f, 0x42, 0x26, 0x4d, 0xb7, 0x8d, 0xc4, 0x35, 0x2e, 0xac, 0x6f, 0x33, 0xf7,
	0x66, 0xce, 0x39, 0x39, 0x73, 0xcf, 0x90, 0x27, 0x90, 0x42, 0x86, 0x0c, 0x62, 0xc8, 0x79, 0x78,
	0x31, 0x83, 0x34, 0x5f, 0xc0, 0x2c, 0x64, 0x60, 0x60, 0x51, 0x9c, 0x87, 0xb9, 0x42, 0xc6, 0x13,
	0xc3, 0xa5, 0xd0, 0xa1, 0x46, 0x75, 0xc1, 0x13, 0xd4, 0x34, 0x57, 0xd2, 0x48, 0x6f, 0x96, 0x48,
	0x61, 0x80, 0x0b, 0x54, 0x1a, 0x38, 0xad, 0x50, 0xe8, 0x06, 0x81, 0x56, 0x08, 0x74, 0x07, 0x61,
	0x34, 0xbb, 0x92, 0x2b, 0x91, 0x59, 0x26, 0x45, 0xb8, 0x2c, 0x50, 0xf1, 0x0d, 0xcb, 0xe8, 0x69,
	0x6b, 0x79, 0x3b, 0xeb, 0xea, 0xec, 0xe3, 0x2b, 0xcf, 0x2a, 0xd4, 0xb2, 0x50, 0x09, 0xea, 0x30,
	0x43, 0x03, 0xb6, 0x5a, 0x1d, 0xbb, 0x33, 0x97, 0x72, 0x9e, 0x62, 0xa8, 0xf2, 0x24, 0xd4, 0x06,
	0x4c, 0x51, 0xe1, 0x4d, 0xbe, 0x38, 0x64, 0x7c, 0xac, 0x10, 0x0c, 0x9e, 0x49, 0x76, 0xb6, 0xa5,
	0x8b, 0x70, 0x59, 0xa0, 0x36, 0x1e, 0x27, 0xc3, 0x5c, 0xb2, 0x78, 0x47, 0x88, 0xef, 0x04, 0x9d,
	0x69, 0xff, 0xe8, 0x05, 0xfd, 0x6b, 0xaf, 0x68, 0x8d, 0x22, 0x1a, 0xe4, 0x35, 0xc6, 0xc9, 0x57,
	0x87, 0xdc, 0x5d, 0x4b, 0x39, 0x95, 0x0c, 0x1b, 0xb4, 0xa4, 0xe4, 0xb6, 0x90, 0x0c, 0x1b, 0xc4,
	0xbc, 0xbc, 0x86, 0x98, 0x3a, 0x49, 0x34, 0x14, 0x75, 0xd2, 0xc9, 0x0f, 0x97, 0x1c, 0xbe, 0xe3,
	0xda, 0x34, 0xfb, 0x82, 0x64, 0x68, 0x2f, 0x75, 0x15, 0x27, 0x52, 0x30, 0x6e, 0x5b, 0xbe, 0x13,
	0x38, 0xd3, 0xfe, 0xd1, 0xb3, 0x96, 0x52, 0xd6, 0x93, 0x41, 0xdf, 0x5b, 0x90, 0xe3, 0x0d, 0x46,
	0x34, 0x58, 0xd6, 0xf6, 0xde, 0x82, 0x0c, 0x05, 0x64, 0xa8, 0x73, 0x48, 0x90, 0xc5, 0x76, 0xe9,
	0xbb, 0x25, 0xcd, 0xf3, 0x96, 0x34, 0x97, 0x13, 0x41, 0x4f, 0x2f, 0x71, 0xec, 0x2a, 0x1a, 0x88,
	0xda, 0xde, 0x0b, 0x48, 0x7f, 0xae, 0x40, 0x14, 0x29, 0x28, 0x6e, 0x56, 0x7e, 0x27, 0x70, 0xa6,
	0x9d, 0x68, 0xb7, 0xe4, 0x8d, 0x49, 0xef, 0x13, 0x4f, 0xd3, 0x98, 0xc1, 0x4a, 0xfb, 0x7b, 0x65,
	0xff, 0xc0, 0x16, 0x5e, 0xc3, 0x4a, 0x7b, 0x87, 0xe4, 0x20, 0x93, 0x0c, 0xd3, 0x98, 0x33, 0x7f,
	0x3f, 0x70, 0xa6, 0xbd, 0xe8, 0xbf, 0x72, 0x7f, 0xc2, 0xbc, 0xfb, 0xe4, 0xd6, 0xd6, 0x77, 0xdb,
	0xef, 0x96, 0xfd, 0xff, 0xb7, 0xc5, 0x13, 0x36, 0xf9, 0xec, 0x92, 0x91, 0x75, 0xfb, 0x37, 0x57,
	0xff, 0x8f, 0xec, 0xbe, 0x47, 0x48, 0x39, 0x61, 0xa5, 0x37, 0xbe, 0x1b, 0x74, 0xa6, 0xbd, 0xa8,
	0x67, 0x2b, 0xa5, 0x79, 0x2d, 0x3c, 0xda, 0xb5, 0x61, 0xef, 0x0f, 0x36, 0xec, 0x37, 0xd8, 0xf0,
	0xcd, 0x59, 0xdb, 0xf0, 0xeb, 0xd0, 0xe9, 0x5c, 0x0a, 0x8d, 0xde, 0x03, 0xd2, 0x5d, 0xa7, 0xb7,
	0xfa, 0x7b, 0x8f, 0xae, 0x73, 0x4d, 0x55, 0x9e, 0xd0, 0x0f, 0x65, 0x27, 0xaa, 0xbe, 0x68, 0x4a,
	0xae, 0x7b, 0x43, 0xc9, 0xfd, 0xee, 0x90, 0x71, 0xe3, 0xe5, 0x5d, 0x43, 0x76, 0x53, 0xc8, 0xdd,
	0x9b, 0x0a, 0xf9, 0xab, 0xb7, 0x1f, 0xdf, 0xcc, 0xb9, 0xa9, 0x06, 0x24, 0xdc, 0xe2, 0x3f, 0x04,
	0x1e, 0xda, 0x17, 0xb6, 0xed, 0x4b, 0x7d, 0xde, 0x2d, 0x9f, 0xd3, 0x47, 0x3f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xfd, 0x19, 0xaa, 0x97, 0x7b, 0x06, 0x00, 0x00,
}