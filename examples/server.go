package main

import (
	"context"
	"github.com/jmz331/nrpc/examples/api"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, req *api.HelloRequest) (*api.HelloReply, error) {
	return &api.HelloReply{Message: "Hello " + req.Name + "!"}, nil
}
