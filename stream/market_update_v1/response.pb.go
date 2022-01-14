// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: sdk/stream/market_update_v1/response.proto

package market_update_v1

import (
	core "github.com/challengerdeep/kaiko-go-sdk/core"
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

// StreamMarketUpdateType is event category for a market update.
type StreamMarketUpdateResponseV1_StreamMarketUpdateType int32

const (
	// Unknown type.
	StreamMarketUpdateResponseV1_UNKNOWN StreamMarketUpdateResponseV1_StreamMarketUpdateType = 0
	// TRADE_BUY type.
	StreamMarketUpdateResponseV1_TRADE_BUY StreamMarketUpdateResponseV1_StreamMarketUpdateType = 1
	// TRADE_SELL type.
	StreamMarketUpdateResponseV1_TRADE_SELL StreamMarketUpdateResponseV1_StreamMarketUpdateType = 2
	// BEST_ASK type.
	StreamMarketUpdateResponseV1_BEST_ASK StreamMarketUpdateResponseV1_StreamMarketUpdateType = 3
	// BEST_BID type.
	StreamMarketUpdateResponseV1_BEST_BID StreamMarketUpdateResponseV1_StreamMarketUpdateType = 4
	// UPDATED_ASK type.
	StreamMarketUpdateResponseV1_UPDATED_ASK StreamMarketUpdateResponseV1_StreamMarketUpdateType = 5
	// UPDATED_BID type.
	StreamMarketUpdateResponseV1_UPDATED_BID StreamMarketUpdateResponseV1_StreamMarketUpdateType = 6
	// SNAPSHOT type.
	StreamMarketUpdateResponseV1_SNAPSHOT StreamMarketUpdateResponseV1_StreamMarketUpdateType = 7
)

// Enum value maps for StreamMarketUpdateResponseV1_StreamMarketUpdateType.
var (
	StreamMarketUpdateResponseV1_StreamMarketUpdateType_name = map[int32]string{
		0: "UNKNOWN",
		1: "TRADE_BUY",
		2: "TRADE_SELL",
		3: "BEST_ASK",
		4: "BEST_BID",
		5: "UPDATED_ASK",
		6: "UPDATED_BID",
		7: "SNAPSHOT",
	}
	StreamMarketUpdateResponseV1_StreamMarketUpdateType_value = map[string]int32{
		"UNKNOWN":     0,
		"TRADE_BUY":   1,
		"TRADE_SELL":  2,
		"BEST_ASK":    3,
		"BEST_BID":    4,
		"UPDATED_ASK": 5,
		"UPDATED_BID": 6,
		"SNAPSHOT":    7,
	}
)

func (x StreamMarketUpdateResponseV1_StreamMarketUpdateType) Enum() *StreamMarketUpdateResponseV1_StreamMarketUpdateType {
	p := new(StreamMarketUpdateResponseV1_StreamMarketUpdateType)
	*p = x
	return p
}

func (x StreamMarketUpdateResponseV1_StreamMarketUpdateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StreamMarketUpdateResponseV1_StreamMarketUpdateType) Descriptor() protoreflect.EnumDescriptor {
	return file_sdk_stream_market_update_v1_response_proto_enumTypes[0].Descriptor()
}

func (StreamMarketUpdateResponseV1_StreamMarketUpdateType) Type() protoreflect.EnumType {
	return &file_sdk_stream_market_update_v1_response_proto_enumTypes[0]
}

func (x StreamMarketUpdateResponseV1_StreamMarketUpdateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StreamMarketUpdateResponseV1_StreamMarketUpdateType.Descriptor instead.
func (StreamMarketUpdateResponseV1_StreamMarketUpdateType) EnumDescriptor() ([]byte, []int) {
	return file_sdk_stream_market_update_v1_response_proto_rawDescGZIP(), []int{0, 0}
}

// StreamMarketUpdateResponseV1
type StreamMarketUpdateResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kind of commodity concerned by the market update.
	Commodity StreamMarketUpdateCommodity `protobuf:"varint,1,opt,name=commodity,proto3,enum=kaikosdk.StreamMarketUpdateCommodity" json:"commodity,omitempty"`
	// Amount / quantity of asset bought or sold, displayed in base currency.
	Amount float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	// Instrument class, internal Kaiko classification denoting whether an instrument is a spot, future, perpetual future, or option.
	Class string `protobuf:"bytes,3,opt,name=class,proto3" json:"class,omitempty"`
	// Instrument code (currency pair), for example btc-usd.
	Code string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	// Instrument exchange code, for example "cbse" (Coinbase).
	Exchange string `protobuf:"bytes,5,opt,name=exchange,proto3" json:"exchange,omitempty"`
	// Sequence ID for event. Sortable in lexicographic order.
	SequenceId string `protobuf:"bytes,6,opt,name=sequence_id,json=sequenceId,proto3" json:"sequence_id,omitempty"`
	// ID from exchange (trades only), empty string when not present.
	Id string `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
	// Price for quote currency.
	Price float64 `protobuf:"fixed64,8,opt,name=price,proto3" json:"price,omitempty"`
	// Timestamp of the event as provided by the exchange.
	TsExchange *core.TimestampValue `protobuf:"bytes,9,opt,name=ts_exchange,json=tsExchange,proto3" json:"ts_exchange,omitempty"`
	// Timestamp of collection (event entered Kaiko's infrastructure), before nomalization.
	TsCollection *core.TimestampValue `protobuf:"bytes,10,opt,name=ts_collection,json=tsCollection,proto3" json:"ts_collection,omitempty"`
	// Event generation timestamp (event created by Kaiko), after normalization.
	TsEvent *timestamp.Timestamp `protobuf:"bytes,11,opt,name=ts_event,json=tsEvent,proto3" json:"ts_event,omitempty"`
	// Event category for this update.
	UpdateType StreamMarketUpdateResponseV1_StreamMarketUpdateType `protobuf:"varint,12,opt,name=update_type,json=updateType,proto3,enum=kaikosdk.StreamMarketUpdateResponseV1_StreamMarketUpdateType" json:"update_type,omitempty"`
	// Snapshot for this update.
	Snapshot *StreamMarketUpdateResponseV1_Snapshot `protobuf:"bytes,13,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
}

func (x *StreamMarketUpdateResponseV1) Reset() {
	*x = StreamMarketUpdateResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_stream_market_update_v1_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamMarketUpdateResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamMarketUpdateResponseV1) ProtoMessage() {}

func (x *StreamMarketUpdateResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_stream_market_update_v1_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamMarketUpdateResponseV1.ProtoReflect.Descriptor instead.
func (*StreamMarketUpdateResponseV1) Descriptor() ([]byte, []int) {
	return file_sdk_stream_market_update_v1_response_proto_rawDescGZIP(), []int{0}
}

func (x *StreamMarketUpdateResponseV1) GetCommodity() StreamMarketUpdateCommodity {
	if x != nil {
		return x.Commodity
	}
	return StreamMarketUpdateCommodity_SMUC_UNKNOWN
}

func (x *StreamMarketUpdateResponseV1) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *StreamMarketUpdateResponseV1) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *StreamMarketUpdateResponseV1) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *StreamMarketUpdateResponseV1) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *StreamMarketUpdateResponseV1) GetSequenceId() string {
	if x != nil {
		return x.SequenceId
	}
	return ""
}

