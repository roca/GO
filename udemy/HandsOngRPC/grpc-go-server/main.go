package main

import (
	"context"
	"fmt"
	"log"
	pb "proto/protogen/go/hello"
	"time"

	"google.golang.org/protobuf/encoding/protojson"

	grpc "google.golang.org/grpc"
)

type logWriter struct{}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Printf(time.Now().Format("15:04:04") + " " + string(bytes))
}

func main() {
	req := &pb.HelloRequest{
		Name: "Joe",
		Age:  32,
	}

	log.Println(req)

	jsonBytes, _ := protojson.Marshal(req)
	log.Println(string(jsonBytes))

	cc := grpc.ClientConn{}
	opts := []grpc.CallOption{}

	client := pb.NewHelloServiceClient(&cc)

	resp, err := client.SayHello(context.Background(), req, opts...)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(&resp)
}
