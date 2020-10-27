package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/lifenglsf/grpc_demo/hw/pb"
	"github.com/lifenglsf/grpc_demo/hw/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Hello world from client")
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server:%s", *serverAddress)
	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("can not dial server:", err)
	}
	laptopClient := pb.NewLaptopServiceClient(conn)
	laptop := sample.NewLaptop()
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}
	//laptop.Id = "04960ef4-6619-4271-b97d-1f235f3bb0ca"
	//laptop.Id = "invalid id"
	//set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			//not a big deal
			log.Printf("laptop already exists")
		} else {
			log.Fatal("can not create laptop:", err)
		}
		return
	}
	log.Printf("create laptop with id:%s", res.Id)
}
