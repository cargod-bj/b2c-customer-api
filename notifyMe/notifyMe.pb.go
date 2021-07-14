// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.4
// source: notifyMe/notifyMe.proto

package notifyMe

import (
	common "github.com/cargod-bj/b2c-proto-common/common"
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

type NotifyMeDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email        string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Color        string `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
	Keywords     string `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords,omitempty"`
	MillageForm  string `protobuf:"bytes,4,opt,name=millageForm,proto3" json:"millageForm,omitempty"`
	MillageTo    string `protobuf:"bytes,5,opt,name=millageTo,proto3" json:"millageTo,omitempty"`
	ModelIds     string `protobuf:"bytes,6,opt,name=modelIds,proto3" json:"modelIds,omitempty"`
	PriceFrom    string `protobuf:"bytes,7,opt,name=priceFrom,proto3" json:"priceFrom,omitempty"`
	PriceTo      string `protobuf:"bytes,8,opt,name=priceTo,proto3" json:"priceTo,omitempty"`
	StoreIds     string `protobuf:"bytes,9,opt,name=storeIds,proto3" json:"storeIds,omitempty"`
	Transmission string `protobuf:"bytes,10,opt,name=transmission,proto3" json:"transmission,omitempty"`
	YearFrom     string `protobuf:"bytes,11,opt,name=yearFrom,proto3" json:"yearFrom,omitempty"`
	YearTo       string `protobuf:"bytes,12,opt,name=yearTo,proto3" json:"yearTo,omitempty"`
}

func (x *NotifyMeDTO) Reset() {
	*x = NotifyMeDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifyMe_notifyMe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyMeDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyMeDTO) ProtoMessage() {}

func (x *NotifyMeDTO) ProtoReflect() protoreflect.Message {
	mi := &file_notifyMe_notifyMe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyMeDTO.ProtoReflect.Descriptor instead.
func (*NotifyMeDTO) Descriptor() ([]byte, []int) {
	return file_notifyMe_notifyMe_proto_rawDescGZIP(), []int{0}
}

func (x *NotifyMeDTO) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *NotifyMeDTO) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *NotifyMeDTO) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *NotifyMeDTO) GetMillageForm() string {
	if x != nil {
		return x.MillageForm
	}
	return ""
}

func (x *NotifyMeDTO) GetMillageTo() string {
	if x != nil {
		return x.MillageTo
	}
	return ""
}

func (x *NotifyMeDTO) GetModelIds() string {
	if x != nil {
		return x.ModelIds
	}
	return ""
}

func (x *NotifyMeDTO) GetPriceFrom() string {
	if x != nil {
		return x.PriceFrom
	}
	return ""
}

func (x *NotifyMeDTO) GetPriceTo() string {
	if x != nil {
		return x.PriceTo
	}
	return ""
}

func (x *NotifyMeDTO) GetStoreIds() string {
	if x != nil {
		return x.StoreIds
	}
	return ""
}

func (x *NotifyMeDTO) GetTransmission() string {
	if x != nil {
		return x.Transmission
	}
	return ""
}

func (x *NotifyMeDTO) GetYearFrom() string {
	if x != nil {
		return x.YearFrom
	}
	return ""
}

func (x *NotifyMeDTO) GetYearTo() string {
	if x != nil {
		return x.YearTo
	}
	return ""
}

var File_notifyMe_notifyMe_proto protoreflect.FileDescriptor

var file_notifyMe_notifyMe_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4d, 0x65, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x4d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x4d, 0x65, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd,
	0x02, 0x0a, 0x0b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4d, 0x65, 0x44, 0x54, 0x4f, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x69, 0x6c, 0x6c, 0x61, 0x67,
	0x65, 0x46, 0x6f, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x69, 0x6c,
	0x6c, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x69, 0x6c, 0x6c,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x69, 0x6c,
	0x6c, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49,
	0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49,
	0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x46, 0x72, 0x6f, 0x6d,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x49, 0x64, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x49, 0x64, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x79, 0x65,
	0x61, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x79, 0x65,
	0x61, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x79, 0x65, 0x61, 0x72, 0x54, 0x6f,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x79, 0x65, 0x61, 0x72, 0x54, 0x6f, 0x32, 0x3c,
	0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4d, 0x65, 0x12, 0x30, 0x0a, 0x03, 0x41, 0x64,
	0x64, 0x12, 0x15, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4d, 0x65, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x4d, 0x65, 0x44, 0x54, 0x4f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f,
	0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4d, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notifyMe_notifyMe_proto_rawDescOnce sync.Once
	file_notifyMe_notifyMe_proto_rawDescData = file_notifyMe_notifyMe_proto_rawDesc
)

func file_notifyMe_notifyMe_proto_rawDescGZIP() []byte {
	file_notifyMe_notifyMe_proto_rawDescOnce.Do(func() {
		file_notifyMe_notifyMe_proto_rawDescData = protoimpl.X.CompressGZIP(file_notifyMe_notifyMe_proto_rawDescData)
	})
	return file_notifyMe_notifyMe_proto_rawDescData
}

var file_notifyMe_notifyMe_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_notifyMe_notifyMe_proto_goTypes = []interface{}{
	(*NotifyMeDTO)(nil),     // 0: notifyMe.NotifyMeDTO
	(*common.Response)(nil), // 1: common.Response
}
var file_notifyMe_notifyMe_proto_depIdxs = []int32{
	0, // 0: notifyMe.NotifyMe.Add:input_type -> notifyMe.NotifyMeDTO
	1, // 1: notifyMe.NotifyMe.Add:output_type -> common.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notifyMe_notifyMe_proto_init() }
func file_notifyMe_notifyMe_proto_init() {
	if File_notifyMe_notifyMe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notifyMe_notifyMe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyMeDTO); i {
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
			RawDescriptor: file_notifyMe_notifyMe_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notifyMe_notifyMe_proto_goTypes,
		DependencyIndexes: file_notifyMe_notifyMe_proto_depIdxs,
		MessageInfos:      file_notifyMe_notifyMe_proto_msgTypes,
	}.Build()
	File_notifyMe_notifyMe_proto = out.File
	file_notifyMe_notifyMe_proto_rawDesc = nil
	file_notifyMe_notifyMe_proto_goTypes = nil
	file_notifyMe_notifyMe_proto_depIdxs = nil
}