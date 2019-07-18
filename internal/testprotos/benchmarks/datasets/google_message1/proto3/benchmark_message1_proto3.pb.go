// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datasets/google_message1/proto3/benchmark_message1_proto3.proto

package proto3

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

type GoogleMessage1 struct {
	state         protoimpl.MessageState
	Field1        string                    `protobuf:"bytes,1,opt,name=field1,proto3" json:"field1,omitempty"`
	Field9        string                    `protobuf:"bytes,9,opt,name=field9,proto3" json:"field9,omitempty"`
	Field18       string                    `protobuf:"bytes,18,opt,name=field18,proto3" json:"field18,omitempty"`
	Field80       bool                      `protobuf:"varint,80,opt,name=field80,proto3" json:"field80,omitempty"`
	Field81       bool                      `protobuf:"varint,81,opt,name=field81,proto3" json:"field81,omitempty"`
	Field2        int32                     `protobuf:"varint,2,opt,name=field2,proto3" json:"field2,omitempty"`
	Field3        int32                     `protobuf:"varint,3,opt,name=field3,proto3" json:"field3,omitempty"`
	Field280      int32                     `protobuf:"varint,280,opt,name=field280,proto3" json:"field280,omitempty"`
	Field6        int32                     `protobuf:"varint,6,opt,name=field6,proto3" json:"field6,omitempty"`
	Field22       int64                     `protobuf:"varint,22,opt,name=field22,proto3" json:"field22,omitempty"`
	Field4        string                    `protobuf:"bytes,4,opt,name=field4,proto3" json:"field4,omitempty"`
	Field5        []uint64                  `protobuf:"fixed64,5,rep,packed,name=field5,proto3" json:"field5,omitempty"`
	Field59       bool                      `protobuf:"varint,59,opt,name=field59,proto3" json:"field59,omitempty"`
	Field7        string                    `protobuf:"bytes,7,opt,name=field7,proto3" json:"field7,omitempty"`
	Field16       int32                     `protobuf:"varint,16,opt,name=field16,proto3" json:"field16,omitempty"`
	Field130      int32                     `protobuf:"varint,130,opt,name=field130,proto3" json:"field130,omitempty"`
	Field12       bool                      `protobuf:"varint,12,opt,name=field12,proto3" json:"field12,omitempty"`
	Field17       bool                      `protobuf:"varint,17,opt,name=field17,proto3" json:"field17,omitempty"`
	Field13       bool                      `protobuf:"varint,13,opt,name=field13,proto3" json:"field13,omitempty"`
	Field14       bool                      `protobuf:"varint,14,opt,name=field14,proto3" json:"field14,omitempty"`
	Field104      int32                     `protobuf:"varint,104,opt,name=field104,proto3" json:"field104,omitempty"`
	Field100      int32                     `protobuf:"varint,100,opt,name=field100,proto3" json:"field100,omitempty"`
	Field101      int32                     `protobuf:"varint,101,opt,name=field101,proto3" json:"field101,omitempty"`
	Field102      string                    `protobuf:"bytes,102,opt,name=field102,proto3" json:"field102,omitempty"`
	Field103      string                    `protobuf:"bytes,103,opt,name=field103,proto3" json:"field103,omitempty"`
	Field29       int32                     `protobuf:"varint,29,opt,name=field29,proto3" json:"field29,omitempty"`
	Field30       bool                      `protobuf:"varint,30,opt,name=field30,proto3" json:"field30,omitempty"`
	Field60       int32                     `protobuf:"varint,60,opt,name=field60,proto3" json:"field60,omitempty"`
	Field271      int32                     `protobuf:"varint,271,opt,name=field271,proto3" json:"field271,omitempty"`
	Field272      int32                     `protobuf:"varint,272,opt,name=field272,proto3" json:"field272,omitempty"`
	Field150      int32                     `protobuf:"varint,150,opt,name=field150,proto3" json:"field150,omitempty"`
	Field23       int32                     `protobuf:"varint,23,opt,name=field23,proto3" json:"field23,omitempty"`
	Field24       bool                      `protobuf:"varint,24,opt,name=field24,proto3" json:"field24,omitempty"`
	Field25       int32                     `protobuf:"varint,25,opt,name=field25,proto3" json:"field25,omitempty"`
	Field15       *GoogleMessage1SubMessage `protobuf:"bytes,15,opt,name=field15,proto3" json:"field15,omitempty"`
	Field78       bool                      `protobuf:"varint,78,opt,name=field78,proto3" json:"field78,omitempty"`
	Field67       int32                     `protobuf:"varint,67,opt,name=field67,proto3" json:"field67,omitempty"`
	Field68       int32                     `protobuf:"varint,68,opt,name=field68,proto3" json:"field68,omitempty"`
	Field128      int32                     `protobuf:"varint,128,opt,name=field128,proto3" json:"field128,omitempty"`
	Field129      string                    `protobuf:"bytes,129,opt,name=field129,proto3" json:"field129,omitempty"`
	Field131      int32                     `protobuf:"varint,131,opt,name=field131,proto3" json:"field131,omitempty"`
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GoogleMessage1) Reset() {
	*x = GoogleMessage1{}
}

