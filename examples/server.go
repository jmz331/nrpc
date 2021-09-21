package main

import (
	"context"
	"github.com/jmz331/nrpc/examples/api"
)

type greeterServiceHandler struct {
}

func (s *greeterServiceHandler) SayHello(ctx context.Context, req *api.HelloRequest) (*api.HelloReply, error) {
	return &api.HelloReply{Message: "Hello " + req.Name + "!"}, nil
}

func NewGreeterServiceHandler() api.GreeterServiceServer {
	return &greeterServiceHandler{}
}
