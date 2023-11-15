package main

import (
	"grpc-go-server/data"
	mygrpc "grpc-go-server/internal/adapter/grpc"
	app "grpc-go-server/internal/application"
	"log"

	"github.com/google/uuid"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	hs := &app.HelloService{}
	bs := &app.BankService{}

	c := initApplication()
	//c.App.ListenAndServe()
	bankAccount := data.BankAccount{
		ID:              uuid.New(),
		AccountNumber:   "1234567890",
		AccountName:     "John Doe",
		Currency:        "PHP",
		CurrencyBalance: 1000.00,
	}

	id, _ := c.Models.BankAccounts.Insert(bankAccount)
	log.Println("Inserted", id)

	grpcadapter := mygrpc.NewGrpcAdapter(hs, bs, 9090)

	grpcadapter.Run()
}
