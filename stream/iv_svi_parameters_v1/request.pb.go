// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: sdk/stream/iv_svi_parameters_v1/request.proto

package iv_svi_parameters_v1

import (
	core "github.com/kaikodata/kaiko-go-sdk/core"
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

// StreamIvSviParametersRequestV1 is IV SVI parameter stream request.
type StreamIvSviParametersRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// base/quote assets. wildcard is supported for both base and quote assets.
	Assets *core.Assets `protobuf:"bytes,1,opt,name=assets,proto3" json:"assets,omitempty"`
	// Exchanges. wildcard is supported.
	Exchanges string `protobuf:"bytes,2,opt,name=exchanges,proto3" json:"exchanges,omitempty"`
}

func (x *StreamIvSviParametersRequestV1) Reset() {
	*x = StreamIvSviParametersRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_stream_iv_svi_parameters_v1_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamIvSviParametersRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamIvSviParametersRequestV1) ProtoMessage() {}

func (x *StreamIvSviParametersRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_stream_iv_svi_parameters_v1_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamIvSviParametersRequestV1.ProtoReflect.Descriptor instead.
func (*StreamIvSviParametersRequestV1) Descriptor() ([]byte, []int) {
	return file_sdk_stream_iv_svi_parameters_v1_request_proto_rawDescGZIP(), []int{0}
}

func (x *StreamIvSviParametersRequestV1) GetAssets() *core.Assets {
	if x != nil {
		return x.Assets
	}
	return nil
}

func (x *StreamIvSviParametersRequestV1) GetExchanges() string {
	if x != nil {
		return x.Exchanges
	}
	return ""
}

var File_sdk_stream_iv_svi_parameters_v1_request_proto protoreflect.FileDescriptor

var file_sdk_stream_iv_svi_parameters_v1_request_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x69, 0x76, 0x5f,
	0x73, 0x76, 0x69, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x1a, 0x15, 0x73, 0x64, 0x6b, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x68, 0x0a, 0x1e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x76, 0x53, 0x76, 0x69, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x31, 0x12, 0x28, 0x0a, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x73, 0x52, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x42, 0xaa, 0x01, 0x0a, 0x29, 0x63,
	0x6f, 0x6d, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x69, 0x76, 0x5f, 0x73, 0x76, 0x69, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2f, 0x69, 0x76, 0x5f, 0x73, 0x76, 0x69, 0x5f, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x76, 0x31, 0x3b, 0x69, 0x76, 0x5f, 0x73, 0x76, 0x69,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x76, 0x31, 0xaa, 0x02,
	0x26, 0x4b, 0x61, 0x69, 0x6b, 0x6f, 0x53, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x76, 0x53, 0x76, 0x69, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sdk_stream_iv_svi_parameters_v1_request_proto_rawDescOnce sync.Once
	file_sdk_stream_iv_svi_parameters_v1_request_proto_rawDescData = file_sdk_stream_iv_svi_parameters_v1_request_proto_rawDesc
)

func file_sdk_stream_iv_svi_parameters_v1_request_proto_rawDescGZIP() []byte {
	file_sdk_stream_iv_svi_parameters_v1_request_proto_rawDescOnce.Do(func() {
		file_sdk_stream_iv_svi_parameters_v1_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdk_stream_iv_svi_parameters_v1_request_proto_rawDescData)
	})
	return file_sdk_stream_iv_svi_parameters_v1_request_proto_rawDescData
}

var file_sdk_stream_iv_svi_parameters_v1_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sdk_stream_iv_svi_parameters_v1_request_proto_goTypes = []interface{}{
	(*StreamIvSviParametersRequestV1)(nil), // 0: kaikosdk.StreamIvSviParametersRequestV1
	(*core.Assets)(nil),                    // 1: kaikosdk.Assets
}
var file_sdk_stream_iv_svi_parameters_v1_request_proto_depIdxs = []int32{
	1, // 0: kaikosdk.StreamIvSviParametersRequestV1.assets:type_name -> kaikosdk.Assets
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sdk_stream_iv_svi_parameters_v1_request_proto_init() }
func file_sdk_stream_iv_svi_parameters_v1_request_proto_init() {
	if File_sdk_stream_iv_svi_parameters_v1_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sdk_stream_iv_svi_parameters_v1_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamIvSviParametersRequestV1); i {
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
			RawDescriptor: file_sdk_stream_iv_svi_parameters_v1_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sdk_stream_iv_svi_parameters_v1_request_proto_goTypes,
		DependencyIndexes: file_sdk_stream_iv_svi_parameters_v1_request_proto_depIdxs,
		MessageInfos:      file_sdk_stream_iv_svi_parameters_v1_request_proto_msgTypes,
	}.Build()
	File_sdk_stream_iv_svi_parameters_v1_request_proto = out.File
	file_sdk_stream_iv_svi_parameters_v1_request_proto_rawDesc = nil
	file_sdk_stream_iv_svi_parameters_v1_request_proto_goTypes = nil
	file_sdk_stream_iv_svi_parameters_v1_request_proto_depIdxs = nil
}
