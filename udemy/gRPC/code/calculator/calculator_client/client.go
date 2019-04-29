package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"udemy.com/gRPC/code/calculator/calculatorpb"
)

func main() {

	fmt.Println("Calculator client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("count no connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	//fmt.Printf("Created client: %f", c)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Sum Unary RPC...")
	nums := &calculatorpb.Nums{
		FirstNum:  3.0,
		SecondNum: 10.0,
	}

	req := &calculatorpb.SumRequest{
		Nums: nums,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.Total)
}
