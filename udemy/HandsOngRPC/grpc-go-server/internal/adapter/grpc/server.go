package grpc

import (
	"fmt"
	"grpc-go-server/data"
	"grpc-go-server/internal/port"
	"log"
	"net"
	"proto/protogen/go/bank"
	"proto/protogen/go/hello"

	"google.golang.org/grpc"
)

type GrpcAdapter struct {
	helloService port.HelloServicePort
	bankService  port.BankServicePort
	grpcPort     int
	server       *grpc.Server
	hello.HelloServiceServer
	bank.BankServiceServer
	data.Models
}

func NewGrpcAdapter(helloService port.HelloServicePort, bankService port.BankServicePort, grpcPort int) *GrpcAdapter {
	//cel := &celeritas.Celeritas{}
	return &GrpcAdapter{
		helloService: helloService,
		bankService:  bankService,
		grpcPort:     grpcPort,
		//Models:       data.New(cel.DB.Pool),
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
