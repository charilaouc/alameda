// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/common/queries.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type TimeRange_AggregateFunction int32

const (
	TimeRange_NONE TimeRange_AggregateFunction = 0
	TimeRange_MAX  TimeRange_AggregateFunction = 1
	TimeRange_AVG  TimeRange_AggregateFunction = 2
)

var TimeRange_AggregateFunction_name = map[int32]string{
	0: "NONE",
	1: "MAX",
	2: "AVG",
}

var TimeRange_AggregateFunction_value = map[string]int32{
	"NONE": 0,
	"MAX":  1,
	"AVG":  2,
}

func (x TimeRange_AggregateFunction) String() string {
	return proto.EnumName(TimeRange_AggregateFunction_name, int32(x))
}

func (TimeRange_AggregateFunction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d602763cab07305c, []int{0, 0}
}

type QueryCondition_Order int32

const (
	QueryCondition_NONE QueryCondition_Order = 0
	QueryCondition_ASC  QueryCondition_Order = 1
	QueryCondition_DESC QueryCondition_Order = 2
)

var QueryCondition_Order_name = map[int32]string{
	0: "NONE",
	1: "ASC",
	2: "DESC",
}

var QueryCondition_Order_value = map[string]int32{
	"NONE": 0,
	"ASC":  1,
	"DESC": 2,
}

func (x QueryCondition_Order) String() string {
	return proto.EnumName(QueryCondition_Order_name, int32(x))
}

func (QueryCondition_Order) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d602763cab07305c, []int{4, 0}
}

//*
// Represents a time range definition
//
type TimeRange struct {
	StartTime            *timestamp.Timestamp        `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp        `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Step                 *duration.Duration          `protobuf:"bytes,3,opt,name=step,proto3" json:"step,omitempty"`
	AggregateFunction    TimeRange_AggregateFunction `protobuf:"varint,4,opt,name=aggregateFunction,proto3,enum=containersai.alameda.v1alpha1.datahub.common.TimeRange_AggregateFunction" json:"aggregateFunction,omitempty"`
	ApplyTime            *timestamp.Timestamp        `protobuf:"bytes,5,opt,name=apply_time,json=applyTime,proto3" json:"apply_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TimeRange) Reset()         { *m = TimeRange{} }
func (m *TimeRange) String() string { return proto.CompactTextString(m) }
func (*TimeRange) ProtoMessage()    {}
func (*TimeRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_d602763cab07305c, []int{0}
}

func (m *TimeRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRange.Unmarshal(m, b)
}
func (m *TimeRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRange.Marshal(b, m, deterministic)
}
func (m *TimeRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRange.Merge(m, src)
}
func (m *TimeRange) XXX_Size() int {
	return xxx_messageInfo_TimeRange.Size(m)
}
func (m *TimeRange) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRange.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRange proto.InternalMessageInfo

func (m *TimeRange) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TimeRange) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *TimeRange) GetStep() *duration.Duration {
	if m != nil {
		return m.Step
	}
	return nil
}

func (m *TimeRange) GetAggregateFunction() TimeRange_AggregateFunction {
	if m != nil {
		return m.AggregateFunction
	}
	return TimeRange_NONE
}

func (m *TimeRange) GetApplyTime() *timestamp.Timestamp {
	if m != nil {
		return m.ApplyTime
	}
	return nil
}

type Condition struct {
	Keys                 []string   `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	Values               []string   `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	Operators            []string   `protobuf:"bytes,3,rep,name=operators,proto3" json:"operators,omitempty"`
	Types                []DataType `protobuf:"varint,4,rep,packed,name=types,proto3,enum=containersai.alameda.v1alpha1.datahub.common.DataType" json:"types,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Condition) Reset()         { *m = Condition{} }
func (m *Condition) String() string { return proto.CompactTextString(m) }
func (*Condition) ProtoMessage()    {}
func (*Condition) Descriptor() ([]byte, []int) {
	return fileDescriptor_d602763cab07305c, []int{1}
}

func (m *Condition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Condition.Unmarshal(m, b)
}
func (m *Condition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Condition.Marshal(b, m, deterministic)
}
func (m *Condition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Condition.Merge(m, src)
}
func (m *Condition) XXX_Size() int {
	return xxx_messageInfo_Condition.Size(m)
}
func (m *Condition) XXX_DiscardUnknown() {
	xxx_messageInfo_Condition.DiscardUnknown(m)
}

var xxx_messageInfo_Condition proto.InternalMessageInfo

func (m *Condition) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Condition) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Condition) GetOperators() []string {
	if m != nil {
		return m.Operators
	}
	return nil
}

