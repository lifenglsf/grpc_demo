package main

import (
	"context"
	"github.com/lifenglsf/grpc_demo/helloworld-json/helloworld"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	helloworld.UnimplementedGreeterServer
}

func (s server) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("received from client:%v", request.GetName())
	return &helloworld.HelloReply{Message: "Hello " + request.GetName(), Code: "3880"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("failed to listen on 9999:%v", err)
	}
	gs := grpc.NewServer()
	helloworld.RegisterGreeterServer(gs, &server{})
	if err := gs.Serve(lis); err != nil {
		log.Fatalf("failed to server:%v", err)
	}
}
