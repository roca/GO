package bank

import (
	"context"
	"grpc-go-client/internal/port"
	"io"
	"log"

	"google.golang.org/grpc"

	pb "proto/protogen/go/bank"
)

type BankAdapter struct {
	bankClient port.BankClientPort
}

func NewBankAdapter(conn *grpc.ClientConn) (*BankAdapter, error) {
	client := pb.NewBankServiceClient(conn)

	return &BankAdapter{
		bankClient: client,
	}, nil
}

func (a *BankAdapter) GetCurrentBalanceWithStatus(ctx context.Context, accountNumber string) (*pb.CurrentBalanceResponse, error) {
	currentBalanceRequest := &pb.CurrentBalanceRequest{
		AccountNumber: accountNumber,
	}

	resp, err := a.bankClient.GetCurrentBalanceWithStatus(ctx, currentBalanceRequest)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *BankAdapter) FetchExchangeRates(ctx context.Context, fromCurrency, toCurrency string) error {
	exchangeRateRequest := &pb.ExchangeRateRequest{
		FromCurrency: fromCurrency,
		ToCurrency:   toCurrency,
	}

	stream, err := a.bankClient.FetchExchangeRates(ctx, exchangeRateRequest)
	if err != nil {
		log.Println("Error on FetchExchangeRates:", err)
		return err
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error on FetchExchangeRates:", err)
			return err
		}
		log.Println(res.Rate)
	}

	return nil
}
func (a *BankAdapter) SummarizeTransactions(ctx context.Context, opts ...grpc.CallOption) (pb.BankService_SummarizeTransactionsClient, error) {
	return nil, nil
}
func (a *BankAdapter) TransferMultiple(ctx context.Context, opts ...grpc.CallOption) (pb.BankService_TransferMultipleClient, error) {
	return nil, nil
}
