package grpc

import (
	"grpc-go-server/internal/port"
	"proto/protogen/hello"

	"google.golang.org/grpc"
)

type GrpcAdapter struct {
	helloService port.HelloServicePort
	grpcPort     int
	server       *grpc.Server
	hello.HelloServiceServer
}
