package grpc

import (
	"fmt"
	"grpc-go-server/internal/port"
	"log"
	"net"
	"proto/protogen/go/bank"
	"proto/protogen/go/hello"

	"google.golang.org/grpc"
)

type GrpcAdapter struct {
	helloService port.HelloServicePort
	BankService  port.BankServicePort
	grpcPort     int
	server       *grpc.Server
	hello.HelloServiceServer
	bank.BankServiceServer
}

func NewGrpcAdapter(helloService port.HelloServicePort, bankService port.BankServicePort, grpcPort int) *GrpcAdapter {
	return &GrpcAdapter{
		helloService: helloService,
		BankService:  bankService,
		grpcPort:     grpcPort,
	}
}

func (a *GrpcAdapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen on port %d : %v\n", a.grpcPort, err)
	}

	log.Printf("Server listening on port %d\n", a.grpcPort)

	grpsServer := grpc.NewServer()
	a.server = grpsServer

	hello.RegisterHelloServiceServer(grpsServer, a)
	bank.RegisterBankServiceServer(grpsServer, a)

	if err := grpsServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %d : %v\n", a.grpcPort, err)
	}
}

func (a *GrpcAdapter) Stop() {
	a.server.Stop()
}
