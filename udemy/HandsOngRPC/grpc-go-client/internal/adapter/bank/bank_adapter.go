package bank

import (
	"context"
	"grpc-go-client/internal/port"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "proto/protogen/go/bank"
)

type BankAdapter struct {
	bankClient port.BankClientPort
}

type BankTransaction struct {
	AccountNumber   string
	Amount          float64
	TransactionType int32
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

func (a *BankAdapter) SummarizeTransactions(ctx context.Context, transactions []BankTransaction) (*pb.TransactionSummary, error) {
	stream, err := a.bankClient.SummarizeTransactions(ctx)
	if err != nil {
		log.Println("Error getting stream on SummarizeTransactions:", err)
		return nil, err
	}

	for _, transaction := range transactions {
		err := stream.Send(&pb.Transaction{
			AccountNumber:   transaction.AccountNumber,
			Amount:          transaction.Amount,
			TransactionType: pb.TransactionType(transaction.TransactionType),
		})
		if err != nil {
			log.Println("Error sending name on stream SummarizeTransactions:", err)
			return nil, err
		}
		time.Sleep(500 * time.Millisecond)
	}

	summary, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("Error closing and receiving response on stream SayHelloToEveryone:", err)
		return nil, err
	}

	return summary, nil
}
func (a *BankAdapter) TransferMultiple(ctx context.Context, opts ...grpc.CallOption) (pb.BankService_TransferMultipleClient, error) {
	return nil, nil
}
