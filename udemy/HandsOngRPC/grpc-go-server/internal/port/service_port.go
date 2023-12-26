package port

import (
	"grpc-go-server/data"
	"time"

	pb "protogen/go/bank"

	"github.com/google/uuid"
)

type TransactionType int32

const (
	TRANSACTION_TYPE_UNSPECIFIED TransactionType = 0
	TRANSACTION_TYPE_DEPOSIT     TransactionType = 1
	TRANSACTION_TYPE_WITHDRAWAL  TransactionType = 2
)

type Transaction struct {
	AccountNumber   string
	TransactionType TransactionType
	Amount          float64
}

var TransactionMap = map[TransactionType]string{
	TRANSACTION_TYPE_UNSPECIFIED: "UNSPECIFIED",
	TRANSACTION_TYPE_DEPOSIT:     "DEPOSIT",
	TRANSACTION_TYPE_WITHDRAWAL:  "WITHDRAWAL",
}

type TransferStatusType int32

const (
	TRANSFER_STATUS_UNSPECIFIED TransferStatusType = 0
	TRANSFER_STATUS_SUCCESS     TransferStatusType = 1
	TRANSFER_STATUS_FAILED      TransferStatusType = 2
)

var TransferStatusTypeMap = map[TransferStatusType]string{
	TRANSFER_STATUS_UNSPECIFIED: "UNSPECIFIED",
	TRANSFER_STATUS_SUCCESS:     "SUCCESS",
	TRANSFER_STATUS_FAILED:      "FAILED",
}

type TransferResponse struct {
	Response *pb.TransferResponse
	Error    error
}

type HelloServicePort interface {
	GenerateHello(name string) string
	GenerateManyHellos(name string, count int) []string
	GenerateHelloToEveryone(names []string) string
	GenerateContinuousHello(name string, count int) <-chan string
}

type BankServicePort interface {
	Save(data data.BankAccount) (uuid.UUID, error)
	FindCurrentBalance(uuid string) (float64, error)
	InsertExchangeRatesAtInterval(exit chan bool, fromCurrency, toCurrency string, interval time.Duration)
	GetExchangeRateAtTimestamp(fromCurrency, toCurrency string, timestamp time.Time) (*data.BankExchangeRate, error)
	StopExchangeRatesAtInterval()
	ExecuteBankTransactions(transactions []*Transaction) (float64, error)
	ExecuteBankTransfers(*pb.TransferRequest) <-chan *TransferResponse
}

type ResiliencyRequest struct {
	MaxDelaySecond int32
	MinDelaySecond int32
	StatusCodes    []uint32
}

type ResiliencyResponse struct {
	Response   string
	StatusCode uint32
	Error      error
}

type ResiliencyServicePort interface {
	GetResiliency(*ResiliencyRequest) (*ResiliencyResponse, error)
	GetResiliencyStream(*ResiliencyRequest) (*ResiliencyResponse, error)
	SendResiliencyStream([]*ResiliencyRequest) (*ResiliencyResponse, error)
	BidirectionalResiliencyStream(*ResiliencyRequest) <-chan *ResiliencyResponse
}
