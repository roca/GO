package port

import (
	"context"

	"google.golang.org/grpc"

	pb "proto/protogen/go/hello"
)

type HelloClientPort interface {
	SayHello(ctx context.Context, in *pb.HelloRequest, opts ...grpc.CallOption) (*pb.HelloResponse, error)
	SayManyHellos(ctx context.Context, in *pb.HelloRequest, opts ...grpc.CallOption) (pb.HelloService_SayManyHellosClient, error)
	SayHelloToEveryone(ctx context.Context, opts ...grpc.CallOption) (pb.HelloService_SayHelloToEveryoneClient, error)
}
