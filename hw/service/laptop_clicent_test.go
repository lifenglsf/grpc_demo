package service_test

import (
	"context"
	"net"
	"testing"

	"github.com/lifenglsf/grpc_demo/hw/pb"
	"github.com/lifenglsf/grpc_demo/hw/sample"
	"github.com/lifenglsf/grpc_demo/hw/serializer"
	"github.com/lifenglsf/grpc_demo/hw/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()
	laptopServer, serverAddress := startTestLaptopServer(t)
	laptopClient := newTestLaptopClient(t, serverAddress)
	laptop := sample.NewLaptop()
	expectedId := laptop.Id
	req := &pb.CreateLaptopRequest{Laptop: laptop}
	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedId, req.Laptop.Id)
	other, err := laptopServer.Store.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, other)
	requireSameLaptop(t, laptop, other)
}

func startTestLaptopServer(t *testing.T) (*service.LaptopServer, string) {
	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())
	grpcServer := grpc.NewServer()
	service := pb.NewLaptopServiceService(laptopServer)
	pb.RegisterLaptopServiceService(grpcServer, service)
	listener, err := net.Listen("tcp", ":0") //random avaliable port
	require.NoError(t, err)
	go grpcServer.Serve(listener)
	return laptopServer, listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewLaptopServiceClient(conn)
}
func requireSameLaptop(t *testing.T, laptop1 *pb.Laptop, laptop2 *pb.Laptop) {
	json1, err := serializer.ProtobufToJSON(laptop1)
	require.NoError(t, err)
	json2, err := serializer.ProtobufToJSON(laptop2)
	require.NoError(t, err)
	require.Equal(t, json1, json2)
}