func (x *StreamMarketUpdateResponseV1) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StreamMarketUpdateResponseV1) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *StreamMarketUpdateResponseV1) GetTsExchange() *core.TimestampValue {
	if x != nil {
		return x.TsExchange
	}
	return nil
}

func (x *StreamMarketUpdateResponseV1) GetTsCollection() *core.TimestampValue {
	if x != nil {
		return x.TsCollection
	}
	return nil
}

func (x *StreamMarketUpdateResponseV1) GetTsEvent() *timestamp.Timestamp {
	if x != nil {
		return x.TsEvent
	}
	return nil
}

func (x *StreamMarketUpdateResponseV1) GetUpdateType() StreamMarketUpdateResponseV1_StreamMarketUpdateType {
	if x != nil {
		return x.UpdateType
	}
	return StreamMarketUpdateResponseV1_UNKNOWN
}

func (x *StreamMarketUpdateResponseV1) GetSnapshot() *StreamMarketUpdateResponseV1_Snapshot {
	if x != nil {
		return x.Snapshot
	}
	return nil
}

// Snapshot is an orderbook snapshot.
type StreamMarketUpdateResponseV1_Snapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Asks is the list of asks of the orderbook.
	Asks []*StreamMarketUpdateResponseV1_Snapshot_Order `protobuf:"bytes,1,rep,name=asks,proto3" json:"asks,omitempty"`
	// Bids is the list of bids of the orderbook.
	Bids []*StreamMarketUpdateResponseV1_Snapshot_Order `protobuf:"bytes,2,rep,name=bids,proto3" json:"bids,omitempty"`
}

