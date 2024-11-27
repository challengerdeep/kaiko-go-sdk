// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: sdk/stream/aggregated_state_price_v1/response.proto

package aggregated_state_price_v1

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

// StreamAggregatedStatePriceResponseV1
type StreamAggregatedStatePriceResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// datetime is the timestamp of the datapoint
	Datetime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=datetime,proto3" json:"datetime,omitempty"`
	// base is the base currency of the pair
	Base string `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	// aggregated_price_usd is the aggregated price in USD
	AggregatedPriceUsd string `protobuf:"bytes,3,opt,name=aggregated_price_usd,json=aggregatedPriceUsd,proto3" json:"aggregated_price_usd,omitempty"`
	// aggregated_price_lst is the aggregated price in LST quote token
	AggregatedPriceLst string `protobuf:"bytes,4,opt,name=aggregated_price_lst,json=aggregatedPriceLst,proto3" json:"aggregated_price_lst,omitempty"`
	// ts_event is the timestamp of the event
	TsEvent *timestamp.Timestamp `protobuf:"bytes,5,opt,name=ts_event,json=tsEvent,proto3" json:"ts_event,omitempty"`
}

func (x *StreamAggregatedStatePriceResponseV1) Reset() {
	*x = StreamAggregatedStatePriceResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_stream_aggregated_state_price_v1_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamAggregatedStatePriceResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamAggregatedStatePriceResponseV1) ProtoMessage() {}

func (x *StreamAggregatedStatePriceResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_stream_aggregated_state_price_v1_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamAggregatedStatePriceResponseV1.ProtoReflect.Descriptor instead.
func (*StreamAggregatedStatePriceResponseV1) Descriptor() ([]byte, []int) {
	return file_sdk_stream_aggregated_state_price_v1_response_proto_rawDescGZIP(), []int{0}
}

func (x *StreamAggregatedStatePriceResponseV1) GetDatetime() *timestamp.Timestamp {
	if x != nil {
		return x.Datetime
	}
	return nil
}

func (x *StreamAggregatedStatePriceResponseV1) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *StreamAggregatedStatePriceResponseV1) GetAggregatedPriceUsd() string {
	if x != nil {
		return x.AggregatedPriceUsd
	}
	return ""
}

func (x *StreamAggregatedStatePriceResponseV1) GetAggregatedPriceLst() string {
	if x != nil {
		return x.AggregatedPriceLst
	}
	return ""
}

func (x *StreamAggregatedStatePriceResponseV1) GetTsEvent() *timestamp.Timestamp {
	if x != nil {
		return x.TsEvent
	}
	return nil
}

var File_sdk_stream_aggregated_state_price_v1_response_proto protoreflect.FileDescriptor

var file_sdk_stream_aggregated_state_price_v1_response_proto_rawDesc = []byte{
	0x0a, 0x33, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x5f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8d, 0x02, 0x0a, 0x24, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x12, 0x36, 0x0a, 0x08, 0x64, 0x61, 0x74,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x55, 0x73, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x73, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x74, 0x73, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x74, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x42, 0xb9, 0x01, 0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x5f, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6b, 0x61, 0x69, 0x6b,
	0x6f, 0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x31, 0x3b, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x5f, 0x76, 0x31, 0xaa, 0x02, 0x26, 0x4b, 0x61, 0x69, 0x6b, 0x6f, 0x53, 0x64, 0x6b, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sdk_stream_aggregated_state_price_v1_response_proto_rawDescOnce sync.Once
	file_sdk_stream_aggregated_state_price_v1_response_proto_rawDescData = file_sdk_stream_aggregated_state_price_v1_response_proto_rawDesc
)

func file_sdk_stream_aggregated_state_price_v1_response_proto_rawDescGZIP() []byte {
	file_sdk_stream_aggregated_state_price_v1_response_proto_rawDescOnce.Do(func() {
		file_sdk_stream_aggregated_state_price_v1_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdk_stream_aggregated_state_price_v1_response_proto_rawDescData)
	})
	return file_sdk_stream_aggregated_state_price_v1_response_proto_rawDescData
}

var file_sdk_stream_aggregated_state_price_v1_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sdk_stream_aggregated_state_price_v1_response_proto_goTypes = []interface{}{
	(*StreamAggregatedStatePriceResponseV1)(nil), // 0: kaikosdk.StreamAggregatedStatePriceResponseV1
	(*timestamp.Timestamp)(nil),                  // 1: google.protobuf.Timestamp
}
var file_sdk_stream_aggregated_state_price_v1_response_proto_depIdxs = []int32{
	1, // 0: kaikosdk.StreamAggregatedStatePriceResponseV1.datetime:type_name -> google.protobuf.Timestamp
	1, // 1: kaikosdk.StreamAggregatedStatePriceResponseV1.ts_event:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sdk_stream_aggregated_state_price_v1_response_proto_init() }
func file_sdk_stream_aggregated_state_price_v1_response_proto_init() {
	if File_sdk_stream_aggregated_state_price_v1_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sdk_stream_aggregated_state_price_v1_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamAggregatedStatePriceResponseV1); i {
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
			RawDescriptor: file_sdk_stream_aggregated_state_price_v1_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sdk_stream_aggregated_state_price_v1_response_proto_goTypes,
		DependencyIndexes: file_sdk_stream_aggregated_state_price_v1_response_proto_depIdxs,
		MessageInfos:      file_sdk_stream_aggregated_state_price_v1_response_proto_msgTypes,
	}.Build()
	File_sdk_stream_aggregated_state_price_v1_response_proto = out.File
	file_sdk_stream_aggregated_state_price_v1_response_proto_rawDesc = nil
	file_sdk_stream_aggregated_state_price_v1_response_proto_goTypes = nil
	file_sdk_stream_aggregated_state_price_v1_response_proto_depIdxs = nil
}
