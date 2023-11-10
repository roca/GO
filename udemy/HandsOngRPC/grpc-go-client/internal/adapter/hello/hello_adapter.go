package hello

import (
	"context"
	"grpc-go-client/internal/port"
	"io"
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

	greet, err := a.helloClient.SayHello(ctx, helloRequest)
	if err != nil {
		log.Println("Error on Sayhello:", err)
		return nil, err
	}

	return greet, nil
}

func (a *HelloAdapter) SayManyHellos(ctx context.Context, name string) error {
	helloRequest := &pb.HelloRequest{
		Name: name,
	}

	greetStream, err := a.helloClient.SayManyHellos(ctx, helloRequest)
	if err != nil {
		log.Println("Error on SayManyHellos:", err)
		return err
	}

	for {
		greet, err := greetStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error on SayManyHellos:", err)
			return err
		}
		log.Println(greet.Greet)
	}

	return nil
}
