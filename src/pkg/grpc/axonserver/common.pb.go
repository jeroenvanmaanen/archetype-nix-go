// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package axonserver

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
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

// An enumeration of possible keys for processing instructions.
type ProcessingKey int32

const (
	// key indicating that the attached value should be used for consistent routing.
	ProcessingKey_ROUTING_KEY ProcessingKey = 0
	// key indicating that the attached value indicates relative priority of this message.
	ProcessingKey_PRIORITY ProcessingKey = 1
	// key indicating that the accompanied message has a finite validity. The attached value contains the number of milliseconds.
	ProcessingKey_TIMEOUT ProcessingKey = 2
	// key indicating that the requester expects at most the given number of results from this message. Use -1 for unlimited.
	ProcessingKey_NR_OF_RESULTS ProcessingKey = 3
)

var ProcessingKey_name = map[int32]string{
	0: "ROUTING_KEY",
	1: "PRIORITY",
	2: "TIMEOUT",
	3: "NR_OF_RESULTS",
}

var ProcessingKey_value = map[string]int32{
	"ROUTING_KEY":   0,
	"PRIORITY":      1,
	"TIMEOUT":       2,
	"NR_OF_RESULTS": 3,
}

func (x ProcessingKey) String() string {
	return proto.EnumName(ProcessingKey_name, int32(x))
}

func (ProcessingKey) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

// Describes a serialized object
type SerializedObject struct {
	// The type identifier of the serialized object.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The revision of the serialized form of the given type.
	Revision string `protobuf:"bytes,2,opt,name=revision,proto3" json:"revision,omitempty"`
	// The actual data representing the object in serialized form.
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SerializedObject) Reset()         { *m = SerializedObject{} }
func (m *SerializedObject) String() string { return proto.CompactTextString(m) }
func (*SerializedObject) ProtoMessage()    {}
func (*SerializedObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *SerializedObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SerializedObject.Unmarshal(m, b)
}
func (m *SerializedObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SerializedObject.Marshal(b, m, deterministic)
}
func (m *SerializedObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SerializedObject.Merge(m, src)
}
func (m *SerializedObject) XXX_Size() int {
	return xxx_messageInfo_SerializedObject.Size(m)
}
func (m *SerializedObject) XXX_DiscardUnknown() {
	xxx_messageInfo_SerializedObject.DiscardUnknown(m)
}

var xxx_messageInfo_SerializedObject proto.InternalMessageInfo

func (m *SerializedObject) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SerializedObject) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *SerializedObject) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// The value of a MetaData entry.
type MetaDataValue struct {
	// The data of the MetaData entry, depending on the type of data it contains.
	//
	// Types that are valid to be assigned to Data:
	//	*MetaDataValue_TextValue
	//	*MetaDataValue_NumberValue
	//	*MetaDataValue_BooleanValue
	//	*MetaDataValue_DoubleValue
	//	*MetaDataValue_BytesValue
	Data                 isMetaDataValue_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MetaDataValue) Reset()         { *m = MetaDataValue{} }
func (m *MetaDataValue) String() string { return proto.CompactTextString(m) }
func (*MetaDataValue) ProtoMessage()    {}
func (*MetaDataValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *MetaDataValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaDataValue.Unmarshal(m, b)
}
func (m *MetaDataValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaDataValue.Marshal(b, m, deterministic)
}
func (m *MetaDataValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaDataValue.Merge(m, src)
}
func (m *MetaDataValue) XXX_Size() int {
	return xxx_messageInfo_MetaDataValue.Size(m)
}
func (m *MetaDataValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaDataValue.DiscardUnknown(m)
}

var xxx_messageInfo_MetaDataValue proto.InternalMessageInfo

type isMetaDataValue_Data interface {
	isMetaDataValue_Data()
}

type MetaDataValue_TextValue struct {
	TextValue string `protobuf:"bytes,1,opt,name=text_value,json=textValue,proto3,oneof"`
}

type MetaDataValue_NumberValue struct {
	NumberValue int64 `protobuf:"zigzag64,2,opt,name=number_value,json=numberValue,proto3,oneof"`
}

