package interceptor

import (
	"context"
	"log"
	"proto/protogen/go/hello"
	"proto/protogen/go/resiliency"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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
		// Modify request message if the request message is HelloRequest. Append a string to request's Name
		switch request := req.(type) {
		case *hello.HelloRequest:
			request.Name = "[MODIFIED REQUEST BY SERVER INTERCEPTOR - 1]" + request.Name
		}
		// Add two new response metadata (any key / value)
		responseMetadata, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			responseMetadata = metadata.New(nil)
		}
		responseMetadata.Append("my-response-metadata-key-1", "my-response-metadata-value-1")
		responseMetadata.Append("my-response-metadata-key-2", "my-response-metadata-value-2")
		// Create new context with new response metadata
		ctx = metadata.NewOutgoingContext(ctx, responseMetadata)

		// Set response header with new response metadata
		grpc.SetHeader(ctx, responseMetadata)

		// handle request with modified context
		res, err := handler(ctx, req)

		if err != nil {
			return res, err
		}

		// modify response
		switch response := res.(type) {
		case *hello.HelloResponse:
			response.Greet = "[MODIFIED RESPONSE BY SERVER INTERCEPTOR - 2]" + response.Greet
		case *resiliency.ResiliencyResponse:
			response.Response = "[MODIFIED RESPONSE BY SERVER INTERCEPTOR - 3]" + response.Response
		}

		return res, nil
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

		// intercept stream
		interceptedServerStream := &InterceptedServerStream{
			ServerStream: ss,
		}

		// add response metadata
		responseMetadata, ok := metadata.FromOutgoingContext(interceptedServerStream.Context())

		if !ok {
			responseMetadata = metadata.New(nil)
		}

		responseMetadata.Append("my-response-metadata-key-1", "my-response-metadata-value-1")
		responseMetadata.Append("my-response-metadata-key-2", "my-response-metadata-value-2")

		interceptedServerStream.SetHeader(responseMetadata)

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
		request.Name = "[MODIFIED REQUEST BY SERVER INTERCEPTOR - 4]" + request.Name
	}

	return nil
}

func (s *InterceptedServerStream) SendMsg(msg interface{}) error {
	switch response := msg.(type) {
	case *hello.HelloResponse:
		response.Greet = "[MODIFIED RESPONSE BY SERVER INTERCEPTOR - 5]" + response.Greet
	case *resiliency.ResiliencyResponse:
		response.Response = "[MODIFIED RESPONSE BY SERVER INTERCEPTOR - 6]" + response.Response
	}

	return s.ServerStream.SendMsg(msg)
}