func (x *GoogleMessage1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleMessage1) ProtoMessage() {}

func (x *GoogleMessage1) ProtoReflect() protoreflect.Message {
	mi := &file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleMessage1.ProtoReflect.Descriptor instead.
func (*GoogleMessage1) Descriptor() ([]byte, []int) {
	return file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_rawDescGZIP(), []int{0}
}

func (x *GoogleMessage1) GetField1() string {
	if x != nil {
		return x.Field1
	}
	return ""
}

func (x *GoogleMessage1) GetField9() string {
	if x != nil {
		return x.Field9
	}
	return ""
}

func (x *GoogleMessage1) GetField18() string {
	if x != nil {
		return x.Field18
	}
	return ""
}

func (x *GoogleMessage1) GetField80() bool {
	if x != nil {
		return x.Field80
	}
	return false
}

func (x *GoogleMessage1) GetField81() bool {
	if x != nil {
		return x.Field81
	}
	return false
}

func (x *GoogleMessage1) GetField2() int32 {
	if x != nil {
		return x.Field2
	}
	return 0
}

func (x *GoogleMessage1) GetField3() int32 {
	if x != nil {
		return x.Field3
	}
	return 0
}

func (x *GoogleMessage1) GetField280() int32 {
	if x != nil {
		return x.Field280
	}
	return 0
}

func (x *GoogleMessage1) GetField6() int32 {
	if x != nil {
		return x.Field6
	}
	return 0
}

func (x *GoogleMessage1) GetField22() int64 {
	if x != nil {
		return x.Field22
	}
	return 0
}

func (x *GoogleMessage1) GetField4() string {
	if x != nil {
		return x.Field4
	}
	return ""
}

func (x *GoogleMessage1) GetField5() []uint64 {
	if x != nil {
		return x.Field5
	}
	return nil
}

func (x *GoogleMessage1) GetField59() bool {
	if x != nil {
		return x.Field59
	}
	return false
}

func (x *GoogleMessage1) GetField7() string {
	if x != nil {
		return x.Field7
	}
	return ""
}

func (x *GoogleMessage1) GetField16() int32 {
	if x != nil {
		return x.Field16
	}
	return 0
}

func (x *GoogleMessage1) GetField130() int32 {
	if x != nil {
		return x.Field130
	}
	return 0
}

func (x *GoogleMessage1) GetField12() bool {
	if x != nil {
		return x.Field12
	}
	return false
}

func (x *GoogleMessage1) GetField17() bool {
	if x != nil {
		return x.Field17
	}
	return false
}

func (x *GoogleMessage1) GetField13() bool {
	if x != nil {
		return x.Field13
	}
	return false
}

func (x *GoogleMessage1) GetField14() bool {
	if x != nil {
		return x.Field14
	}
	return false
}

func (x *GoogleMessage1) GetField104() int32 {
	if x != nil {
		return x.Field104
	}
	return 0
}

func (x *GoogleMessage1) GetField100() int32 {
	if x != nil {
		return x.Field100
	}
	return 0
}

func (x *GoogleMessage1) GetField101() int32 {
	if x != nil {
		return x.Field101
	}
	return 0
}

func (x *GoogleMessage1) GetField102() string {
	if x != nil {
		return x.Field102
	}
	return ""
}

func (x *GoogleMessage1) GetField103() string {
	if x != nil {
		return x.Field103
	}
	return ""
}

func (x *GoogleMessage1) GetField29() int32 {
	if x != nil {
		return x.Field29
	}
	return 0
}

func (x *GoogleMessage1) GetField30() bool {
	if x != nil {
		return x.Field30
	}
	return false
}

func (x *GoogleMessage1) GetField60() int32 {
	if x != nil {
		return x.Field60
	}
	return 0
}

func (x *GoogleMessage1) GetField271() int32 {
	if x != nil {
		return x.Field271
	}
	return 0
}

func (x *GoogleMessage1) GetField272() int32 {
	if x != nil {
		return x.Field272
	}
	return 0
}

func (x *GoogleMessage1) GetField150() int32 {
	if x != nil {
		return x.Field150
	}
	return 0
}

func (x *GoogleMessage1) GetField23() int32 {
	if x != nil {
		return x.Field23
	}
	return 0
}

func (x *GoogleMessage1) GetField24() bool {
	if x != nil {
		return x.Field24
	}
	return false
}

func (x *GoogleMessage1) GetField25() int32 {
	if x != nil {
		return x.Field25
	}
	return 0
}

func (x *GoogleMessage1) GetField15() *GoogleMessage1SubMessage {
	if x != nil {
		return x.Field15
	}
	return nil
}

func (x *GoogleMessage1) GetField78() bool {
	if x != nil {
		return x.Field78
	}
	return false
}

func (x *GoogleMessage1) GetField67() int32 {
	if x != nil {
		return x.Field67
	}
	return 0
}

func (x *GoogleMessage1) GetField68() int32 {
	if x != nil {
		return x.Field68
	}
	return 0
}

func (x *GoogleMessage1) GetField128() int32 {
	if x != nil {
		return x.Field128
	}
	return 0
}

func (x *GoogleMessage1) GetField129() string {
	if x != nil {
		return x.Field129
	}
	return ""
}

func (x *GoogleMessage1) GetField131() int32 {
	if x != nil {
		return x.Field131
	}
	return 0
}

type GoogleMessage1SubMessage struct {
	state         protoimpl.MessageState
	Field1        int32  `protobuf:"varint,1,opt,name=field1,proto3" json:"field1,omitempty"`
	Field2        int32  `protobuf:"varint,2,opt,name=field2,proto3" json:"field2,omitempty"`
	Field3        int32  `protobuf:"varint,3,opt,name=field3,proto3" json:"field3,omitempty"`
	Field15       string `protobuf:"bytes,15,opt,name=field15,proto3" json:"field15,omitempty"`
	Field12       bool   `protobuf:"varint,12,opt,name=field12,proto3" json:"field12,omitempty"`
	Field13       int64  `protobuf:"varint,13,opt,name=field13,proto3" json:"field13,omitempty"`
	Field14       int64  `protobuf:"varint,14,opt,name=field14,proto3" json:"field14,omitempty"`
	Field16       int32  `protobuf:"varint,16,opt,name=field16,proto3" json:"field16,omitempty"`
	Field19       int32  `protobuf:"varint,19,opt,name=field19,proto3" json:"field19,omitempty"`
	Field20       bool   `protobuf:"varint,20,opt,name=field20,proto3" json:"field20,omitempty"`
	Field28       bool   `protobuf:"varint,28,opt,name=field28,proto3" json:"field28,omitempty"`
	Field21       uint64 `protobuf:"fixed64,21,opt,name=field21,proto3" json:"field21,omitempty"`
	Field22       int32  `protobuf:"varint,22,opt,name=field22,proto3" json:"field22,omitempty"`
	Field23       bool   `protobuf:"varint,23,opt,name=field23,proto3" json:"field23,omitempty"`
	Field206      bool   `protobuf:"varint,206,opt,name=field206,proto3" json:"field206,omitempty"`
	Field203      uint32 `protobuf:"fixed32,203,opt,name=field203,proto3" json:"field203,omitempty"`
	Field204      int32  `protobuf:"varint,204,opt,name=field204,proto3" json:"field204,omitempty"`
	Field205      string `protobuf:"bytes,205,opt,name=field205,proto3" json:"field205,omitempty"`
	Field207      uint64 `protobuf:"varint,207,opt,name=field207,proto3" json:"field207,omitempty"`
	Field300      uint64 `protobuf:"varint,300,opt,name=field300,proto3" json:"field300,omitempty"`
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GoogleMessage1SubMessage) Reset() {
	*x = GoogleMessage1SubMessage{}
}

func (x *GoogleMessage1SubMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleMessage1SubMessage) ProtoMessage() {}

func (x *GoogleMessage1SubMessage) ProtoReflect() protoreflect.Message {
	mi := &file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleMessage1SubMessage.ProtoReflect.Descriptor instead.
func (*GoogleMessage1SubMessage) Descriptor() ([]byte, []int) {
	return file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_rawDescGZIP(), []int{1}
}

func (x *GoogleMessage1SubMessage) GetField1() int32 {
	if x != nil {
		return x.Field1
	}
	return 0
}

func (x *GoogleMessage1SubMessage) GetField2() int32 {
	if x != nil {
		return x.Field2
	}
	return 0
}

func (x *GoogleMessage1SubMessage) GetField3() int32 {
	if x != nil {
		return x.Field3
	}
	return 0
}

func (x *GoogleMessage1SubMessage) GetField15() string {
	if x != nil {
		return x.Field15
	}
	return ""
}

func (x *GoogleMessage1SubMessage) GetField12() bool {
	if x != nil {
		return x.Field12
	}
	return false
}

func (x *GoogleMessage1SubMessage) GetField13() int64 {
	if x != nil {
		return x.Field13
	}
	return 0
}

func (x *GoogleMessage1SubMessage) GetField14() int64 {
	if x != nil {
		return x.Field14
	}
	return 0
}

func (x *GoogleMessage1SubMessage) GetField16() int32 {
	if x != nil {
		return x.Field16
	}
	return 0
}

func (x *GoogleMessage1SubMessage) GetField19() int32 {
	if x != nil {
		return x.Field19
	}
	return 0
}

func (x *GoogleMessage1SubMessage) GetField20() bool {
	if x != nil {
		return x.Field20
	}
	return false
}

func (x *GoogleMessage1SubMessage) GetField28() bool {
	if x != nil {
		return x.Field28
	}
	return false
}

func (x *GoogleMessage1SubMessage) GetField21() uint64 {
	if x != nil {
		return x.Field21
	}
	return 0
}

func (x *GoogleMessage1SubMessage) GetField22() int32 {
	if x != nil {
		return x.Field22
	}
	return 0
}

func (x *GoogleMessage1SubMessage) GetField23() bool {
	if x != nil {
		return x.Field23
	}
	return false
}

func (x *GoogleMessage1SubMessage) GetField206() bool {
	if x != nil {
		return x.Field206
	}
	return false
}

func (x *GoogleMessage1SubMessage) GetField203() uint32 {
	if x != nil {
		return x.Field203
	}
	return 0
}

func (x *GoogleMessage1SubMessage) GetField204() int32 {
	if x != nil {
		return x.Field204
	}
	return 0
}

func (x *GoogleMessage1SubMessage) GetField205() string {
	if x != nil {
		return x.Field205
	}
	return ""
}

func (x *GoogleMessage1SubMessage) GetField207() uint64 {
	if x != nil {
		return x.Field207
	}
	return 0
}

func (x *GoogleMessage1SubMessage) GetField300() uint64 {
	if x != nil {
		return x.Field300
	}
	return 0
}

var File_datasets_google_message1_proto3_benchmark_message1_proto3_proto protoreflect.FileDescriptor

var file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x31, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33, 0x22, 0xf9, 0x08, 0x0a, 0x0e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x12,
	0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x39, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x39, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x31, 0x38, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31,
	0x38, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x38, 0x30, 0x18, 0x50, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x38, 0x30, 0x12, 0x18, 0x0a, 0x07, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x38, 0x31, 0x18, 0x51, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x38, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x33, 0x12, 0x1b, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x38,
	0x30, 0x18, 0x98, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32,
	0x38, 0x30, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x32, 0x32, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x32, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x34, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x34, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x35, 0x18, 0x05, 0x20, 0x03, 0x28, 0x06, 0x52, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x35, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x35, 0x39, 0x18,
	0x3b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x35, 0x39, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x37, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x37, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31,
	0x36, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x36,
	0x12, 0x1b, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x33, 0x30, 0x18, 0x82, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x33, 0x30, 0x12, 0x18, 0x0a,
	0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x32, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x31, 0x37, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31,
	0x37, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x33, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x33, 0x12, 0x18, 0x0a, 0x07, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x31, 0x34, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x31, 0x34, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30,
	0x34, 0x18, 0x68, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30,
	0x34, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x30, 0x18, 0x64, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x30, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x31, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x31, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x31, 0x30, 0x32, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x31, 0x30, 0x32, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30,
	0x33, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30,
	0x33, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x39, 0x18, 0x1d, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x39, 0x12, 0x18, 0x0a, 0x07, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x33, 0x30, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x33, 0x30, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x30,
	0x18, 0x3c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x30, 0x12,
	0x1b, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x37, 0x31, 0x18, 0x8f, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x37, 0x31, 0x12, 0x1b, 0x0a, 0x08,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x37, 0x32, 0x18, 0x90, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x37, 0x32, 0x12, 0x1b, 0x0a, 0x08, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x31, 0x35, 0x30, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x31, 0x35, 0x30, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32,
	0x33, 0x18, 0x17, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x33,
	0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x34, 0x18, 0x18, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x34, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x32, 0x35, 0x18, 0x19, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x32, 0x35, 0x12, 0x45, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x35, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72,
	0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x35, 0x12, 0x18, 0x0a, 0x07, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x37, 0x38, 0x18, 0x4e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x37, 0x38, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x37,
	0x18, 0x43, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x37, 0x12,
	0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x38, 0x18, 0x44, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x38, 0x12, 0x1b, 0x0a, 0x08, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x31, 0x32, 0x38, 0x18, 0x80, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x31, 0x32, 0x38, 0x12, 0x1b, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31,
	0x32, 0x39, 0x18, 0x81, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x31, 0x32, 0x39, 0x12, 0x1b, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x33, 0x31, 0x18,
	0x83, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x33, 0x31,
	0x22, 0xae, 0x04, 0x0a, 0x18, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x31, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x33, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x35,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x35, 0x12,
	0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x32, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x31, 0x33, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x31, 0x33, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x34, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x34, 0x12, 0x18, 0x0a,
	0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x36, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x36, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x31, 0x39, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31,
	0x39, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x30, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x30, 0x12, 0x18, 0x0a, 0x07, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x32, 0x38, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x32, 0x38, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x31,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x06, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x31, 0x12,
	0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x32, 0x18, 0x16, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x32, 0x33, 0x18, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x32, 0x33, 0x12, 0x1b, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x30, 0x36, 0x18,
	0xce, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x30, 0x36,
	0x12, 0x1b, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x30, 0x33, 0x18, 0xcb, 0x01, 0x20,
	0x01, 0x28, 0x07, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x30, 0x33, 0x12, 0x1b, 0x0a,
	0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x30, 0x34, 0x18, 0xcc, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x30, 0x34, 0x12, 0x1b, 0x0a, 0x08, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x32, 0x30, 0x35, 0x18, 0xcd, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x32, 0x30, 0x35, 0x12, 0x1b, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x32, 0x30, 0x37, 0x18, 0xcf, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x32, 0x30, 0x37, 0x12, 0x1b, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x30, 0x30,
	0x18, 0xac, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x30,
	0x30, 0x42, 0x80, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d,
	0x61, 0x72, 0x6b, 0x73, 0x48, 0x01, 0x5a, 0x59, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b,
	0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_rawDescOnce sync.Once
	file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_rawDescData = file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_rawDesc
)

func file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_rawDescGZIP() []byte {
	file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_rawDescOnce.Do(func() {
		file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_rawDescData = protoimpl.X.CompressGZIP(file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_rawDescData)
	})
	return file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_rawDescData
}

var file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_goTypes = []interface{}{
	(*GoogleMessage1)(nil),           // 0: benchmarks.proto3.GoogleMessage1
	(*GoogleMessage1SubMessage)(nil), // 1: benchmarks.proto3.GoogleMessage1SubMessage
}
var file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_depIdxs = []int32{
	1, // benchmarks.proto3.GoogleMessage1.field15:type_name -> benchmarks.proto3.GoogleMessage1SubMessage
	1, // starting offset of method output_type sub-list
	1, // starting offset of method input_type sub-list
	1, // starting offset of extension type_name sub-list
	1, // starting offset of extension extendee sub-list
	0, // starting offset of field type_name sub-list
}

func init() { file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_init() }
func file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_init() {
	if File_datasets_google_message1_proto3_benchmark_message1_proto3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleMessage1); i {
			case 0:
				return &v.state
			case 42:
				return &v.sizeCache
			case 43:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleMessage1SubMessage); i {
			case 0:
				return &v.state
			case 21:
				return &v.sizeCache
			case 22:
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
			RawDescriptor: file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_goTypes,
		DependencyIndexes: file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_depIdxs,
		MessageInfos:      file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_msgTypes,
	}.Build()
	File_datasets_google_message1_proto3_benchmark_message1_proto3_proto = out.File
	file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_rawDesc = nil
	file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_goTypes = nil
	file_datasets_google_message1_proto3_benchmark_message1_proto3_proto_depIdxs = nil
}