type MetaDataValue_BooleanValue struct {
	BooleanValue bool `protobuf:"varint,3,opt,name=boolean_value,json=booleanValue,proto3,oneof"`
}

type MetaDataValue_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,4,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type MetaDataValue_BytesValue struct {
	BytesValue *SerializedObject `protobuf:"bytes,5,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}

func (*MetaDataValue_TextValue) isMetaDataValue_Data() {}

func (*MetaDataValue_NumberValue) isMetaDataValue_Data() {}

func (*MetaDataValue_BooleanValue) isMetaDataValue_Data() {}

func (*MetaDataValue_DoubleValue) isMetaDataValue_Data() {}

func (*MetaDataValue_BytesValue) isMetaDataValue_Data() {}

func (m *MetaDataValue) GetData() isMetaDataValue_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MetaDataValue) GetTextValue() string {
	if x, ok := m.GetData().(*MetaDataValue_TextValue); ok {
		return x.TextValue
	}
	return ""
}

func (m *MetaDataValue) GetNumberValue() int64 {
	if x, ok := m.GetData().(*MetaDataValue_NumberValue); ok {
		return x.NumberValue
	}
	return 0
}

func (m *MetaDataValue) GetBooleanValue() bool {
	if x, ok := m.GetData().(*MetaDataValue_BooleanValue); ok {
		return x.BooleanValue
	}
	return false
}

func (m *MetaDataValue) GetDoubleValue() float64 {
	if x, ok := m.GetData().(*MetaDataValue_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *MetaDataValue) GetBytesValue() *SerializedObject {
	if x, ok := m.GetData().(*MetaDataValue_BytesValue); ok {
		return x.BytesValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MetaDataValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MetaDataValue_TextValue)(nil),
		(*MetaDataValue_NumberValue)(nil),
		(*MetaDataValue_BooleanValue)(nil),
		(*MetaDataValue_DoubleValue)(nil),
		(*MetaDataValue_BytesValue)(nil),
	}
}

// An instruction for routing components when routing or processing a message.
type ProcessingInstruction struct {
	// The type of processing message.
	Key ProcessingKey `protobuf:"varint,1,opt,name=key,proto3,enum=io.axoniq.axonserver.grpc.ProcessingKey" json:"key,omitempty"`
	// The value associated with the processing key.
	Value                *MetaDataValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ProcessingInstruction) Reset()         { *m = ProcessingInstruction{} }
func (m *ProcessingInstruction) String() string { return proto.CompactTextString(m) }
func (*ProcessingInstruction) ProtoMessage()    {}
func (*ProcessingInstruction) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *ProcessingInstruction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessingInstruction.Unmarshal(m, b)
}
func (m *ProcessingInstruction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessingInstruction.Marshal(b, m, deterministic)
}
func (m *ProcessingInstruction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessingInstruction.Merge(m, src)
}
func (m *ProcessingInstruction) XXX_Size() int {
	return xxx_messageInfo_ProcessingInstruction.Size(m)
}
func (m *ProcessingInstruction) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessingInstruction.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessingInstruction proto.InternalMessageInfo

func (m *ProcessingInstruction) GetKey() ProcessingKey {
	if m != nil {
		return m.Key
	}
	return ProcessingKey_ROUTING_KEY
}

func (m *ProcessingInstruction) GetValue() *MetaDataValue {
	if m != nil {
		return m.Value
	}
	return nil
}

// Message containing details of an error
type ErrorMessage struct {
	// A human readable message explaining the error
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// A description of the location (client component, server) where the error occurred
	Location string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	// A collection of messages providing more details about root causes of the error
	Details []string `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
	// An Error Code identifying the type of error
	ErrorCode            string   `protobuf:"bytes,4,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorMessage) Reset()         { *m = ErrorMessage{} }
func (m *ErrorMessage) String() string { return proto.CompactTextString(m) }
func (*ErrorMessage) ProtoMessage()    {}
func (*ErrorMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

func (m *ErrorMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorMessage.Unmarshal(m, b)
}
func (m *ErrorMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorMessage.Marshal(b, m, deterministic)
}
func (m *ErrorMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorMessage.Merge(m, src)
}
func (m *ErrorMessage) XXX_Size() int {
	return xxx_messageInfo_ErrorMessage.Size(m)
}
func (m *ErrorMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorMessage proto.InternalMessageInfo

func (m *ErrorMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ErrorMessage) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *ErrorMessage) GetDetails() []string {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *ErrorMessage) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

// Message used for Flow Control instruction, providing the counterpart with additional permits for sending messages
type FlowControl struct {
	// The ClientID of the component providing additional permits
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// The number of permits to provide
	Permits              int64    `protobuf:"varint,3,opt,name=permits,proto3" json:"permits,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowControl) Reset()         { *m = FlowControl{} }
func (m *FlowControl) String() string { return proto.CompactTextString(m) }
func (*FlowControl) ProtoMessage()    {}
func (*FlowControl) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4}
}

