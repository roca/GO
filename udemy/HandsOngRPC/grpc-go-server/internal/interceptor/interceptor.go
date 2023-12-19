package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

var BasicUnaryClientInterceptorFunc grpc.UnaryServerInterceptor = func(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {

	// TODO:
	// 1. Make modifications of the req
	// 2. Make modifications of the response

	return nil, nil
}

type InterceptedServerStream struct {
	grpc.ServerStream
}

var BasicClientStreamInterceptorFunc grpc.StreamServerInterceptor = func(
	srv interface{},
	ss grpc.ServerStream, info *grpc.StreamServerInfo,
	handler grpc.StreamHandler) error {

	return nil
}

func BasicUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return BasicUnaryClientInterceptorFunc
}

func BasicStreamServerInterceptor() grpc.StreamServerInterceptor {
	return BasicClientStreamInterceptorFunc
}
