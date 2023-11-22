package port

import (
	"grpc-go-server/data"
	"time"

	"github.com/google/uuid"
)

type TransactionType int32

const (
	TransactionType_TRANSACTION_TYPE_UNSPECIFIED TransactionType = 0
	TransactionType_TRANSACTION_TYPE_DEPOSIT     TransactionType = 1
	TransactionType_TRANSACTION_TYPE_WITHDRAWAL  TransactionType = 2
)


type Transaction struct {
	AccountNumber   string
	TransactionType TransactionType
	Amount          float64
}

type HelloServicePort interface {
	GenerateHello(name string) string
	GenerateManyHellos(name string, count int) []string
	GenerateHelloToEveryone(names []string) string
	GenerateContinuousHello(name string, count int) <-chan string
}

type BankServicePort interface {
	Save(data data.BankAccount) (uuid.UUID, error)
	FindCurrentBalance(uuid string) float64
	InsertExchangeRatesAtInterval(exit chan bool, fromCurrency, toCurrency string, interval time.Duration)
	GetExchangeRateAtTimestamp(fromCurrency, toCurrency string, timestamp time.Time) (*data.BankExchangeRate, error)
	StopExchangeRatesAtInterval()
	ExecuteBankTransactions(transactions []*Transaction) (float64, error)
}
