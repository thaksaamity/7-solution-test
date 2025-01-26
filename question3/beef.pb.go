// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.0--rc3
// source: beef.proto

package question3

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SummaryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SummaryRequest) Reset() {
	*x = SummaryRequest{}
	mi := &file_beef_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryRequest) ProtoMessage() {}

func (x *SummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_beef_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryRequest.ProtoReflect.Descriptor instead.
func (*SummaryRequest) Descriptor() ([]byte, []int) {
	return file_beef_proto_rawDescGZIP(), []int{0}
}

type SummaryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Beef          map[string]int32       `protobuf:"bytes,1,rep,name=beef,proto3" json:"beef,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SummaryResponse) Reset() {
	*x = SummaryResponse{}
	mi := &file_beef_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SummaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryResponse) ProtoMessage() {}

func (x *SummaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_beef_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryResponse.ProtoReflect.Descriptor instead.
func (*SummaryResponse) Descriptor() ([]byte, []int) {
	return file_beef_proto_rawDescGZIP(), []int{1}
}

func (x *SummaryResponse) GetBeef() map[string]int32 {
	if x != nil {
		return x.Beef
	}
	return nil
}

var File_beef_proto protoreflect.FileDescriptor

var file_beef_proto_rawDesc = string([]byte{
	0x0a, 0x0a, 0x62, 0x65, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x33, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x0f, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x62,
	0x65, 0x65, 0x66, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x33, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x65, 0x65, 0x66, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x04, 0x62, 0x65, 0x65, 0x66, 0x1a, 0x37, 0x0a, 0x09, 0x42, 0x65, 0x65, 0x66, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x67,
	0x0a, 0x0b, 0x42, 0x65, 0x65, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x19, 0x2e, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x33, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x33, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_beef_proto_rawDescOnce sync.Once
	file_beef_proto_rawDescData []byte
)

func file_beef_proto_rawDescGZIP() []byte {
	file_beef_proto_rawDescOnce.Do(func() {
		file_beef_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_beef_proto_rawDesc), len(file_beef_proto_rawDesc)))
	})
	return file_beef_proto_rawDescData
}

var file_beef_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_beef_proto_goTypes = []any{
	(*SummaryRequest)(nil),  // 0: question3.SummaryRequest
	(*SummaryResponse)(nil), // 1: question3.SummaryResponse
	nil,                     // 2: question3.SummaryResponse.BeefEntry
}
var file_beef_proto_depIdxs = []int32{
	2, // 0: question3.SummaryResponse.beef:type_name -> question3.SummaryResponse.BeefEntry
	0, // 1: question3.BeefService.GetSummary:input_type -> question3.SummaryRequest
	1, // 2: question3.BeefService.GetSummary:output_type -> question3.SummaryResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_beef_proto_init() }
func file_beef_proto_init() {
	if File_beef_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_beef_proto_rawDesc), len(file_beef_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_beef_proto_goTypes,
		DependencyIndexes: file_beef_proto_depIdxs,
		MessageInfos:      file_beef_proto_msgTypes,
	}.Build()
	File_beef_proto = out.File
	file_beef_proto_goTypes = nil
	file_beef_proto_depIdxs = nil
}
