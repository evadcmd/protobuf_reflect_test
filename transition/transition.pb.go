// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: transition.proto

package transition

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

type Param_Cond int32

const (
	Param_NONE Param_Cond = 0
	Param_MIN  Param_Cond = 1
	Param_MAX  Param_Cond = 2
)

// Enum value maps for Param_Cond.
var (
	Param_Cond_name = map[int32]string{
		0: "NONE",
		1: "MIN",
		2: "MAX",
	}
	Param_Cond_value = map[string]int32{
		"NONE": 0,
		"MIN":  1,
		"MAX":  2,
	}
)

func (x Param_Cond) Enum() *Param_Cond {
	p := new(Param_Cond)
	*p = x
	return p
}

func (x Param_Cond) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Param_Cond) Descriptor() protoreflect.EnumDescriptor {
	return file_transition_proto_enumTypes[0].Descriptor()
}

func (Param_Cond) Type() protoreflect.EnumType {
	return &file_transition_proto_enumTypes[0]
}

func (x Param_Cond) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Param_Cond.Descriptor instead.
func (Param_Cond) EnumDescriptor() ([]byte, []int) {
	return file_transition_proto_rawDescGZIP(), []int{0, 0}
}

type Res_Tag int32

const (
	Res_SUCCESS Res_Tag = 0
	Res_WARNING Res_Tag = 1
	Res_FAILED  Res_Tag = 2
)

// Enum value maps for Res_Tag.
var (
	Res_Tag_name = map[int32]string{
		0: "SUCCESS",
		1: "WARNING",
		2: "FAILED",
	}
	Res_Tag_value = map[string]int32{
		"SUCCESS": 0,
		"WARNING": 1,
		"FAILED":  2,
	}
)

func (x Res_Tag) Enum() *Res_Tag {
	p := new(Res_Tag)
	*p = x
	return p
}

func (x Res_Tag) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Res_Tag) Descriptor() protoreflect.EnumDescriptor {
	return file_transition_proto_enumTypes[1].Descriptor()
}

func (Res_Tag) Type() protoreflect.EnumType {
	return &file_transition_proto_enumTypes[1]
}

func (x Res_Tag) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Res_Tag.Descriptor instead.
func (Res_Tag) EnumDescriptor() ([]byte, []int) {
	return file_transition_proto_rawDescGZIP(), []int{1, 0}
}

type Param struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X    int32      `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y    int32      `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Z    int32      `protobuf:"varint,3,opt,name=z,proto3" json:"z,omitempty"`
	Cond Param_Cond `protobuf:"varint,4,opt,name=cond,proto3,enum=transition.Param_Cond" json:"cond,omitempty"`
}

func (x *Param) Reset() {
	*x = Param{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Param) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Param) ProtoMessage() {}

func (x *Param) ProtoReflect() protoreflect.Message {
	mi := &file_transition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Param.ProtoReflect.Descriptor instead.
func (*Param) Descriptor() ([]byte, []int) {
	return file_transition_proto_rawDescGZIP(), []int{0}
}

func (x *Param) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Param) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Param) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *Param) GetCond() Param_Cond {
	if x != nil {
		return x.Cond
	}
	return Param_NONE
}

type Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Tags  []Res_Tag `protobuf:"varint,2,rep,packed,name=tags,proto3,enum=transition.Res_Tag" json:"tags,omitempty"`
}

func (x *Res) Reset() {
	*x = Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Res) ProtoMessage() {}

func (x *Res) ProtoReflect() protoreflect.Message {
	mi := &file_transition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Res.ProtoReflect.Descriptor instead.
func (*Res) Descriptor() ([]byte, []int) {
	return file_transition_proto_rawDescGZIP(), []int{1}
}

func (x *Res) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Res) GetTags() []Res_Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_transition_proto protoreflect.FileDescriptor

var file_transition_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x81,
	0x01, 0x0a, 0x05, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x01, 0x7a, 0x12, 0x2a, 0x0a, 0x04, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x52, 0x04, 0x63, 0x6f, 0x6e, 0x64, 0x22, 0x22,
	0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x4d, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x58,
	0x10, 0x02, 0x22, 0x71, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x27, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x13, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x2e, 0x54,
	0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x2b, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x10, 0x02, 0x32, 0xdd, 0x01, 0x0a, 0x0a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x06, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x11,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x1a, 0x0f, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x11, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x0f, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x33, 0x0a, 0x09, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x11, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x0f, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x28, 0x01, 0x12,
	0x34, 0x0a, 0x08, 0x42, 0x69, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x11, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x0f,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x20, 0x5a, 0x1e, 0x70, 0x72, 0x6f, 0x62, 0x75, 0x66, 0x5f,
	0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transition_proto_rawDescOnce sync.Once
	file_transition_proto_rawDescData = file_transition_proto_rawDesc
)

func file_transition_proto_rawDescGZIP() []byte {
	file_transition_proto_rawDescOnce.Do(func() {
		file_transition_proto_rawDescData = protoimpl.X.CompressGZIP(file_transition_proto_rawDescData)
	})
	return file_transition_proto_rawDescData
}

var file_transition_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_transition_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transition_proto_goTypes = []interface{}{
	(Param_Cond)(0), // 0: transition.Param.Cond
	(Res_Tag)(0),    // 1: transition.Res.Tag
	(*Param)(nil),   // 2: transition.Param
	(*Res)(nil),     // 3: transition.Res
}
var file_transition_proto_depIdxs = []int32{
	0, // 0: transition.Param.cond:type_name -> transition.Param.Cond
	1, // 1: transition.Res.tags:type_name -> transition.Res.Tag
	2, // 2: transition.Transition.Simple:input_type -> transition.Param
	2, // 3: transition.Transition.StreamResp:input_type -> transition.Param
	2, // 4: transition.Transition.StreamReq:input_type -> transition.Param
	2, // 5: transition.Transition.BiStream:input_type -> transition.Param
	3, // 6: transition.Transition.Simple:output_type -> transition.Res
	3, // 7: transition.Transition.StreamResp:output_type -> transition.Res
	3, // 8: transition.Transition.StreamReq:output_type -> transition.Res
	3, // 9: transition.Transition.BiStream:output_type -> transition.Res
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_transition_proto_init() }
func file_transition_proto_init() {
	if File_transition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Param); i {
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
		file_transition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Res); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transition_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transition_proto_goTypes,
		DependencyIndexes: file_transition_proto_depIdxs,
		EnumInfos:         file_transition_proto_enumTypes,
		MessageInfos:      file_transition_proto_msgTypes,
	}.Build()
	File_transition_proto = out.File
	file_transition_proto_rawDesc = nil
	file_transition_proto_goTypes = nil
	file_transition_proto_depIdxs = nil
}
