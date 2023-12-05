package grpc

import (
	"context"
	"fmt"
	"grpc-go-server/data"
	"grpc-go-server/internal/port"
	"io"
	"log"
	pb "proto/protogen/go/bank"
	"time"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *GrpcAdapter) GetCurrentBalance(ctx context.Context, req *pb.CurrentBalanceRequest) (*pb.CurrentBalanceResponse, error) {
	balance, _ := a.BankService.FindCurrentBalance(req.AccountNumber)

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

func (a *GrpcAdapter) GetCurrentBalanceWithStatus(ctx context.Context, req *pb.CurrentBalanceRequest) (*pb.CurrentBalanceResponse, error) {
	balance, err := a.BankService.FindCurrentBalance(req.AccountNumber)
	if err != nil {
		s := NewAccountNotFoundError(req.AccountNumber)
		fmt.Errorf("Error getting current balance: %s\n", s.Err())
		return nil, s.Err()
	}

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
			a.BankService.StopExchangeRatesAtInterval()
			return nil
		default:
			now := time.Now().Truncate(time.Second)
			rate, err := a.BankService.GetExchangeRateAtTimestamp(req.FromCurrency, req.ToCurrency, now)
			if err != nil {
				s := status.New(codes.InvalidArgument, fmt.Sprintf("DB table BankExchangeRates error: %s", err))
				s, _ = s.WithDetails(&errdetails.ErrorInfo{
					Reason: "INVALID_CURRENCY",
					Domain: "DB table BankExchangeRates",
					Metadata: map[string]string{
						"from_currency": req.FromCurrency,
						"to_currency":   req.ToCurrency,
					},
				})
				s, _ = s.WithDetails(&errdetails.DebugInfo{})
				return s.Err()
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
			return nil
		}
	}
}

func (a *GrpcAdapter) SummarizeTransactions(stream pb.BankService_SummarizeTransactionsServer) error {
	var portTransactions []*port.Transaction = nil
	var pbTransactions []*pb.Transaction = nil
	var accountNumber string = ""
	var amount float64 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			balance, err := a.BankService.ExecuteBankTransactions(portTransactions)
			if err != nil {
				switch err {
				case data.ErrAccountNotFound:
					s := NewAccountNotFoundError(accountNumber)
					return s.Err()
				case data.ErrInsufficientBalance:
					s := NewInsufficientBalanceError(accountNumber, amount)
					return s.Err()
				default:
					s := NewInternalTableError("BankTransactions", data.ErrInternalTable)
					return s.Err()
				}
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
		amount = req.Amount
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

			for transferResponse := range a.BankService.ExecuteBankTransfers(req) {
				if transferResponse.Error != nil {
					switch transferResponse.Error {
					case data.ErrAccountNotFound:
						s := NewAccountNotFoundError(transferResponse.Response.FromAccountNumber)
						return s.Err()
					case data.ErrInsufficientBalance:
						s := NewInsufficientBalanceError(transferResponse.Response.FromAccountNumber, transferResponse.Response.Amount)
						return s.Err()
					default:
						s := NewInternalTableError("BankTransfers", data.ErrInternalTable)
						return s.Err()
					}
				}
				_ = stream.Send(transferResponse.Response)
				// if err != nil {
				// 	return err
				// }
				time.Sleep(500 * time.Millisecond)
			}
		}
	}
}

func (a *GrpcAdapter) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "No yet implemented")
}