func (m *FlowControl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowControl.Unmarshal(m, b)
}
func (m *FlowControl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowControl.Marshal(b, m, deterministic)
}
func (m *FlowControl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowControl.Merge(m, src)
}
func (m *FlowControl) XXX_Size() int {
	return xxx_messageInfo_FlowControl.Size(m)
}
func (m *FlowControl) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowControl.DiscardUnknown(m)
}

var xxx_messageInfo_FlowControl proto.InternalMessageInfo

func (m *FlowControl) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *FlowControl) GetPermits() int64 {
	if m != nil {
		return m.Permits
	}
	return 0
}

// Message describing instruction acknowledgement
type InstructionAck struct {
	// The identifier of the instruction
	InstructionId string `protobuf:"bytes,1,opt,name=instruction_id,json=instructionId,proto3" json:"instruction_id,omitempty"`
	// Indicator whether the instruction was acknowledged successfully
	Success bool `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	// Set if instruction acknowledgement failed.
	Error                *ErrorMessage `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *InstructionAck) Reset()         { *m = InstructionAck{} }
func (m *InstructionAck) String() string { return proto.CompactTextString(m) }
func (*InstructionAck) ProtoMessage()    {}
func (*InstructionAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{5}
}

func (m *InstructionAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstructionAck.Unmarshal(m, b)
}
func (m *InstructionAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstructionAck.Marshal(b, m, deterministic)
}
func (m *InstructionAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstructionAck.Merge(m, src)
}
func (m *InstructionAck) XXX_Size() int {
	return xxx_messageInfo_InstructionAck.Size(m)
}
func (m *InstructionAck) XXX_DiscardUnknown() {
	xxx_messageInfo_InstructionAck.DiscardUnknown(m)
}

var xxx_messageInfo_InstructionAck proto.InternalMessageInfo

func (m *InstructionAck) GetInstructionId() string {
	if m != nil {
		return m.InstructionId
	}
	return ""
}

