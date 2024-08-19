// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: sdk/stream/derivatives_instrument_metrics_v1/response.proto

package derivatives_instrument_metrics_v1

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// StreamDerivativesInstrumentMetricsResponseV1 derivatives v1 response.
type StreamDerivativesInstrumentMetricsResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value of the commodity.
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// Legacy symbol.
	LegacySymbol string `protobuf:"bytes,2,opt,name=legacy_symbol,json=legacySymbol,proto3" json:"legacy_symbol,omitempty"`
	// Exchange.
	Exchange string `protobuf:"bytes,3,opt,name=exchange,proto3" json:"exchange,omitempty"`
	// Commodity.
	// See https://docs.kaiko.com/v/kaiko-rest-api/market-data/derivatives/derivatives-metrics/risk#fields
	// and https://docs.kaiko.com/v/kaiko-rest-api/market-data/derivatives/derivatives-metrics/price#fields.
	Commodity string `protobuf:"bytes,4,opt,name=commodity,proto3" json:"commodity,omitempty"`
	// Kind of commodity (price of risk).
	CommodityKind StreamDerivativesInstrumentCommodityKindV1 `protobuf:"varint,5,opt,name=commodity_kind,json=commodityKind,proto3,enum=kaikosdk.StreamDerivativesInstrumentCommodityKindV1" json:"commodity_kind,omitempty"`
	// Timestamp of collection (event entered Kaiko's infrastructure), before nomalization.
	TsCollection *timestamp.Timestamp `protobuf:"bytes,6,opt,name=ts_collection,json=tsCollection,proto3" json:"ts_collection,omitempty"`
	// Event generation timestamp (event created by Kaiko), after normalization.
	TsEvent *timestamp.Timestamp `protobuf:"bytes,7,opt,name=ts_event,json=tsEvent,proto3" json:"ts_event,omitempty"`
}

func (x *StreamDerivativesInstrumentMetricsResponseV1) Reset() {
	*x = StreamDerivativesInstrumentMetricsResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamDerivativesInstrumentMetricsResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamDerivativesInstrumentMetricsResponseV1) ProtoMessage() {}

func (x *StreamDerivativesInstrumentMetricsResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamDerivativesInstrumentMetricsResponseV1.ProtoReflect.Descriptor instead.
func (*StreamDerivativesInstrumentMetricsResponseV1) Descriptor() ([]byte, []int) {
	return file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_rawDescGZIP(), []int{0}
}

func (x *StreamDerivativesInstrumentMetricsResponseV1) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *StreamDerivativesInstrumentMetricsResponseV1) GetLegacySymbol() string {
	if x != nil {
		return x.LegacySymbol
	}
	return ""
}

func (x *StreamDerivativesInstrumentMetricsResponseV1) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *StreamDerivativesInstrumentMetricsResponseV1) GetCommodity() string {
	if x != nil {
		return x.Commodity
	}
	return ""
}

func (x *StreamDerivativesInstrumentMetricsResponseV1) GetCommodityKind() StreamDerivativesInstrumentCommodityKindV1 {
	if x != nil {
		return x.CommodityKind
	}
	return StreamDerivativesInstrumentCommodityKindV1_SDICK_UNKNOWN
}

func (x *StreamDerivativesInstrumentMetricsResponseV1) GetTsCollection() *timestamp.Timestamp {
	if x != nil {
		return x.TsCollection
	}
	return nil
}

func (x *StreamDerivativesInstrumentMetricsResponseV1) GetTsEvent() *timestamp.Timestamp {
	if x != nil {
		return x.TsEvent
	}
	return nil
}

var File_sdk_stream_derivatives_instrument_metrics_v1_response_proto protoreflect.FileDescriptor

var file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x64, 0x65, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x76, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6b,
	0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2f, 0x64, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x76, 0x65, 0x73,
	0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x5f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79,
	0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x02, 0x0a, 0x2c,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x73, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x65, 0x67, 0x61, 0x63,
	0x79, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74,
	0x79, 0x12, 0x5b, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x6b,
	0x69, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x6b, 0x61, 0x69, 0x6b,
	0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x65, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x73, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x4b, 0x69, 0x6e, 0x64, 0x56, 0x31, 0x52,
	0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x3f,
	0x0a, 0x0d, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0c, 0x74, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x35, 0x0a, 0x08, 0x74, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x74,
	0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0xd9, 0x01, 0x0a, 0x36, 0x63, 0x6f, 0x6d, 0x2e, 0x6b,
	0x61, 0x69, 0x6b, 0x6f, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x64, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x76, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x6c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x2d,
	0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x64, 0x65,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x76, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x76, 0x31, 0x3b,
	0x64, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x76, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x76,
	0x31, 0xaa, 0x02, 0x2e, 0x4b, 0x61, 0x69, 0x6b, 0x6f, 0x53, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x44, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x76, 0x65, 0x73, 0x49,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_rawDescOnce sync.Once
	file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_rawDescData = file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_rawDesc
)

func file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_rawDescGZIP() []byte {
	file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_rawDescOnce.Do(func() {
		file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_rawDescData)
	})
	return file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_rawDescData
}

var file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_goTypes = []interface{}{
	(*StreamDerivativesInstrumentMetricsResponseV1)(nil), // 0: kaikosdk.StreamDerivativesInstrumentMetricsResponseV1
	(StreamDerivativesInstrumentCommodityKindV1)(0),      // 1: kaikosdk.StreamDerivativesInstrumentCommodityKindV1
	(*timestamp.Timestamp)(nil),                          // 2: google.protobuf.Timestamp
}
var file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_depIdxs = []int32{
	1, // 0: kaikosdk.StreamDerivativesInstrumentMetricsResponseV1.commodity_kind:type_name -> kaikosdk.StreamDerivativesInstrumentCommodityKindV1
	2, // 1: kaikosdk.StreamDerivativesInstrumentMetricsResponseV1.ts_collection:type_name -> google.protobuf.Timestamp
	2, // 2: kaikosdk.StreamDerivativesInstrumentMetricsResponseV1.ts_event:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_init() }
func file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_init() {
	if File_sdk_stream_derivatives_instrument_metrics_v1_response_proto != nil {
		return
	}
	file_sdk_stream_derivatives_instrument_metrics_v1_commodity_kind_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamDerivativesInstrumentMetricsResponseV1); i {
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
			RawDescriptor: file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_goTypes,
		DependencyIndexes: file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_depIdxs,
		MessageInfos:      file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_msgTypes,
	}.Build()
	File_sdk_stream_derivatives_instrument_metrics_v1_response_proto = out.File
	file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_rawDesc = nil
	file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_goTypes = nil
	file_sdk_stream_derivatives_instrument_metrics_v1_response_proto_depIdxs = nil
}
