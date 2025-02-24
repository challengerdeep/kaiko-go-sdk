// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: sdk/stream/aggregates_direct_exchange_rate_v2/request.proto

package aggregates_direct_exchange_rate_v2

import (
	duration "github.com/golang/protobuf/ptypes/duration"
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

// StreamAggregatesDirectExchangeRateV1RequestV1
type StreamAggregatesDirectExchangeRateV2RequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// base/quote assets.
	Assets *core.Assets `protobuf:"bytes,1,opt,name=assets,proto3" json:"assets,omitempty"`
	// Sliding Window duration (duration string, ex: 1m).
	Window *duration.Duration `protobuf:"bytes,2,opt,name=window,proto3" json:"window,omitempty"`
	// list of exchange code to include.
	IncludeExchanges []string `protobuf:"bytes,3,rep,name=include_exchanges,json=includeExchanges,proto3" json:"include_exchanges,omitempty"`
	// list of exchange code to exclude.
	ExcludeExchanges []string `protobuf:"bytes,4,rep,name=exclude_exchanges,json=excludeExchanges,proto3" json:"exclude_exchanges,omitempty"`
	// Update frequency (duration string, ex: 1m).
	UpdateFrequency *duration.Duration `protobuf:"bytes,5,opt,name=update_frequency,json=updateFrequency,proto3" json:"update_frequency,omitempty"`
}

func (x *StreamAggregatesDirectExchangeRateV2RequestV1) Reset() {
	*x = StreamAggregatesDirectExchangeRateV2RequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamAggregatesDirectExchangeRateV2RequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamAggregatesDirectExchangeRateV2RequestV1) ProtoMessage() {}

func (x *StreamAggregatesDirectExchangeRateV2RequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamAggregatesDirectExchangeRateV2RequestV1.ProtoReflect.Descriptor instead.
func (*StreamAggregatesDirectExchangeRateV2RequestV1) Descriptor() ([]byte, []int) {
	return file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_rawDescGZIP(), []int{0}
}

func (x *StreamAggregatesDirectExchangeRateV2RequestV1) GetAssets() *core.Assets {
	if x != nil {
		return x.Assets
	}
	return nil
}

func (x *StreamAggregatesDirectExchangeRateV2RequestV1) GetWindow() *duration.Duration {
	if x != nil {
		return x.Window
	}
	return nil
}

func (x *StreamAggregatesDirectExchangeRateV2RequestV1) GetIncludeExchanges() []string {
	if x != nil {
		return x.IncludeExchanges
	}
	return nil
}

func (x *StreamAggregatesDirectExchangeRateV2RequestV1) GetExcludeExchanges() []string {
	if x != nil {
		return x.ExcludeExchanges
	}
	return nil
}

func (x *StreamAggregatesDirectExchangeRateV2RequestV1) GetUpdateFrequency() *duration.Duration {
	if x != nil {
		return x.UpdateFrequency
	}
	return nil
}

var File_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto protoreflect.FileDescriptor

var file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x32, 0x2f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6b,
	0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x73, 0x64, 0x6b, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac,
	0x02, 0x0a, 0x2d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x73, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x61, 0x74, 0x65, 0x56, 0x32, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31,
	0x12, 0x28, 0x0a, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x73, 0x52, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x77, 0x69,
	0x6e, 0x64, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x2b, 0x0a,
	0x11, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x10, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x42, 0xdc, 0x01,
	0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x73, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x32, 0x50, 0x01, 0x5a, 0x6e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73,
	0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x32, 0x3b, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x73, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x32, 0xaa, 0x02, 0x2e, 0x4b, 0x61,
	0x69, 0x6b, 0x6f, 0x53, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_rawDescOnce sync.Once
	file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_rawDescData = file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_rawDesc
)

func file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_rawDescGZIP() []byte {
	file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_rawDescOnce.Do(func() {
		file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_rawDescData)
	})
	return file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_rawDescData
}

var file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_goTypes = []interface{}{
	(*StreamAggregatesDirectExchangeRateV2RequestV1)(nil), // 0: kaikosdk.StreamAggregatesDirectExchangeRateV2RequestV1
	(*core.Assets)(nil),       // 1: kaikosdk.Assets
	(*duration.Duration)(nil), // 2: google.protobuf.Duration
}
var file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_depIdxs = []int32{
	1, // 0: kaikosdk.StreamAggregatesDirectExchangeRateV2RequestV1.assets:type_name -> kaikosdk.Assets
	2, // 1: kaikosdk.StreamAggregatesDirectExchangeRateV2RequestV1.window:type_name -> google.protobuf.Duration
	2, // 2: kaikosdk.StreamAggregatesDirectExchangeRateV2RequestV1.update_frequency:type_name -> google.protobuf.Duration
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_init() }
func file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_init() {
	if File_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamAggregatesDirectExchangeRateV2RequestV1); i {
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
			RawDescriptor: file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_goTypes,
		DependencyIndexes: file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_depIdxs,
		MessageInfos:      file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_msgTypes,
	}.Build()
	File_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto = out.File
	file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_rawDesc = nil
	file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_goTypes = nil
	file_sdk_stream_aggregates_direct_exchange_rate_v2_request_proto_depIdxs = nil
}
