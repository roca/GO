package port

import (
	"context"

	"google.golang.org/grpc"

	pb "proto/protogen/go/resiliency"
)


type ResiliencyClientPort interface {
	GetResiliency(ctx context.Context, in *pb.ResiliencyRequest, opts ...grpc.CallOption) (*pb.ResiliencyResponse, error)
	GetResiliencyStream(ctx context.Context, in *pb.ResiliencyRequest, opts ...grpc.CallOption) (pb.ResiliencyService_GetResiliencyStreamClient, error)
	SendResiliencyStream(ctx context.Context, opts ...grpc.CallOption) (pb.ResiliencyService_SendResiliencyStreamClient, error)
	BidirectionalResiliencyStream(ctx context.Context, opts ...grpc.CallOption) (pb.ResiliencyService_BidirectionalResiliencyStreamClient, error)
}