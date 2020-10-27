package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/lifenglsf/grpc_demo/helloworld/helloworld"
	"google.golang.org/grpc"
)

type server struct {
}

func (s server) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("received from client:%v", request.GetName())
	return &helloworld.HelloReply{Message: "Hello " + request.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("failed to listen on 9999:%v", err)
	}
	gs := grpc.NewServer()
	s := helloworld.NewGreeterService(&server{})
	helloworld.RegisterGreeterService(gs, s)
	if err := gs.Serve(lis); err != nil {
		log.Fatalf("failed to server:%v", err)
	}
	time.Now().Format(time.ANSIC)
}
