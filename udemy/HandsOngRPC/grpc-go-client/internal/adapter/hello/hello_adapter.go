package hello

import (
	"context"
	"grpc-go-client/internal/port"
	"io"
	"log"
	"time"

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

func (a *HelloAdapter) SayHelloToEveryone(ctx context.Context, names []string) error {

	stream, err := a.helloClient.SayHelloToEveryone(ctx)
	if err != nil {
		log.Println("Error getting stream on SayHelloToEveryone:", err)
		return err
	}

	for _, name := range names {
		err := stream.Send(&pb.HelloRequest{
			Name: name,
		})
		if err != nil {
			log.Println("Error sending name on stream SayHelloToEveryone:", err)
			return err
		}
		time.Sleep(500 * time.Millisecond)
	}

	greet, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("Error closing and receiving response on stream SayHelloToEveryone:", err)
		return err
	}

	log.Println(greet.Greet)
	return nil
}

func (a *HelloAdapter) SayHelloContinuous(ctx context.Context, names []string) error {
	stream, err := a.helloClient.SayHelloContinuous(ctx)
	if err != nil {
		log.Println("Error getting stream on SayHelloContinuoue:", err)
		return err
	}

	greetChan := make(chan struct{})

	go func() {

		for _, name := range names {
			err := stream.Send(&pb.HelloRequest{
				Name: name,
			})
			if err != nil {
				log.Println("Error sending name on stream SayHelloContinuous:", err)
			}
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			greet, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println("Error on receiving on SayHelloContinuous:", err)
			}
			log.Println(greet.Greet)
		}
		close(greetChan)
	}()

	<-greetChan

	return nil
}
