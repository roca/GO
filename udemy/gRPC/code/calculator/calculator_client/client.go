package main

import (
	"context"
	"fmt"
	"io"
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

	//doUnary(c)
	doServerStreaming(c)
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

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &calculatorpb.PrimeNumbersRequest{
		Number: 12390392840,
	}

	resStream, err := c.PrimeNumbers(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling PrimeNumbers RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from PrimeNumbers: %v", msg.GetResult())
	}

}
