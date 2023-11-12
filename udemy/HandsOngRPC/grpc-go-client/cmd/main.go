package main

import (
	"context"
	"grpc-go-client/internal/adapter/hello"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("localhost:9090", opts...)
	if err != nil {
		log.Fatalln("Can not connect to gRPC server:", err)
	}
	defer conn.Close()

	helloAdapter, err := hello.NewHelloAdapter(conn)
	if err != nil {
		log.Fatalln("Can not create HelloAdapter:", err)
	}

	// runSayHello(helloAdapter, "Bruce Wayne")
	// runSayManyHellos(helloAdapter, "Bruce Wayne")
	runSayHelloToEveryone(helloAdapter, []string{"Bruce Wayne", "Clark Kent", "Diana Prince"})
}
func runSayHello(adapter *hello.HelloAdapter, name string) {
	greet, err := adapter.SayHello(context.Background(), name)
	if err != nil {
		log.Fatalln("Can not invoke SayHello on the HelloAdapter:", err)
	}

	log.Println(greet.Greet)
}

func runSayManyHellos(adapter *hello.HelloAdapter, name string) {
	err := adapter.SayManyHellos(context.Background(), name)
	if err != nil {
		log.Fatalln("Can not invoke SayManyHellos on the HelloAdapter:", err)
	}
}

func runSayHelloToEveryone(adapter *hello.HelloAdapter, names []string) {
	err := adapter.SayHelloToEveryone(context.Background(), names)
	if err != nil {
		log.Fatalln("Can not invoke SayHelloToEveryone on the HelloAdapter:", err)
	}
}
