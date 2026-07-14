package interceptor

import (
	"context"
	"log"
	"protogen/go/hello"
	"protogen/go/resiliency"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Unary Interceptors

func LogUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption) error {

		log.Println("[LOGGED BY CLIENT INTERCEPTOR]", req)

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func BasicUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption) error {

		//modify request
		switch request := req.(type) {
		case *hello.HelloRequest:
			request.Name = "[MODIFIED REQUEST BY CLIENT INTERCEPTOR - 1]" + request.Name
		}
		// add request metadata
		ctx = metadata.AppendToOutgoingContext(ctx,
			"my-request-metadata-key-1", "my-request-metadata-value-1")
		ctx = metadata.AppendToOutgoingContext(ctx,
			"my-request-metadata-key-2", "my-request-metadata-value-2")

		// invoke grpc method
		err := invoker(ctx, method, req, reply, cc, opts...)

		if err != nil {
			return err
		}

		// modify response
		switch response := reply.(type) {
		case *hello.HelloResponse:
			response.Greet = "[MODIFIED RESPONSE BY CLIENT INTERCEPTOR - 2]" + response.Greet
		case *resiliency.ResiliencyResponse:
			response.Response = "[MODIFIED RESPONSE BY CLIENT INTERCEPTOR - 3]" + response.Response
		}

		return nil
	}

}

func TimeoutUnaryClientInterceptor(timeout time.Duration) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		newCtx, _ := context.WithTimeout(ctx, timeout)

		return invoker(newCtx, method, req, reply, cc, opts...)
	}
}

// Streaming Interceptors

type InterceptedClientStream struct {
	grpc.ClientStream
}

func LogStreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption) (grpc.ClientStream, error) {

		log.Println("[LOGGED BY CLIENT INTERCEPTOR]", method)

		return streamer(ctx, desc, cc, method, opts...)
	}
}

func BasicClientStreamInterceptor() grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption) (grpc.ClientStream, error) {

		ctx = metadata.AppendToOutgoingContext(ctx,
			"my-request-metadata-key-1", "my-request-metadata-value-1")
		ctx = metadata.AppendToOutgoingContext(ctx,
			"my-request-metadata-key-2", "my-request-metadata-value-2")

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
}

func (s *InterceptedClientStream) SendMsg(msg interface{}) error {
	switch request := msg.(type) {
	case *hello.HelloRequest:
		request.Name = "[MODIFIED REQUEST BY CLIENT INTERCEPTOR - 4]" + request.Name
	}

	return s.ClientStream.SendMsg(msg)
}

func (s *InterceptedClientStream) RecvMsg(msg interface{}) error {
	err := s.ClientStream.RecvMsg(msg)

	if err != nil {
		return err
	}

	switch response := msg.(type) {
	case *hello.HelloResponse:
		response.Greet = "[MODIFIED RESPONSE BY CLIENT INTERCEPTOR - 5]" + response.Greet
	case *resiliency.ResiliencyResponse:
		response.Response = "[MODIFIED RESPONSE BY CLIENT INTERCEPTOR - 6]" + response.Response
	}

	return nil
}

func TimeoutStreamClientInterceptor(timeout time.Duration) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string,
		streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		newCtx, _ := context.WithTimeout(ctx, timeout)

		return streamer(newCtx, desc, cc, method, opts...)
	}
}
