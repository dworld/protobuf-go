// Code generated by protoc-gen-go. DO NOT EDIT.
// source: import_public/sub2/a.proto

package sub2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 0)
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(0 - protoimpl.MinVersion)
)

type Sub2Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Sub2Message) Reset() {
	*x = Sub2Message{}
}

func (x *Sub2Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sub2Message) ProtoMessage() {}

func (x *Sub2Message) ProtoReflect() protoreflect.Message {
	mi := &file_import_public_sub2_a_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sub2Message.ProtoReflect.Descriptor instead.
func (*Sub2Message) Descriptor() ([]byte, []int) {
	return file_import_public_sub2_a_proto_rawDescGZIP(), []int{0}
}

var File_import_public_sub2_a_proto protoreflect.FileDescriptor

var file_import_public_sub2_a_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f,
	0x73, 0x75, 0x62, 0x32, 0x2f, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x69, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x73, 0x75, 0x62, 0x32, 0x22,
	0x0d, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x32, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x4a,
	0x5a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6d, 0x64,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x73, 0x75, 0x62, 0x32,
}

var (
	file_import_public_sub2_a_proto_rawDescOnce sync.Once
	file_import_public_sub2_a_proto_rawDescData = file_import_public_sub2_a_proto_rawDesc
)

func file_import_public_sub2_a_proto_rawDescGZIP() []byte {
	file_import_public_sub2_a_proto_rawDescOnce.Do(func() {
		file_import_public_sub2_a_proto_rawDescData = protoimpl.X.CompressGZIP(file_import_public_sub2_a_proto_rawDescData)
	})
	return file_import_public_sub2_a_proto_rawDescData
}

var file_import_public_sub2_a_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_import_public_sub2_a_proto_goTypes = []interface{}{
	(*Sub2Message)(nil), // 0: goproto.protoc.import_public.sub2.Sub2Message
}
var file_import_public_sub2_a_proto_depIdxs = []int32{
	0, // starting offset of method output_type sub-list
	0, // starting offset of method input_type sub-list
	0, // starting offset of extension type_name sub-list
	0, // starting offset of extension extendee sub-list
	0, // starting offset of field type_name sub-list
}

func init() { file_import_public_sub2_a_proto_init() }
func file_import_public_sub2_a_proto_init() {
	if File_import_public_sub2_a_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_import_public_sub2_a_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sub2Message); i {
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
			RawDescriptor: file_import_public_sub2_a_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_import_public_sub2_a_proto_goTypes,
		DependencyIndexes: file_import_public_sub2_a_proto_depIdxs,
		MessageInfos:      file_import_public_sub2_a_proto_msgTypes,
	}.Build()
	File_import_public_sub2_a_proto = out.File
	file_import_public_sub2_a_proto_rawDesc = nil
	file_import_public_sub2_a_proto_goTypes = nil
	file_import_public_sub2_a_proto_depIdxs = nil
}
