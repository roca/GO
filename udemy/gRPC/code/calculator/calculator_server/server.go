package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"udemy.com/gRPC/code/calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Recieved Sum RPC: %v\n", req)
	firstNum := req.GetNums().GetFirstNum()
	secondNum := req.GetNums().GetSecondNum()
	total := firstNum + secondNum
	res := &calculatorpb.SumResponse{
		Total: total,
	}
	return res, nil
}

func main() {
	fmt.Println("Calculator Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
