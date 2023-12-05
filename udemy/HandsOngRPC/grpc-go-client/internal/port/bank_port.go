package port

import (
	"context"

	"google.golang.org/grpc"

	pb "proto/protogen/go/bank"
)

type BankClientPort interface {
	GetCurrentBalanceWithStatus(ctx context.Context, in *pb.CurrentBalanceRequest, opts ...grpc.CallOption) (*pb.CurrentBalanceResponse, error)
}
