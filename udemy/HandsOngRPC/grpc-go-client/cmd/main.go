package main

import (
	"context"
	"grpc-go-client/internal/adapter/bank"
	"grpc-go-client/internal/adapter/hello"
	"grpc-go-client/internal/adapter/resiliency"
	"log"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithUnaryInterceptor(
		grpc_retry.UnaryClientInterceptor(
			grpc_retry.WithCodes(codes.Unknown, codes.Internal),
			grpc_retry.WithMax(4),
			grpc_retry.WithBackoff(
				grpc_retry.BackoffExponential(2*time.Second),
			),
		),
	))

	opts = append(opts, grpc.WithStreamInterceptor(
		grpc_retry.StreamClientInterceptor(
			grpc_retry.WithCodes(codes.Unknown, codes.Internal),
			grpc_retry.WithMax(4),
			grpc_retry.WithBackoff(
				grpc_retry.BackoffLinear(4*time.Second),
			),
		),
	))

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("localhost:9090", opts...)
	if err != nil {
		log.Fatalln("Can not connect to gRPC server:", err)
	}
	defer conn.Close()

	// helloAdapter, err := hello.NewHelloAdapter(conn)
	// if err != nil {
	// 	log.Fatalln("Can not create HelloAdapter:", err)
	// }

	// bankAdapter, err := bank.NewBankAdapter(conn)
	// if err != nil {
	// 	log.Fatalln("Can not create HelloAdapter:", err)
	// }

	resiliencyAdapter, err := resiliency.NewResiliencyAdapter(conn)
	if err != nil {
		log.Fatalln("Can not create ResiliencyAdapter:", err)
	}

	// runSayHello(helloAdapter, "Bruce Wayne")
	// runSayManyHellos(helloAdapter, "Bruce Wayne")
	// runSayHelloToEveryone(helloAdapter, []string{"Bruce Wayne", "Clark Kent", "Diana Prince"})
	// runSayHelloContinuous(helloAdapter, []string{"Anna", "Bella", "Carol", "Diana", "Emma"})

	//runGetCurrentBalance(bankAdapter, "1e9230bd-4264-4526-a9cd-2a86d3ca9590")
	// runFetchExchangeRates(bankAdapter, "USD", "IDR")
	// runSummarizeTransactions(bankAdapter, []bank.BankTransaction{
	// 	{AccountNumber: "1e9230bd-4264-4526-a9cd-2a86d3ca9594", Amount: 5, TransactionType: 1},
	// 	{AccountNumber: "1e9230bd-4264-4526-a9cd-2a86d3ca9594", Amount: 5, TransactionType: 2},
	// })

	// runTransferMultiple(bankAdapter, []bank.BankTransfer{
	// 	{
	// 		Amount:            1,
	// 		Currency:          "USD",
	// 		FromAccountNumber: "1e9230bd-4264-4526-a9cd-2a86d3ca9594",
	// 		ToAccountNumber:   "2a7d5f68-baa1-4264-bf41-facba0414c59",
	// 	},
	// 	{
	// 		Amount:            1,
	// 		Currency:          "USD",
	// 		FromAccountNumber: "1e9230bd-4264-4526-a9cd-2a86d3ca9594",
	// 		ToAccountNumber:   "2a7d5f68-baa1-4264-bf41-facba0414c59",
	// 	},
	// })

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	runGetResiliency(ctx, resiliencyAdapter, &resiliency.ResiliencyRequest{
		MaxDelaySecond: 5,
		MinDelaySecond: 4,
		StatusCodes:    []uint32{0},
	})

	// runGetResiliencyStream(resiliencyAdapter, &resiliency.ResiliencyRequest{
	// 	MaxDelaySecond: 5000,
	// 	MinDelaySecond: 4000,
	// 	StatusCodes:    []uint32{0},
	// })
}

func runSayHello(adapter *hello.HelloAdapter, name string) {
	greet, err := adapter.SayHello(context.Background(), name)
	if err != nil {
		log.Fatalln("Can not invoke SayHello on the HelloAdapter:", err)
	}

	log.Println(greet.Greet)
}

func runSayManyHellos(adapter *hello.HelloAdapter, name string) {
	err := adapter.SayManyHellos(context.Background(), name)
	if err != nil {
		log.Fatalln("Can not invoke SayManyHellos on the HelloAdapter:", err)
	}
}

