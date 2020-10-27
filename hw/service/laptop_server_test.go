package service_test

import (
	"context"
	"testing"

	"github.com/lifenglsf/grpc_demo/hw/pb"
	"github.com/lifenglsf/grpc_demo/hw/sample"
	"github.com/lifenglsf/grpc_demo/hw/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()
	laptopNoId := sample.NewLaptop()
	laptopNoId.Id = ""

	laptopInvalidId := sample.NewLaptop()
	laptopInvalidId.Id = "invalid-uuid"

	laptopDuplicateId := sample.NewLaptop()
	storeDuplicateId := service.NewInMemoryLaptopStore()
	err := storeDuplicateId.Save(laptopDuplicateId)
	require.Nil(t, err)
	testCases := []struct {
		name   string
		laptop *pb.Laptop
		store  service.LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		}, {
			name:   "success_no_id",
			laptop: laptopNoId,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		}, {
			name:   "failure_invalid_id",
			laptop: laptopInvalidId,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "failure_duplicate_id",
			laptop: laptopDuplicateId,
			store:  storeDuplicateId,
			code:   codes.AlreadyExists,
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			req := &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}
			server := service.NewLaptopServer(tc.store)
			res, err := server.CreateLaptop(context.Background(), req)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				if len(tc.laptop.Id) > 0 {
					require.Equal(t, tc.laptop.Id, res.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}
		})
	}
}
