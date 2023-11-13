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
	bs := &app.BankService{}

	grpcadapter := mygrpc.NewGrpcAdapter(hs, bs, 9090)

	grpcadapter.Run()
}
