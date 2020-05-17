package transaction

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
	"github.com/xmlking/logger/log"

	transactionPB "github.com/xmlking/micro-starter-kit/service/recorder/proto/transaction"
	"github.com/xmlking/micro-starter-kit/shared/constants"
)

func publish(ctx context.Context, publisher micro.Event, req, rsp proto.Message) (err error) {
	reqB, err := proto.Marshal(req)
	if err != nil {
		log.WithError(err).Error("marshaling error for req")
		return
	}
	resB, err := proto.Marshal(rsp)
	if err != nil {
		log.WithError(err).Error("marshaling error for rsp")
		return
	}
	event := &transactionPB.TransactionEvent{Req: reqB, Rsp: resB}
	if err = publisher.Publish(ctx, event); err != nil {
		log.WithError(err).Error("Publisher: Failed publishing transation")
	}
	return
}

// NewHandlerWrapper return HandlerWrapper which publish transaction event
func NewHandlerWrapper(p micro.Event) server.HandlerWrapper {
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) (err error) {
			// add TranID to context if not present
			// ctx = metadata.Set(ctx, constants.TransID, uuid.New().String())
			ctx = metadata.MergeContext(ctx, map[string]string{constants.TransID: uuid.New().String()}, false)
			err = fn(ctx, req, rsp)
			// we already logged error in Publish. lets ignore error here. # Note: this is blocking call..
			_ = publish(ctx, p, req.Body().(proto.Message), rsp.(proto.Message))
			// go publish(ctx, p, req.Body().(proto.Message), rsp.(proto.Message))
			return
		}
	}
}

// NewSubscriberWrapper return SubscriberWrapper which publish transaction event
func NewSubscriberWrapper(p micro.Event) server.SubscriberWrapper {
	return func(fn server.SubscriberFunc) server.SubscriberFunc {
		return func(ctx context.Context, req server.Message) (err error) {
			// add TranID to context if not present
			ctx = metadata.MergeContext(ctx, map[string]string{constants.TransID: uuid.New().String()}, false)
			err = fn(ctx, req)
			// we already logged error in Publish. lets ignore error here.
			// FIXME: `Micro-From-Service` is not replaced
			_ = publish(ctx, p, req.Payload().(proto.Message), &empty.Empty{})
			// go publish(ctx, p, req.Payload().(proto.Message), &empty.Empty{})
			return
		}
	}
}
