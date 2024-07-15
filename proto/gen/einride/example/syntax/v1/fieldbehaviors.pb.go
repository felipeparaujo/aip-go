// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: einride/example/syntax/v1/fieldbehaviors.proto

package syntaxv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type FieldBehaviorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field                       string                           `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	OutputOnlyField             string                           `protobuf:"bytes,2,opt,name=output_only_field,json=outputOnlyField,proto3" json:"output_only_field,omitempty"`
	OptionalField               string                           `protobuf:"bytes,3,opt,name=optional_field,json=optionalField,proto3" json:"optional_field,omitempty"`
	MessageWithoutFieldBehavior *FieldBehaviorMessage            `protobuf:"bytes,12,opt,name=message_without_field_behavior,json=messageWithoutFieldBehavior,proto3" json:"message_without_field_behavior,omitempty"`
	OutputOnlyMessage           *FieldBehaviorMessage            `protobuf:"bytes,13,opt,name=output_only_message,json=outputOnlyMessage,proto3" json:"output_only_message,omitempty"`
	OptionalMessage             *FieldBehaviorMessage            `protobuf:"bytes,14,opt,name=optional_message,json=optionalMessage,proto3" json:"optional_message,omitempty"`
	RepeatedMessage             []*FieldBehaviorMessage          `protobuf:"bytes,4,rep,name=repeated_message,json=repeatedMessage,proto3" json:"repeated_message,omitempty"`
	RepeatedOutputOnlyMessage   []*FieldBehaviorMessage          `protobuf:"bytes,5,rep,name=repeated_output_only_message,json=repeatedOutputOnlyMessage,proto3" json:"repeated_output_only_message,omitempty"`
	RepeatedOptionalMessage     []*FieldBehaviorMessage          `protobuf:"bytes,6,rep,name=repeated_optional_message,json=repeatedOptionalMessage,proto3" json:"repeated_optional_message,omitempty"`
	MapMessage                  map[string]*FieldBehaviorMessage `protobuf:"bytes,7,rep,name=map_message,json=mapMessage,proto3" json:"map_message,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapOutputOnlyMessage        map[string]*FieldBehaviorMessage `protobuf:"bytes,8,rep,name=map_output_only_message,json=mapOutputOnlyMessage,proto3" json:"map_output_only_message,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapOptionalMessage          map[string]*FieldBehaviorMessage `protobuf:"bytes,9,rep,name=map_optional_message,json=mapOptionalMessage,proto3" json:"map_optional_message,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	StringMap                   map[string]string                `protobuf:"bytes,10,rep,name=string_map,json=stringMap,proto3" json:"string_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	StringList                  []string                         `protobuf:"bytes,11,rep,name=string_list,json=stringList,proto3" json:"string_list,omitempty"`
	// Types that are assignable to Oneof:
	//
	//	*FieldBehaviorMessage_FieldBehaviorMessage
	//	*FieldBehaviorMessage_SmallFieldBehaviorMessage
	Oneof isFieldBehaviorMessage_Oneof `protobuf_oneof:"oneof"`
}

func (x *FieldBehaviorMessage) Reset() {
	*x = FieldBehaviorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_example_syntax_v1_fieldbehaviors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldBehaviorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldBehaviorMessage) ProtoMessage() {}

func (x *FieldBehaviorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_einride_example_syntax_v1_fieldbehaviors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldBehaviorMessage.ProtoReflect.Descriptor instead.
func (*FieldBehaviorMessage) Descriptor() ([]byte, []int) {
	return file_einride_example_syntax_v1_fieldbehaviors_proto_rawDescGZIP(), []int{0}
}

func (x *FieldBehaviorMessage) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *FieldBehaviorMessage) GetOutputOnlyField() string {
	if x != nil {
		return x.OutputOnlyField
	}
	return ""
}

func (x *FieldBehaviorMessage) GetOptionalField() string {
	if x != nil {
		return x.OptionalField
	}
	return ""
}

func (x *FieldBehaviorMessage) GetMessageWithoutFieldBehavior() *FieldBehaviorMessage {
	if x != nil {
		return x.MessageWithoutFieldBehavior
	}
	return nil
}

func (x *FieldBehaviorMessage) GetOutputOnlyMessage() *FieldBehaviorMessage {
	if x != nil {
		return x.OutputOnlyMessage
	}
	return nil
}

