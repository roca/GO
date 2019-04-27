package main

import (
	"context"
	"fmt"
	"log"

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

	greating := &greetpb.Greeting{
		FirstName: "Stephane",
		LastName:  "Maarek",
	}

	req := &greetpb.GreetRequest{
		Greeting: greating,
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
