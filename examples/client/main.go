package main

import (
	"fmt"
	"github.com/jmz331/nrpc/examples/api"
	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	client := api.NewGreeterServiceClient(nc)

	res, err := client.SayHello(&api.HelloRequest{Name: "jmz331"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Message)
}