func (x *FieldBehaviorMessage) GetOptionalMessage() *FieldBehaviorMessage {
	if x != nil {
		return x.OptionalMessage
	}
	return nil
}

func (x *FieldBehaviorMessage) GetRepeatedMessage() []*FieldBehaviorMessage {
	if x != nil {
		return x.RepeatedMessage
	}
	return nil
}

func (x *FieldBehaviorMessage) GetRepeatedOutputOnlyMessage() []*FieldBehaviorMessage {
	if x != nil {
		return x.RepeatedOutputOnlyMessage
	}
	return nil
}

func (x *FieldBehaviorMessage) GetRepeatedOptionalMessage() []*FieldBehaviorMessage {
	if x != nil {
		return x.RepeatedOptionalMessage
	}
	return nil
}

func (x *FieldBehaviorMessage) GetMapMessage() map[string]*FieldBehaviorMessage {
	if x != nil {
		return x.MapMessage
	}
	return nil
}

func (x *FieldBehaviorMessage) GetMapOutputOnlyMessage() map[string]*FieldBehaviorMessage {
	if x != nil {
		return x.MapOutputOnlyMessage
	}
	return nil
}

func (x *FieldBehaviorMessage) GetMapOptionalMessage() map[string]*FieldBehaviorMessage {
	if x != nil {
		return x.MapOptionalMessage
	}
	return nil
}

func (x *FieldBehaviorMessage) GetStringMap() map[string]string {
	if x != nil {
		return x.StringMap
	}
	return nil
}

func (x *FieldBehaviorMessage) GetStringList() []string {
	if x != nil {
		return x.StringList
	}
	return nil
}

func (m *FieldBehaviorMessage) GetOneof() isFieldBehaviorMessage_Oneof {
	if m != nil {
		return m.Oneof
	}
	return nil
}

func (x *FieldBehaviorMessage) GetFieldBehaviorMessage() *FieldBehaviorMessage {
	if x, ok := x.GetOneof().(*FieldBehaviorMessage_FieldBehaviorMessage); ok {
		return x.FieldBehaviorMessage
	}
	return nil
}

func (x *FieldBehaviorMessage) GetSmallFieldBehaviorMessage() *SmallFieldBehaviorMessage {
	if x, ok := x.GetOneof().(*FieldBehaviorMessage_SmallFieldBehaviorMessage); ok {
		return x.SmallFieldBehaviorMessage
	}
	return nil
}

type isFieldBehaviorMessage_Oneof interface {
	isFieldBehaviorMessage_Oneof()
}

type FieldBehaviorMessage_FieldBehaviorMessage struct {
	FieldBehaviorMessage *FieldBehaviorMessage `protobuf:"bytes,15,opt,name=field_behavior_message,json=fieldBehaviorMessage,proto3,oneof"`
}

type FieldBehaviorMessage_SmallFieldBehaviorMessage struct {
	SmallFieldBehaviorMessage *SmallFieldBehaviorMessage `protobuf:"bytes,16,opt,name=small_field_behavior_message,json=smallFieldBehaviorMessage,proto3,oneof"`
}

func (*FieldBehaviorMessage_FieldBehaviorMessage) isFieldBehaviorMessage_Oneof() {}

func (*FieldBehaviorMessage_SmallFieldBehaviorMessage) isFieldBehaviorMessage_Oneof() {}

type SmallFieldBehaviorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field           string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	OutputOnlyField string `protobuf:"bytes,2,opt,name=output_only_field,json=outputOnlyField,proto3" json:"output_only_field,omitempty"`
	OptionalField   string `protobuf:"bytes,3,opt,name=optional_field,json=optionalField,proto3" json:"optional_field,omitempty"`
}

func (x *SmallFieldBehaviorMessage) Reset() {
	*x = SmallFieldBehaviorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_example_syntax_v1_fieldbehaviors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmallFieldBehaviorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmallFieldBehaviorMessage) ProtoMessage() {}

func (x *SmallFieldBehaviorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_einride_example_syntax_v1_fieldbehaviors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmallFieldBehaviorMessage.ProtoReflect.Descriptor instead.
func (*SmallFieldBehaviorMessage) Descriptor() ([]byte, []int) {
	return file_einride_example_syntax_v1_fieldbehaviors_proto_rawDescGZIP(), []int{1}
}

func (x *SmallFieldBehaviorMessage) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *SmallFieldBehaviorMessage) GetOutputOnlyField() string {
	if x != nil {
		return x.OutputOnlyField
	}
	return ""
}

