// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: cur_rates.proto

package rates

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Currencies int32

const (
	Currencies_AUD Currencies = 0
	Currencies_AZN Currencies = 1
	Currencies_GBP Currencies = 2
	Currencies_AMD Currencies = 3
	Currencies_BYN Currencies = 4
	Currencies_BGN Currencies = 5
	Currencies_BRL Currencies = 6
	Currencies_HUF Currencies = 7
	Currencies_VND Currencies = 8
	Currencies_HKD Currencies = 9
	Currencies_GEL Currencies = 10
	Currencies_DKK Currencies = 11
	Currencies_AED Currencies = 12
	Currencies_USD Currencies = 13
	Currencies_EUR Currencies = 14
	Currencies_EGP Currencies = 15
	Currencies_INR Currencies = 16
	Currencies_IDR Currencies = 17
	Currencies_KZT Currencies = 18
	Currencies_CAD Currencies = 19
	Currencies_QAR Currencies = 20
	Currencies_KGS Currencies = 21
	Currencies_CNY Currencies = 22
	Currencies_MDL Currencies = 23
	Currencies_NZD Currencies = 24
	Currencies_NOK Currencies = 25
	Currencies_PLN Currencies = 26
	Currencies_RON Currencies = 27
	Currencies_XDR Currencies = 28
	Currencies_SGD Currencies = 29
	Currencies_TJS Currencies = 30
	Currencies_THB Currencies = 31
	Currencies_TRY Currencies = 32
	Currencies_TMT Currencies = 33
	Currencies_UZS Currencies = 34
	Currencies_UAH Currencies = 35
	Currencies_CZK Currencies = 36
	Currencies_SEK Currencies = 37
	Currencies_CHF Currencies = 38
	Currencies_RSD Currencies = 39
	Currencies_ZAR Currencies = 40
	Currencies_KRW Currencies = 41
	Currencies_JPY Currencies = 42
)

// Enum value maps for Currencies.
var (
	Currencies_name = map[int32]string{
		0:  "AUD",
		1:  "AZN",
		2:  "GBP",
		3:  "AMD",
		4:  "BYN",
		5:  "BGN",
		6:  "BRL",
		7:  "HUF",
		8:  "VND",
		9:  "HKD",
		10: "GEL",
		11: "DKK",
		12: "AED",
		13: "USD",
		14: "EUR",
		15: "EGP",
		16: "INR",
		17: "IDR",
		18: "KZT",
		19: "CAD",
		20: "QAR",
		21: "KGS",
		22: "CNY",
		23: "MDL",
		24: "NZD",
		25: "NOK",
		26: "PLN",
		27: "RON",
		28: "XDR",
		29: "SGD",
		30: "TJS",
		31: "THB",
		32: "TRY",
		33: "TMT",
		34: "UZS",
		35: "UAH",
		36: "CZK",
		37: "SEK",
		38: "CHF",
		39: "RSD",
		40: "ZAR",
		41: "KRW",
		42: "JPY",
	}
	Currencies_value = map[string]int32{
		"AUD": 0,
		"AZN": 1,
		"GBP": 2,
		"AMD": 3,
		"BYN": 4,
		"BGN": 5,
		"BRL": 6,
		"HUF": 7,
		"VND": 8,
		"HKD": 9,
		"GEL": 10,
		"DKK": 11,
		"AED": 12,
		"USD": 13,
		"EUR": 14,
		"EGP": 15,
		"INR": 16,
		"IDR": 17,
		"KZT": 18,
		"CAD": 19,
		"QAR": 20,
		"KGS": 21,
		"CNY": 22,
		"MDL": 23,
		"NZD": 24,
		"NOK": 25,
		"PLN": 26,
		"RON": 27,
		"XDR": 28,
		"SGD": 29,
		"TJS": 30,
		"THB": 31,
		"TRY": 32,
		"TMT": 33,
		"UZS": 34,
		"UAH": 35,
		"CZK": 36,
		"SEK": 37,
		"CHF": 38,
		"RSD": 39,
		"ZAR": 40,
		"KRW": 41,
		"JPY": 42,
	}
)

func (x Currencies) Enum() *Currencies {
	p := new(Currencies)
	*p = x
	return p
}

func (x Currencies) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Currencies) Descriptor() protoreflect.EnumDescriptor {
	return file_cur_rates_proto_enumTypes[0].Descriptor()
}

func (Currencies) Type() protoreflect.EnumType {
	return &file_cur_rates_proto_enumTypes[0]
}

func (x Currencies) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Currencies.Descriptor instead.
func (Currencies) EnumDescriptor() ([]byte, []int) {
	return file_cur_rates_proto_rawDescGZIP(), []int{0}
}

type RatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base Currencies `protobuf:"varint,1,opt,name=Base,proto3,enum=Currencies" json:"Base,omitempty"`
}

func (x *RatesRequest) Reset() {
	*x = RatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cur_rates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatesRequest) ProtoMessage() {}

