package serializer_test

import (
	"testing"

	"github.com/lifenglsf/grpc_demo/hw/pb"
	"github.com/lifenglsf/grpc_demo/hw/sample"
	"github.com/lifenglsf/grpc_demo/hw/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()
	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/latop.json"
	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))
	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}
