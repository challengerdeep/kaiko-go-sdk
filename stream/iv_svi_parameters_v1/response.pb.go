// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: sdk/stream/iv_svi_parameters_v1/response.proto

package iv_svi_parameters_v1

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

// StreamIvSviParametersResponseV1 is IV SVI paramater stream response.
type StreamIvSviParametersResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Start time of the calculation window used.
	DataStartTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=data_start_time,json=dataStartTime,proto3" json:"data_start_time,omitempty"`
	// End time of the calculation window used.
	DataEndTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=data_end_time,json=dataEndTime,proto3" json:"data_end_time,omitempty"`
	// Expiry date of the instrument.
	Expiry *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expiry,proto3" json:"expiry,omitempty"`
	// Exchanges included.
	Exchanges []string `protobuf:"bytes,4,rep,name=exchanges,proto3" json:"exchanges,omitempty"`
	// Base asset.
	Base string `protobuf:"bytes,5,opt,name=base,proto3" json:"base,omitempty"`
	// Quote asset.
	Quote string `protobuf:"bytes,6,opt,name=quote,proto3" json:"quote,omitempty"`
	// Time to expiry.
	TimeToExpiry string `protobuf:"bytes,7,opt,name=time_to_expiry,json=timeToExpiry,proto3" json:"time_to_expiry,omitempty"`
	// ATM implied variance.
	AtmImpliedVariance string `protobuf:"bytes,8,opt,name=atm_implied_variance,json=atmImpliedVariance,proto3" json:"atm_implied_variance,omitempty"`
	// ATM skew.
	AtmSkew string `protobuf:"bytes,9,opt,name=atm_skew,json=atmSkew,proto3" json:"atm_skew,omitempty"`
	// Left slope of IV smile.
	LeftSlope string `protobuf:"bytes,10,opt,name=left_slope,json=leftSlope,proto3" json:"left_slope,omitempty"`
	// Right slope of IV smile.
	RightSlope string `protobuf:"bytes,11,opt,name=right_slope,json=rightSlope,proto3" json:"right_slope,omitempty"`
	// Min implied variance.
	MinImpliedVariance string `protobuf:"bytes,12,opt,name=min_implied_variance,json=minImpliedVariance,proto3" json:"min_implied_variance,omitempty"`
	// Current spot.
	CurrentSpot string `protobuf:"bytes,13,opt,name=current_spot,json=currentSpot,proto3" json:"current_spot,omitempty"`
	// Interest rate.
	InterestRate string `protobuf:"bytes,14,opt,name=interest_rate,json=interestRate,proto3" json:"interest_rate,omitempty"`
	// Event generation timestamp (event created by Kaiko), after normalization.
	TsEvent *timestamp.Timestamp `protobuf:"bytes,15,opt,name=ts_event,json=tsEvent,proto3" json:"ts_event,omitempty"`
}

func (x *StreamIvSviParametersResponseV1) Reset() {
	*x = StreamIvSviParametersResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_stream_iv_svi_parameters_v1_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamIvSviParametersResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamIvSviParametersResponseV1) ProtoMessage() {}

func (x *StreamIvSviParametersResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_stream_iv_svi_parameters_v1_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamIvSviParametersResponseV1.ProtoReflect.Descriptor instead.
func (*StreamIvSviParametersResponseV1) Descriptor() ([]byte, []int) {
	return file_sdk_stream_iv_svi_parameters_v1_response_proto_rawDescGZIP(), []int{0}
}

func (x *StreamIvSviParametersResponseV1) GetDataStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.DataStartTime
	}
	return nil
}

func (x *StreamIvSviParametersResponseV1) GetDataEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.DataEndTime
	}
	return nil
}

func (x *StreamIvSviParametersResponseV1) GetExpiry() *timestamp.Timestamp {
	if x != nil {
		return x.Expiry
	}
	return nil
}

func (x *StreamIvSviParametersResponseV1) GetExchanges() []string {
	if x != nil {
		return x.Exchanges
	}
	return nil
}

func (x *StreamIvSviParametersResponseV1) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *StreamIvSviParametersResponseV1) GetQuote() string {
	if x != nil {
		return x.Quote
	}
	return ""
}

func (x *StreamIvSviParametersResponseV1) GetTimeToExpiry() string {
	if x != nil {
		return x.TimeToExpiry
	}
	return ""
}

func (x *StreamIvSviParametersResponseV1) GetAtmImpliedVariance() string {
	if x != nil {
		return x.AtmImpliedVariance
	}
	return ""
}

func (x *StreamIvSviParametersResponseV1) GetAtmSkew() string {
	if x != nil {
		return x.AtmSkew
	}
	return ""
}

func (x *StreamIvSviParametersResponseV1) GetLeftSlope() string {
	if x != nil {
		return x.LeftSlope
	}
	return ""
}

func (x *StreamIvSviParametersResponseV1) GetRightSlope() string {
	if x != nil {
		return x.RightSlope
	}
	return ""
}

func (x *StreamIvSviParametersResponseV1) GetMinImpliedVariance() string {
	if x != nil {
		return x.MinImpliedVariance
	}
	return ""
}

