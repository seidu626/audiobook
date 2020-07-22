// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: services/language/proto/word/word_service.proto

package word

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/seidu626/audiobook/language/proto/entities"
	math "math"
)

import (
	context "context"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for WordService service

type WordService interface {
	Exist(ctx context.Context, in *ExistRequest, opts ...client.CallOption) (*ExistResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
}

type wordService struct {
	c    client.Client
	name string
}

func NewWordService(name string, c client.Client) WordService {
	return &wordService{
		c:    c,
		name: name,
	}
}

func (c *wordService) Exist(ctx context.Context, in *ExistRequest, opts ...client.CallOption) (*ExistResponse, error) {
	req := c.c.NewRequest(c.name, "WordService.Exist", in)
	out := new(ExistResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wordService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "WordService.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wordService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "WordService.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wordService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "WordService.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wordService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "WordService.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wordService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "WordService.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WordService service

type WordServiceHandler interface {
	Exist(context.Context, *ExistRequest, *ExistResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
	Get(context.Context, *GetRequest, *GetResponse) error
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
}

func RegisterWordServiceHandler(s server.Server, hdlr WordServiceHandler, opts ...server.HandlerOption) error {
	type wordService interface {
		Exist(ctx context.Context, in *ExistRequest, out *ExistResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
		Get(ctx context.Context, in *GetRequest, out *GetResponse) error
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
	}
	type WordService struct {
		wordService
	}
	h := &wordServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&WordService{h}, opts...))
}

type wordServiceHandler struct {
	WordServiceHandler
}

func (h *wordServiceHandler) Exist(ctx context.Context, in *ExistRequest, out *ExistResponse) error {
	return h.WordServiceHandler.Exist(ctx, in, out)
}

func (h *wordServiceHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.WordServiceHandler.List(ctx, in, out)
}

func (h *wordServiceHandler) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.WordServiceHandler.Get(ctx, in, out)
}

func (h *wordServiceHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.WordServiceHandler.Create(ctx, in, out)
}

func (h *wordServiceHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.WordServiceHandler.Update(ctx, in, out)
}

func (h *wordServiceHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.WordServiceHandler.Delete(ctx, in, out)
}