func runSayHelloToEveryone(adapter *hello.HelloAdapter, names []string) {
	err := adapter.SayHelloToEveryone(context.Background(), names)
	if err != nil {
		log.Fatalln("Can not invoke SayHelloToEveryone on the HelloAdapter:", err)
	}
}

func runSayHelloContinuous(adapter *hello.HelloAdapter, names []string) {
	err := adapter.SayHelloContinuous(context.Background(), names)
	if err != nil {
		log.Fatalln("Can not invoke SayHelloContinuouse on the HelloAdapter:", err)
	}
}

func runGetCurrentBalance(adapter *bank.BankAdapter, accountNumber string) {
	resp, err := adapter.GetCurrentBalanceWithStatus(context.Background(), accountNumber)
	if err != nil {
		s := status.Convert(err)
		log.Fatalln("Can not invoke GetCurrentBalance on the BankAdapter:", "\nCode:", s.Code(), "\nMessage:", s.Message(), "\nDetails:", s.Details())
	}

	log.Println("CurrentBalance:", resp.CurrentBalance)
}

func runFetchExchangeRates(adapter *bank.BankAdapter, fromCurrency, toCurrency string) {
	err := adapter.FetchExchangeRates(context.Background(), fromCurrency, toCurrency)
	if err != nil {
		s := status.Convert(err)
		log.Fatalln("Can not invoke FetchExchangeRates on the BankAdapter:", "\nCode:", s.Code(), "\nMessage:", s.Message(), "\nDetails:", s.Details())
	}
}

func runSummarizeTransactions(adapter *bank.BankAdapter, transactions []bank.BankTransaction) {
	resp, err := adapter.SummarizeTransactions(context.Background(), transactions)
	if err != nil {
		s := status.Convert(err)
		log.Fatalln("Can not invoke SummarizeTransactions on the BankAdapter:", "\nCode:", s.Code(), "\nMessage:", s.Message(), "\nDetails:", s.Details())
	}

	log.Println("CurrentBalance:", resp.Balance)
}

func runTransferMultiple(adapter *bank.BankAdapter, transfers []bank.BankTransfer) {
	err := adapter.TransferMultiple(context.Background(), transfers)
	if err != nil {
		s := status.Convert(err)
		log.Fatalln("Can not invoke SummarizeTransactions on the BankAdapter:", "\nCode:", s.Code(), "\nMessage:", s.Message(), "\nDetails:", s.Details())
	}
}

func runGetResiliency(ctx context.Context, adapter *resiliency.ResiliencyAdapter, resiliencyRequest *resiliency.ResiliencyRequest) {
	resp, err := adapter.GetResiliency(ctx, resiliencyRequest)
	if err != nil {
		s := status.Convert(err)
		log.Fatalln("Can not invoke GetResiliency on the ResiliencyAdapter:", "\nCode:", s.Code(), "\nMessage:", s.Message(), "\nDetails:", s.Details())
	}
	log.Println(resp)
}

func runGetResiliencyStream(ctx context.Context, adapter *resiliency.ResiliencyAdapter, resiliencyRequest *resiliency.ResiliencyRequest) {
	err := adapter.GetResiliencyStream(ctx, resiliencyRequest)
	if err != nil {
		s := status.Convert(err)
		log.Fatalln("Can not invoke GetResiliency on the ResiliencyAdapter:", "\nCode:", s.Code(), "\nMessage:", s.Message(), "\nDetails:", s.Details())
	}
}

func runSendResiliencyStream(ctx context.Context, adapter *resiliency.ResiliencyAdapter, resiliencyRequests []*resiliency.ResiliencyRequest) {
	resp, err := adapter.SendResiliencyStream(ctx, resiliencyRequests)
	if err != nil {
		s := status.Convert(err)
		log.Fatalln("Can not invoke GetResiliency on the ResiliencyAdapter:", "\nCode:", s.Code(), "\nMessage:", s.Message(), "\nDetails:", s.Details())
	}
	log.Println(resp)
}

func runBidirectionalResiliencyStream(ctx context.Context, adapter *resiliency.ResiliencyAdapter, resiliencyRequests []*resiliency.ResiliencyRequest) {
	err := adapter.BidirectionalResiliencyStream(ctx, resiliencyRequests)
	if err != nil {
		s := status.Convert(err)
		log.Fatalln("Can not invoke SummarizeTransactions on the BankAdapter:", "\nCode:", s.Code(), "\nMessage:", s.Message(), "\nDetails:", s.Details())
	}
}
