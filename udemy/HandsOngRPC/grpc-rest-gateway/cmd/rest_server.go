package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"protogen/gateway/go/bank"
	"protogen/gateway/go/hello"
	"protogen/gateway/go/resiliency"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	//opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	var opts []grpc.DialOption

	creds, err := credentials.NewClientTLSFromFile("ssl/ca.crt", "")
	if err != nil {
		log.Fatalln("Can't create client credentials :", err)
	}

	opts = append(opts, grpc.WithTransportCredentials(creds))

	if err := hello.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}

	if err := resiliency.RegisterResiliencyServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}

	if err := bank.RegisterBankServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

// func main() {
// 	flag.Parse()
// 	defer glog.Flush()

// 	if err := run(); err != nil {
// 		glog.Fatal(err)
// 	}
// }