func (x *StreamMarketUpdateResponseV1_Snapshot) Reset() {
	*x = StreamMarketUpdateResponseV1_Snapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_stream_market_update_v1_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamMarketUpdateResponseV1_Snapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamMarketUpdateResponseV1_Snapshot) ProtoMessage() {}

func (x *StreamMarketUpdateResponseV1_Snapshot) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_stream_market_update_v1_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamMarketUpdateResponseV1_Snapshot.ProtoReflect.Descriptor instead.
func (*StreamMarketUpdateResponseV1_Snapshot) Descriptor() ([]byte, []int) {
	return file_sdk_stream_market_update_v1_response_proto_rawDescGZIP(), []int{0, 0}
}

func (x *StreamMarketUpdateResponseV1_Snapshot) GetAsks() []*StreamMarketUpdateResponseV1_Snapshot_Order {
	if x != nil {
		return x.Asks
	}
	return nil
}

func (x *StreamMarketUpdateResponseV1_Snapshot) GetBids() []*StreamMarketUpdateResponseV1_Snapshot_Order {
	if x != nil {
		return x.Bids
	}
	return nil
}

// Order is a amount at a price level.
type StreamMarketUpdateResponseV1_Snapshot_Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Amount / quantity of asset bought or sold, displayed in base currency.
	Amount float64 `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// Price for quote currency.
	Price float64 `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *StreamMarketUpdateResponseV1_Snapshot_Order) Reset() {
	*x = StreamMarketUpdateResponseV1_Snapshot_Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_stream_market_update_v1_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamMarketUpdateResponseV1_Snapshot_Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamMarketUpdateResponseV1_Snapshot_Order) ProtoMessage() {}

func (x *StreamMarketUpdateResponseV1_Snapshot_Order) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_stream_market_update_v1_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamMarketUpdateResponseV1_Snapshot_Order.ProtoReflect.Descriptor instead.
func (*StreamMarketUpdateResponseV1_Snapshot_Order) Descriptor() ([]byte, []int) {
	return file_sdk_stream_market_update_v1_response_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *StreamMarketUpdateResponseV1_Snapshot_Order) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *StreamMarketUpdateResponseV1_Snapshot_Order) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_sdk_stream_market_update_v1_response_proto protoreflect.FileDescriptor

var file_sdk_stream_market_update_v1_response_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6b, 0x61,
	0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x73, 0x64, 0x6b, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2b, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x07,
	0x0a, 0x1c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x12, 0x43,
	0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x25, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x74, 0x73, 0x5f, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x74, 0x73, 0x45, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x61, 0x69,
	0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x74, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x08, 0x74, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x74, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x5e, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3d,
	0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4b, 0x0a, 0x08, 0x73, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6b, 0x61,
	0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x56, 0x31, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x08, 0x73, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x1a, 0xd7, 0x01, 0x0a, 0x08, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x12, 0x49, 0x0a, 0x04, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x35, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x04, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x49,
	0x0a, 0x04, 0x62, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x6b,
	0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x56, 0x31, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x04, 0x62, 0x69, 0x64, 0x73, 0x1a, 0x35, 0x0a, 0x05, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x22, 0x90, 0x01, 0x0a, 0x16, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x52, 0x41, 0x44,
	0x45, 0x5f, 0x42, 0x55, 0x59, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x52, 0x41, 0x44, 0x45,
	0x5f, 0x53, 0x45, 0x4c, 0x4c, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x45, 0x53, 0x54, 0x5f,
	0x41, 0x53, 0x4b, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x45, 0x53, 0x54, 0x5f, 0x42, 0x49,
	0x44, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x41,
	0x53, 0x4b, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x5f,
	0x42, 0x49, 0x44, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x4e, 0x41, 0x50, 0x53, 0x48, 0x4f,
	0x54, 0x10, 0x07, 0x42, 0x9b, 0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x61, 0x69, 0x6b,
	0x6f, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x50, 0x01, 0x5a,
	0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x64, 0x65, 0x65, 0x70, 0x2f, 0x6b, 0x61, 0x69, 0x6b, 0x6f,
	0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x3b,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x31,
	0xaa, 0x02, 0x1e, 0x4b, 0x61, 0x69, 0x6b, 0x6f, 0x53, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sdk_stream_market_update_v1_response_proto_rawDescOnce sync.Once
	file_sdk_stream_market_update_v1_response_proto_rawDescData = file_sdk_stream_market_update_v1_response_proto_rawDesc
)

