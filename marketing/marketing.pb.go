// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.14.0
// source: marketing/marketing.proto

package marketing

import (
	common "github.com/cargod-bj/b2c-proto-common/common"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetMarketingByIdDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMarketingByIdDto) Reset() {
	*x = GetMarketingByIdDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marketing_marketing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMarketingByIdDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMarketingByIdDto) ProtoMessage() {}

func (x *GetMarketingByIdDto) ProtoReflect() protoreflect.Message {
	mi := &file_marketing_marketing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMarketingByIdDto.ProtoReflect.Descriptor instead.
func (*GetMarketingByIdDto) Descriptor() ([]byte, []int) {
	return file_marketing_marketing_proto_rawDescGZIP(), []int{0}
}

func (x *GetMarketingByIdDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CustomerResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Status           uint64 `protobuf:"varint,3,opt,name=Status,proto3" json:"Status,omitempty"`
	IssuanceCoupon   uint64 `protobuf:"varint,4,opt,name=IssuanceCoupon,proto3" json:"IssuanceCoupon,omitempty"`
	CouponTemplateId uint64 `protobuf:"varint,5,opt,name=CouponTemplateId,proto3" json:"CouponTemplateId,omitempty"`
	ImageWeb         string `protobuf:"bytes,6,opt,name=ImageWeb,proto3" json:"ImageWeb,omitempty"`
	ImageH5          string `protobuf:"bytes,7,opt,name=ImageH5,proto3" json:"ImageH5,omitempty"`
	UtmSource        string `protobuf:"bytes,8,opt,name=UtmSource,proto3" json:"UtmSource,omitempty"`
	UtmMedium        string `protobuf:"bytes,9,opt,name=UtmMedium,proto3" json:"UtmMedium,omitempty"`
	UtmCampaign      string `protobuf:"bytes,10,opt,name=UtmCampaign,proto3" json:"UtmCampaign,omitempty"`
	UtmContent       string `protobuf:"bytes,11,opt,name=UtmContent,proto3" json:"UtmContent,omitempty"`
	UtmTerm          string `protobuf:"bytes,12,opt,name=UtmTerm,proto3" json:"UtmTerm,omitempty"`
	CreateTime       uint64 `protobuf:"varint,13,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	UpdateTime       uint64 `protobuf:"varint,14,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
}

func (x *CustomerResult) Reset() {
	*x = CustomerResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marketing_marketing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerResult) ProtoMessage() {}

func (x *CustomerResult) ProtoReflect() protoreflect.Message {
	mi := &file_marketing_marketing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerResult.ProtoReflect.Descriptor instead.
func (*CustomerResult) Descriptor() ([]byte, []int) {
	return file_marketing_marketing_proto_rawDescGZIP(), []int{1}
}

func (x *CustomerResult) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CustomerResult) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomerResult) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CustomerResult) GetIssuanceCoupon() uint64 {
	if x != nil {
		return x.IssuanceCoupon
	}
	return 0
}

func (x *CustomerResult) GetCouponTemplateId() uint64 {
	if x != nil {
		return x.CouponTemplateId
	}
	return 0
}

func (x *CustomerResult) GetImageWeb() string {
	if x != nil {
		return x.ImageWeb
	}
	return ""
}

func (x *CustomerResult) GetImageH5() string {
	if x != nil {
		return x.ImageH5
	}
	return ""
}

func (x *CustomerResult) GetUtmSource() string {
	if x != nil {
		return x.UtmSource
	}
	return ""
}

func (x *CustomerResult) GetUtmMedium() string {
	if x != nil {
		return x.UtmMedium
	}
	return ""
}

func (x *CustomerResult) GetUtmCampaign() string {
	if x != nil {
		return x.UtmCampaign
	}
	return ""
}

func (x *CustomerResult) GetUtmContent() string {
	if x != nil {
		return x.UtmContent
	}
	return ""
}

func (x *CustomerResult) GetUtmTerm() string {
	if x != nil {
		return x.UtmTerm
	}
	return ""
}

func (x *CustomerResult) GetCreateTime() uint64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *CustomerResult) GetUpdateTime() uint64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

var File_marketing_marketing_proto protoreflect.FileDescriptor

var file_marketing_marketing_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x42, 0x79, 0x49, 0x64, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0xae, 0x03, 0x0a, 0x0e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x49, 0x73, 0x73, 0x75, 0x61,
	0x6e, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0e, 0x49, 0x73, 0x73, 0x75, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12,
	0x2a, 0x0a, 0x10, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x43, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x57, 0x65, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x57, 0x65, 0x62, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x48, 0x35, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x48,
	0x35, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x74, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x74, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x55, 0x74, 0x6d, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x55, 0x74, 0x6d, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x12, 0x20, 0x0a,
	0x0b, 0x55, 0x74, 0x6d, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x55, 0x74, 0x6d, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12,
	0x1e, 0x0a, 0x0a, 0x55, 0x74, 0x6d, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x55, 0x74, 0x6d, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x55, 0x74, 0x6d, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x55, 0x74, 0x6d, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x52, 0x0a, 0x09, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x45, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x2e, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x42, 0x79, 0x49, 0x64, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31, 0x5a,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67,
	0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_marketing_marketing_proto_rawDescOnce sync.Once
	file_marketing_marketing_proto_rawDescData = file_marketing_marketing_proto_rawDesc
)

func file_marketing_marketing_proto_rawDescGZIP() []byte {
	file_marketing_marketing_proto_rawDescOnce.Do(func() {
		file_marketing_marketing_proto_rawDescData = protoimpl.X.CompressGZIP(file_marketing_marketing_proto_rawDescData)
	})
	return file_marketing_marketing_proto_rawDescData
}

var file_marketing_marketing_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_marketing_marketing_proto_goTypes = []interface{}{
	(*GetMarketingByIdDto)(nil), // 0: marketing.GetMarketingByIdDto
	(*CustomerResult)(nil),      // 1: marketing.CustomerResult
	(*common.Response)(nil),     // 2: common.Response
}
var file_marketing_marketing_proto_depIdxs = []int32{
	0, // 0: marketing.Marketing.GetCustomerById:input_type -> marketing.GetMarketingByIdDto
	2, // 1: marketing.Marketing.GetCustomerById:output_type -> common.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_marketing_marketing_proto_init() }
func file_marketing_marketing_proto_init() {
	if File_marketing_marketing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_marketing_marketing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMarketingByIdDto); i {
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
		file_marketing_marketing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerResult); i {
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
			RawDescriptor: file_marketing_marketing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_marketing_marketing_proto_goTypes,
		DependencyIndexes: file_marketing_marketing_proto_depIdxs,
		MessageInfos:      file_marketing_marketing_proto_msgTypes,
	}.Build()
	File_marketing_marketing_proto = out.File
	file_marketing_marketing_proto_rawDesc = nil
	file_marketing_marketing_proto_goTypes = nil
	file_marketing_marketing_proto_depIdxs = nil
}
