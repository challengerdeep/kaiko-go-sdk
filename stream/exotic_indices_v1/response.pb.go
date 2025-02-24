// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: sdk/stream/exotic_indices_v1/response.proto

package exotic_indices_v1

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	core "github.com/kaikodata/kaiko-go-sdk/core"
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

// Position is the position of the indice
type StreamExoticIndicesPosition int32

const (
	// Position unknown.
	StreamExoticIndicesPosition_SEIP_POSITION_UNKNOWN StreamExoticIndicesPosition = 0
	// Long position.
	StreamExoticIndicesPosition_SEIP_LONG StreamExoticIndicesPosition = 1
	// Short position.
	StreamExoticIndicesPosition_SEIP_SHORT StreamExoticIndicesPosition = 2
)

// Enum value maps for StreamExoticIndicesPosition.
var (
	StreamExoticIndicesPosition_name = map[int32]string{
		0: "SEIP_POSITION_UNKNOWN",
		1: "SEIP_LONG",
		2: "SEIP_SHORT",
	}
	StreamExoticIndicesPosition_value = map[string]int32{
		"SEIP_POSITION_UNKNOWN": 0,
		"SEIP_LONG":             1,
		"SEIP_SHORT":            2,
	}
)

func (x StreamExoticIndicesPosition) Enum() *StreamExoticIndicesPosition {
	p := new(StreamExoticIndicesPosition)
	*p = x
	return p
}

func (x StreamExoticIndicesPosition) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StreamExoticIndicesPosition) Descriptor() protoreflect.EnumDescriptor {
	return file_sdk_stream_exotic_indices_v1_response_proto_enumTypes[0].Descriptor()
}

func (StreamExoticIndicesPosition) Type() protoreflect.EnumType {
	return &file_sdk_stream_exotic_indices_v1_response_proto_enumTypes[0]
}

func (x StreamExoticIndicesPosition) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StreamExoticIndicesPosition.Descriptor instead.
func (StreamExoticIndicesPosition) EnumDescriptor() ([]byte, []int) {
	return file_sdk_stream_exotic_indices_v1_response_proto_rawDescGZIP(), []int{0}
}

