package nrpc

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
	"log"
	"time"
)

//go:generate protoc --go_out=. --go_opt=paths=source_relative nrpc.proto

type NatsConn interface {
	Publish(subj string, data []byte) error
	PublishRequest(subj, reply string, data []byte) error
	Request(subj string, data []byte, timeout time.Duration) (*nats.Msg, error)

	ChanSubscribe(subj string, ch chan *nats.Msg) (*nats.Subscription, error)
	Subscribe(subj string, handler nats.MsgHandler) (*nats.Subscription, error)
	SubscribeSync(subj string) (*nats.Subscription, error)
}

type RPCHandler interface {
	Subscribe() error
	Unsubscribe() error
}

type Request struct {
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s error: %s", Error_Type_name[int32(e.Type)], e.Message)
}

func Call(req proto.Message, rep proto.Message, nc NatsConn, subject string, timeout time.Duration) error {
	// encode request
	rawRequest, err := proto.Marshal(req)
	if err != nil {
		//todo 处理异常情况日志输出
		log.Printf("nrpc: inner request marshal failed: %v", err)
		return err
	}

	// call
	if _, isNoReply := rep.(*NoReply); isNoReply {
		err := nc.Publish(subject, rawRequest)
		if err != nil {
			//todo 处理异常情况日志输出
			log.Printf("nrpc: nats publish failed: %v", err)
		}
		return err
	}
	msg, err := nc.Request(subject, rawRequest, timeout)
	if err != nil {
		//todo 处理异常情况日志输出
		log.Printf("nrpc: nats request failed: %v", err)
		return err
	}

	data := msg.Data
	if err := proto.Unmarshal(data, rep); err != nil {
		if _, isError := err.(*Error); !isError {
			log.Printf("nrpc: response unmarshal failed: %v", err)
		}
		return err
	}
	return nil
}
