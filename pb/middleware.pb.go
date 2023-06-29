// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: pb/middleware.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MiddlewareCaller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the middleware.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The arguments of the middleware.
	Arguments []string `protobuf:"bytes,2,rep,name=arguments,proto3" json:"arguments,omitempty"`
}

func (x *MiddlewareCaller) Reset() {
	*x = MiddlewareCaller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_middleware_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MiddlewareCaller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiddlewareCaller) ProtoMessage() {}

func (x *MiddlewareCaller) ProtoReflect() protoreflect.Message {
	mi := &file_pb_middleware_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiddlewareCaller.ProtoReflect.Descriptor instead.
func (*MiddlewareCaller) Descriptor() ([]byte, []int) {
	return file_pb_middleware_proto_rawDescGZIP(), []int{0}
}

func (x *MiddlewareCaller) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MiddlewareCaller) GetArguments() []string {
	if x != nil {
		return x.Arguments
	}
	return nil
}

var file_pb_middleware_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: ([]*MiddlewareCaller)(nil),
		Field:         72295830,
		Name:          "pb.middleware",
		Tag:           "bytes,72295830,rep,name=middleware",
		Filename:      "pb/middleware.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// See `Middleware`,
	//
	// repeated pb.MiddlewareCaller middleware = 72295830;
	E_Middleware = &file_pb_middleware_proto_extTypes[0]
)

var File_pb_middleware_proto protoreflect.FileDescriptor

var file_pb_middleware_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x62, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x10, 0x4d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x3a, 0x57, 0x0a, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12,
	0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x96, 0xcb, 0xbc, 0x22, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x0a,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x6f,
	0x70, 0x6b, 0x67, 0x2e, 0x69, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x69, 0x78, 0x65, 0x64, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_middleware_proto_rawDescOnce sync.Once
	file_pb_middleware_proto_rawDescData = file_pb_middleware_proto_rawDesc
)

func file_pb_middleware_proto_rawDescGZIP() []byte {
	file_pb_middleware_proto_rawDescOnce.Do(func() {
		file_pb_middleware_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_middleware_proto_rawDescData)
	})
	return file_pb_middleware_proto_rawDescData
}

var file_pb_middleware_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_middleware_proto_goTypes = []interface{}{
	(*MiddlewareCaller)(nil),           // 0: pb.MiddlewareCaller
	(*descriptorpb.MethodOptions)(nil), // 1: google.protobuf.MethodOptions
}
var file_pb_middleware_proto_depIdxs = []int32{
	1, // 0: pb.middleware:extendee -> google.protobuf.MethodOptions
	0, // 1: pb.middleware:type_name -> pb.MiddlewareCaller
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_middleware_proto_init() }
func file_pb_middleware_proto_init() {
	if File_pb_middleware_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_middleware_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MiddlewareCaller); i {
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
			RawDescriptor: file_pb_middleware_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_pb_middleware_proto_goTypes,
		DependencyIndexes: file_pb_middleware_proto_depIdxs,
		MessageInfos:      file_pb_middleware_proto_msgTypes,
		ExtensionInfos:    file_pb_middleware_proto_extTypes,
	}.Build()
	File_pb_middleware_proto = out.File
	file_pb_middleware_proto_rawDesc = nil
	file_pb_middleware_proto_goTypes = nil
	file_pb_middleware_proto_depIdxs = nil
}