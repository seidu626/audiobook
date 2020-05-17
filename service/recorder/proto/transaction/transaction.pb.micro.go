// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: service/recorder/proto/transaction/transaction.proto

package transaction

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// Client API for TransactionService service

type TransactionService interface {
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*TransactionEvent, error)
	Write(ctx context.Context, in *WriteRequest, opts ...client.CallOption) (*empty.Empty, error)
}

type transactionService struct {
	c    client.Client
	name string
}

func NewTransactionService(name string, c client.Client) TransactionService {
	return &transactionService{
		c:    c,
		name: name,
	}
}

func (c *transactionService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*TransactionEvent, error) {
	req := c.c.NewRequest(c.name, "TransactionService.Read", in)
	out := new(TransactionEvent)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionService) Write(ctx context.Context, in *WriteRequest, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "TransactionService.Write", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TransactionService service

type TransactionServiceHandler interface {
	Read(context.Context, *ReadRequest, *TransactionEvent) error
	Write(context.Context, *WriteRequest, *empty.Empty) error
}

func RegisterTransactionServiceHandler(s server.Server, hdlr TransactionServiceHandler, opts ...server.HandlerOption) error {
	type transactionService interface {
		Read(ctx context.Context, in *ReadRequest, out *TransactionEvent) error
		Write(ctx context.Context, in *WriteRequest, out *empty.Empty) error
	}
	type TransactionService struct {
		transactionService
	}
	h := &transactionServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TransactionService{h}, opts...))
}

type transactionServiceHandler struct {
	TransactionServiceHandler
}

func (h *transactionServiceHandler) Read(ctx context.Context, in *ReadRequest, out *TransactionEvent) error {
	return h.TransactionServiceHandler.Read(ctx, in, out)
}

func (h *transactionServiceHandler) Write(ctx context.Context, in *WriteRequest, out *empty.Empty) error {
	return h.TransactionServiceHandler.Write(ctx, in, out)
}
