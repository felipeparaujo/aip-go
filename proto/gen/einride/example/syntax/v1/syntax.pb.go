// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: einride/example/syntax/v1/syntax.proto

package syntaxv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Enum int32

const (
	Enum_ENUM_UNSPECIFIED Enum = 0
	Enum_ENUM_ONE         Enum = 1
	Enum_ENUM_TWO         Enum = 2
)

// Enum value maps for Enum.
var (
	Enum_name = map[int32]string{
		0: "ENUM_UNSPECIFIED",
		1: "ENUM_ONE",
		2: "ENUM_TWO",
	}
	Enum_value = map[string]int32{
		"ENUM_UNSPECIFIED": 0,
		"ENUM_ONE":         1,
		"ENUM_TWO":         2,
	}
)

func (x Enum) Enum() *Enum {
	p := new(Enum)
	*p = x
	return p
}

func (x Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_example_syntax_v1_syntax_proto_enumTypes[0].Descriptor()
}

func (Enum) Type() protoreflect.EnumType {
	return &file_einride_example_syntax_v1_syntax_proto_enumTypes[0]
}

func (x Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Enum.Descriptor instead.
func (Enum) EnumDescriptor() ([]byte, []int) {
	return file_einride_example_syntax_v1_syntax_proto_rawDescGZIP(), []int{0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Double           float64             `protobuf:"fixed64,1,opt,name=double,proto3" json:"double,omitempty"`
	Float            float32             `protobuf:"fixed32,2,opt,name=float,proto3" json:"float,omitempty"`
	Int32            int32               `protobuf:"varint,3,opt,name=int32,proto3" json:"int32,omitempty"`
	Int64            int64               `protobuf:"varint,4,opt,name=int64,proto3" json:"int64,omitempty"`
	Uint32           uint32              `protobuf:"varint,5,opt,name=uint32,proto3" json:"uint32,omitempty"`
	Uint64           uint64              `protobuf:"varint,6,opt,name=uint64,proto3" json:"uint64,omitempty"`
	Sint32           int32               `protobuf:"zigzag32,7,opt,name=sint32,proto3" json:"sint32,omitempty"`
	Sint64           int64               `protobuf:"zigzag64,8,opt,name=sint64,proto3" json:"sint64,omitempty"`
	Fixed32          uint32              `protobuf:"fixed32,9,opt,name=fixed32,proto3" json:"fixed32,omitempty"`
	Fixed64          uint64              `protobuf:"fixed64,10,opt,name=fixed64,proto3" json:"fixed64,omitempty"`
	Sfixed32         int32               `protobuf:"fixed32,11,opt,name=sfixed32,proto3" json:"sfixed32,omitempty"`
	Sfixed64         int64               `protobuf:"fixed64,12,opt,name=sfixed64,proto3" json:"sfixed64,omitempty"`
	Bool             bool                `protobuf:"varint,13,opt,name=bool,proto3" json:"bool,omitempty"`
	String_          string              `protobuf:"bytes,14,opt,name=string,proto3" json:"string,omitempty"`
	Bytes            []byte              `protobuf:"bytes,15,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Enum             Enum                `protobuf:"varint,16,opt,name=enum,proto3,enum=einride.example.syntax.v1.Enum" json:"enum,omitempty"`
	Message          *Message            `protobuf:"bytes,17,opt,name=message,proto3" json:"message,omitempty"`
	RepeatedDouble   []float64           `protobuf:"fixed64,18,rep,packed,name=repeated_double,json=repeatedDouble,proto3" json:"repeated_double,omitempty"`
	RepeatedFloat    []float32           `protobuf:"fixed32,19,rep,packed,name=repeated_float,json=repeatedFloat,proto3" json:"repeated_float,omitempty"`
	RepeatedInt32    []int32             `protobuf:"varint,20,rep,packed,name=repeated_int32,json=repeatedInt32,proto3" json:"repeated_int32,omitempty"`
	RepeatedInt64    []int64             `protobuf:"varint,21,rep,packed,name=repeated_int64,json=repeatedInt64,proto3" json:"repeated_int64,omitempty"`
	RepeatedUint32   []uint32            `protobuf:"varint,22,rep,packed,name=repeated_uint32,json=repeatedUint32,proto3" json:"repeated_uint32,omitempty"`
	RepeatedUint64   []uint64            `protobuf:"varint,23,rep,packed,name=repeated_uint64,json=repeatedUint64,proto3" json:"repeated_uint64,omitempty"`
	RepeatedSint32   []int32             `protobuf:"zigzag32,24,rep,packed,name=repeated_sint32,json=repeatedSint32,proto3" json:"repeated_sint32,omitempty"`
	RepeatedSint64   []int64             `protobuf:"zigzag64,25,rep,packed,name=repeated_sint64,json=repeatedSint64,proto3" json:"repeated_sint64,omitempty"`
	RepeatedFixed32  []uint32            `protobuf:"fixed32,26,rep,packed,name=repeated_fixed32,json=repeatedFixed32,proto3" json:"repeated_fixed32,omitempty"`
	RepeatedFixed64  []uint64            `protobuf:"fixed64,27,rep,packed,name=repeated_fixed64,json=repeatedFixed64,proto3" json:"repeated_fixed64,omitempty"`
	RepeatedSfixed32 []int32             `protobuf:"fixed32,28,rep,packed,name=repeated_sfixed32,json=repeatedSfixed32,proto3" json:"repeated_sfixed32,omitempty"`
	RepeatedSfixed64 []int64             `protobuf:"fixed64,29,rep,packed,name=repeated_sfixed64,json=repeatedSfixed64,proto3" json:"repeated_sfixed64,omitempty"`
	RepeatedBool     []bool              `protobuf:"varint,30,rep,packed,name=repeated_bool,json=repeatedBool,proto3" json:"repeated_bool,omitempty"`
	RepeatedString   []string            `protobuf:"bytes,31,rep,name=repeated_string,json=repeatedString,proto3" json:"repeated_string,omitempty"`
	RepeatedBytes    [][]byte            `protobuf:"bytes,32,rep,name=repeated_bytes,json=repeatedBytes,proto3" json:"repeated_bytes,omitempty"`
	RepeatedEnum     []Enum              `protobuf:"varint,33,rep,packed,name=repeated_enum,json=repeatedEnum,proto3,enum=einride.example.syntax.v1.Enum" json:"repeated_enum,omitempty"`
	RepeatedMessage  []*Message          `protobuf:"bytes,34,rep,name=repeated_message,json=repeatedMessage,proto3" json:"repeated_message,omitempty"`
	MapStringString  map[string]string   `protobuf:"bytes,35,rep,name=map_string_string,json=mapStringString,proto3" json:"map_string_string,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapStringMessage map[string]*Message `protobuf:"bytes,36,rep,name=map_string_message,json=mapStringMessage,proto3" json:"map_string_message,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Types that are assignable to Oneof:
	//
	//	*Message_OneofString
	//	*Message_OneofEnum
	//	*Message_OneofMessage1
	//	*Message_OneofMessage2
	Oneof isMessage_Oneof `protobuf_oneof:"oneof"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_example_syntax_v1_syntax_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_einride_example_syntax_v1_syntax_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_einride_example_syntax_v1_syntax_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetDouble() float64 {
	if x != nil {
		return x.Double
	}
	return 0
}

func (x *Message) GetFloat() float32 {
	if x != nil {
		return x.Float
	}
	return 0
}

func (x *Message) GetInt32() int32 {
	if x != nil {
		return x.Int32
	}
	return 0
}

func (x *Message) GetInt64() int64 {
	if x != nil {
		return x.Int64
	}
	return 0
}

func (x *Message) GetUint32() uint32 {
	if x != nil {
		return x.Uint32
	}
	return 0
}

func (x *Message) GetUint64() uint64 {
	if x != nil {
		return x.Uint64
	}
	return 0
}

func (x *Message) GetSint32() int32 {
	if x != nil {
		return x.Sint32
	}
	return 0
}

func (x *Message) GetSint64() int64 {
	if x != nil {
		return x.Sint64
	}
	return 0
}

func (x *Message) GetFixed32() uint32 {
	if x != nil {
		return x.Fixed32
	}
	return 0
}

func (x *Message) GetFixed64() uint64 {
	if x != nil {
		return x.Fixed64
	}
	return 0
}

func (x *Message) GetSfixed32() int32 {
	if x != nil {
		return x.Sfixed32
	}
	return 0
}

func (x *Message) GetSfixed64() int64 {
	if x != nil {
		return x.Sfixed64
	}
	return 0
}

func (x *Message) GetBool() bool {
	if x != nil {
		return x.Bool
	}
	return false
}

func (x *Message) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

func (x *Message) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *Message) GetEnum() Enum {
	if x != nil {
		return x.Enum
	}
	return Enum_ENUM_UNSPECIFIED
}

func (x *Message) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Message) GetRepeatedDouble() []float64 {
	if x != nil {
		return x.RepeatedDouble
	}
	return nil
}

func (x *Message) GetRepeatedFloat() []float32 {
	if x != nil {
		return x.RepeatedFloat
	}
	return nil
}

func (x *Message) GetRepeatedInt32() []int32 {
	if x != nil {
		return x.RepeatedInt32
	}
	return nil
}

func (x *Message) GetRepeatedInt64() []int64 {
	if x != nil {
		return x.RepeatedInt64
	}
	return nil
}

func (x *Message) GetRepeatedUint32() []uint32 {
	if x != nil {
		return x.RepeatedUint32
	}
	return nil
}

func (x *Message) GetRepeatedUint64() []uint64 {
	if x != nil {
		return x.RepeatedUint64
	}
	return nil
}

func (x *Message) GetRepeatedSint32() []int32 {
	if x != nil {
		return x.RepeatedSint32
	}
	return nil
}

func (x *Message) GetRepeatedSint64() []int64 {
	if x != nil {
		return x.RepeatedSint64
	}
	return nil
}

func (x *Message) GetRepeatedFixed32() []uint32 {
	if x != nil {
		return x.RepeatedFixed32
	}
	return nil
}

func (x *Message) GetRepeatedFixed64() []uint64 {
	if x != nil {
		return x.RepeatedFixed64
	}
	return nil
}

func (x *Message) GetRepeatedSfixed32() []int32 {
	if x != nil {
		return x.RepeatedSfixed32
	}
	return nil
}

func (x *Message) GetRepeatedSfixed64() []int64 {
	if x != nil {
		return x.RepeatedSfixed64
	}
	return nil
}

func (x *Message) GetRepeatedBool() []bool {
	if x != nil {
		return x.RepeatedBool
	}
	return nil
}

func (x *Message) GetRepeatedString() []string {
	if x != nil {
		return x.RepeatedString
	}
	return nil
}

func (x *Message) GetRepeatedBytes() [][]byte {
	if x != nil {
		return x.RepeatedBytes
	}
	return nil
}

func (x *Message) GetRepeatedEnum() []Enum {
	if x != nil {
		return x.RepeatedEnum
	}
	return nil
}

func (x *Message) GetRepeatedMessage() []*Message {
	if x != nil {
		return x.RepeatedMessage
	}
	return nil
}

func (x *Message) GetMapStringString() map[string]string {
	if x != nil {
		return x.MapStringString
	}
	return nil
}

func (x *Message) GetMapStringMessage() map[string]*Message {
	if x != nil {
		return x.MapStringMessage
	}
	return nil
}

func (m *Message) GetOneof() isMessage_Oneof {
	if m != nil {
		return m.Oneof
	}
	return nil
}

func (x *Message) GetOneofString() string {
	if x, ok := x.GetOneof().(*Message_OneofString); ok {
		return x.OneofString
	}
	return ""
}

func (x *Message) GetOneofEnum() Enum {
	if x, ok := x.GetOneof().(*Message_OneofEnum); ok {
		return x.OneofEnum
	}
	return Enum_ENUM_UNSPECIFIED
}

func (x *Message) GetOneofMessage1() *Message {
	if x, ok := x.GetOneof().(*Message_OneofMessage1); ok {
		return x.OneofMessage1
	}
	return nil
}

func (x *Message) GetOneofMessage2() *Message {
	if x, ok := x.GetOneof().(*Message_OneofMessage2); ok {
		return x.OneofMessage2
	}
	return nil
}

type isMessage_Oneof interface {
	isMessage_Oneof()
}

type Message_OneofString struct {
	OneofString string `protobuf:"bytes,37,opt,name=oneof_string,json=oneofString,proto3,oneof"`
}

type Message_OneofEnum struct {
	OneofEnum Enum `protobuf:"varint,38,opt,name=oneof_enum,json=oneofEnum,proto3,enum=einride.example.syntax.v1.Enum,oneof"`
}

type Message_OneofMessage1 struct {
	OneofMessage1 *Message `protobuf:"bytes,39,opt,name=oneof_message1,json=oneofMessage1,proto3,oneof"`
}

type Message_OneofMessage2 struct {
	OneofMessage2 *Message `protobuf:"bytes,40,opt,name=oneof_message2,json=oneofMessage2,proto3,oneof"`
}

func (*Message_OneofString) isMessage_Oneof() {}

func (*Message_OneofEnum) isMessage_Oneof() {}

func (*Message_OneofMessage1) isMessage_Oneof() {}

func (*Message_OneofMessage2) isMessage_Oneof() {}

var File_einride_example_syntax_v1_syntax_proto protoreflect.FileDescriptor

var file_einride_example_syntax_v1_syntax_proto_rawDesc = []byte{
	0x0a, 0x26, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x6e, 0x74,
	0x61, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78,
	0x2e, 0x76, 0x31, 0x22, 0xe4, 0x0e, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x18, 0x07, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x12, 0x52, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x33, 0x32, 0x18, 0x09, 0x20, 0x01, 0x28, 0x07, 0x52, 0x07, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x33, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x06, 0x52, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0f, 0x52,
	0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x10, 0x52, 0x08, 0x73, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x3c, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x12, 0x20,
	0x03, 0x28, 0x01, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x44, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x13, 0x20, 0x03, 0x28, 0x02, 0x52, 0x0d, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x14, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x18, 0x15, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x16, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x55, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x18, 0x17, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x18, 0x20,
	0x03, 0x28, 0x11, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x19, 0x20, 0x03, 0x28, 0x12, 0x52, 0x0e, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x29, 0x0a, 0x10,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32,
	0x18, 0x1a, 0x20, 0x03, 0x28, 0x07, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x1b, 0x20, 0x03, 0x28,
	0x06, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x69, 0x78, 0x65, 0x64,
	0x36, 0x34, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x1c, 0x20, 0x03, 0x28, 0x0f, 0x52, 0x10, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12,
	0x2b, 0x0a, 0x11, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x36, 0x34, 0x18, 0x1d, 0x20, 0x03, 0x28, 0x10, 0x52, 0x10, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x1e, 0x20,
	0x03, 0x28, 0x08, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x6f, 0x6f,
	0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x18, 0x1f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x20, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x44, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6e,
	0x75, 0x6d, 0x18, 0x21, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x61,
	0x78, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x4d, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x22, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x63, 0x0a, 0x11, 0x6d, 0x61, 0x70, 0x5f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x23, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x6d, 0x61, 0x70, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x66, 0x0a, 0x12, 0x6d,
	0x61, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x24, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x10, 0x6d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x18, 0x25, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x40, 0x0a, 0x0a, 0x6f, 0x6e, 0x65, 0x6f,
	0x66, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x26, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x65,
	0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73,
	0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x48, 0x00, 0x52,
	0x09, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x4b, 0x0a, 0x0e, 0x6f, 0x6e,
	0x65, 0x6f, 0x66, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x18, 0x27, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x12, 0x4b, 0x0a, 0x0e, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0x1a, 0x42, 0x0a, 0x14, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x67, 0x0a, 0x15, 0x4d, 0x61, 0x70, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x07, 0x0a, 0x05, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2a, 0x38, 0x0a, 0x04, 0x45, 0x6e,
	0x75, 0x6d, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x4e, 0x55, 0x4d,
	0x5f, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x54,
	0x57, 0x4f, 0x10, 0x02, 0x42, 0xf5, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x6e,
	0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x79, 0x6e,
	0x74, 0x61, 0x78, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x6f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x61, 0x69, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2f, 0x76, 0x31, 0x3b, 0x73,
	0x79, 0x6e, 0x74, 0x61, 0x78, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x45, 0x53, 0xaa, 0x02, 0x19,
	0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x19, 0x45, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x5c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x53, 0x79, 0x6e, 0x74,
	0x61, 0x78, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x25, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c,
	0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x3a, 0x3a, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x3a, 0x3a, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_example_syntax_v1_syntax_proto_rawDescOnce sync.Once
	file_einride_example_syntax_v1_syntax_proto_rawDescData = file_einride_example_syntax_v1_syntax_proto_rawDesc
)

func file_einride_example_syntax_v1_syntax_proto_rawDescGZIP() []byte {
	file_einride_example_syntax_v1_syntax_proto_rawDescOnce.Do(func() {
		file_einride_example_syntax_v1_syntax_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_example_syntax_v1_syntax_proto_rawDescData)
	})
	return file_einride_example_syntax_v1_syntax_proto_rawDescData
}

var file_einride_example_syntax_v1_syntax_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_einride_example_syntax_v1_syntax_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_einride_example_syntax_v1_syntax_proto_goTypes = []interface{}{
	(Enum)(0),       // 0: einride.example.syntax.v1.Enum
	(*Message)(nil), // 1: einride.example.syntax.v1.Message
	nil,             // 2: einride.example.syntax.v1.Message.MapStringStringEntry
	nil,             // 3: einride.example.syntax.v1.Message.MapStringMessageEntry
}
var file_einride_example_syntax_v1_syntax_proto_depIdxs = []int32{
	0,  // 0: einride.example.syntax.v1.Message.enum:type_name -> einride.example.syntax.v1.Enum
	1,  // 1: einride.example.syntax.v1.Message.message:type_name -> einride.example.syntax.v1.Message
	0,  // 2: einride.example.syntax.v1.Message.repeated_enum:type_name -> einride.example.syntax.v1.Enum
	1,  // 3: einride.example.syntax.v1.Message.repeated_message:type_name -> einride.example.syntax.v1.Message
	2,  // 4: einride.example.syntax.v1.Message.map_string_string:type_name -> einride.example.syntax.v1.Message.MapStringStringEntry
	3,  // 5: einride.example.syntax.v1.Message.map_string_message:type_name -> einride.example.syntax.v1.Message.MapStringMessageEntry
	0,  // 6: einride.example.syntax.v1.Message.oneof_enum:type_name -> einride.example.syntax.v1.Enum
	1,  // 7: einride.example.syntax.v1.Message.oneof_message1:type_name -> einride.example.syntax.v1.Message
	1,  // 8: einride.example.syntax.v1.Message.oneof_message2:type_name -> einride.example.syntax.v1.Message
	1,  // 9: einride.example.syntax.v1.Message.MapStringMessageEntry.value:type_name -> einride.example.syntax.v1.Message
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_einride_example_syntax_v1_syntax_proto_init() }
func file_einride_example_syntax_v1_syntax_proto_init() {
	if File_einride_example_syntax_v1_syntax_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_einride_example_syntax_v1_syntax_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_einride_example_syntax_v1_syntax_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Message_OneofString)(nil),
		(*Message_OneofEnum)(nil),
		(*Message_OneofMessage1)(nil),
		(*Message_OneofMessage2)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_einride_example_syntax_v1_syntax_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_example_syntax_v1_syntax_proto_goTypes,
		DependencyIndexes: file_einride_example_syntax_v1_syntax_proto_depIdxs,
		EnumInfos:         file_einride_example_syntax_v1_syntax_proto_enumTypes,
		MessageInfos:      file_einride_example_syntax_v1_syntax_proto_msgTypes,
	}.Build()
	File_einride_example_syntax_v1_syntax_proto = out.File
	file_einride_example_syntax_v1_syntax_proto_rawDesc = nil
	file_einride_example_syntax_v1_syntax_proto_goTypes = nil
	file_einride_example_syntax_v1_syntax_proto_depIdxs = nil
}
