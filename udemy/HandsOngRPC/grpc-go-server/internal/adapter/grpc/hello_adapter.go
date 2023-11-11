package grpc

import (
	"context"
	"io"
	pb "proto/protogen/go/hello"
	"time"
)

func (a *GrpcAdapter) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	greet := a.helloService.GenerateHello(req.Name)

	return &pb.HelloResponse{
		Greet: greet,
	}, nil
}

func (a *GrpcAdapter) SayManyHellos(req *pb.HelloRequest, stream pb.HelloService_SayManyHellosServer) error {
	hellos := a.helloService.GenerateManyHellos(req.Name, 10)
	for _, hello := range hellos {
		err := stream.Send(&pb.HelloResponse{
			Greet: hello,
		})
		if err != nil {
			return err
		}
		time.Sleep(500 * time.Millisecond)
	}

	return nil
}

func (a *GrpcAdapter) SayHelloToEveryone(stream pb.HelloService_SayHelloToEveryoneServer) error {
	names := []string{}
	res := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			
			return stream.SendAndClose(&pb.HelloResponse{
				Greet: res,
			},
			)
		}

		if err != nil {
			return err
		}

		names := append(names, req.Name)
		res = a.helloService.GenerateHelloToEveryone(names)

		
	}
}
