package main

import (
	"context"
	"grpc-go-client/internal/adapter/bank"
	"grpc-go-client/internal/adapter/hello"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	var opts []grpc.DialOption
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

	bankAdapter, err := bank.NewBankAdapter(conn)
	if err != nil {
		log.Fatalln("Can not create HelloAdapter:", err)
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

	runTransferMultiple(bankAdapter, []bank.BankTransfer{
		// TODO
	})
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
