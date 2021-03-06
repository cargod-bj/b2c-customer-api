// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: marketingChannel/marketingChannel.proto

package marketingChannel

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

// Api Endpoints for MarketingChannel service

func NewMarketingChannelEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MarketingChannel service

type MarketingChannelService interface {
	//根据ID获取活动信息
	GetMarketingById(ctx context.Context, in *GetMarketingByIdDto, opts ...client.CallOption) (*common.Response, error)
}

type marketingChannelService struct {
	c    client.Client
	name string
}

func NewMarketingChannelService(name string, c client.Client) MarketingChannelService {
	return &marketingChannelService{
		c:    c,
		name: name,
	}
}

func (c *marketingChannelService) GetMarketingById(ctx context.Context, in *GetMarketingByIdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "MarketingChannel.GetMarketingById", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MarketingChannel service

type MarketingChannelHandler interface {
	//根据ID获取活动信息
	GetMarketingById(context.Context, *GetMarketingByIdDto, *common.Response) error
}

func RegisterMarketingChannelHandler(s server.Server, hdlr MarketingChannelHandler, opts ...server.HandlerOption) error {
	type marketingChannel interface {
		GetMarketingById(ctx context.Context, in *GetMarketingByIdDto, out *common.Response) error
	}
	type MarketingChannel struct {
		marketingChannel
	}
	h := &marketingChannelHandler{hdlr}
	return s.Handle(s.NewHandler(&MarketingChannel{h}, opts...))
}

type marketingChannelHandler struct {
	MarketingChannelHandler
}

func (h *marketingChannelHandler) GetMarketingById(ctx context.Context, in *GetMarketingByIdDto, out *common.Response) error {
	return h.MarketingChannelHandler.GetMarketingById(ctx, in, out)
}
