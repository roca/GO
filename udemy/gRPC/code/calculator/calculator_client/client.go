package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	//doServerStreaming(c)
	//doClientStreaming(c)
	doBiDiStreaming(c)
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

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")

	numbers := []int64{1, 2, 3, 4}

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	// we iterate over our slice and send each message individually
	for _, number := range numbers {
		fmt.Printf("Sending number: %v\n", number)
		stream.Send(&calculatorpb.ComputeAverageRequest{Number: number})
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while recieving response from LongGreet: %v", err)
	}

	fmt.Printf("ComputeAverage Response: %v\n", res)

}

func doBiDiStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a BiDi Streaming RPC...")

	requests := []*calculatorpb.FindMaximumRequest{}
	numbers := []int64{1, 5, 3, 6, 2, 20, 17, 10, 30}

	for _, n := range numbers {
		requests = append(requests, &calculatorpb.FindMaximumRequest{Number: n})
	}
	// we create a stream by invoking the client
	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
	}

	waitc := make(chan struct{})

	// we send a bunch of messages to the server (go routine)
	go func() {
		// function to send a bunch of messages
		for _, req := range requests {
			fmt.Printf("Sending message %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// we receive a bund of messages from the server (go routine)
	go func() {
		// function to receive a bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("Received %v\n", res.GetResult())
		}
		close(waitc)
	}()

	// block until everything in done
	<-waitc

}
