package main

import (
	"context"
	"github.com/lifenglsf/grpc_demo/helloworld/helloworld"
	"google.golang.org/grpc"
	"log"
)

type client struct {
}

func main() {
	conn, err := grpc.Dial(":9999", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to port 9999,%v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)
	response, err := c.SayHello(context.Background(), &helloworld.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("error when calling sayhello:%s", err)
	}
	log.Printf("Response form server:%s", response.Message)
}
