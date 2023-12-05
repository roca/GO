package bank

import (
	"context"
	"grpc-go-client/internal/port"

	"google.golang.org/grpc"

	pb "proto/protogen/go/bank"
)

type BankAdapter struct {
	helloClient port.BankClientPort
}

func NewBankAdapter(conn *grpc.ClientConn) (*BankAdapter, error) {
	client := pb.NewBankServiceClient(conn)

	return &BankAdapter{
		helloClient: client,
	}, nil
}

func (a *BankAdapter) GetCurrentBalanceWithStatus(ctx context.Context, accountNumber string) (*pb.CurrentBalanceResponse, error) {
	currentBalanceRequest := &pb.CurrentBalanceRequest{
		AccountNumber: accountNumber,
	}

	resp, err := a.helloClient.GetCurrentBalanceWithStatus(ctx, currentBalanceRequest)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
