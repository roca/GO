package grpc

import (
	"context"
	pb "proto/protogen/go/bank"
	"time"

	"google.golang.org/genproto/googleapis/type/date"
)

func (a *GrpcAdapter) GetCurrentBalance(ctx context.Context, req *pb.CurrentBalanceRequest) (*pb.CurrentBalanceResponse, error) {
	cb := &pb.CurrentBalanceResponse{
		Amount: a.bankService.GetCurrentBalance(int(req.AccountNumber)),
		CurrentDate: &date.Date{
			Year:  int32(time.Now().Year()),
			Month: int32(time.Now().Month()),
			Day:   int32(time.Now().Day()),
		},
	}
	return cb, nil
}
