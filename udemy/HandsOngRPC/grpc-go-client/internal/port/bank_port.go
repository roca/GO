package port

import (
	"context"

	"google.golang.org/grpc"

	pb "proto/protogen/go/bank"
)

type BankClientPort interface {
	GetCurrentBalance(ctx context.Context, in *pb.CurrentBalanceRequest, opts ...grpc.CallOption) (*pb.CurrentBalanceResponse, error)
}