// StreamExoticIndicesServiceResponseComposition is the composition used to compute the index.
type StreamExoticIndicesServiceResponseComposition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Underlying rate of the indices.
	UnderlyingInstrument string `protobuf:"bytes,1,opt,name=underlying_instrument,json=underlyingInstrument,proto3" json:"underlying_instrument,omitempty"`
	// Base.
	Base string `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	// Quote.
	Quote string `protobuf:"bytes,3,opt,name=quote,proto3" json:"quote,omitempty"`
	// Currency conversion.
	CurrencyConversion string `protobuf:"bytes,6,opt,name=currency_conversion,json=currencyConversion,proto3" json:"currency_conversion,omitempty"`
	// Timestamp (tick) of underlying rate ts.
	TsEvent *timestamp.Timestamp `protobuf:"bytes,7,opt,name=ts_event,json=tsEvent,proto3" json:"ts_event,omitempty"`
}

func (x *StreamExoticIndicesServiceResponseComposition) Reset() {
	*x = StreamExoticIndicesServiceResponseComposition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_stream_exotic_indices_v1_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamExoticIndicesServiceResponseComposition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamExoticIndicesServiceResponseComposition) ProtoMessage() {}

func (x *StreamExoticIndicesServiceResponseComposition) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_stream_exotic_indices_v1_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamExoticIndicesServiceResponseComposition.ProtoReflect.Descriptor instead.
func (*StreamExoticIndicesServiceResponseComposition) Descriptor() ([]byte, []int) {
	return file_sdk_stream_exotic_indices_v1_response_proto_rawDescGZIP(), []int{0}
}

func (x *StreamExoticIndicesServiceResponseComposition) GetUnderlyingInstrument() string {
	if x != nil {
		return x.UnderlyingInstrument
	}
	return ""
}

func (x *StreamExoticIndicesServiceResponseComposition) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *StreamExoticIndicesServiceResponseComposition) GetQuote() string {
	if x != nil {
		return x.Quote
	}
	return ""
}

func (x *StreamExoticIndicesServiceResponseComposition) GetCurrencyConversion() string {
	if x != nil {
		return x.CurrencyConversion
	}
	return ""
}

func (x *StreamExoticIndicesServiceResponseComposition) GetTsEvent() *timestamp.Timestamp {
	if x != nil {
		return x.TsEvent
	}
	return nil
}

// StreamExoticIndicesServiceResponsePair is the pair information for the rates used
type StreamExoticIndicesServiceResponsePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Underlying instrument of the indices.
	UnderlyingInstrument string `protobuf:"bytes,1,opt,name=underlying_instrument,json=underlyingInstrument,proto3" json:"underlying_instrument,omitempty"`
	// Underlying price of the instrument.
	UnderlyingPrice *wrappers.DoubleValue `protobuf:"bytes,2,opt,name=underlying_price,json=underlyingPrice,proto3" json:"underlying_price,omitempty"`
	// Weighting factor of the instrument.
	WeightingFactor float64 `protobuf:"fixed64,3,opt,name=weighting_factor,json=weightingFactor,proto3" json:"weighting_factor,omitempty"`
	// Capping factor of the instrument.
	CappingFactor float64 `protobuf:"fixed64,4,opt,name=capping_factor,json=cappingFactor,proto3" json:"capping_factor,omitempty"`
	// Currency conversion factor of the instrument.
	CurrencyConversionFactor float64 `protobuf:"fixed64,5,opt,name=currency_conversion_factor,json=currencyConversionFactor,proto3" json:"currency_conversion_factor,omitempty"`
	// Position of the instrument.
	Position StreamExoticIndicesPosition `protobuf:"varint,6,opt,name=position,proto3,enum=kaikosdk.StreamExoticIndicesPosition" json:"position,omitempty"`
}

func (x *StreamExoticIndicesServiceResponsePair) Reset() {
	*x = StreamExoticIndicesServiceResponsePair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_stream_exotic_indices_v1_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamExoticIndicesServiceResponsePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamExoticIndicesServiceResponsePair) ProtoMessage() {}

func (x *StreamExoticIndicesServiceResponsePair) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_stream_exotic_indices_v1_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamExoticIndicesServiceResponsePair.ProtoReflect.Descriptor instead.
func (*StreamExoticIndicesServiceResponsePair) Descriptor() ([]byte, []int) {
	return file_sdk_stream_exotic_indices_v1_response_proto_rawDescGZIP(), []int{1}
}

func (x *StreamExoticIndicesServiceResponsePair) GetUnderlyingInstrument() string {
	if x != nil {
		return x.UnderlyingInstrument
	}
	return ""
}

func (x *StreamExoticIndicesServiceResponsePair) GetUnderlyingPrice() *wrappers.DoubleValue {
	if x != nil {
		return x.UnderlyingPrice
	}
	return nil
}

func (x *StreamExoticIndicesServiceResponsePair) GetWeightingFactor() float64 {
	if x != nil {
		return x.WeightingFactor
	}
	return 0
}

func (x *StreamExoticIndicesServiceResponsePair) GetCappingFactor() float64 {
	if x != nil {
		return x.CappingFactor
	}
	return 0
}

func (x *StreamExoticIndicesServiceResponsePair) GetCurrencyConversionFactor() float64 {
	if x != nil {
		return x.CurrencyConversionFactor
	}
	return 0
}

func (x *StreamExoticIndicesServiceResponsePair) GetPosition() StreamExoticIndicesPosition {
	if x != nil {
		return x.Position
	}
	return StreamExoticIndicesPosition_SEIP_POSITION_UNKNOWN
}

// StreamExoticIndicesServiceResponsePrices is the prices informations on the pair used
type StreamExoticIndicesServiceResponsePrices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Index value.
	IndexValue float64 `protobuf:"fixed64,1,opt,name=index_value,json=indexValue,proto3" json:"index_value,omitempty"`
	// Divisor.
	Divisor float64 `protobuf:"fixed64,2,opt,name=divisor,proto3" json:"divisor,omitempty"`
	// StreamExoticIndicesServiceResponsePair is the pair information.
	Pairs []*StreamExoticIndicesServiceResponsePair `protobuf:"bytes,3,rep,name=pairs,proto3" json:"pairs,omitempty"`
}

func (x *StreamExoticIndicesServiceResponsePrices) Reset() {
	*x = StreamExoticIndicesServiceResponsePrices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_stream_exotic_indices_v1_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamExoticIndicesServiceResponsePrices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamExoticIndicesServiceResponsePrices) ProtoMessage() {}

func (x *StreamExoticIndicesServiceResponsePrices) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_stream_exotic_indices_v1_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamExoticIndicesServiceResponsePrices.ProtoReflect.Descriptor instead.
func (*StreamExoticIndicesServiceResponsePrices) Descriptor() ([]byte, []int) {
	return file_sdk_stream_exotic_indices_v1_response_proto_rawDescGZIP(), []int{2}
}

func (x *StreamExoticIndicesServiceResponsePrices) GetIndexValue() float64 {
	if x != nil {
		return x.IndexValue
	}
	return 0
}

func (x *StreamExoticIndicesServiceResponsePrices) GetDivisor() float64 {
	if x != nil {
		return x.Divisor
	}
	return 0
}

func (x *StreamExoticIndicesServiceResponsePrices) GetPairs() []*StreamExoticIndicesServiceResponsePair {
	if x != nil {
		return x.Pairs
	}
	return nil
}

// StreamIndexServiceResponseV1 is the response for the exotic indices.
type StreamExoticIndicesServiceResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Event type.
	Commodity index_v1.StreamIndexCommodity `protobuf:"varint,1,opt,name=commodity,proto3,enum=kaikosdk.StreamIndexCommodity" json:"commodity,omitempty"`
	// Index code.
	IndexCode string `protobuf:"bytes,2,opt,name=index_code,json=indexCode,proto3" json:"index_code,omitempty"`
	// Data interval.
	Interval *core.DataInterval `protobuf:"bytes,3,opt,name=interval,proto3" json:"interval,omitempty"`
	// Quote.
	MainQuote string `protobuf:"bytes,4,opt,name=main_quote,json=mainQuote,proto3" json:"main_quote,omitempty"`
	// List of indices compositions.
	Compositions []*StreamExoticIndicesServiceResponseComposition `protobuf:"bytes,5,rep,name=compositions,proto3" json:"compositions,omitempty"`
	// Price of the indice.
	Price *StreamExoticIndicesServiceResponsePrices `protobuf:"bytes,6,opt,name=price,proto3" json:"price,omitempty"`
	// Event generation timestamp (event created by Kaiko).
	TsEvent *timestamp.Timestamp `protobuf:"bytes,7,opt,name=ts_event,json=tsEvent,proto3" json:"ts_event,omitempty"`
	// Timestamp of computation.
	TsCompute *timestamp.Timestamp `protobuf:"bytes,8,opt,name=ts_compute,json=tsCompute,proto3" json:"ts_compute,omitempty"`
}

func (x *StreamExoticIndicesServiceResponseV1) Reset() {
	*x = StreamExoticIndicesServiceResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_stream_exotic_indices_v1_response_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamExoticIndicesServiceResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamExoticIndicesServiceResponseV1) ProtoMessage() {}

func (x *StreamExoticIndicesServiceResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_stream_exotic_indices_v1_response_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamExoticIndicesServiceResponseV1.ProtoReflect.Descriptor instead.
func (*StreamExoticIndicesServiceResponseV1) Descriptor() ([]byte, []int) {
	return file_sdk_stream_exotic_indices_v1_response_proto_rawDescGZIP(), []int{3}
}

func (x *StreamExoticIndicesServiceResponseV1) GetCommodity() index_v1.StreamIndexCommodity {
	if x != nil {
		return x.Commodity
	}
	return index_v1.StreamIndexCommodity(0)
}

func (x *StreamExoticIndicesServiceResponseV1) GetIndexCode() string {
	if x != nil {
		return x.IndexCode
	}
	return ""
}

func (x *StreamExoticIndicesServiceResponseV1) GetInterval() *core.DataInterval {
	if x != nil {
		return x.Interval
	}
	return nil
}

func (x *StreamExoticIndicesServiceResponseV1) GetMainQuote() string {
	if x != nil {
		return x.MainQuote
	}
	return ""
}

func (x *StreamExoticIndicesServiceResponseV1) GetCompositions() []*StreamExoticIndicesServiceResponseComposition {
	if x != nil {
		return x.Compositions
	}
	return nil
}

func (x *StreamExoticIndicesServiceResponseV1) GetPrice() *StreamExoticIndicesServiceResponsePrices {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *StreamExoticIndicesServiceResponseV1) GetTsEvent() *timestamp.Timestamp {
	if x != nil {
		return x.TsEvent
	}
	return nil
}

func (x *StreamExoticIndicesServiceResponseV1) GetTsCompute() *timestamp.Timestamp {
	if x != nil {
		return x.TsCompute
	}
	return nil
}

var File_sdk_stream_exotic_indices_v1_response_proto protoreflect.FileDescriptor

var file_sdk_stream_exotic_indices_v1_response_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x65, 0x78, 0x6f,
	0x74, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6b,
	0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x73, 0x64, 0x6b, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x01, 0x0a, 0x2d,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x6f, 0x74, 0x69, 0x63, 0x49, 0x6e, 0x64, 0x69,
	0x63, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a,
	0x15, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x75, 0x6e,
	0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x13,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a,
	0x08, 0x74, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x74, 0x73, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x22, 0xf9, 0x02, 0x0a, 0x26, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45,
	0x78, 0x6f, 0x74, 0x69, 0x63, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x69, 0x72, 0x12,
	0x33, 0x0a, 0x15, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x10, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69,
	0x6e, 0x67, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x75, 0x6e,
	0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a,
	0x10, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x69,
	0x6e, 0x67, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0d, 0x63, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x3c, 0x0a, 0x1a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x18, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x41, 0x0a,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x25, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x45, 0x78, 0x6f, 0x74, 0x69, 0x63, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xad, 0x01, 0x0a, 0x28, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x6f, 0x74, 0x69,
	0x63, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0a, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x69, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x07, 0x64, 0x69, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x12, 0x46, 0x0a, 0x05, 0x70, 0x61, 0x69, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73,
	0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x6f, 0x74, 0x69, 0x63, 0x49,
	0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05, 0x70, 0x61, 0x69, 0x72, 0x73,
	0x22, 0xef, 0x03, 0x0a, 0x24, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x6f, 0x74, 0x69,
	0x63, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x12, 0x3c, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6b,
	0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f,
	0x73, 0x64, 0x6b, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x61, 0x69, 0x6e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x5b, 0x0a, 0x0c, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x37, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x45, 0x78, 0x6f, 0x74, 0x69, 0x63, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x48, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x6f, 0x74, 0x69, 0x63, 0x49, 0x6e, 0x64,
	0x69, 0x63, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x35, 0x0a, 0x08, 0x74, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x74, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x74, 0x73, 0x5f, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x2a, 0x57, 0x0a, 0x1b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x6f, 0x74,
	0x69, 0x63, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x45, 0x49, 0x50, 0x5f, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x53, 0x45, 0x49, 0x50, 0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x53,
	0x45, 0x49, 0x50, 0x5f, 0x53, 0x48, 0x4f, 0x52, 0x54, 0x10, 0x02, 0x42, 0x9a, 0x01, 0x0a, 0x26,
	0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2e, 0x65, 0x78, 0x6f, 0x74, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x64, 0x69,
	0x63, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6b,
	0x61, 0x69, 0x6b, 0x6f, 0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2f, 0x65, 0x78, 0x6f, 0x74, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65,
	0x73, 0x5f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x6f, 0x74, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x64, 0x69,
	0x63, 0x65, 0x73, 0x5f, 0x76, 0x31, 0xaa, 0x02, 0x1f, 0x4b, 0x61, 0x69, 0x6b, 0x6f, 0x53, 0x64,
	0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x45, 0x78, 0x6f, 0x74, 0x69, 0x63, 0x49,
	0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sdk_stream_exotic_indices_v1_response_proto_rawDescOnce sync.Once
	file_sdk_stream_exotic_indices_v1_response_proto_rawDescData = file_sdk_stream_exotic_indices_v1_response_proto_rawDesc
)

func file_sdk_stream_exotic_indices_v1_response_proto_rawDescGZIP() []byte {
	file_sdk_stream_exotic_indices_v1_response_proto_rawDescOnce.Do(func() {
		file_sdk_stream_exotic_indices_v1_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdk_stream_exotic_indices_v1_response_proto_rawDescData)
	})
	return file_sdk_stream_exotic_indices_v1_response_proto_rawDescData
}

var file_sdk_stream_exotic_indices_v1_response_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sdk_stream_exotic_indices_v1_response_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sdk_stream_exotic_indices_v1_response_proto_goTypes = []interface{}{
	(StreamExoticIndicesPosition)(0),                      // 0: kaikosdk.StreamExoticIndicesPosition
	(*StreamExoticIndicesServiceResponseComposition)(nil), // 1: kaikosdk.StreamExoticIndicesServiceResponseComposition
	(*StreamExoticIndicesServiceResponsePair)(nil),        // 2: kaikosdk.StreamExoticIndicesServiceResponsePair
	(*StreamExoticIndicesServiceResponsePrices)(nil),      // 3: kaikosdk.StreamExoticIndicesServiceResponsePrices
	(*StreamExoticIndicesServiceResponseV1)(nil),          // 4: kaikosdk.StreamExoticIndicesServiceResponseV1
	(*timestamp.Timestamp)(nil),                           // 5: google.protobuf.Timestamp
	(*wrappers.DoubleValue)(nil),                          // 6: google.protobuf.DoubleValue
	(index_v1.StreamIndexCommodity)(0),                    // 7: kaikosdk.StreamIndexCommodity
	(*core.DataInterval)(nil),                             // 8: kaikosdk.DataInterval
}
var file_sdk_stream_exotic_indices_v1_response_proto_depIdxs = []int32{
	5,  // 0: kaikosdk.StreamExoticIndicesServiceResponseComposition.ts_event:type_name -> google.protobuf.Timestamp
	6,  // 1: kaikosdk.StreamExoticIndicesServiceResponsePair.underlying_price:type_name -> google.protobuf.DoubleValue
	0,  // 2: kaikosdk.StreamExoticIndicesServiceResponsePair.position:type_name -> kaikosdk.StreamExoticIndicesPosition
	2,  // 3: kaikosdk.StreamExoticIndicesServiceResponsePrices.pairs:type_name -> kaikosdk.StreamExoticIndicesServiceResponsePair
	7,  // 4: kaikosdk.StreamExoticIndicesServiceResponseV1.commodity:type_name -> kaikosdk.StreamIndexCommodity
	8,  // 5: kaikosdk.StreamExoticIndicesServiceResponseV1.interval:type_name -> kaikosdk.DataInterval
	1,  // 6: kaikosdk.StreamExoticIndicesServiceResponseV1.compositions:type_name -> kaikosdk.StreamExoticIndicesServiceResponseComposition
	3,  // 7: kaikosdk.StreamExoticIndicesServiceResponseV1.price:type_name -> kaikosdk.StreamExoticIndicesServiceResponsePrices
	5,  // 8: kaikosdk.StreamExoticIndicesServiceResponseV1.ts_event:type_name -> google.protobuf.Timestamp
	5,  // 9: kaikosdk.StreamExoticIndicesServiceResponseV1.ts_compute:type_name -> google.protobuf.Timestamp
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_sdk_stream_exotic_indices_v1_response_proto_init() }
func file_sdk_stream_exotic_indices_v1_response_proto_init() {
	if File_sdk_stream_exotic_indices_v1_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sdk_stream_exotic_indices_v1_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamExoticIndicesServiceResponseComposition); i {
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
		file_sdk_stream_exotic_indices_v1_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamExoticIndicesServiceResponsePair); i {
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
		file_sdk_stream_exotic_indices_v1_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamExoticIndicesServiceResponsePrices); i {
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
		file_sdk_stream_exotic_indices_v1_response_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamExoticIndicesServiceResponseV1); i {
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
			RawDescriptor: file_sdk_stream_exotic_indices_v1_response_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sdk_stream_exotic_indices_v1_response_proto_goTypes,
		DependencyIndexes: file_sdk_stream_exotic_indices_v1_response_proto_depIdxs,
		EnumInfos:         file_sdk_stream_exotic_indices_v1_response_proto_enumTypes,
		MessageInfos:      file_sdk_stream_exotic_indices_v1_response_proto_msgTypes,
	}.Build()
	File_sdk_stream_exotic_indices_v1_response_proto = out.File
	file_sdk_stream_exotic_indices_v1_response_proto_rawDesc = nil
	file_sdk_stream_exotic_indices_v1_response_proto_goTypes = nil
	file_sdk_stream_exotic_indices_v1_response_proto_depIdxs = nil
}
