package grpc

import (
	"context"
	pb "proto/protogen/go/bank"
	"time"

	"google.golang.org/genproto/googleapis/type/date"
)

func (a *GrpcAdapter) GetCurrentBalance(ctx context.Context, req *pb.CurrentBalanceRequest) (*pb.CurrentBalanceResponse, error) {
	balance := a.BankService.FindCurrentBalance(req.AccountNumber)

	now := time.Now()

	return &pb.CurrentBalanceResponse{
		CurrentBalance: balance,
		CurrentDate: &date.Date{
			Year:  int32(now.Year()),
			Month: int32(now.Month()),
			Day:   int32(now.Day()),
		},
	}, nil
}
