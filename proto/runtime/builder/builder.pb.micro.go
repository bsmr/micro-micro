// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/micro/proto/runtime/builder/builder.proto

package builder

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v3/api"
	client "github.com/micro/go-micro/v3/client"
	server "github.com/micro/go-micro/v3/server"
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

// Api Endpoints for Builder service

func NewBuilderEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Builder service

type BuilderService interface {
	Build(ctx context.Context, opts ...client.CallOption) (Builder_BuildService, error)
}

type builderService struct {
	c    client.Client
	name string
}

func NewBuilderService(name string, c client.Client) BuilderService {
	return &builderService{
		c:    c,
		name: name,
	}
}

func (c *builderService) Build(ctx context.Context, opts ...client.CallOption) (Builder_BuildService, error) {
	req := c.c.NewRequest(c.name, "Builder.Build", &BuildRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &builderServiceBuild{stream}, nil
}

type Builder_BuildService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BuildRequest) error
	Recv() (*Result, error)
}

type builderServiceBuild struct {
	stream client.Stream
}

func (x *builderServiceBuild) Close() error {
	return x.stream.Close()
}

func (x *builderServiceBuild) Context() context.Context {
	return x.stream.Context()
}

func (x *builderServiceBuild) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *builderServiceBuild) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *builderServiceBuild) Send(m *BuildRequest) error {
	return x.stream.Send(m)
}

func (x *builderServiceBuild) Recv() (*Result, error) {
	m := new(Result)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Builder service

type BuilderHandler interface {
	Build(context.Context, Builder_BuildStream) error
}

func RegisterBuilderHandler(s server.Server, hdlr BuilderHandler, opts ...server.HandlerOption) error {
	type builder interface {
		Build(ctx context.Context, stream server.Stream) error
	}
	type Builder struct {
		builder
	}
	h := &builderHandler{hdlr}
	return s.Handle(s.NewHandler(&Builder{h}, opts...))
}

type builderHandler struct {
	BuilderHandler
}

func (h *builderHandler) Build(ctx context.Context, stream server.Stream) error {
	return h.BuilderHandler.Build(ctx, &builderBuildStream{stream})
}

type Builder_BuildStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Result) error
	Recv() (*BuildRequest, error)
}

type builderBuildStream struct {
	stream server.Stream
}

func (x *builderBuildStream) Close() error {
	return x.stream.Close()
}

func (x *builderBuildStream) Context() context.Context {
	return x.stream.Context()
}

func (x *builderBuildStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *builderBuildStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *builderBuildStream) Send(m *Result) error {
	return x.stream.Send(m)
}

func (x *builderBuildStream) Recv() (*BuildRequest, error) {
	m := new(BuildRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}