func (x *SmallFieldBehaviorMessage) GetOptionalField() string {
	if x != nil {
		return x.OptionalField
	}
	return ""
}

var File_einride_example_syntax_v1_fieldbehaviors_proto protoreflect.FileDescriptor

var file_einride_example_syntax_v1_fieldbehaviors_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x0f, 0x0a,
	0x14, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x30, 0x0a, 0x11, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0f, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2b, 0x0a,
	0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x01, 0x52, 0x0d, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x74, 0x0a, 0x1e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x1b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68,
	0x6f, 0x75, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x12, 0x65, 0x0a, 0x13, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x04,
	0xe2, 0x41, 0x01, 0x03, 0x52, 0x11, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4f, 0x6e, 0x6c, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x60, 0x0a, 0x10, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x01, 0x52, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x5a, 0x0a, 0x10, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x76, 0x0a, 0x1c, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x79,
	0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x04, 0xe2, 0x41,
	0x01, 0x03, 0x52, 0x19, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x71, 0x0a,
	0x19, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x01, 0x52, 0x17, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x60, 0x0a, 0x0b, 0x6d, 0x61, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x6d, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x17, 0x6d, 0x61, 0x70, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4f,
	0x6e, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42,
	0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x14, 0x6d, 0x61, 0x70, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x4f, 0x6e, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x7f, 0x0a, 0x14, 0x6d,
	0x61, 0x70, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x65, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x79, 0x6e, 0x74,
	0x61, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x01, 0x52, 0x12, 0x6d, 0x61, 0x70, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x5d, 0x0a, 0x0a,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3e, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x09, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x67, 0x0a, 0x16,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65,
	0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73,
	0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52,
	0x14, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x77, 0x0a, 0x1c, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x65, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x79,
	0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x48, 0x00, 0x52, 0x19, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x6e,
	0x0a, 0x0f, 0x4d, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x45, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x78,
	0x0a, 0x19, 0x4d, 0x61, 0x70, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x45, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65,
	0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73,
	0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x76, 0x0a, 0x17, 0x4d, 0x61, 0x70, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x45, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x3c, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x07,
	0x0a, 0x05, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x22, 0x90, 0x01, 0x0a, 0x19, 0x53, 0x6d, 0x61, 0x6c,
	0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x30, 0x0a, 0x11, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0f, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2b, 0x0a,
	0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x01, 0x52, 0x0d, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0xfd, 0x01, 0x0a, 0x1d, 0x63,
	0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x6f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e,
	0x74, 0x65, 0x63, 0x68, 0x2f, 0x61, 0x69, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x79, 0x6e,
	0x74, 0x61, 0x78, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x45, 0x53, 0xaa, 0x02, 0x19, 0x45, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53, 0x79,
	0x6e, 0x74, 0x61, 0x78, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x19, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x5c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x25, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x45, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x3a, 0x3a, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3a, 0x3a,
	0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_einride_example_syntax_v1_fieldbehaviors_proto_rawDescOnce sync.Once
	file_einride_example_syntax_v1_fieldbehaviors_proto_rawDescData = file_einride_example_syntax_v1_fieldbehaviors_proto_rawDesc
)

func file_einride_example_syntax_v1_fieldbehaviors_proto_rawDescGZIP() []byte {
	file_einride_example_syntax_v1_fieldbehaviors_proto_rawDescOnce.Do(func() {
		file_einride_example_syntax_v1_fieldbehaviors_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_example_syntax_v1_fieldbehaviors_proto_rawDescData)
	})
	return file_einride_example_syntax_v1_fieldbehaviors_proto_rawDescData
}

