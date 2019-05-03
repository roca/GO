package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"udemy.com/gRPC/code/greet/greetpb"
)

func main() {

	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("count no connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Created client: %f", c)

	//doUnary(c)
	//doServerStreaming(c)
	//doClientStreaming(c)
	doBiDiStreaming(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	greating := &greetpb.Greeting{
		FirstName: "Stephane",
		LastName:  "Maarek",
	}

	req := &greetpb.GreetRequest{
		Greeting: greating,
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")
	greating := &greetpb.Greeting{
		FirstName: "Stephane",
		LastName:  "Maarek",
	}

	req := &greetpb.GreetManyTimesRequest{
		Greeting: greating,
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
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
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}

}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{Greeting: &greetpb.Greeting{FirstName: "Stephane"}},
		&greetpb.LongGreetRequest{Greeting: &greetpb.Greeting{FirstName: "John"}},
		&greetpb.LongGreetRequest{Greeting: &greetpb.Greeting{FirstName: "Lucy"}},
		&greetpb.LongGreetRequest{Greeting: &greetpb.Greeting{FirstName: "Mark"}},
		&greetpb.LongGreetRequest{Greeting: &greetpb.Greeting{FirstName: "Piper"}},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
		return
	}

	// we iterate over our slice and send each message individually
	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while recieving response from LongGreet: %v", err)
	}

	fmt.Printf("LongGreet Response: %v\n", res)
}

func doBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a BiDi Streaming RPC...")

	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{Greeting: &greetpb.Greeting{FirstName: "Stephane"}},
		&greetpb.GreetEveryoneRequest{Greeting: &greetpb.Greeting{FirstName: "John"}},
		&greetpb.GreetEveryoneRequest{Greeting: &greetpb.Greeting{FirstName: "Lucy"}},
		&greetpb.GreetEveryoneRequest{Greeting: &greetpb.Greeting{FirstName: "Mark"}},
		&greetpb.GreetEveryoneRequest{Greeting: &greetpb.Greeting{FirstName: "Piper"}},
	}

	// we create a stream by invoking the client
	stream, err := c.GreetEveryone(context.Background())
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
