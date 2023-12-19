package interceptor

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

var BasicUnaryClientInterceptorFunc grpc.UnaryClientInterceptor = func(
	ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {

	// TODO:
	// 1. Make modifications of the req
	// 2. Make modifications of the reply

	return nil
}

type InterceptedClientStream struct {
	grpc.ClientStream
}

var BasicClientStreamInterceptorFunc grpc.StreamClientInterceptor = func(
	ctx context.Context,
	desc *grpc.StreamDesc,
	cc *grpc.ClientConn,
	method string,
	streamer grpc.Streamer,
	opts ...grpc.CallOption) (grpc.ClientStream, error) {

	clientStream, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		log.Printf("Failed to start %v streaming call to %v : %v\n", desc.StreamName, method, err)
		return nil, err
	}

	interceptedClientStream := &InterceptedClientStream{
		ClientStream: clientStream,
	}

	return interceptedClientStream, nil
}

func BasicUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return BasicUnaryClientInterceptorFunc
}

func BasicClientStreamInterceptor() grpc.StreamClientInterceptor {
	return BasicClientStreamInterceptorFunc
}
