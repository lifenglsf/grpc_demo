package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/lifenglsf/grpc_demo/hw/pb"
	"github.com/lifenglsf/grpc_demo/hw/service"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello world from server")
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	fmt.Println("port:", port)
	log.Printf("start server on port %d", *port)
	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())
	grpcServer := grpc.NewServer()
	s := pb.NewLaptopServiceService(laptopServer)
	pb.RegisterLaptopServiceService(grpcServer, s)
	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("can not start server:", err)
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("can not start server:", err)
	}
}
