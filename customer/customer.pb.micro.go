// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: customer/customer.proto

package customer

import (
	fmt "fmt"
	common "github.com/cargod-bj/b2c-proto-common/common"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Customer service

func NewCustomerEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Customer service

type CustomerService interface {
	//新增客户，返回data.nil
	Add(ctx context.Context, in *CustomerDTO, opts ...client.CallOption) (*common.Response, error)
	//删除客户，这里是假删除，设置status为1表示已删除，返回data.nil
	Delete(ctx context.Context, in *DeleteId, opts ...client.CallOption) (*common.Response, error)
	//更新客户，返回data.nil
	Update(ctx context.Context, in *CustomerDTO, opts ...client.CallOption) (*common.Response, error)
	//获取客户根据登录名(手机号)
	GetCustomer(ctx context.Context, in *CustomerDTO, opts ...client.CallOption) (*common.Response, error)
	//根据分页信息+查询条件查询客户信息
	GetList(ctx context.Context, in *CustCondDto, opts ...client.CallOption) (*common.Response, error)
	//根据入参条件模糊搜索用户信息 不分页
	GetCustomerByCond(ctx context.Context, in *CustomerDTO, opts ...client.CallOption) (*common.Response, error)
	//根据 userid批量获取用户信息
	GetCustomerList(ctx context.Context, in *IdListDto, opts ...client.CallOption) (*common.Response, error)
	//批量分配用户到PIC
	AssignCustomer(ctx context.Context, in *AssignDto, opts ...client.CallOption) (*common.Response, error)
	//根据手机号精准搜索客户信息 分页
	GetCustomerByContactNo(ctx context.Context, in *CustCondDto, opts ...client.CallOption) (*common.Response, error)
	//根据userId搜索客户信息
	GetCustomerByIds(ctx context.Context, in *IdList, opts ...client.CallOption) (*common.Response, error)
	//回收邮件
	GetMail(ctx context.Context, in *DeleteId, opts ...client.CallOption) (*common.Response, error)
	//拉取C2B拓客列表(CMS前端展示用)
	GetC2BCustomerList(ctx context.Context, in *C2BCondDto, opts ...client.CallOption) (*common.Response, error)
	//存储营销客户相关验证数据
	AddMarketVerifyData(ctx context.Context, in *MarketVerifyData, opts ...client.CallOption) (*common.Response, error)
	//定时任务拉取营销客户数据
	ScheduleMarketingUser(ctx context.Context, in *MarketingCondDto, opts ...client.CallOption) (*common.Response, error)
	//更新客户营销记录表
	UpdateMarketingRecord(ctx context.Context, in *MarketingRecordDto, opts ...client.CallOption) (*common.Response, error)
	//通过门店和customerId获取人员的PIC
	GetCustomerPIC(ctx context.Context, in *QueryPicDTO, opts ...client.CallOption) (*common.Response, error)
	//website更新用户信息
	UpdateCustomerInfo(ctx context.Context, in *CustomerUpdateDto, opts ...client.CallOption) (*common.Response, error)
	//根据用户ID获取用户信息
	GetCustomerById(ctx context.Context, in *GetCustomerByIdDto, opts ...client.CallOption) (*common.Response, error)
	//  my customer
	GetMyCustomer(ctx context.Context, in *GetMyCustomerCond, opts ...client.CallOption) (*common.Response, error)
	ListCustomerByConditon(ctx context.Context, in *CustomerCondition, opts ...client.CallOption) (*common.Response, error)
}

type customerService struct {
	c    client.Client
	name string
}

func NewCustomerService(name string, c client.Client) CustomerService {
	return &customerService{
		c:    c,
		name: name,
	}
}

func (c *customerService) Add(ctx context.Context, in *CustomerDTO, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Customer.Add", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) Delete(ctx context.Context, in *DeleteId, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Customer.Delete", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) Update(ctx context.Context, in *CustomerDTO, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Customer.Update", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetCustomer(ctx context.Context, in *CustomerDTO, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Customer.GetCustomer", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetList(ctx context.Context, in *CustCondDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Customer.GetList", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetCustomerByCond(ctx context.Context, in *CustomerDTO, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Customer.GetCustomerByCond", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetCustomerList(ctx context.Context, in *IdListDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Customer.GetCustomerList", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) AssignCustomer(ctx context.Context, in *AssignDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Customer.AssignCustomer", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetCustomerByContactNo(ctx context.Context, in *CustCondDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Customer.GetCustomerByContactNo", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetCustomerByIds(ctx context.Context, in *IdList, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Customer.GetCustomerByIds", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetMail(ctx context.Context, in *DeleteId, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Customer.GetMail", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetC2BCustomerList(ctx context.Context, in *C2BCondDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Customer.GetC2BCustomerList", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) AddMarketVerifyData(ctx context.Context, in *MarketVerifyData, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Customer.AddMarketVerifyData", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) ScheduleMarketingUser(ctx context.Context, in *MarketingCondDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Customer.ScheduleMarketingUser", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) UpdateMarketingRecord(ctx context.Context, in *MarketingRecordDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Customer.UpdateMarketingRecord", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetCustomerPIC(ctx context.Context, in *QueryPicDTO, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Customer.GetCustomerPIC", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) UpdateCustomerInfo(ctx context.Context, in *CustomerUpdateDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Customer.UpdateCustomerInfo", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetCustomerById(ctx context.Context, in *GetCustomerByIdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Customer.GetCustomerById", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) GetMyCustomer(ctx context.Context, in *GetMyCustomerCond, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Customer.GetMyCustomer", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerService) ListCustomerByConditon(ctx context.Context, in *CustomerCondition, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Customer.ListCustomerByConditon", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Customer service

type CustomerHandler interface {
	//新增客户，返回data.nil
	Add(context.Context, *CustomerDTO, *common.Response) error
	//删除客户，这里是假删除，设置status为1表示已删除，返回data.nil
	Delete(context.Context, *DeleteId, *common.Response) error
	//更新客户，返回data.nil
	Update(context.Context, *CustomerDTO, *common.Response) error
	//获取客户根据登录名(手机号)
	GetCustomer(context.Context, *CustomerDTO, *common.Response) error
	//根据分页信息+查询条件查询客户信息
	GetList(context.Context, *CustCondDto, *common.Response) error
	//根据入参条件模糊搜索用户信息 不分页
	GetCustomerByCond(context.Context, *CustomerDTO, *common.Response) error
	//根据 userid批量获取用户信息
	GetCustomerList(context.Context, *IdListDto, *common.Response) error
	//批量分配用户到PIC
	AssignCustomer(context.Context, *AssignDto, *common.Response) error
	//根据手机号精准搜索客户信息 分页
	GetCustomerByContactNo(context.Context, *CustCondDto, *common.Response) error
	//根据userId搜索客户信息
	GetCustomerByIds(context.Context, *IdList, *common.Response) error
	//回收邮件
	GetMail(context.Context, *DeleteId, *common.Response) error
	//拉取C2B拓客列表(CMS前端展示用)
	GetC2BCustomerList(context.Context, *C2BCondDto, *common.Response) error
	//存储营销客户相关验证数据
	AddMarketVerifyData(context.Context, *MarketVerifyData, *common.Response) error
	//定时任务拉取营销客户数据
	ScheduleMarketingUser(context.Context, *MarketingCondDto, *common.Response) error
	//更新客户营销记录表
	UpdateMarketingRecord(context.Context, *MarketingRecordDto, *common.Response) error
	//通过门店和customerId获取人员的PIC
	GetCustomerPIC(context.Context, *QueryPicDTO, *common.Response) error
	//website更新用户信息
	UpdateCustomerInfo(context.Context, *CustomerUpdateDto, *common.Response) error
	//根据用户ID获取用户信息
	GetCustomerById(context.Context, *GetCustomerByIdDto, *common.Response) error
	//  my customer
	GetMyCustomer(context.Context, *GetMyCustomerCond, *common.Response) error
	ListCustomerByConditon(context.Context, *CustomerCondition, *common.Response) error
}

func RegisterCustomerHandler(s server.Server, hdlr CustomerHandler, opts ...server.HandlerOption) error {
	type customer interface {
		Add(ctx context.Context, in *CustomerDTO, out *common.Response) error
		Delete(ctx context.Context, in *DeleteId, out *common.Response) error
		Update(ctx context.Context, in *CustomerDTO, out *common.Response) error
		GetCustomer(ctx context.Context, in *CustomerDTO, out *common.Response) error
		GetList(ctx context.Context, in *CustCondDto, out *common.Response) error
		GetCustomerByCond(ctx context.Context, in *CustomerDTO, out *common.Response) error
		GetCustomerList(ctx context.Context, in *IdListDto, out *common.Response) error
		AssignCustomer(ctx context.Context, in *AssignDto, out *common.Response) error
		GetCustomerByContactNo(ctx context.Context, in *CustCondDto, out *common.Response) error
		GetCustomerByIds(ctx context.Context, in *IdList, out *common.Response) error
		GetMail(ctx context.Context, in *DeleteId, out *common.Response) error
		GetC2BCustomerList(ctx context.Context, in *C2BCondDto, out *common.Response) error
		AddMarketVerifyData(ctx context.Context, in *MarketVerifyData, out *common.Response) error
		ScheduleMarketingUser(ctx context.Context, in *MarketingCondDto, out *common.Response) error
		UpdateMarketingRecord(ctx context.Context, in *MarketingRecordDto, out *common.Response) error
		GetCustomerPIC(ctx context.Context, in *QueryPicDTO, out *common.Response) error
		UpdateCustomerInfo(ctx context.Context, in *CustomerUpdateDto, out *common.Response) error
		GetCustomerById(ctx context.Context, in *GetCustomerByIdDto, out *common.Response) error
		GetMyCustomer(ctx context.Context, in *GetMyCustomerCond, out *common.Response) error
		ListCustomerByConditon(ctx context.Context, in *CustomerCondition, out *common.Response) error
	}
	type Customer struct {
		customer
	}
	h := &customerHandler{hdlr}
	return s.Handle(s.NewHandler(&Customer{h}, opts...))
}

type customerHandler struct {
	CustomerHandler
}

func (h *customerHandler) Add(ctx context.Context, in *CustomerDTO, out *common.Response) error {
	return h.CustomerHandler.Add(ctx, in, out)
}

func (h *customerHandler) Delete(ctx context.Context, in *DeleteId, out *common.Response) error {
	return h.CustomerHandler.Delete(ctx, in, out)
}

func (h *customerHandler) Update(ctx context.Context, in *CustomerDTO, out *common.Response) error {
	return h.CustomerHandler.Update(ctx, in, out)
}

func (h *customerHandler) GetCustomer(ctx context.Context, in *CustomerDTO, out *common.Response) error {
	return h.CustomerHandler.GetCustomer(ctx, in, out)
}

func (h *customerHandler) GetList(ctx context.Context, in *CustCondDto, out *common.Response) error {
	return h.CustomerHandler.GetList(ctx, in, out)
}

func (h *customerHandler) GetCustomerByCond(ctx context.Context, in *CustomerDTO, out *common.Response) error {
	return h.CustomerHandler.GetCustomerByCond(ctx, in, out)
}

func (h *customerHandler) GetCustomerList(ctx context.Context, in *IdListDto, out *common.Response) error {
	return h.CustomerHandler.GetCustomerList(ctx, in, out)
}

func (h *customerHandler) AssignCustomer(ctx context.Context, in *AssignDto, out *common.Response) error {
	return h.CustomerHandler.AssignCustomer(ctx, in, out)
}

func (h *customerHandler) GetCustomerByContactNo(ctx context.Context, in *CustCondDto, out *common.Response) error {
	return h.CustomerHandler.GetCustomerByContactNo(ctx, in, out)
}

func (h *customerHandler) GetCustomerByIds(ctx context.Context, in *IdList, out *common.Response) error {
	return h.CustomerHandler.GetCustomerByIds(ctx, in, out)
}

func (h *customerHandler) GetMail(ctx context.Context, in *DeleteId, out *common.Response) error {
	return h.CustomerHandler.GetMail(ctx, in, out)
}

func (h *customerHandler) GetC2BCustomerList(ctx context.Context, in *C2BCondDto, out *common.Response) error {
	return h.CustomerHandler.GetC2BCustomerList(ctx, in, out)
}

func (h *customerHandler) AddMarketVerifyData(ctx context.Context, in *MarketVerifyData, out *common.Response) error {
	return h.CustomerHandler.AddMarketVerifyData(ctx, in, out)
}

func (h *customerHandler) ScheduleMarketingUser(ctx context.Context, in *MarketingCondDto, out *common.Response) error {
	return h.CustomerHandler.ScheduleMarketingUser(ctx, in, out)
}

func (h *customerHandler) UpdateMarketingRecord(ctx context.Context, in *MarketingRecordDto, out *common.Response) error {
	return h.CustomerHandler.UpdateMarketingRecord(ctx, in, out)
}

func (h *customerHandler) GetCustomerPIC(ctx context.Context, in *QueryPicDTO, out *common.Response) error {
	return h.CustomerHandler.GetCustomerPIC(ctx, in, out)
}

func (h *customerHandler) UpdateCustomerInfo(ctx context.Context, in *CustomerUpdateDto, out *common.Response) error {
	return h.CustomerHandler.UpdateCustomerInfo(ctx, in, out)
}

func (h *customerHandler) GetCustomerById(ctx context.Context, in *GetCustomerByIdDto, out *common.Response) error {
	return h.CustomerHandler.GetCustomerById(ctx, in, out)
}

func (h *customerHandler) GetMyCustomer(ctx context.Context, in *GetMyCustomerCond, out *common.Response) error {
	return h.CustomerHandler.GetMyCustomer(ctx, in, out)
}

func (h *customerHandler) ListCustomerByConditon(ctx context.Context, in *CustomerCondition, out *common.Response) error {
	return h.CustomerHandler.ListCustomerByConditon(ctx, in, out)
}
