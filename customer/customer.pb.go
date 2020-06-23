// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: customer/customer.proto

package customer

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

type CustomerDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`           //用户姓名
	LoginName  string `protobuf:"bytes,3,opt,name=loginName,proto3" json:"loginName,omitempty"` //用户名
	Email      string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Icon       string `protobuf:"bytes,5,opt,name=icon,proto3" json:"icon,omitempty"`
	ContactNo  string `protobuf:"bytes,6,opt,name=ContactNo,proto3" json:"ContactNo,omitempty"`
	Password   string `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
	Gender     uint32 `protobuf:"varint,8,opt,name=gender,proto3" json:"gender,omitempty"` //性别1是女，2是男
	Birthday   string `protobuf:"bytes,9,opt,name=birthday,proto3" json:"birthday,omitempty"`
	CreateTime uint64 `protobuf:"varint,10,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime uint64 `protobuf:"varint,11,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	Source     string `protobuf:"bytes,12,opt,name=source,proto3" json:"source,omitempty"`
	CustomerId string `protobuf:"bytes,13,opt,name=customerId,proto3" json:"customerId,omitempty"`
}

func (x *CustomerDTO) Reset() {
	*x = CustomerDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerDTO) ProtoMessage() {}

func (x *CustomerDTO) ProtoReflect() protoreflect.Message {
	mi := &file_customer_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerDTO.ProtoReflect.Descriptor instead.
func (*CustomerDTO) Descriptor() ([]byte, []int) {
	return file_customer_customer_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerDTO) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CustomerDTO) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomerDTO) GetLoginName() string {
	if x != nil {
		return x.LoginName
	}
	return ""
}

func (x *CustomerDTO) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CustomerDTO) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *CustomerDTO) GetContactNo() string {
	if x != nil {
		return x.ContactNo
	}
	return ""
}

func (x *CustomerDTO) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CustomerDTO) GetGender() uint32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *CustomerDTO) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *CustomerDTO) GetCreateTime() uint64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *CustomerDTO) GetUpdateTime() uint64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *CustomerDTO) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *CustomerDTO) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type CustomerList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNo    uint32         `protobuf:"varint,1,opt,name=pageNo,proto3" json:"pageNo,omitempty"`
	PageSize  uint32         `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Customers []*CustomerDTO `protobuf:"bytes,3,rep,name=customers,proto3" json:"customers,omitempty"`
	Count     uint32         `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CustomerList) Reset() {
	*x = CustomerList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerList) ProtoMessage() {}

func (x *CustomerList) ProtoReflect() protoreflect.Message {
	mi := &file_customer_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerList.ProtoReflect.Descriptor instead.
func (*CustomerList) Descriptor() ([]byte, []int) {
	return file_customer_customer_proto_rawDescGZIP(), []int{1}
}

func (x *CustomerList) GetPageNo() uint32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *CustomerList) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CustomerList) GetCustomers() []*CustomerDTO {
	if x != nil {
		return x.Customers
	}
	return nil
}

func (x *CustomerList) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type AddCustomerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddCustomerData) Reset() {
	*x = AddCustomerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCustomerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCustomerData) ProtoMessage() {}