func (x *RatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cur_rates_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatesRequest.ProtoReflect.Descriptor instead.
func (*RatesRequest) Descriptor() ([]byte, []int) {
	return file_cur_rates_proto_rawDescGZIP(), []int{0}
}

func (x *RatesRequest) GetBase() Currencies {
	if x != nil {
		return x.Base
	}
	return Currencies_AUD
}

type RatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base    Currencies `protobuf:"varint,1,opt,name=Base,proto3,enum=Currencies" json:"Base,omitempty"`
	Title   string     `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	NumCode int32      `protobuf:"varint,3,opt,name=NumCode,proto3" json:"NumCode,omitempty"`
	Rate    float64    `protobuf:"fixed64,4,opt,name=Rate,proto3" json:"Rate,omitempty"`
}

func (x *RatesResponse) Reset() {
	*x = RatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cur_rates_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatesResponse) ProtoMessage() {}

func (x *RatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cur_rates_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatesResponse.ProtoReflect.Descriptor instead.
func (*RatesResponse) Descriptor() ([]byte, []int) {
	return file_cur_rates_proto_rawDescGZIP(), []int{1}
}

func (x *RatesResponse) GetBase() Currencies {
	if x != nil {
		return x.Base
	}
	return Currencies_AUD
}

func (x *RatesResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RatesResponse) GetNumCode() int32 {
	if x != nil {
		return x.NumCode
	}
	return 0
}

func (x *RatesResponse) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

var File_cur_rates_proto protoreflect.FileDescriptor

var file_cur_rates_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x75, 0x72, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2f, 0x0a, 0x0c, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0b, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x04, 0x42, 0x61,
	0x73, 0x65, 0x22, 0x74, 0x0a, 0x0d, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0b, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x04,
	0x42, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x75,
	0x6d, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x4e, 0x75, 0x6d,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x52, 0x61, 0x74, 0x65, 0x2a, 0x8f, 0x03, 0x0a, 0x0a, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x55, 0x44, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x41, 0x5a, 0x4e, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x42, 0x50,
	0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4d, 0x44, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x42,
	0x59, 0x4e, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x47, 0x4e, 0x10, 0x05, 0x12, 0x07, 0x0a,
	0x03, 0x42, 0x52, 0x4c, 0x10, 0x06, 0x12, 0x07, 0x0a, 0x03, 0x48, 0x55, 0x46, 0x10, 0x07, 0x12,
	0x07, 0x0a, 0x03, 0x56, 0x4e, 0x44, 0x10, 0x08, 0x12, 0x07, 0x0a, 0x03, 0x48, 0x4b, 0x44, 0x10,
	0x09, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x4c, 0x10, 0x0a, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x4b,
	0x4b, 0x10, 0x0b, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x45, 0x44, 0x10, 0x0c, 0x12, 0x07, 0x0a, 0x03,
	0x55, 0x53, 0x44, 0x10, 0x0d, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x55, 0x52, 0x10, 0x0e, 0x12, 0x07,
	0x0a, 0x03, 0x45, 0x47, 0x50, 0x10, 0x0f, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4e, 0x52, 0x10, 0x10,
	0x12, 0x07, 0x0a, 0x03, 0x49, 0x44, 0x52, 0x10, 0x11, 0x12, 0x07, 0x0a, 0x03, 0x4b, 0x5a, 0x54,
	0x10, 0x12, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x41, 0x44, 0x10, 0x13, 0x12, 0x07, 0x0a, 0x03, 0x51,
	0x41, 0x52, 0x10, 0x14, 0x12, 0x07, 0x0a, 0x03, 0x4b, 0x47, 0x53, 0x10, 0x15, 0x12, 0x07, 0x0a,
	0x03, 0x43, 0x4e, 0x59, 0x10, 0x16, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x44, 0x4c, 0x10, 0x17, 0x12,
	0x07, 0x0a, 0x03, 0x4e, 0x5a, 0x44, 0x10, 0x18, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x4f, 0x4b, 0x10,
	0x19, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x4c, 0x4e, 0x10, 0x1a, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x4f,
	0x4e, 0x10, 0x1b, 0x12, 0x07, 0x0a, 0x03, 0x58, 0x44, 0x52, 0x10, 0x1c, 0x12, 0x07, 0x0a, 0x03,
	0x53, 0x47, 0x44, 0x10, 0x1d, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x4a, 0x53, 0x10, 0x1e, 0x12, 0x07,
	0x0a, 0x03, 0x54, 0x48, 0x42, 0x10, 0x1f, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x52, 0x59, 0x10, 0x20,
	0x12, 0x07, 0x0a, 0x03, 0x54, 0x4d, 0x54, 0x10, 0x21, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x5a, 0x53,
	0x10, 0x22, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x41, 0x48, 0x10, 0x23, 0x12, 0x07, 0x0a, 0x03, 0x43,
	0x5a, 0x4b, 0x10, 0x24, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x45, 0x4b, 0x10, 0x25, 0x12, 0x07, 0x0a,
	0x03, 0x43, 0x48, 0x46, 0x10, 0x26, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x53, 0x44, 0x10, 0x27, 0x12,
	0x07, 0x0a, 0x03, 0x5a, 0x41, 0x52, 0x10, 0x28, 0x12, 0x07, 0x0a, 0x03, 0x4b, 0x52, 0x57, 0x10,
	0x29, 0x12, 0x07, 0x0a, 0x03, 0x4a, 0x50, 0x59, 0x10, 0x2a, 0x32, 0x74, 0x0a, 0x0d, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12,
	0x0d, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x0d, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0e, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cur_rates_proto_rawDescOnce sync.Once
	file_cur_rates_proto_rawDescData = file_cur_rates_proto_rawDesc
)

func file_cur_rates_proto_rawDescGZIP() []byte {
	file_cur_rates_proto_rawDescOnce.Do(func() {
		file_cur_rates_proto_rawDescData = protoimpl.X.CompressGZIP(file_cur_rates_proto_rawDescData)
	})
	return file_cur_rates_proto_rawDescData
}

var file_cur_rates_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cur_rates_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cur_rates_proto_goTypes = []interface{}{
	(Currencies)(0),       // 0: Currencies
	(*RatesRequest)(nil),  // 1: RatesRequest
	(*RatesResponse)(nil), // 2: RatesResponse
}
var file_cur_rates_proto_depIdxs = []int32{
	0, // 0: RatesRequest.Base:type_name -> Currencies
	0, // 1: RatesResponse.Base:type_name -> Currencies
	1, // 2: CurrencyRates.GetCurrencyRates:input_type -> RatesRequest
	1, // 3: CurrencyRates.GetCurrencyRate:input_type -> RatesRequest
	2, // 4: CurrencyRates.GetCurrencyRates:output_type -> RatesResponse
	2, // 5: CurrencyRates.GetCurrencyRate:output_type -> RatesResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cur_rates_proto_init() }
func file_cur_rates_proto_init() {
	if File_cur_rates_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cur_rates_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatesRequest); i {
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
		file_cur_rates_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatesResponse); i {
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
			RawDescriptor: file_cur_rates_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cur_rates_proto_goTypes,
		DependencyIndexes: file_cur_rates_proto_depIdxs,
		EnumInfos:         file_cur_rates_proto_enumTypes,
		MessageInfos:      file_cur_rates_proto_msgTypes,
	}.Build()
	File_cur_rates_proto = out.File
	file_cur_rates_proto_rawDesc = nil
	file_cur_rates_proto_goTypes = nil
	file_cur_rates_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CurrencyRatesClient is the client API for CurrencyRates service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CurrencyRatesClient interface {
	GetCurrencyRates(ctx context.Context, in *RatesRequest, opts ...grpc.CallOption) (*RatesResponse, error)
	GetCurrencyRate(ctx context.Context, in *RatesRequest, opts ...grpc.CallOption) (*RatesResponse, error)
}

type currencyRatesClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyRatesClient(cc grpc.ClientConnInterface) CurrencyRatesClient {
	return &currencyRatesClient{cc}
}

func (c *currencyRatesClient) GetCurrencyRates(ctx context.Context, in *RatesRequest, opts ...grpc.CallOption) (*RatesResponse, error) {
	out := new(RatesResponse)
	err := c.cc.Invoke(ctx, "/CurrencyRates/GetCurrencyRates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyRatesClient) GetCurrencyRate(ctx context.Context, in *RatesRequest, opts ...grpc.CallOption) (*RatesResponse, error) {
	out := new(RatesResponse)
	err := c.cc.Invoke(ctx, "/CurrencyRates/GetCurrencyRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrencyRatesServer is the server API for CurrencyRates service.
type CurrencyRatesServer interface {
	GetCurrencyRates(context.Context, *RatesRequest) (*RatesResponse, error)
	GetCurrencyRate(context.Context, *RatesRequest) (*RatesResponse, error)
}

// UnimplementedCurrencyRatesServer can be embedded to have forward compatible implementations.
type UnimplementedCurrencyRatesServer struct {
}

func (*UnimplementedCurrencyRatesServer) GetCurrencyRates(context.Context, *RatesRequest) (*RatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrencyRates not implemented")
}
func (*UnimplementedCurrencyRatesServer) GetCurrencyRate(context.Context, *RatesRequest) (*RatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrencyRate not implemented")
}

func RegisterCurrencyRatesServer(s *grpc.Server, srv CurrencyRatesServer) {
	s.RegisterService(&_CurrencyRates_serviceDesc, srv)
}

func _CurrencyRates_GetCurrencyRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyRatesServer).GetCurrencyRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CurrencyRates/GetCurrencyRates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyRatesServer).GetCurrencyRates(ctx, req.(*RatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyRates_GetCurrencyRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyRatesServer).GetCurrencyRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CurrencyRates/GetCurrencyRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyRatesServer).GetCurrencyRate(ctx, req.(*RatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CurrencyRates_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CurrencyRates",
	HandlerType: (*CurrencyRatesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrencyRates",
			Handler:    _CurrencyRates_GetCurrencyRates_Handler,
		},
		{
			MethodName: "GetCurrencyRate",
			Handler:    _CurrencyRates_GetCurrencyRate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cur_rates.proto",
}
