package main

import (
	mygrpc "grpc-go-server/internal/adapter/grpc"
	app "grpc-go-server/internal/application"
	"log"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	hs := &app.HelloService{}

	grpcadapter := mygrpc.NewGrpcAdapter(hs, 9090)

	grpcadapter.Run()
}
