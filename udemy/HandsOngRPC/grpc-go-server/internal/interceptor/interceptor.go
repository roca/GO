package interceptor

import (
	"context"
	"log"
	"proto/protogen/go/hello"
	"proto/protogen/go/resiliency"

	"google.golang.org/grpc"
)

// Unary Interceptors

func LogUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Println("[LOGGED BY SERVER INTERCEPTOR]", req)

		return handler(ctx, req)
	}
}

func BasicUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {

		// TODO:
		// 1. Make modifications of the req
		// 2. Make modifications of the response

		return nil, nil
	}
}

// Streaming Interceptors

type InterceptedServerStream struct {
	grpc.ServerStream
}

func LogStreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo,
		handler grpc.StreamHandler) error {
		log.Println("[LOGGED BY SERVER INTERCEPTOR]", info)

		return handler(srv, ss)
	}
}

func BasicStreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream, info *grpc.StreamServerInfo,
		handler grpc.StreamHandler) error {

		interceptedServerStream := &InterceptedServerStream{
			ServerStream: ss,
		}

		// handle request
		return handler(srv, interceptedServerStream)
	}
}

func (s *InterceptedServerStream) RecvMsg(msg interface{}) error {
	if err := s.ServerStream.RecvMsg(msg); err != nil {
		return err
	}

	switch request := msg.(type) {
	case *hello.HelloRequest:
		request.Name = "[MODIFIED BY SERVER INTERCEPTOR - 4]" + request.Name
	}

	return nil
}

func (s *InterceptedServerStream) SendMsg(msg interface{}) error {
	switch response := msg.(type) {
	case *hello.HelloResponse:
		response.Greet = "[MODIFIED BY SERVER INTERCEPTOR - 5]" + response.Greet
	case *resiliency.ResiliencyResponse:
		response.Response = "[MODIFIED BY SERVER INTERCEPTOR - 6]" + response.Response
	}

	return s.ServerStream.SendMsg(msg)
}
