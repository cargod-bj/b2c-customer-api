// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.14.0
// source: verify/verify.proto

package verify

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

type VerifyDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Phone        string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	VerifyCode   string `protobuf:"bytes,3,opt,name=verifyCode,proto3" json:"verifyCode,omitempty"`
	CreateTime   uint64 `protobuf:"varint,4,opt,name=createTime,proto3" json:"createTime,omitempty"`
	RequestCount uint32 `protobuf:"varint,5,opt,name=requestCount,proto3" json:"requestCount,omitempty"`
	UpdateTime   uint64 `protobuf:"varint,6,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *VerifyDto) Reset() {
	*x = VerifyDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verify_verify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyDto) ProtoMessage() {}

func (x *VerifyDto) ProtoReflect() protoreflect.Message {
	mi := &file_verify_verify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyDto.ProtoReflect.Descriptor instead.
func (*VerifyDto) Descriptor() ([]byte, []int) {
	return file_verify_verify_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VerifyDto) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *VerifyDto) GetVerifyCode() string {
	if x != nil {
		return x.VerifyCode
	}
	return ""
}

func (x *VerifyDto) GetCreateTime() uint64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *VerifyDto) GetRequestCount() uint32 {
	if x != nil {
		return x.RequestCount
	}
	return 0
}

func (x *VerifyDto) GetUpdateTime() uint64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

type LoginDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone   string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Captcha string `protobuf:"bytes,2,opt,name=captcha,proto3" json:"captcha,omitempty"`
}

func (x *LoginDto) Reset() {
	*x = LoginDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verify_verify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginDto) ProtoMessage() {}

func (x *LoginDto) ProtoReflect() protoreflect.Message {
	mi := &file_verify_verify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginDto.ProtoReflect.Descriptor instead.
func (*LoginDto) Descriptor() ([]byte, []int) {
	return file_verify_verify_proto_rawDescGZIP(), []int{1}
}

func (x *LoginDto) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *LoginDto) GetCaptcha() string {
	if x != nil {
		return x.Captcha
	}
	return ""
}

type SmsDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From  string `protobuf:"bytes,1,opt,name=From,proto3" json:"From,omitempty"`
	Phone string `protobuf:"bytes,2,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Text  string `protobuf:"bytes,3,opt,name=Text,proto3" json:"Text,omitempty"`
}

func (x *SmsDto) Reset() {
	*x = SmsDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verify_verify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmsDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsDto) ProtoMessage() {}

func (x *SmsDto) ProtoReflect() protoreflect.Message {
	mi := &file_verify_verify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsDto.ProtoReflect.Descriptor instead.
func (*SmsDto) Descriptor() ([]byte, []int) {
	return file_verify_verify_proto_rawDescGZIP(), []int{2}
}

func (x *SmsDto) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SmsDto) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SmsDto) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_verify_verify_proto protoreflect.FileDescriptor

var file_verify_verify_proto_rawDesc = []byte{
	0x0a, 0x13, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x1a, 0x39, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d,
	0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x09, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x3a, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x22, 0x46, 0x0a, 0x06, 0x53,
	0x6d, 0x73, 0x44, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54,
	0x65, 0x78, 0x74, 0x32, 0xda, 0x01, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x33,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x0f,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x74, 0x6f, 0x1a,
	0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x42, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x10, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x44, 0x74, 0x6f, 0x1a, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x2c, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x4d, 0x53, 0x12, 0x0d, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x53, 0x6d, 0x73, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_verify_verify_proto_rawDescOnce sync.Once
	file_verify_verify_proto_rawDescData = file_verify_verify_proto_rawDesc
)

func file_verify_verify_proto_rawDescGZIP() []byte {
	file_verify_verify_proto_rawDescOnce.Do(func() {
		file_verify_verify_proto_rawDescData = protoimpl.X.CompressGZIP(file_verify_verify_proto_rawDescData)
	})
	return file_verify_verify_proto_rawDescData
}

var file_verify_verify_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_verify_verify_proto_goTypes = []interface{}{
	(*VerifyDto)(nil),       // 0: login.VerifyDto
	(*LoginDto)(nil),        // 1: login.LoginDto
	(*SmsDto)(nil),          // 2: login.SmsDto
	(*common.Response)(nil), // 3: common.Response
}
var file_verify_verify_proto_depIdxs = []int32{
	1, // 0: login.Verify.CreateVerify:input_type -> login.LoginDto
	1, // 1: login.Verify.GetVerifyByPhone:input_type -> login.LoginDto
	0, // 2: login.Verify.UpdateVerify:input_type -> login.VerifyDto
	2, // 3: login.Verify.SendSMS:input_type -> login.SmsDto
	3, // 4: login.Verify.CreateVerify:output_type -> common.Response
	3, // 5: login.Verify.GetVerifyByPhone:output_type -> common.Response
	3, // 6: login.Verify.UpdateVerify:output_type -> common.Response
	3, // 7: login.Verify.SendSMS:output_type -> common.Response
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_verify_verify_proto_init() }
func file_verify_verify_proto_init() {
	if File_verify_verify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_verify_verify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyDto); i {
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
		file_verify_verify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginDto); i {
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
		file_verify_verify_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmsDto); i {
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
			RawDescriptor: file_verify_verify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_verify_verify_proto_goTypes,
		DependencyIndexes: file_verify_verify_proto_depIdxs,
		MessageInfos:      file_verify_verify_proto_msgTypes,
	}.Build()
	File_verify_verify_proto = out.File
	file_verify_verify_proto_rawDesc = nil
	file_verify_verify_proto_goTypes = nil
	file_verify_verify_proto_depIdxs = nil
}
