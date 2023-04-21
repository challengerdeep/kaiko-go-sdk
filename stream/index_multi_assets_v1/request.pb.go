// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: sdk/stream/index_multi_assets_v1/request.proto

package index_multi_assets_v1

import (
	index_v1 "github.com/kaikodata/kaiko-go-sdk/stream/index_v1"
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

// StreamIndexServiceRequestV1
type StreamIndexMultiAssetsServiceRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Index code.
	IndexCode string `protobuf:"bytes,1,opt,name=index_code,json=indexCode,proto3" json:"index_code,omitempty"`
	// Enum indicating type of feed.
	Commodities []index_v1.StreamIndexCommodity `protobuf:"varint,2,rep,packed,name=commodities,proto3,enum=kaikosdk.StreamIndexCommodity" json:"commodities,omitempty"`
}

func (x *StreamIndexMultiAssetsServiceRequestV1) Reset() {
	*x = StreamIndexMultiAssetsServiceRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_stream_index_multi_assets_v1_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamIndexMultiAssetsServiceRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamIndexMultiAssetsServiceRequestV1) ProtoMessage() {}

func (x *StreamIndexMultiAssetsServiceRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_stream_index_multi_assets_v1_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamIndexMultiAssetsServiceRequestV1.ProtoReflect.Descriptor instead.
func (*StreamIndexMultiAssetsServiceRequestV1) Descriptor() ([]byte, []int) {
	return file_sdk_stream_index_multi_assets_v1_request_proto_rawDescGZIP(), []int{0}
}

func (x *StreamIndexMultiAssetsServiceRequestV1) GetIndexCode() string {
	if x != nil {
		return x.IndexCode
	}
	return ""
}

func (x *StreamIndexMultiAssetsServiceRequestV1) GetCommodities() []index_v1.StreamIndexCommodity {
	if x != nil {
		return x.Commodities
	}
	return nil
}

var File_sdk_stream_index_multi_assets_v1_request_proto protoreflect.FileDescriptor

var file_sdk_stream_index_multi_assets_v1_request_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x5f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x1a, 0x23, 0x73, 0x64, 0x6b, 0x2f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x89, 0x01, 0x0a, 0x26, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1e,
	0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x52, 0x0b,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0xa9, 0x01, 0x0a, 0x2a,
	0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x5f, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x54, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x5f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x5f,
	0x76, 0x31, 0xaa, 0x02, 0x22, 0x4b, 0x61, 0x69, 0x6b, 0x6f, 0x53, 0x64, 0x6b, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sdk_stream_index_multi_assets_v1_request_proto_rawDescOnce sync.Once
	file_sdk_stream_index_multi_assets_v1_request_proto_rawDescData = file_sdk_stream_index_multi_assets_v1_request_proto_rawDesc
)

func file_sdk_stream_index_multi_assets_v1_request_proto_rawDescGZIP() []byte {
	file_sdk_stream_index_multi_assets_v1_request_proto_rawDescOnce.Do(func() {
		file_sdk_stream_index_multi_assets_v1_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdk_stream_index_multi_assets_v1_request_proto_rawDescData)
	})
	return file_sdk_stream_index_multi_assets_v1_request_proto_rawDescData
}

var file_sdk_stream_index_multi_assets_v1_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sdk_stream_index_multi_assets_v1_request_proto_goTypes = []interface{}{
	(*StreamIndexMultiAssetsServiceRequestV1)(nil), // 0: kaikosdk.StreamIndexMultiAssetsServiceRequestV1
	(index_v1.StreamIndexCommodity)(0),             // 1: kaikosdk.StreamIndexCommodity
}
var file_sdk_stream_index_multi_assets_v1_request_proto_depIdxs = []int32{
	1, // 0: kaikosdk.StreamIndexMultiAssetsServiceRequestV1.commodities:type_name -> kaikosdk.StreamIndexCommodity
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sdk_stream_index_multi_assets_v1_request_proto_init() }
func file_sdk_stream_index_multi_assets_v1_request_proto_init() {
	if File_sdk_stream_index_multi_assets_v1_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sdk_stream_index_multi_assets_v1_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamIndexMultiAssetsServiceRequestV1); i {
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
			RawDescriptor: file_sdk_stream_index_multi_assets_v1_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sdk_stream_index_multi_assets_v1_request_proto_goTypes,
		DependencyIndexes: file_sdk_stream_index_multi_assets_v1_request_proto_depIdxs,
		MessageInfos:      file_sdk_stream_index_multi_assets_v1_request_proto_msgTypes,
	}.Build()
	File_sdk_stream_index_multi_assets_v1_request_proto = out.File
	file_sdk_stream_index_multi_assets_v1_request_proto_rawDesc = nil
	file_sdk_stream_index_multi_assets_v1_request_proto_goTypes = nil
	file_sdk_stream_index_multi_assets_v1_request_proto_depIdxs = nil
}
