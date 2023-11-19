package port

import (
	"grpc-go-server/data"

	"github.com/google/uuid"
)

type HelloServicePort interface {
	GenerateHello(name string) string
	GenerateManyHellos(name string, count int) []string
	GenerateHelloToEveryone(names []string) string
	GenerateContinuousHello(name string, count int) <-chan string
}

type BankServicePort interface {
	Save(data data.BankAccount) (uuid.UUID, error)
	FindCurrentBalance(uuid string) float64
}