func (m *Condition) GetTypes() []DataType {
	if m != nil {
		return m.Types
	}
	return nil
}

type Function struct {
	Type                 FunctionType `protobuf:"varint,1,opt,name=type,proto3,enum=containersai.alameda.v1alpha1.datahub.common.FunctionType" json:"type,omitempty"`
	Fields               []string     `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	Tags                 []string     `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	Target               string       `protobuf:"bytes,4,opt,name=target,proto3" json:"target,omitempty"`
	RegularExpression    string       `protobuf:"bytes,5,opt,name=regular_expression,json=regularExpression,proto3" json:"regular_expression,omitempty"`
	Unit                 string       `protobuf:"bytes,6,opt,name=unit,proto3" json:"unit,omitempty"`
	Number               int64        `protobuf:"varint,7,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Function) Reset()         { *m = Function{} }
func (m *Function) String() string { return proto.CompactTextString(m) }
func (*Function) ProtoMessage()    {}
func (*Function) Descriptor() ([]byte, []int) {
	return fileDescriptor_d602763cab07305c, []int{2}
}

func (m *Function) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Function.Unmarshal(m, b)
}
func (m *Function) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Function.Marshal(b, m, deterministic)
}
func (m *Function) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Function.Merge(m, src)
}
func (m *Function) XXX_Size() int {
	return xxx_messageInfo_Function.Size(m)
}
func (m *Function) XXX_DiscardUnknown() {
	xxx_messageInfo_Function.DiscardUnknown(m)
}

var xxx_messageInfo_Function proto.InternalMessageInfo

func (m *Function) GetType() FunctionType {
	if m != nil {
		return m.Type
	}
	return FunctionType_FUNCTIONTYPE_UNDEFINED
}

func (m *Function) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Function) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Function) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *Function) GetRegularExpression() string {
	if m != nil {
		return m.RegularExpression
	}
	return ""
}

func (m *Function) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *Function) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type Into struct {
	Database                 string   `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	RetentionPolicy          string   `protobuf:"bytes,2,opt,name=retention_policy,json=retentionPolicy,proto3" json:"retention_policy,omitempty"`
	Measurement              string   `protobuf:"bytes,3,opt,name=measurement,proto3" json:"measurement,omitempty"`
	IsDefaultRetentionPolicy bool     `protobuf:"varint,4,opt,name=is_default_retention_policy,json=isDefaultRetentionPolicy,proto3" json:"is_default_retention_policy,omitempty"`
	IsAllMeasurements        bool     `protobuf:"varint,5,opt,name=is_all_measurements,json=isAllMeasurements,proto3" json:"is_all_measurements,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *Into) Reset()         { *m = Into{} }
func (m *Into) String() string { return proto.CompactTextString(m) }
func (*Into) ProtoMessage()    {}
func (*Into) Descriptor() ([]byte, []int) {
	return fileDescriptor_d602763cab07305c, []int{3}
}

func (m *Into) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Into.Unmarshal(m, b)
}
func (m *Into) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Into.Marshal(b, m, deterministic)
}
func (m *Into) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Into.Merge(m, src)
}
func (m *Into) XXX_Size() int {
	return xxx_messageInfo_Into.Size(m)
}
func (m *Into) XXX_DiscardUnknown() {
	xxx_messageInfo_Into.DiscardUnknown(m)
}

var xxx_messageInfo_Into proto.InternalMessageInfo

func (m *Into) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *Into) GetRetentionPolicy() string {
	if m != nil {
		return m.RetentionPolicy
	}
	return ""
}

func (m *Into) GetMeasurement() string {
	if m != nil {
		return m.Measurement
	}
	return ""
}

func (m *Into) GetIsDefaultRetentionPolicy() bool {
	if m != nil {
		return m.IsDefaultRetentionPolicy
	}
	return false
}

func (m *Into) GetIsAllMeasurements() bool {
	if m != nil {
		return m.IsAllMeasurements
	}
	return false
}

type QueryCondition struct {
	TimeRange            *TimeRange           `protobuf:"bytes,1,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	Order                QueryCondition_Order `protobuf:"varint,2,opt,name=order,proto3,enum=containersai.alameda.v1alpha1.datahub.common.QueryCondition_Order" json:"order,omitempty"`
	Function             *Function            `protobuf:"bytes,3,opt,name=function,proto3" json:"function,omitempty"`
	Into                 *Into                `protobuf:"bytes,4,opt,name=into,proto3" json:"into,omitempty"`
	WhereClause          string               `protobuf:"bytes,5,opt,name=where_clause,json=whereClause,proto3" json:"where_clause,omitempty"`
	WhereCondition       []*Condition         `protobuf:"bytes,6,rep,name=where_condition,json=whereCondition,proto3" json:"where_condition,omitempty"`
	Selects              []string             `protobuf:"bytes,7,rep,name=selects,proto3" json:"selects,omitempty"`
	Groups               []string             `protobuf:"bytes,8,rep,name=groups,proto3" json:"groups,omitempty"`
	Limit                uint64               `protobuf:"varint,9,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *QueryCondition) Reset()         { *m = QueryCondition{} }
func (m *QueryCondition) String() string { return proto.CompactTextString(m) }
func (*QueryCondition) ProtoMessage()    {}
func (*QueryCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_d602763cab07305c, []int{4}
}

func (m *QueryCondition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryCondition.Unmarshal(m, b)
}
func (m *QueryCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryCondition.Marshal(b, m, deterministic)
}
func (m *QueryCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCondition.Merge(m, src)
}
func (m *QueryCondition) XXX_Size() int {
	return xxx_messageInfo_QueryCondition.Size(m)
}
func (m *QueryCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCondition.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCondition proto.InternalMessageInfo

func (m *QueryCondition) GetTimeRange() *TimeRange {
	if m != nil {
		return m.TimeRange
	}
	return nil
}

func (m *QueryCondition) GetOrder() QueryCondition_Order {
	if m != nil {
		return m.Order
	}
	return QueryCondition_NONE
}

func (m *QueryCondition) GetFunction() *Function {
	if m != nil {
		return m.Function
	}
	return nil
}

func (m *QueryCondition) GetInto() *Into {
	if m != nil {
		return m.Into
	}
	return nil
}

func (m *QueryCondition) GetWhereClause() string {
	if m != nil {
		return m.WhereClause
	}
	return ""
}

func (m *QueryCondition) GetWhereCondition() []*Condition {
	if m != nil {
		return m.WhereCondition
	}
	return nil
}

func (m *QueryCondition) GetSelects() []string {
	if m != nil {
		return m.Selects
	}
	return nil
}

func (m *QueryCondition) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *QueryCondition) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func init() {
	proto.RegisterEnum("containersai.alameda.v1alpha1.datahub.common.TimeRange_AggregateFunction", TimeRange_AggregateFunction_name, TimeRange_AggregateFunction_value)
	proto.RegisterEnum("containersai.alameda.v1alpha1.datahub.common.QueryCondition_Order", QueryCondition_Order_name, QueryCondition_Order_value)
	proto.RegisterType((*TimeRange)(nil), "containersai.alameda.v1alpha1.datahub.common.TimeRange")
	proto.RegisterType((*Condition)(nil), "containersai.alameda.v1alpha1.datahub.common.Condition")
	proto.RegisterType((*Function)(nil), "containersai.alameda.v1alpha1.datahub.common.Function")
	proto.RegisterType((*Into)(nil), "containersai.alameda.v1alpha1.datahub.common.Into")
	proto.RegisterType((*QueryCondition)(nil), "containersai.alameda.v1alpha1.datahub.common.QueryCondition")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/common/queries.proto", fileDescriptor_d602763cab07305c)
}

var fileDescriptor_d602763cab07305c = []byte{
	// 814 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0x8d, 0xd3, 0xc4, 0xa7, 0x28, 0x9b, 0x0e, 0x08, 0x99, 0x82, 0xc0, 0x44, 0x5c, 0x04,
	0x89, 0xda, 0x6a, 0x10, 0x8b, 0x40, 0x42, 0x22, 0xfd, 0x59, 0xb4, 0x12, 0xdb, 0x85, 0xd9, 0x6a,
	0xb5, 0xe2, 0xc6, 0x4c, 0xe2, 0x53, 0x77, 0xc4, 0xd8, 0x63, 0x66, 0xc6, 0xbb, 0xe4, 0x31, 0x78,
	0x06, 0x9e, 0x88, 0x17, 0xe0, 0x3d, 0xb8, 0x43, 0x33, 0xfe, 0x69, 0xb7, 0xbd, 0xd8, 0x4d, 0xef,
	0xe6, 0x9c, 0xf3, 0x9d, 0x6f, 0xf2, 0x7d, 0x39, 0x73, 0x0c, 0x47, 0x4c, 0xb0, 0x02, 0x33, 0x96,
	0xb2, 0x8a, 0x27, 0x2f, 0x8f, 0x98, 0xa8, 0xae, 0xd8, 0x51, 0x92, 0x31, 0xc3, 0xae, 0xea, 0x55,
	0xb2, 0x96, 0x45, 0x21, 0xcb, 0xe4, 0x8f, 0x1a, 0x15, 0x47, 0x1d, 0x57, 0x4a, 0x1a, 0x49, 0xbe,
	0x5c, 0xcb, 0xd2, 0x30, 0x5e, 0xa2, 0xd2, 0x8c, 0xc7, 0x6d, 0x7f, 0xdc, 0xf5, 0xc6, 0x6d, 0x6f,
	0xdc, 0xf4, 0x1e, 0x24, 0x6f, 0x73, 0x81, 0xd9, 0x54, 0x1d, 0xfd, 0xc1, 0x27, 0xb9, 0x94, 0xb9,
	0xc0, 0xc4, 0x45, 0xab, 0xfa, 0x32, 0xc9, 0x6a, 0xc5, 0x0c, 0x97, 0x65, 0x5b, 0xff, 0xf4, 0x76,
	0xdd, 0xf0, 0x02, 0xb5, 0x61, 0x45, 0xd5, 0x00, 0x66, 0x7f, 0x0d, 0x20, 0xb8, 0xe0, 0x05, 0x52,
	0x56, 0xe6, 0x48, 0xbe, 0x05, 0xd0, 0x86, 0x29, 0x93, 0x5a, 0x58, 0xe8, 0x45, 0xde, 0x7c, 0x6f,
	0x71, 0x10, 0x37, 0x1c, 0x71, 0xc7, 0x11, 0x5f, 0x74, 0x1c, 0x34, 0x70, 0x68, 0x1b, 0x93, 0xaf,
	0x61, 0x8c, 0x65, 0xd6, 0x34, 0xee, 0xbc, 0xb1, 0x71, 0x84, 0x65, 0xe6, 0xda, 0x0e, 0xc1, 0xd7,
	0x06, 0xab, 0x70, 0xe0, 0x5a, 0x3e, 0xbc, 0xd3, 0x72, 0xda, 0xea, 0xa1, 0x0e, 0x46, 0x5e, 0xc1,
	0x3e, 0xcb, 0x73, 0x85, 0x39, 0x33, 0xf8, 0xa8, 0x2e, 0xd7, 0xb6, 0x14, 0xfa, 0x91, 0x37, 0x9f,
	0x2c, 0x1e, 0xc7, 0xdb, 0x58, 0x1d, 0xf7, 0xa2, 0xe3, 0xe5, 0x6d, 0x42, 0x7a, 0xf7, 0x0e, 0xeb,
	0x0c, 0xab, 0x2a, 0xb1, 0x69, 0x04, 0x0e, 0xdf, 0xec, 0x8c, 0x43, 0xdb, 0x78, 0x96, 0xc0, 0xfe,
	0x9d, 0x2b, 0xc8, 0x18, 0xfc, 0xf3, 0xa7, 0xe7, 0x67, 0xd3, 0x77, 0xc8, 0x08, 0x06, 0x4f, 0x96,
	0x2f, 0xa6, 0x9e, 0x3d, 0x2c, 0x9f, 0xff, 0x38, 0xdd, 0x99, 0xfd, 0xed, 0x41, 0x70, 0x22, 0xcb,
	0x8c, 0x3b, 0x24, 0x01, 0xff, 0x77, 0xdc, 0xe8, 0xd0, 0x8b, 0x06, 0xf3, 0x80, 0xba, 0x33, 0xf9,
	0x00, 0x76, 0x5f, 0x32, 0x51, 0xa3, 0x0e, 0x77, 0x5c, 0xb6, 0x8d, 0xc8, 0xc7, 0x10, 0xc8, 0x0a,
	0x15, 0x33, 0x52, 0xe9, 0x70, 0xe0, 0x4a, 0xd7, 0x09, 0xf2, 0x13, 0x0c, 0xdd, 0xec, 0x84, 0x7e,
	0x34, 0x98, 0x4f, 0x16, 0x0f, 0xb7, 0x33, 0xec, 0x94, 0x19, 0x76, 0xb1, 0xa9, 0x90, 0x36, 0x24,
	0xb3, 0xff, 0x3c, 0x18, 0xf7, 0x72, 0xce, 0xc1, 0xb7, 0x59, 0x37, 0x32, 0x93, 0xc5, 0x77, 0xdb,
	0x31, 0x77, 0x2c, 0x8e, 0xdd, 0xf1, 0x58, 0x81, 0x97, 0x1c, 0x45, 0xd6, 0x0b, 0x6c, 0x22, 0x6b,
	0x86, 0x61, 0x79, 0xa7, 0xcd, 0x9d, 0x2d, 0xd6, 0x30, 0x95, 0xa3, 0x71, 0x83, 0x10, 0xd0, 0x36,
	0x22, 0x87, 0x40, 0x14, 0xe6, 0xb5, 0x60, 0x2a, 0xc5, 0x3f, 0x2b, 0x85, 0x5a, 0xdb, 0x61, 0x19,
	0x3a, 0xcc, 0x7e, 0x5b, 0x39, 0xeb, 0x0b, 0x96, 0xba, 0x2e, 0xb9, 0x09, 0x77, 0x1d, 0xc0, 0x9d,
	0x2d, 0x75, 0x59, 0x17, 0x2b, 0x54, 0xe1, 0x28, 0xf2, 0xe6, 0x03, 0xda, 0x46, 0xb3, 0x7f, 0x3d,
	0xf0, 0x1f, 0x97, 0x46, 0x92, 0x03, 0x18, 0x5b, 0x31, 0x2b, 0xa6, 0x1b, 0xed, 0x01, 0xed, 0x63,
	0xf2, 0x05, 0x4c, 0x15, 0x1a, 0x2c, 0xad, 0xb4, 0xb4, 0x92, 0x82, 0xaf, 0x37, 0xee, 0x65, 0x04,
	0xf4, 0x41, 0x9f, 0xff, 0xd9, 0xa5, 0x49, 0x04, 0x7b, 0x05, 0x32, 0x5d, 0x2b, 0x2c, 0xb0, 0x34,
	0xee, 0x31, 0x04, 0xf4, 0x66, 0x8a, 0x7c, 0x0f, 0x1f, 0x71, 0x9d, 0x66, 0x78, 0xc9, 0x6a, 0x61,
	0xd2, 0x3b, 0xbc, 0x56, 0xf9, 0x98, 0x86, 0x5c, 0x9f, 0x36, 0x08, 0x7a, 0xeb, 0x82, 0x18, 0xde,
	0xe3, 0x3a, 0x65, 0x42, 0xa4, 0x37, 0x48, 0xb5, 0x33, 0x63, 0x4c, 0xf7, 0xb9, 0x5e, 0x0a, 0xf1,
	0xe4, 0x46, 0x61, 0xf6, 0x8f, 0x0f, 0x93, 0x5f, 0x6a, 0x54, 0x9b, 0xeb, 0x39, 0x7c, 0x0e, 0x60,
	0x67, 0x3f, 0x55, 0xf6, 0xd1, 0xb4, 0xbb, 0xe1, 0x9b, 0x7b, 0xbe, 0x39, 0x1a, 0x98, 0x7e, 0xe7,
	0xbc, 0x80, 0xa1, 0x54, 0x19, 0x2a, 0xe7, 0xcd, 0x64, 0x71, 0xbc, 0x1d, 0xe5, 0xeb, 0x3f, 0x32,
	0x7e, 0x6a, 0x99, 0x68, 0x43, 0x48, 0x28, 0x8c, 0x2f, 0xbb, 0x1d, 0xd1, 0xec, 0x97, 0x87, 0xf7,
	0x1b, 0x4c, 0xda, 0xf3, 0x90, 0x47, 0xe0, 0xf3, 0xd2, 0x48, 0x67, 0xf8, 0xde, 0x62, 0xb1, 0x1d,
	0x9f, 0x1d, 0x19, 0xea, 0xfa, 0xc9, 0x67, 0xf0, 0xee, 0xab, 0x2b, 0x54, 0x98, 0xae, 0x05, 0xab,
	0x35, 0xb6, 0x63, 0xb9, 0xe7, 0x72, 0x27, 0x2e, 0x45, 0x7e, 0x83, 0x07, 0x2d, 0xa4, 0x93, 0x17,
	0xee, 0x46, 0x83, 0xed, 0x5d, 0xef, 0xdd, 0xa1, 0x93, 0x86, 0xbe, 0xff, 0x4b, 0x43, 0x18, 0x69,
	0x14, 0xb8, 0x36, 0x3a, 0x1c, 0xb9, 0x07, 0xd5, 0x85, 0x76, 0xf0, 0x73, 0x25, 0xeb, 0x4a, 0x87,
	0xe3, 0xe6, 0xfd, 0x35, 0x11, 0x79, 0x1f, 0x86, 0x82, 0x17, 0xdc, 0x84, 0x41, 0xe4, 0xcd, 0x7d,
	0xda, 0x04, 0xb3, 0xcf, 0x61, 0xe8, 0x8c, 0x7f, 0x7d, 0xab, 0x2d, 0x9f, 0x9d, 0x4c, 0x3d, 0x9b,
	0x3a, 0x3d, 0x7b, 0x76, 0x32, 0xdd, 0x39, 0x3e, 0xfe, 0xf5, 0x87, 0x9c, 0x9b, 0xf6, 0x87, 0x25,
	0xd7, 0x12, 0x0e, 0x19, 0x4f, 0xec, 0xf7, 0xee, 0x2d, 0xbe, 0x7d, 0xab, 0x5d, 0xb7, 0x6a, 0xbf,
	0xfa, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x59, 0x15, 0xc4, 0x8a, 0x07, 0x00, 0x00,
}
