package main

import (
	"fmt"
	"grpc-go-server/data"
	mygrpc "grpc-go-server/internal/adapter/grpc"
	app "grpc-go-server/internal/application"
	"log"
	"os"
	"time"

	"github.com/roca/celeritas"
)

var cel *celeritas.Celeritas

func init() {
	cel = &celeritas.Celeritas{}
	cel.AppName = "grpc-go-server"
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	log.SetFlags(0)
	log.SetOutput(logWriter{})
}

func main() {
	exit := make(chan bool)
	defer func() {
		exit <- true
		fmt.Println("Finished InsertExchangeRatesAtInterval")
	}()

	hs := &app.HelloService{}
	bs := &app.BankService{
		Models:   data.New(cel.DB.Pool),
		ExitChan: exit,
	}

	grpcadapter := mygrpc.NewGrpcAdapter(hs, bs, 9090)

	go bs.InsertExchangeRatesAtInterval(exit, "USD", "IDR", 5*time.Second)

	grpcadapter.Run()
}