func file_sdk_stream_market_update_v1_response_proto_rawDescGZIP() []byte {
	file_sdk_stream_market_update_v1_response_proto_rawDescOnce.Do(func() {
		file_sdk_stream_market_update_v1_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdk_stream_market_update_v1_response_proto_rawDescData)
	})
	return file_sdk_stream_market_update_v1_response_proto_rawDescData
}

var file_sdk_stream_market_update_v1_response_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sdk_stream_market_update_v1_response_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sdk_stream_market_update_v1_response_proto_goTypes = []interface{}{
	(StreamMarketUpdateResponseV1_StreamMarketUpdateType)(0), // 0: kaikosdk.StreamMarketUpdateResponseV1.StreamMarketUpdateType
	(*StreamMarketUpdateResponseV1)(nil),                     // 1: kaikosdk.StreamMarketUpdateResponseV1
	(*StreamMarketUpdateResponseV1_Snapshot)(nil),            // 2: kaikosdk.StreamMarketUpdateResponseV1.Snapshot
	(*StreamMarketUpdateResponseV1_Snapshot_Order)(nil),      // 3: kaikosdk.StreamMarketUpdateResponseV1.Snapshot.Order
	(StreamMarketUpdateCommodity)(0),                         // 4: kaikosdk.StreamMarketUpdateCommodity
	(*core.TimestampValue)(nil),                              // 5: kaikosdk.TimestampValue
	(*timestamp.Timestamp)(nil),                              // 6: google.protobuf.Timestamp
}
var file_sdk_stream_market_update_v1_response_proto_depIdxs = []int32{
	4, // 0: kaikosdk.StreamMarketUpdateResponseV1.commodity:type_name -> kaikosdk.StreamMarketUpdateCommodity
	5, // 1: kaikosdk.StreamMarketUpdateResponseV1.ts_exchange:type_name -> kaikosdk.TimestampValue
	5, // 2: kaikosdk.StreamMarketUpdateResponseV1.ts_collection:type_name -> kaikosdk.TimestampValue
	6, // 3: kaikosdk.StreamMarketUpdateResponseV1.ts_event:type_name -> google.protobuf.Timestamp
	0, // 4: kaikosdk.StreamMarketUpdateResponseV1.update_type:type_name -> kaikosdk.StreamMarketUpdateResponseV1.StreamMarketUpdateType
	2, // 5: kaikosdk.StreamMarketUpdateResponseV1.snapshot:type_name -> kaikosdk.StreamMarketUpdateResponseV1.Snapshot
	3, // 6: kaikosdk.StreamMarketUpdateResponseV1.Snapshot.asks:type_name -> kaikosdk.StreamMarketUpdateResponseV1.Snapshot.Order
	3, // 7: kaikosdk.StreamMarketUpdateResponseV1.Snapshot.bids:type_name -> kaikosdk.StreamMarketUpdateResponseV1.Snapshot.Order
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_sdk_stream_market_update_v1_response_proto_init() }
func file_sdk_stream_market_update_v1_response_proto_init() {
	if File_sdk_stream_market_update_v1_response_proto != nil {
		return
	}
	file_sdk_stream_market_update_v1_commodity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sdk_stream_market_update_v1_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamMarketUpdateResponseV1); i {
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
		file_sdk_stream_market_update_v1_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamMarketUpdateResponseV1_Snapshot); i {
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
		file_sdk_stream_market_update_v1_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamMarketUpdateResponseV1_Snapshot_Order); i {
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
			RawDescriptor: file_sdk_stream_market_update_v1_response_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sdk_stream_market_update_v1_response_proto_goTypes,
		DependencyIndexes: file_sdk_stream_market_update_v1_response_proto_depIdxs,
		EnumInfos:         file_sdk_stream_market_update_v1_response_proto_enumTypes,
		MessageInfos:      file_sdk_stream_market_update_v1_response_proto_msgTypes,
	}.Build()
	File_sdk_stream_market_update_v1_response_proto = out.File
	file_sdk_stream_market_update_v1_response_proto_rawDesc = nil
	file_sdk_stream_market_update_v1_response_proto_goTypes = nil
	file_sdk_stream_market_update_v1_response_proto_depIdxs = nil
}
