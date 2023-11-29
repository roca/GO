package grpc

import (
	"context"
	"grpc-go-server/internal/port"
	"io"
	"log"
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

func (a *GrpcAdapter) FetchExchangeRates(req *pb.ExchangeRateRequest, stream pb.BankService_FetchExchangeRatesServer) error {
	context := stream.Context()
	for {
		select {
		case <-context.Done():
			log.Println("Client has cancelled stream")
			//a.BankService.StopExchangeRatesAtInterval()
		default:
			now := time.Now().Truncate(time.Second)
			rate, err := a.BankService.GetExchangeRateAtTimestamp(req.FromCurrency, req.ToCurrency, now)
			if err != nil {
				return err
			}
			stream.Send(
				&pb.ExchangeRateResponse{
					FromCurrency: req.FromCurrency,
					ToCurrency:   req.ToCurrency,
					Rate:         rate.Rate,
					Timestamp:    now.Format(time.RFC3339),
				},
			)
			log.Printf("Exchange rate sent to client, %v to %v : %v\n", req.FromCurrency, req.ToCurrency, rate.Rate)
			time.Sleep(3 * time.Second)
		}
	}
}

func (a *GrpcAdapter) SummarizeTransactions(stream pb.BankService_SummarizeTransactionsServer) error {
	var portTransactions []*port.Transaction = nil
	var pbTransactions []*pb.Transaction = nil
	var accountNumber string = ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			balance, err := a.BankService.ExecuteBankTransactions(portTransactions)
			if err != nil {
				return err
			}
			return stream.SendAndClose(&pb.TransactionSummary{
				AccountNumber: accountNumber,
				Balance:       balance,
				Transactions:  pbTransactions,
			},
			)
		}
		if err != nil {
			return err
		}
		accountNumber = req.AccountNumber
		portTransactions = append(portTransactions, &port.Transaction{
			AccountNumber:   req.AccountNumber,
			TransactionType: port.TransactionType(req.TransactionType),
			Amount:          req.Amount,
		})
		pbTransactions = append(pbTransactions, &pb.Transaction{
			AccountNumber:   req.AccountNumber,
			TransactionType: req.TransactionType,
			Amount:          req.Amount,
		})
	}
}

func (a *GrpcAdapter) TransferMultiple(stream pb.BankService_TransferMultipleServer) error {
	context := stream.Context()
	for {
		select {
		case <-context.Done():
			log.Println("Client has cancelled stream")
			return nil
		default:
			req, err := stream.Recv()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}

			for resp := range a.BankService.ExecuteBankTransfers(req) {
				_ = stream.Send(resp)
				// if err != nil {
				// 	return err
				// }
				time.Sleep(500 * time.Millisecond)
			}
		}
	}
}
