// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: sdk/sdk.proto

package kaikosdk

import (
	aggregated_price_v1 "github.com/kaikodata/kaiko-go-sdk/stream/aggregated_price_v1"
	aggregated_quote_v2 "github.com/kaikodata/kaiko-go-sdk/stream/aggregated_quote_v2"
	aggregated_state_price_v1 "github.com/kaikodata/kaiko-go-sdk/stream/aggregated_state_price_v1"
	aggregates_direct_exchange_rate_v2 "github.com/kaikodata/kaiko-go-sdk/stream/aggregates_direct_exchange_rate_v2"
	aggregates_ohlcv_v1 "github.com/kaikodata/kaiko-go-sdk/stream/aggregates_ohlcv_v1"
	aggregates_spot_exchange_rate_v2 "github.com/kaikodata/kaiko-go-sdk/stream/aggregates_spot_exchange_rate_v2"
	aggregates_vwap_v1 "github.com/kaikodata/kaiko-go-sdk/stream/aggregates_vwap_v1"
	derivatives_instrument_metrics_v1 "github.com/kaikodata/kaiko-go-sdk/stream/derivatives_instrument_metrics_v1"
	exotic_indices_v1 "github.com/kaikodata/kaiko-go-sdk/stream/exotic_indices_v1"
	index_forex_rate_v1 "github.com/kaikodata/kaiko-go-sdk/stream/index_forex_rate_v1"
	index_multi_assets_v1 "github.com/kaikodata/kaiko-go-sdk/stream/index_multi_assets_v1"
	index_v1 "github.com/kaikodata/kaiko-go-sdk/stream/index_v1"
	iv_svi_parameters_v1 "github.com/kaikodata/kaiko-go-sdk/stream/iv_svi_parameters_v1"
	market_update_v1 "github.com/kaikodata/kaiko-go-sdk/stream/market_update_v1"
	orderbookl2_v1 "github.com/kaikodata/kaiko-go-sdk/stream/orderbookl2_v1"
	trades_v1 "github.com/kaikodata/kaiko-go-sdk/stream/trades_v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_sdk_sdk_proto protoreflect.FileDescriptor

var file_sdk_sdk_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x1a, 0x2c, 0x73, 0x64, 0x6b, 0x2f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x71, 0x75,
	0x6f, 0x74, 0x65, 0x5f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x5f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f,
	0x76, 0x32, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x3c, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x32, 0x2f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c,
	0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x6f, 0x68, 0x6c, 0x63, 0x76, 0x5f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x73, 0x5f, 0x6f, 0x68, 0x6c, 0x63, 0x76, 0x5f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x73, 0x64, 0x6b,
	0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x73, 0x5f, 0x73, 0x70, 0x6f, 0x74, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x73, 0x70,
	0x6f, 0x74, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65,
	0x5f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2b, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x76, 0x77, 0x61, 0x70, 0x5f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2c, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x76, 0x77, 0x61, 0x70, 0x5f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x73,
	0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x64, 0x65, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x73, 0x64, 0x6b, 0x2f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x64, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x73, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x5f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2f, 0x65, 0x78, 0x6f, 0x74, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65,
	0x73, 0x5f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2b, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x65,
	0x78, 0x6f, 0x74, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2c, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f,
	0x66, 0x6f, 0x72, 0x65, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x73, 0x64,
	0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x5f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x73, 0x64,
	0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x5f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x73,
	0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x22, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x5f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2f, 0x69, 0x76, 0x5f, 0x73, 0x76, 0x69, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x5f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f,
	0x69, 0x76, 0x5f, 0x73, 0x76, 0x69, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x5f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a,
	0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x73, 0x64, 0x6b, 0x2f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x6f, 0x6b,
	0x6c, 0x32, 0x5f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x6f, 0x6b, 0x6c, 0x32, 0x5f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x73,
	0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x73,
	0x5f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x23, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x73, 0x64, 0x6b, 0x2f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x86, 0x01, 0x0a, 0x1e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x56, 0x32, 0x12, 0x64, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12,
	0x28, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x32, 0x1a, 0x29, 0x2e, 0x6b, 0x61, 0x69, 0x6b,
	0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x56, 0x32, 0x22, 0x00, 0x30, 0x01, 0x32, 0x89, 0x01, 0x0a, 0x1e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x12, 0x67, 0x0a, 0x09, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x28, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f,
	0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x31, 0x1a, 0x29, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x03, 0x88,
	0x02, 0x01, 0x30, 0x01, 0x32, 0x86, 0x01, 0x0a, 0x1e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x4f, 0x48, 0x4c, 0x43, 0x56, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x12, 0x64, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x12, 0x28, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73,
	0x4f, 0x48, 0x4c, 0x43, 0x56, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x29,
	0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x4f, 0x48, 0x4c, 0x43, 0x56, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x00, 0x30, 0x01, 0x32, 0x6b, 0x0a,
	0x15, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x12, 0x52, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x12, 0x1f, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x56, 0x31, 0x1a, 0x20, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x00, 0x30, 0x01, 0x32, 0x83, 0x01, 0x0a, 0x1d, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x56,
	0x57, 0x41, 0x50, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x12, 0x62, 0x0a, 0x09,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x27, 0x2e, 0x6b, 0x61, 0x69, 0x6b,
	0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x73, 0x56, 0x57, 0x41, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x31, 0x1a, 0x28, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x56, 0x57,
	0x41, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x00, 0x30, 0x01,
	0x32, 0x76, 0x0a, 0x14, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x12, 0x5e, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x25, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x26, 0x2e, 0x6b,
	0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x56, 0x31, 0x22, 0x00, 0x30, 0x01, 0x32, 0x97, 0x01, 0x0a, 0x1f, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x12, 0x74, 0x0a, 0x09,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x30, 0x2e, 0x6b, 0x61, 0x69, 0x6b,
	0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x31, 0x2e, 0x6b, 0x61,
	0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x00,
	0x30, 0x01, 0x32, 0x7d, 0x0a, 0x1b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56,
	0x31, 0x12, 0x5e, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x25,
	0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x26, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x00, 0x30,
	0x01, 0x32, 0xad, 0x01, 0x0a, 0x2b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x53, 0x70, 0x6f, 0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x56, 0x32, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56,
	0x31, 0x12, 0x7e, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x35,
	0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x53, 0x70, 0x6f, 0x74, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x56, 0x32, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x36, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x73, 0x53, 0x70, 0x6f, 0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x56, 0x32, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x00, 0x30,
	0x01, 0x32, 0xb8, 0x01, 0x0a, 0x31, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x53, 0x70, 0x6f, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x56, 0x32, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x12, 0x82, 0x01, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x37, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x73, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x61, 0x74, 0x65, 0x56, 0x32, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x38,
	0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x56, 0x32, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x00, 0x30, 0x01, 0x32, 0x91, 0x01, 0x0a,
	0x1d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x6f, 0x72, 0x65,
	0x78, 0x52, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x12, 0x70,
	0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x2e, 0x2e, 0x6b, 0x61,
	0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x46, 0x6f, 0x72, 0x65, 0x78, 0x52, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x2f, 0x2e, 0x6b, 0x61,
	0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x46, 0x6f, 0x72, 0x65, 0x78, 0x52, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x00, 0x30, 0x01,
	0x32, 0xad, 0x01, 0x0a, 0x2b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x65, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x73, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31,
	0x12, 0x7e, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x35, 0x2e,
	0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44,
	0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x76, 0x65, 0x73, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x56, 0x31, 0x1a, 0x36, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x73, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x00, 0x30, 0x01,
	0x32, 0x86, 0x01, 0x0a, 0x1e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x76, 0x53, 0x76, 0x69,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x56, 0x31, 0x12, 0x64, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x12, 0x28, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x49, 0x76, 0x53, 0x76, 0x69, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x29, 0x2e, 0x6b, 0x61, 0x69,
	0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x76, 0x53, 0x76,
	0x69, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x00, 0x30, 0x01, 0x32, 0x8e, 0x01, 0x0a, 0x1c, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x6f, 0x74, 0x69, 0x63, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x12, 0x6e, 0x0a, 0x09, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x2d, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73,
	0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x6f, 0x74, 0x69, 0x63, 0x49,
	0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x2e, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64,
	0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x6f, 0x74, 0x69, 0x63, 0x49, 0x6e,
	0x64, 0x69, 0x63, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x00, 0x30, 0x01, 0x32, 0x7a, 0x0a, 0x1a, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x6f, 0x6b, 0x4c, 0x32, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x12, 0x5c, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x24, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b,
	0x4c, 0x32, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x25, 0x2e, 0x6b, 0x61,
	0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x32, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x56, 0x31, 0x22, 0x00, 0x30, 0x01, 0x32, 0x95, 0x01, 0x0a, 0x23, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x12, 0x6e,
	0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x2d, 0x2e, 0x6b, 0x61,
	0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x2e, 0x2e, 0x6b, 0x61, 0x69,
	0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x00, 0x30, 0x01, 0x32, 0x86,
	0x01, 0x0a, 0x20, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x6f,
	0x6f, 0x6b, 0x4c, 0x32, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x56, 0x31, 0x12, 0x62, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x12, 0x2a, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x32, 0x52, 0x65, 0x70,
	0x6c, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x25, 0x2e, 0x6b,
	0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x32, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x56, 0x31, 0x22, 0x00, 0x30, 0x01, 0x42, 0x52, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x6b,
	0x61, 0x69, 0x6b, 0x6f, 0x2e, 0x73, 0x64, 0x6b, 0x42, 0x08, 0x53, 0x64, 0x6b, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6b, 0x61, 0x69, 0x6b, 0x6f,
	0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x3b, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b,
	0xaa, 0x02, 0x08, 0x4b, 0x61, 0x69, 0x6b, 0x6f, 0x53, 0x64, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_sdk_sdk_proto_goTypes = []interface{}{
	(*aggregated_quote_v2.StreamAggregatedQuoteRequestV2)(nil),                                // 0: kaikosdk.StreamAggregatedQuoteRequestV2
	(*aggregated_price_v1.StreamAggregatedPriceRequestV1)(nil),                                // 1: kaikosdk.StreamAggregatedPriceRequestV1
	(*aggregates_ohlcv_v1.StreamAggregatesOHLCVRequestV1)(nil),                                // 2: kaikosdk.StreamAggregatesOHLCVRequestV1
	(*trades_v1.StreamTradesRequestV1)(nil),                                                   // 3: kaikosdk.StreamTradesRequestV1
	(*aggregates_vwap_v1.StreamAggregatesVWAPRequestV1)(nil),                                  // 4: kaikosdk.StreamAggregatesVWAPRequestV1
	(*index_v1.StreamIndexServiceRequestV1)(nil),                                              // 5: kaikosdk.StreamIndexServiceRequestV1
	(*index_multi_assets_v1.StreamIndexMultiAssetsServiceRequestV1)(nil),                      // 6: kaikosdk.StreamIndexMultiAssetsServiceRequestV1
	(*market_update_v1.StreamMarketUpdateRequestV1)(nil),                                      // 7: kaikosdk.StreamMarketUpdateRequestV1
	(*aggregates_spot_exchange_rate_v2.StreamAggregatesSpotExchangeRateV2RequestV1)(nil),      // 8: kaikosdk.StreamAggregatesSpotExchangeRateV2RequestV1
	(*aggregates_direct_exchange_rate_v2.StreamAggregatesDirectExchangeRateV2RequestV1)(nil),  // 9: kaikosdk.StreamAggregatesDirectExchangeRateV2RequestV1
	(*index_forex_rate_v1.StreamIndexForexRateServiceRequestV1)(nil),                          // 10: kaikosdk.StreamIndexForexRateServiceRequestV1
	(*derivatives_instrument_metrics_v1.StreamDerivativesInstrumentMetricsRequestV1)(nil),     // 11: kaikosdk.StreamDerivativesInstrumentMetricsRequestV1
	(*iv_svi_parameters_v1.StreamIvSviParametersRequestV1)(nil),                               // 12: kaikosdk.StreamIvSviParametersRequestV1
	(*exotic_indices_v1.StreamExoticIndicesServiceRequestV1)(nil),                             // 13: kaikosdk.StreamExoticIndicesServiceRequestV1
	(*orderbookl2_v1.StreamOrderBookL2RequestV1)(nil),                                         // 14: kaikosdk.StreamOrderBookL2RequestV1
	(*aggregated_state_price_v1.StreamAggregatedStatePriceRequestV1)(nil),                     // 15: kaikosdk.StreamAggregatedStatePriceRequestV1
	(*orderbookl2_v1.StreamOrderBookL2ReplayRequestV1)(nil),                                   // 16: kaikosdk.StreamOrderBookL2ReplayRequestV1
	(*aggregated_quote_v2.StreamAggregatedQuoteResponseV2)(nil),                               // 17: kaikosdk.StreamAggregatedQuoteResponseV2
	(*aggregated_price_v1.StreamAggregatedPriceResponseV1)(nil),                               // 18: kaikosdk.StreamAggregatedPriceResponseV1
	(*aggregates_ohlcv_v1.StreamAggregatesOHLCVResponseV1)(nil),                               // 19: kaikosdk.StreamAggregatesOHLCVResponseV1
	(*trades_v1.StreamTradesResponseV1)(nil),                                                  // 20: kaikosdk.StreamTradesResponseV1
	(*aggregates_vwap_v1.StreamAggregatesVWAPResponseV1)(nil),                                 // 21: kaikosdk.StreamAggregatesVWAPResponseV1
	(*index_v1.StreamIndexServiceResponseV1)(nil),                                             // 22: kaikosdk.StreamIndexServiceResponseV1
	(*index_multi_assets_v1.StreamIndexMultiAssetsServiceResponseV1)(nil),                     // 23: kaikosdk.StreamIndexMultiAssetsServiceResponseV1
	(*market_update_v1.StreamMarketUpdateResponseV1)(nil),                                     // 24: kaikosdk.StreamMarketUpdateResponseV1
	(*aggregates_spot_exchange_rate_v2.StreamAggregatesSpotExchangeRateV2ResponseV1)(nil),     // 25: kaikosdk.StreamAggregatesSpotExchangeRateV2ResponseV1
	(*aggregates_direct_exchange_rate_v2.StreamAggregatesDirectExchangeRateV2ResponseV1)(nil), // 26: kaikosdk.StreamAggregatesDirectExchangeRateV2ResponseV1
	(*index_forex_rate_v1.StreamIndexForexRateServiceResponseV1)(nil),                         // 27: kaikosdk.StreamIndexForexRateServiceResponseV1
	(*derivatives_instrument_metrics_v1.StreamDerivativesInstrumentMetricsResponseV1)(nil),    // 28: kaikosdk.StreamDerivativesInstrumentMetricsResponseV1
	(*iv_svi_parameters_v1.StreamIvSviParametersResponseV1)(nil),                              // 29: kaikosdk.StreamIvSviParametersResponseV1
	(*exotic_indices_v1.StreamExoticIndicesServiceResponseV1)(nil),                            // 30: kaikosdk.StreamExoticIndicesServiceResponseV1
	(*orderbookl2_v1.StreamOrderBookL2ResponseV1)(nil),                                        // 31: kaikosdk.StreamOrderBookL2ResponseV1
	(*aggregated_state_price_v1.StreamAggregatedStatePriceResponseV1)(nil),                    // 32: kaikosdk.StreamAggregatedStatePriceResponseV1
}
var file_sdk_sdk_proto_depIdxs = []int32{
	0,  // 0: kaikosdk.StreamAggregatedQuoteServiceV2.Subscribe:input_type -> kaikosdk.StreamAggregatedQuoteRequestV2
	1,  // 1: kaikosdk.StreamAggregatedPriceServiceV1.Subscribe:input_type -> kaikosdk.StreamAggregatedPriceRequestV1
	2,  // 2: kaikosdk.StreamAggregatesOHLCVServiceV1.Subscribe:input_type -> kaikosdk.StreamAggregatesOHLCVRequestV1
	3,  // 3: kaikosdk.StreamTradesServiceV1.Subscribe:input_type -> kaikosdk.StreamTradesRequestV1
	4,  // 4: kaikosdk.StreamAggregatesVWAPServiceV1.Subscribe:input_type -> kaikosdk.StreamAggregatesVWAPRequestV1
	5,  // 5: kaikosdk.StreamIndexServiceV1.Subscribe:input_type -> kaikosdk.StreamIndexServiceRequestV1
	6,  // 6: kaikosdk.StreamIndexMultiAssetsServiceV1.Subscribe:input_type -> kaikosdk.StreamIndexMultiAssetsServiceRequestV1
	7,  // 7: kaikosdk.StreamMarketUpdateServiceV1.Subscribe:input_type -> kaikosdk.StreamMarketUpdateRequestV1
	8,  // 8: kaikosdk.StreamAggregatesSpotExchangeRateV2ServiceV1.Subscribe:input_type -> kaikosdk.StreamAggregatesSpotExchangeRateV2RequestV1
	9,  // 9: kaikosdk.StreamAggregatesSpotDirectExchangeRateV2ServiceV1.Subscribe:input_type -> kaikosdk.StreamAggregatesDirectExchangeRateV2RequestV1
	10, // 10: kaikosdk.StreamIndexForexRateServiceV1.Subscribe:input_type -> kaikosdk.StreamIndexForexRateServiceRequestV1
	11, // 11: kaikosdk.StreamDerivativesInstrumentMetricsServiceV1.Subscribe:input_type -> kaikosdk.StreamDerivativesInstrumentMetricsRequestV1
	12, // 12: kaikosdk.StreamIvSviParametersServiceV1.Subscribe:input_type -> kaikosdk.StreamIvSviParametersRequestV1
	13, // 13: kaikosdk.StreamExoticIndicesServiceV1.Subscribe:input_type -> kaikosdk.StreamExoticIndicesServiceRequestV1
	14, // 14: kaikosdk.StreamOrderbookL2ServiceV1.Subscribe:input_type -> kaikosdk.StreamOrderBookL2RequestV1
	15, // 15: kaikosdk.StreamAggregatedStatePriceServiceV1.Subscribe:input_type -> kaikosdk.StreamAggregatedStatePriceRequestV1
	16, // 16: kaikosdk.StreamOrderbookL2ReplayServiceV1.Subscribe:input_type -> kaikosdk.StreamOrderBookL2ReplayRequestV1
	17, // 17: kaikosdk.StreamAggregatedQuoteServiceV2.Subscribe:output_type -> kaikosdk.StreamAggregatedQuoteResponseV2
	18, // 18: kaikosdk.StreamAggregatedPriceServiceV1.Subscribe:output_type -> kaikosdk.StreamAggregatedPriceResponseV1
	19, // 19: kaikosdk.StreamAggregatesOHLCVServiceV1.Subscribe:output_type -> kaikosdk.StreamAggregatesOHLCVResponseV1
	20, // 20: kaikosdk.StreamTradesServiceV1.Subscribe:output_type -> kaikosdk.StreamTradesResponseV1
	21, // 21: kaikosdk.StreamAggregatesVWAPServiceV1.Subscribe:output_type -> kaikosdk.StreamAggregatesVWAPResponseV1
	22, // 22: kaikosdk.StreamIndexServiceV1.Subscribe:output_type -> kaikosdk.StreamIndexServiceResponseV1
	23, // 23: kaikosdk.StreamIndexMultiAssetsServiceV1.Subscribe:output_type -> kaikosdk.StreamIndexMultiAssetsServiceResponseV1
	24, // 24: kaikosdk.StreamMarketUpdateServiceV1.Subscribe:output_type -> kaikosdk.StreamMarketUpdateResponseV1
	25, // 25: kaikosdk.StreamAggregatesSpotExchangeRateV2ServiceV1.Subscribe:output_type -> kaikosdk.StreamAggregatesSpotExchangeRateV2ResponseV1
	26, // 26: kaikosdk.StreamAggregatesSpotDirectExchangeRateV2ServiceV1.Subscribe:output_type -> kaikosdk.StreamAggregatesDirectExchangeRateV2ResponseV1
	27, // 27: kaikosdk.StreamIndexForexRateServiceV1.Subscribe:output_type -> kaikosdk.StreamIndexForexRateServiceResponseV1
	28, // 28: kaikosdk.StreamDerivativesInstrumentMetricsServiceV1.Subscribe:output_type -> kaikosdk.StreamDerivativesInstrumentMetricsResponseV1
	29, // 29: kaikosdk.StreamIvSviParametersServiceV1.Subscribe:output_type -> kaikosdk.StreamIvSviParametersResponseV1
	30, // 30: kaikosdk.StreamExoticIndicesServiceV1.Subscribe:output_type -> kaikosdk.StreamExoticIndicesServiceResponseV1
	31, // 31: kaikosdk.StreamOrderbookL2ServiceV1.Subscribe:output_type -> kaikosdk.StreamOrderBookL2ResponseV1
	32, // 32: kaikosdk.StreamAggregatedStatePriceServiceV1.Subscribe:output_type -> kaikosdk.StreamAggregatedStatePriceResponseV1
	31, // 33: kaikosdk.StreamOrderbookL2ReplayServiceV1.Subscribe:output_type -> kaikosdk.StreamOrderBookL2ResponseV1
	17, // [17:34] is the sub-list for method output_type
	0,  // [0:17] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_sdk_sdk_proto_init() }
func file_sdk_sdk_proto_init() {
	if File_sdk_sdk_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sdk_sdk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   17,
		},
		GoTypes:           file_sdk_sdk_proto_goTypes,
		DependencyIndexes: file_sdk_sdk_proto_depIdxs,
	}.Build()
	File_sdk_sdk_proto = out.File
	file_sdk_sdk_proto_rawDesc = nil
	file_sdk_sdk_proto_goTypes = nil
	file_sdk_sdk_proto_depIdxs = nil
}
