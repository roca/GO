package grpc

import (
	"context"
	pb "proto/protogen/go/hello"
)

func (a *GrpcAdapter) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	greet := a.helloService.GenerateHello(req.Name)

	return &pb.HelloResponse{
		Greet: greet,
	}, nil
}
