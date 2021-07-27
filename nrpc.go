package nrpc

import (
	"github.com/nats-io/nats.go"
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
