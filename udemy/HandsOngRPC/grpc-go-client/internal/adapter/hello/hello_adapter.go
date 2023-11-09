package hello

import (
	"context"
	"grpc-go-client/internal/port"
	"log"

	"google.golang.org/grpc"

	pb "proto/protogen/go/hello"
)

type HelloAdapter struct {
	helloClient port.HelloClientPort
}

func NewHelloAdapter(conn *grpc.ClientConn) (*HelloAdapter, error) {
	client := pb.NewHelloServiceClient(conn)

	return &HelloAdapter{
		helloClient: client,
	}, nil
}

func (a *HelloAdapter) SayHello(ctx context.Context, name string) (*pb.HelloResponse, error) {
	helloRequest := &pb.HelloRequest{
		Name: name,
	}

	greet,err := a.helloClient.SayHello(ctx, helloRequest)
	if err != nil {
		log.Println("Error on Sayhello:",err)
		return nil,err
	}

	return greet,nil
}
