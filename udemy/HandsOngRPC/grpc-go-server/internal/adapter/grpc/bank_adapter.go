package grpc

import (
	"context"
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
			return nil
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
