package main

import (
	"context"
	"grpc-go-client/internal/adapter/bank"
	"grpc-go-client/internal/adapter/hello"
	"grpc-go-client/internal/adapter/resiliency"
	"log"
	"time"

	dresl "grpc-go-client/internal/application/domain/resiliency"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"github.com/sony/gobreaker"
	breaker "github.com/sony/gobreaker"
)

var cb *breaker.CircuitBreaker

func init() {
	var st breaker.Settings
	st.Name = "Unary GetResiliency"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		counts.TotalFailures = 6
		counts.Requests = 10
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		log.Printf("Circuit breaker failure is %v, requests is %v, means failure ratio : %v\n",
			counts.TotalFailures, counts.Requests, failureRatio)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}

	st.Timeout = 4 * time.Second
	st.MaxRequests = 3
	st.OnStateChange = func(name string, from gobreaker.State, to gobreaker.State) {
		log.Printf("Circuit breaker state changed from %v to %v\n", from, to)
	}

	cb = breaker.NewCircuitBreaker(st)
}

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	var opts []grpc.DialOption
	// opts = append(opts, grpc.WithUnaryInterceptor(
	// 	grpc_retry.UnaryClientInterceptor(
	// 		grpc_retry.WithCodes(codes.Unknown, codes.Internal),
	// 		grpc_retry.WithMax(4),
	// 		grpc_retry.WithBackoff(
	// 			grpc_retry.BackoffExponential(2*time.Second),
	// 		),
	// 	),
	// ))

	// opts = append(opts, grpc.WithStreamInterceptor(
	// 	grpc_retry.StreamClientInterceptor(
	// 		grpc_retry.WithCodes(codes.Unknown, codes.Internal),
	// 		grpc_retry.WithMax(4),
	// 		grpc_retry.WithBackoff(
	// 			grpc_retry.BackoffLinear(3*time.Second),
	// 		),
	// 	),
	// ))

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

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	ctx := context.Background()

	// runGetResiliency(ctx, resiliencyAdapter, &resiliency.ResiliencyRequest{
	// 	MaxDelaySecond: 10,
	// 	MinDelaySecond: 8,
	// 	StatusCodes:    []uint32{dresl.StatusCode_OK},
	// })

	// runGetResiliencyStream(ctx, resiliencyAdapter, &resiliency.ResiliencyRequest{
	// 	MaxDelaySecond: 3,
	// 	MinDelaySecond: 0,
	// 	StatusCodes:    []uint32{dresl.StatusCode_OK},
	// })

	for i := 0; i < 300; i++ {
		runGetResiliencyWithCiruitBreaker(ctx, resiliencyAdapter, &resiliency.ResiliencyRequest{
			MaxDelaySecond: 0,
			MinDelaySecond: 0,
			StatusCodes:    []uint32{dresl.StatusCode_Unknown, dresl.StatusCode_OK},
		})
		time.Sleep(time.Second)
	}
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

func runGetResiliencyWithCiruitBreaker(ctx context.Context, adapter *resiliency.ResiliencyAdapter, resiliencyRequest *resiliency.ResiliencyRequest) {
	cbreakerRes, cbreakerErr := cb.Execute(func() (interface{}, error) {
		return adapter.GetResiliency(ctx, resiliencyRequest)
	})
	if cbreakerErr != nil {
		s := status.Convert(cbreakerErr)
		log.Println("Can not invoke GetResiliency on the ResiliencyAdapter:", "\nCode:", s.Code(), "\nMessage:", s.Message(), "\nDetails:", s.Details())
		return
	}
	if s, ok := cbreakerRes.(*resiliency.ResiliencyResponse); ok {
		log.Println(s.Response)
	} else {
		log.Println(string(cbreakerRes.([]byte)))
	}
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
