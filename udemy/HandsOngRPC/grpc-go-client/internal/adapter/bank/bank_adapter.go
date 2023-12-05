package bank

import (
	"context"
	"grpc-go-client/internal/port"
	"log"

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

func (a *BankAdapter) GetCurrentBalance(ctx context.Context, accountNumber string) (*pb.CurrentBalanceResponse, error) {
	currentBalanceRequest := &pb.CurrentBalanceRequest{
		AccountNumber: accountNumber,
	}

	resp, err := a.helloClient.GetCurrentBalance(ctx, currentBalanceRequest)
	if err != nil {
		log.Println("Error on GetCurrentBalanc:", err)
		return nil, err
	}

	return resp, nil
}
