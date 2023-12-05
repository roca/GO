package port

import (
	"context"

	"google.golang.org/grpc"

	pb "proto/protogen/go/bank"
)

type BankClientPort interface {
	GetCurrentBalanceWithStatus(ctx context.Context, in *pb.CurrentBalanceRequest, opts ...grpc.CallOption) (*pb.CurrentBalanceResponse, error)
	FetchExchangeRates(ctx context.Context, in *pb.ExchangeRateRequest, opts ...grpc.CallOption) (pb.BankService_FetchExchangeRatesClient, error)
	SummarizeTransactions(ctx context.Context, opts ...grpc.CallOption) (pb.BankService_SummarizeTransactionsClient, error)
	TransferMultiple(ctx context.Context, opts ...grpc.CallOption) (pb.BankService_TransferMultipleClient, error)
}
