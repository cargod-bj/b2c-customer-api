// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: verify/verify.proto

package verify

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

// Api Endpoints for Verify service

func NewVerifyEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Verify service

type VerifyService interface {
	//创建验证码，返回data.nil
	CreateVerify(ctx context.Context, in *LoginDto, opts ...client.CallOption) (*common.Response, error)
	//通过phone拿到验证码
	GetVerifyByPhone(ctx context.Context, in *LoginDto, opts ...client.CallOption) (*common.Response, error)
	//更改验证码信息
	UpdateVerify(ctx context.Context, in *VerifyDto, opts ...client.CallOption) (*common.Response, error)
}

type verifyService struct {
	c    client.Client
	name string
}

func NewVerifyService(name string, c client.Client) VerifyService {
	return &verifyService{
		c:    c,
		name: name,
	}
}

func (c *verifyService) CreateVerify(ctx context.Context, in *LoginDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Verify.CreateVerify", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifyService) GetVerifyByPhone(ctx context.Context, in *LoginDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Verify.GetVerifyByPhone", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifyService) UpdateVerify(ctx context.Context, in *VerifyDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Verify.UpdateVerify", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Verify service

type VerifyHandler interface {
	//创建验证码，返回data.nil
	CreateVerify(context.Context, *LoginDto, *common.Response) error
	//通过phone拿到验证码
	GetVerifyByPhone(context.Context, *LoginDto, *common.Response) error
	//更改验证码信息
	UpdateVerify(context.Context, *VerifyDto, *common.Response) error
}

func RegisterVerifyHandler(s server.Server, hdlr VerifyHandler, opts ...server.HandlerOption) error {
	type verify interface {
		CreateVerify(ctx context.Context, in *LoginDto, out *common.Response) error
		GetVerifyByPhone(ctx context.Context, in *LoginDto, out *common.Response) error
		UpdateVerify(ctx context.Context, in *VerifyDto, out *common.Response) error
	}
	type Verify struct {
		verify
	}
	h := &verifyHandler{hdlr}
	return s.Handle(s.NewHandler(&Verify{h}, opts...))
}

type verifyHandler struct {
	VerifyHandler
}

func (h *verifyHandler) CreateVerify(ctx context.Context, in *LoginDto, out *common.Response) error {
	return h.VerifyHandler.CreateVerify(ctx, in, out)
}

func (h *verifyHandler) GetVerifyByPhone(ctx context.Context, in *LoginDto, out *common.Response) error {
	return h.VerifyHandler.GetVerifyByPhone(ctx, in, out)
}

func (h *verifyHandler) UpdateVerify(ctx context.Context, in *VerifyDto, out *common.Response) error {
	return h.VerifyHandler.UpdateVerify(ctx, in, out)
}
