package main

import (
	"context"
	"fmt"
	"log"
	pb "protogen/go/hello"
	"time"

	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	grpc "google.golang.org/grpc"
)

type logWriter struct{}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Printf(time.Now().Format("15:04:04") + " " + string(bytes))
}

func main() {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, _ := grpc.Dial("localhost:9090", opts...)
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	runHello(client)

}

func runHello(client pb.HelloServiceClient) {
	req := &pb.HelloRequest{
		Name: "Joe",
		Age:  32,
	}

	log.Println(req)

	jsonBytes, _ := protojson.Marshal(req)
	log.Println(string(jsonBytes))

	opts := []grpc.CallOption{}

	resp, err := client.SayHello(context.Background(), req, opts...)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(&resp)
}
