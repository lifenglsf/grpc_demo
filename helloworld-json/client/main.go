package main

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	"github.com/lifenglsf/grpc_demo/helloworld-json/helloworld"
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
	m := jsonpb.Marshaler{}
	res, _ := m.MarshalToString(response)
	log.Printf("Response form server:%v", res)
}
