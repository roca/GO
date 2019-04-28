package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"udemy.com/gRPC/code/calculator/calculatorpb"
	"udemy.com/gRPC/code/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invocked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello" + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum function was invocked with %v\n", req)
	firstNum := req.GetNums().GetFirstNum()
	secondNum := req.GetNums().GetSecondNum()
	result := firstNum + secondNum
	res := &calculatorpb.SumResponse{
		Total: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
