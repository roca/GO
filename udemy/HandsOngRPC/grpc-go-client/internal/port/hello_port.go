package port

import (
	"context"

	"google.golang.org/grpc"

	pb "proto/protogen/go/hello"
)

type HelloClientPort interface {
	SayHello(ctx context.Context, in *pb.HelloRequest, opts ...grpc.CallOption) (*pb.HelloResponse, error)
}
