package main

import (
	"context"
	"fmt"
	"github.com/jmz331/nrpc"
	"github.com/jmz331/nrpc/examples/api"
	"github.com/nats-io/nats.go"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	err = run(api.NewGreeterServiceHandler(ctx, nc, &server{}))
	if err != nil {
		panic(err)
	}
}

func run(handlers ...nrpc.RPCHandler) error {
	for _, h := range handlers {
		err := h.Subscribe()
		if err != nil {
			return err
		}
	}
	defer func() {
		for _, h := range handlers {
			_ = h.Unsubscribe()
		}
	}()

	// Keep running until ^C.
	fmt.Println("server is running, ^C quits.")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	close(c)

	return nil
}
