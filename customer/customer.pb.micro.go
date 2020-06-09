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
	//获取客户列表，返回客户列表
	GetList(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error)
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

func (c *customerService) GetList(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Customer.GetList", in)
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
	//获取客户列表，返回客户列表
	GetList(context.Context, *common.Page, *common.Response) error
}

func RegisterCustomerHandler(s server.Server, hdlr CustomerHandler, opts ...server.HandlerOption) error {
	type customer interface {
		Add(ctx context.Context, in *CustomerDTO, out *common.Response) error
		Delete(ctx context.Context, in *DeleteId, out *common.Response) error
		Update(ctx context.Context, in *CustomerDTO, out *common.Response) error
		GetList(ctx context.Context, in *common.Page, out *common.Response) error
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

func (h *customerHandler) GetList(ctx context.Context, in *common.Page, out *common.Response) error {
	return h.CustomerHandler.GetList(ctx, in, out)
}