func (x *StreamIvSviParametersResponseV1) GetCurrentSpot() string {
	if x != nil {
		return x.CurrentSpot
	}
	return ""
}

func (x *StreamIvSviParametersResponseV1) GetInterestRate() string {
	if x != nil {
		return x.InterestRate
	}
	return ""
}

func (x *StreamIvSviParametersResponseV1) GetTsEvent() *timestamp.Timestamp {
	if x != nil {
		return x.TsEvent
	}
	return nil
}

var File_sdk_stream_iv_svi_parameters_v1_response_proto protoreflect.FileDescriptor

var file_sdk_stream_iv_svi_parameters_v1_response_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x69, 0x76, 0x5f,
	0x73, 0x76, 0x69, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x05, 0x0a, 0x1f,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x76, 0x53, 0x76, 0x69, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x12,
	0x42, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x65, 0x6e, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x12,
	0x24, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x6f, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x6f, 0x45,
	0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x74, 0x6d, 0x5f, 0x69, 0x6d, 0x70,
	0x6c, 0x69, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x61, 0x74, 0x6d, 0x49, 0x6d, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x74, 0x6d, 0x5f, 0x73,
	0x6b, 0x65, 0x77, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x74, 0x6d, 0x53, 0x6b,
	0x65, 0x77, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x65, 0x66, 0x74, 0x5f, 0x73, 0x6c, 0x6f, 0x70, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x65, 0x66, 0x74, 0x53, 0x6c, 0x6f, 0x70,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x73, 0x6c, 0x6f, 0x70, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x69, 0x67, 0x68, 0x74, 0x53, 0x6c, 0x6f,
	0x70, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x65,
	0x64, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x6d, 0x69, 0x6e, 0x49, 0x6d, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x70, 0x6f, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x53, 0x70, 0x6f, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x65, 0x73, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x08,
	0x74, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x74, 0x73, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x42, 0xa4, 0x01, 0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x61, 0x69, 0x6b,
	0x6f, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x69, 0x76, 0x5f,
	0x73, 0x76, 0x69, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x2d,
	0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x69, 0x76,
	0x5f, 0x73, 0x76, 0x69, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x5f,
	0x76, 0x31, 0x3b, 0x69, 0x76, 0x5f, 0x73, 0x76, 0x69, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x5f, 0x76, 0x31, 0xaa, 0x02, 0x20, 0x4b, 0x61, 0x69, 0x6b, 0x6f, 0x53,
	0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x49, 0x76, 0x53, 0x76, 0x69, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sdk_stream_iv_svi_parameters_v1_response_proto_rawDescOnce sync.Once
	file_sdk_stream_iv_svi_parameters_v1_response_proto_rawDescData = file_sdk_stream_iv_svi_parameters_v1_response_proto_rawDesc
)

func file_sdk_stream_iv_svi_parameters_v1_response_proto_rawDescGZIP() []byte {
	file_sdk_stream_iv_svi_parameters_v1_response_proto_rawDescOnce.Do(func() {
		file_sdk_stream_iv_svi_parameters_v1_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdk_stream_iv_svi_parameters_v1_response_proto_rawDescData)
	})
	return file_sdk_stream_iv_svi_parameters_v1_response_proto_rawDescData
}

var file_sdk_stream_iv_svi_parameters_v1_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sdk_stream_iv_svi_parameters_v1_response_proto_goTypes = []interface{}{
	(*StreamIvSviParametersResponseV1)(nil), // 0: kaikosdk.StreamIvSviParametersResponseV1
	(*timestamp.Timestamp)(nil),             // 1: google.protobuf.Timestamp
}
var file_sdk_stream_iv_svi_parameters_v1_response_proto_depIdxs = []int32{
	1, // 0: kaikosdk.StreamIvSviParametersResponseV1.data_start_time:type_name -> google.protobuf.Timestamp
	1, // 1: kaikosdk.StreamIvSviParametersResponseV1.data_end_time:type_name -> google.protobuf.Timestamp
	1, // 2: kaikosdk.StreamIvSviParametersResponseV1.expiry:type_name -> google.protobuf.Timestamp
	1, // 3: kaikosdk.StreamIvSviParametersResponseV1.ts_event:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_sdk_stream_iv_svi_parameters_v1_response_proto_init() }
func file_sdk_stream_iv_svi_parameters_v1_response_proto_init() {
	if File_sdk_stream_iv_svi_parameters_v1_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sdk_stream_iv_svi_parameters_v1_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamIvSviParametersResponseV1); i {
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
			RawDescriptor: file_sdk_stream_iv_svi_parameters_v1_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sdk_stream_iv_svi_parameters_v1_response_proto_goTypes,
		DependencyIndexes: file_sdk_stream_iv_svi_parameters_v1_response_proto_depIdxs,
		MessageInfos:      file_sdk_stream_iv_svi_parameters_v1_response_proto_msgTypes,
	}.Build()
	File_sdk_stream_iv_svi_parameters_v1_response_proto = out.File
	file_sdk_stream_iv_svi_parameters_v1_response_proto_rawDesc = nil
	file_sdk_stream_iv_svi_parameters_v1_response_proto_goTypes = nil
	file_sdk_stream_iv_svi_parameters_v1_response_proto_depIdxs = nil
}