func (m *InstructionAck) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *InstructionAck) GetError() *ErrorMessage {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterEnum("io.axoniq.axonserver.grpc.ProcessingKey", ProcessingKey_name, ProcessingKey_value)
	proto.RegisterType((*SerializedObject)(nil), "io.axoniq.axonserver.grpc.SerializedObject")
	proto.RegisterType((*MetaDataValue)(nil), "io.axoniq.axonserver.grpc.MetaDataValue")
	proto.RegisterType((*ProcessingInstruction)(nil), "io.axoniq.axonserver.grpc.ProcessingInstruction")
	proto.RegisterType((*ErrorMessage)(nil), "io.axoniq.axonserver.grpc.ErrorMessage")
	proto.RegisterType((*FlowControl)(nil), "io.axoniq.axonserver.grpc.FlowControl")
	proto.RegisterType((*InstructionAck)(nil), "io.axoniq.axonserver.grpc.InstructionAck")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xef, 0x6b, 0xd3, 0x40,
	0x18, 0x6e, 0xd6, 0xb4, 0x4b, 0xdf, 0xb4, 0x33, 0x1e, 0x88, 0xdd, 0x86, 0x58, 0x22, 0xc3, 0xa2,
	0xd0, 0xc2, 0xfc, 0x26, 0x28, 0xb8, 0xd9, 0xd9, 0x6e, 0x6e, 0x1d, 0xb7, 0x6e, 0x30, 0xbf, 0xd4,
	0x4b, 0xf2, 0x5a, 0xe2, 0xd2, 0x5c, 0xbd, 0x5c, 0xe6, 0x2a, 0xf8, 0x3f, 0x08, 0xfe, 0xb5, 0x7e,
	0x93, 0xbb, 0x4b, 0xd7, 0x4e, 0x70, 0x9f, 0x72, 0xef, 0xc3, 0xf3, 0xe3, 0xee, 0x79, 0x09, 0xd4,
	0x43, 0x3e, 0x9d, 0xf2, 0xb4, 0x33, 0x13, 0x5c, 0x72, 0xb2, 0x19, 0xf3, 0x0e, 0xbb, 0xe1, 0x69,
	0xfc, 0x4d, 0x7f, 0x32, 0x14, 0xd7, 0x28, 0x3a, 0x13, 0x31, 0x0b, 0xb7, 0x36, 0x27, 0x9c, 0x4f,
	0x12, 0xec, 0x6a, 0x62, 0x90, 0x7f, 0xe9, 0xb2, 0x74, 0x6e, 0x54, 0xfe, 0x05, 0x78, 0x67, 0x28,
	0x62, 0x96, 0xc4, 0x3f, 0x30, 0x1a, 0x06, 0x5f, 0x31, 0x94, 0x84, 0x80, 0x2d, 0xe7, 0x33, 0x6c,
	0x5a, 0x2d, 0xab, 0x5d, 0xa3, 0xfa, 0x4c, 0xb6, 0xc0, 0x11, 0x78, 0x1d, 0x67, 0x31, 0x4f, 0x9b,
	0x6b, 0x1a, 0xbf, 0x9d, 0x15, 0x3f, 0x62, 0x92, 0x35, 0xcb, 0x2d, 0xab, 0x5d, 0xa7, 0xfa, 0xec,
	0xff, 0xb1, 0xa0, 0x71, 0x8c, 0x92, 0xbd, 0x67, 0x92, 0x5d, 0xb0, 0x24, 0x47, 0xf2, 0x14, 0x40,
	0xe2, 0x8d, 0x1c, 0x5f, 0xab, 0xc9, 0x78, 0xf7, 0x4b, 0xb4, 0xa6, 0x30, 0x43, 0x78, 0x06, 0xf5,
	0x34, 0x9f, 0x06, 0x28, 0x0a, 0x8a, 0x8a, 0x21, 0xfd, 0x12, 0x75, 0x0d, 0x6a, 0x48, 0x3b, 0xd0,
	0x08, 0x38, 0x4f, 0x90, 0xa5, 0x05, 0x4b, 0x85, 0x3a, 0xfd, 0x12, 0xad, 0x17, 0xf0, 0xad, 0x57,
	0xc4, 0xf3, 0x20, 0xc1, 0x82, 0x65, 0xb7, 0xac, 0xb6, 0xa5, 0xbc, 0x0c, 0x6a, 0x48, 0x27, 0xe0,
	0x06, 0x73, 0x89, 0x59, 0xc1, 0xa9, 0xb4, 0xac, 0xb6, 0xbb, 0xfb, 0xb2, 0xf3, 0xdf, 0x1e, 0x3b,
	0xff, 0x36, 0xd5, 0x2f, 0x51, 0xd0, 0x0e, 0xda, 0x6f, 0xaf, 0x6a, 0x7a, 0xf0, 0x7f, 0x5b, 0xf0,
	0xe8, 0x54, 0xf0, 0x10, 0xb3, 0x2c, 0x4e, 0x27, 0x83, 0x34, 0x93, 0x22, 0x0f, 0xa5, 0x6a, 0xea,
	0x35, 0x94, 0xaf, 0x70, 0xae, 0x1f, 0xbf, 0xb1, 0xdb, 0xbe, 0x27, 0x69, 0x29, 0x3f, 0xc2, 0x39,
	0x55, 0x22, 0xf2, 0x16, 0x2a, 0xcb, 0x5e, 0xdc, 0x7b, 0xd5, 0x77, 0x8a, 0xa7, 0x46, 0xe6, 0xff,
	0x84, 0x7a, 0x4f, 0x08, 0x2e, 0x8e, 0x31, 0xcb, 0xd8, 0x04, 0x49, 0x13, 0xd6, 0xa7, 0xe6, 0x58,
	0x2c, 0x7a, 0x31, 0xaa, 0x5d, 0x27, 0x3c, 0x64, 0x72, 0x65, 0xd7, 0x8b, 0x59, 0xa9, 0x22, 0x94,
	0x2c, 0x4e, 0xb2, 0x66, 0xb9, 0x55, 0x56, 0xaa, 0x62, 0x24, 0x4f, 0x00, 0x50, 0xf9, 0x8f, 0x43,
	0x1e, 0x99, 0xc2, 0x6b, 0xb4, 0xa6, 0x91, 0x7d, 0x1e, 0xa1, 0x7f, 0x08, 0xee, 0x41, 0xc2, 0xbf,
	0xef, 0xf3, 0x54, 0x0a, 0x9e, 0x90, 0x6d, 0xa8, 0x85, 0x49, 0x8c, 0xa9, 0x1c, 0xc7, 0xd1, 0x22,
	0xc4, 0x00, 0x83, 0x48, 0x85, 0xcc, 0x50, 0x4c, 0x63, 0x99, 0xe9, 0xf5, 0x96, 0xe9, 0x62, 0x3c,
	0xb4, 0x1d, 0xcb, 0x5b, 0xf3, 0x7f, 0x59, 0xb0, 0xb1, 0x52, 0xeb, 0xbb, 0xf0, 0x8a, 0xec, 0xc0,
	0x46, 0xbc, 0x44, 0x94, 0xa9, 0x79, 0x54, 0x63, 0x05, 0x35, 0xce, 0x59, 0x1e, 0xaa, 0x6a, 0x75,
	0xa8, 0x43, 0x17, 0x23, 0x79, 0x03, 0x15, 0x7d, 0x59, 0x9d, 0xe8, 0xee, 0x3e, 0xbf, 0xa7, 0xde,
	0xd5, 0x1a, 0xa9, 0x51, 0xbd, 0xf8, 0x0c, 0x8d, 0x3b, 0x3b, 0x23, 0x0f, 0xc0, 0xa5, 0xc3, 0xf3,
	0xd1, 0xe0, 0xe4, 0xc3, 0xf8, 0xa8, 0x77, 0xe9, 0x95, 0x48, 0x1d, 0x9c, 0x53, 0x3a, 0x18, 0xd2,
	0xc1, 0xe8, 0xd2, 0xb3, 0x88, 0x0b, 0xeb, 0xa3, 0xc1, 0x71, 0x6f, 0x78, 0x3e, 0xf2, 0xd6, 0xc8,
	0x43, 0x68, 0x9c, 0xd0, 0xf1, 0xf0, 0x60, 0x4c, 0x7b, 0x67, 0xe7, 0x1f, 0x47, 0x67, 0x5e, 0xd9,
	0xb7, 0x1d, 0xdb, 0xb3, 0x7d, 0xdb, 0xa9, 0x78, 0x15, 0xdf, 0x76, 0xaa, 0x5e, 0x75, 0x6f, 0xfb,
	0xd4, 0xfa, 0xf4, 0x38, 0x13, 0x61, 0x77, 0x76, 0x35, 0xe9, 0xaa, 0x7b, 0x74, 0x97, 0xf7, 0x0a,
	0xaa, 0xfa, 0x6f, 0x7e, 0xf5, 0x37, 0x00, 0x00, 0xff, 0xff, 0x79, 0xc1, 0xfa, 0x93, 0x13, 0x04,
	0x00, 0x00,
}
