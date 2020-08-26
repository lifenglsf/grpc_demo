package main

import (
	"context"
	"encoding/json"
	"github.com/lifenglsf/grpc_demo/helloworld-json/helloworld"
	"google.golang.org/grpc"
	"log"
	"net/http"
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
	http.HandleFunc("/helloworld", func(writer http.ResponseWriter, request *http.Request) {
		response, err := c.SayHello(context.Background(), &helloworld.HelloRequest{Name: "World"})
		if err != nil {
			log.Fatalf("error when calling sayhello:%s", err)
		}
		writer.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(response)
		writer.Write(data)
	})
	http.ListenAndServe(":8888", nil)
}