var file_einride_example_syntax_v1_fieldbehaviors_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_einride_example_syntax_v1_fieldbehaviors_proto_goTypes = []any{
	(*FieldBehaviorMessage)(nil),      // 0: einride.example.syntax.v1.FieldBehaviorMessage
	(*SmallFieldBehaviorMessage)(nil), // 1: einride.example.syntax.v1.SmallFieldBehaviorMessage
	nil,                               // 2: einride.example.syntax.v1.FieldBehaviorMessage.MapMessageEntry
	nil,                               // 3: einride.example.syntax.v1.FieldBehaviorMessage.MapOutputOnlyMessageEntry
	nil,                               // 4: einride.example.syntax.v1.FieldBehaviorMessage.MapOptionalMessageEntry
	nil,                               // 5: einride.example.syntax.v1.FieldBehaviorMessage.StringMapEntry
}
var file_einride_example_syntax_v1_fieldbehaviors_proto_depIdxs = []int32{
	0,  // 0: einride.example.syntax.v1.FieldBehaviorMessage.message_without_field_behavior:type_name -> einride.example.syntax.v1.FieldBehaviorMessage
	0,  // 1: einride.example.syntax.v1.FieldBehaviorMessage.output_only_message:type_name -> einride.example.syntax.v1.FieldBehaviorMessage
	0,  // 2: einride.example.syntax.v1.FieldBehaviorMessage.optional_message:type_name -> einride.example.syntax.v1.FieldBehaviorMessage
	0,  // 3: einride.example.syntax.v1.FieldBehaviorMessage.repeated_message:type_name -> einride.example.syntax.v1.FieldBehaviorMessage
	0,  // 4: einride.example.syntax.v1.FieldBehaviorMessage.repeated_output_only_message:type_name -> einride.example.syntax.v1.FieldBehaviorMessage
	0,  // 5: einride.example.syntax.v1.FieldBehaviorMessage.repeated_optional_message:type_name -> einride.example.syntax.v1.FieldBehaviorMessage
	2,  // 6: einride.example.syntax.v1.FieldBehaviorMessage.map_message:type_name -> einride.example.syntax.v1.FieldBehaviorMessage.MapMessageEntry
	3,  // 7: einride.example.syntax.v1.FieldBehaviorMessage.map_output_only_message:type_name -> einride.example.syntax.v1.FieldBehaviorMessage.MapOutputOnlyMessageEntry
	4,  // 8: einride.example.syntax.v1.FieldBehaviorMessage.map_optional_message:type_name -> einride.example.syntax.v1.FieldBehaviorMessage.MapOptionalMessageEntry
	5,  // 9: einride.example.syntax.v1.FieldBehaviorMessage.string_map:type_name -> einride.example.syntax.v1.FieldBehaviorMessage.StringMapEntry
	0,  // 10: einride.example.syntax.v1.FieldBehaviorMessage.field_behavior_message:type_name -> einride.example.syntax.v1.FieldBehaviorMessage
	1,  // 11: einride.example.syntax.v1.FieldBehaviorMessage.small_field_behavior_message:type_name -> einride.example.syntax.v1.SmallFieldBehaviorMessage
	0,  // 12: einride.example.syntax.v1.FieldBehaviorMessage.MapMessageEntry.value:type_name -> einride.example.syntax.v1.FieldBehaviorMessage
	0,  // 13: einride.example.syntax.v1.FieldBehaviorMessage.MapOutputOnlyMessageEntry.value:type_name -> einride.example.syntax.v1.FieldBehaviorMessage
	0,  // 14: einride.example.syntax.v1.FieldBehaviorMessage.MapOptionalMessageEntry.value:type_name -> einride.example.syntax.v1.FieldBehaviorMessage
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_einride_example_syntax_v1_fieldbehaviors_proto_init() }
func file_einride_example_syntax_v1_fieldbehaviors_proto_init() {
	if File_einride_example_syntax_v1_fieldbehaviors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_einride_example_syntax_v1_fieldbehaviors_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FieldBehaviorMessage); i {
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
		file_einride_example_syntax_v1_fieldbehaviors_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SmallFieldBehaviorMessage); i {
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
	file_einride_example_syntax_v1_fieldbehaviors_proto_msgTypes[0].OneofWrappers = []any{
		(*FieldBehaviorMessage_FieldBehaviorMessage)(nil),
		(*FieldBehaviorMessage_SmallFieldBehaviorMessage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_einride_example_syntax_v1_fieldbehaviors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_example_syntax_v1_fieldbehaviors_proto_goTypes,
		DependencyIndexes: file_einride_example_syntax_v1_fieldbehaviors_proto_depIdxs,
		MessageInfos:      file_einride_example_syntax_v1_fieldbehaviors_proto_msgTypes,
	}.Build()
	File_einride_example_syntax_v1_fieldbehaviors_proto = out.File
	file_einride_example_syntax_v1_fieldbehaviors_proto_rawDesc = nil
	file_einride_example_syntax_v1_fieldbehaviors_proto_goTypes = nil
	file_einride_example_syntax_v1_fieldbehaviors_proto_depIdxs = nil
}