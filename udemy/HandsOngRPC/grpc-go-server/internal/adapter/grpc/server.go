package grpc

import (
	"fmt"
	"grpc-go-server/internal/port"
	"log"
	"net"
	"protogen/go/bank"
	"protogen/go/hello"
	"protogen/go/resiliency"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GrpcAdapter struct {
	helloService      port.HelloServicePort
	BankService       port.BankServicePort
	ResiliencyService port.ResiliencyServicePort
	grpcPort          int
	server            *grpc.Server
	hello.HelloServiceServer
	bank.BankServiceServer
	resiliency.ResiliencyServiceServer
}

func NewGrpcAdapter(helloService port.HelloServicePort, bankService port.BankServicePort, resiliencyService port.ResiliencyServicePort, grpcPort int) *GrpcAdapter {
	return &GrpcAdapter{
		helloService:      helloService,
		BankService:       bankService,
		ResiliencyService: resiliencyService,
		grpcPort:          grpcPort,
	}
}

func (a *GrpcAdapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen on port %d : %v\n", a.grpcPort, err)
	}

	log.Printf("Server listening on port %d\n", a.grpcPort)

	creds, err := credentials.NewServerTLSFromFile("ssl/server.crt", "ssl/server.pem")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(creds),
	// grpc.ChainUnaryInterceptor(
	// 	interceptor.LogUnaryServerInterceptor(),
	// 	interceptor.BasicUnaryServerInterceptor(),
	// ),
	// grpc.ChainStreamInterceptor(
	// 	interceptor.LogStreamServerInterceptor(),
	// 	interceptor.BasicStreamServerInterceptor(),
	// ),
	)
	a.server = grpcServer

	hello.RegisterHelloServiceServer(grpcServer, a)
	bank.RegisterBankServiceServer(grpcServer, a)
	resiliency.RegisterResiliencyServiceServer(grpcServer, a)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %d : %v\n", a.grpcPort, err)
	}
}

func (a *GrpcAdapter) Stop() {
	a.server.Stop()
}