func (x *AddCustomerData) ProtoReflect() protoreflect.Message {
	mi := &file_customer_customer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCustomerData.ProtoReflect.Descriptor instead.
func (*AddCustomerData) Descriptor() ([]byte, []int) {
	return file_customer_customer_proto_rawDescGZIP(), []int{2}
}

func (x *AddCustomerData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteId) Reset() {
	*x = DeleteId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_customer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteId) ProtoMessage() {}

func (x *DeleteId) ProtoReflect() protoreflect.Message {
	mi := &file_customer_customer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteId.ProtoReflect.Descriptor instead.
func (*DeleteId) Descriptor() ([]byte, []int) {
	return file_customer_customer_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteId) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IdListDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdListDto) Reset() {
	*x = IdListDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_customer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdListDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdListDto) ProtoMessage() {}

func (x *IdListDto) ProtoReflect() protoreflect.Message {
	mi := &file_customer_customer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdListDto.ProtoReflect.Descriptor instead.
func (*IdListDto) Descriptor() ([]byte, []int) {
	return file_customer_customer_proto_rawDescGZIP(), []int{4}
}

func (x *IdListDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_customer_customer_proto protoreflect.FileDescriptor

var file_customer_customer_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf,
	0x02, 0x0a, 0x0b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x8d, 0x01, 0x0a, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x52,
	0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x21, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x1b, 0x0a, 0x09, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x32, 0x86, 0x03, 0x0a,
	0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x03, 0x41, 0x64, 0x64,
	0x12, 0x15, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x1a, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x12, 0x15, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x12, 0x15,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x44, 0x54, 0x4f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x74,
	0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32,
	0x63, 0x2d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customer_customer_proto_rawDescOnce sync.Once
	file_customer_customer_proto_rawDescData = file_customer_customer_proto_rawDesc
)

func file_customer_customer_proto_rawDescGZIP() []byte {
	file_customer_customer_proto_rawDescOnce.Do(func() {
		file_customer_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_customer_customer_proto_rawDescData)
	})
	return file_customer_customer_proto_rawDescData
}

var file_customer_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_customer_customer_proto_goTypes = []interface{}{
	(*CustomerDTO)(nil),     // 0: customer.CustomerDTO
	(*CustomerList)(nil),    // 1: customer.CustomerList
	(*AddCustomerData)(nil), // 2: customer.AddCustomerData
	(*DeleteId)(nil),        // 3: customer.DeleteId
	(*IdListDto)(nil),       // 4: customer.IdListDto
	(*common.Page)(nil),     // 5: common.Page
	(*common.Response)(nil), // 6: common.Response
}
var file_customer_customer_proto_depIdxs = []int32{
	0, // 0: customer.CustomerList.customers:type_name -> customer.CustomerDTO
	0, // 1: customer.Customer.Add:input_type -> customer.CustomerDTO
	3, // 2: customer.Customer.Delete:input_type -> customer.DeleteId
	0, // 3: customer.Customer.Update:input_type -> customer.CustomerDTO
	0, // 4: customer.Customer.GetCustomer:input_type -> customer.CustomerDTO
	5, // 5: customer.Customer.GetList:input_type -> common.Page
	0, // 6: customer.Customer.GetCustomerByCond:input_type -> customer.CustomerDTO
	4, // 7: customer.Customer.GetCustomerList:input_type -> customer.IdListDto
	6, // 8: customer.Customer.Add:output_type -> common.Response
	6, // 9: customer.Customer.Delete:output_type -> common.Response
	6, // 10: customer.Customer.Update:output_type -> common.Response
	6, // 11: customer.Customer.GetCustomer:output_type -> common.Response
	6, // 12: customer.Customer.GetList:output_type -> common.Response
	6, // 13: customer.Customer.GetCustomerByCond:output_type -> common.Response
	6, // 14: customer.Customer.GetCustomerList:output_type -> common.Response
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_customer_customer_proto_init() }
func file_customer_customer_proto_init() {
	if File_customer_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customer_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerDTO); i {
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
		file_customer_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerList); i {
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
		file_customer_customer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCustomerData); i {
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
		file_customer_customer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteId); i {
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
		file_customer_customer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdListDto); i {
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
			RawDescriptor: file_customer_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_customer_customer_proto_goTypes,
		DependencyIndexes: file_customer_customer_proto_depIdxs,
		MessageInfos:      file_customer_customer_proto_msgTypes,
	}.Build()
	File_customer_customer_proto = out.File
	file_customer_customer_proto_rawDesc = nil
	file_customer_customer_proto_goTypes = nil
	file_customer_customer_proto_depIdxs = nil
}
