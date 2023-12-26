package grpc

import (
	"context"
	"io"
	pb "protogen/go/hello"
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

var names []string
var res string

func (a *GrpcAdapter) SayHelloToEveryone(stream pb.HelloService_SayHelloToEveryoneServer) error {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			res = a.helloService.GenerateHelloToEveryone(names)
			names = nil
			return stream.SendAndClose(&pb.HelloResponse{
				Greet: res,
			},
			)
		}
		if err != nil {
			return err
		}

		names = append(names, req.Name)
	}
}

func (a *GrpcAdapter) SayHelloContinuous(stream pb.HelloService_SayHelloContinuousServer) error {
	i := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		for hello := range a.helloService.GenerateContinuousHello(req.Name, i) {
			err := stream.Send(&pb.HelloResponse{
				Greet: hello,
			})
			if err != nil {
				return err
			}
			time.Sleep(500 * time.Millisecond)
		}
		i++
	}
}
