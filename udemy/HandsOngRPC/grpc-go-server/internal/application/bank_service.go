package application

import (
	"grpc-go-server/data"

	"github.com/google/uuid"
)

type BankService struct {
	Models data.Models
}

func (b *BankService) Save(data data.BankAccount) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

func (b *BankService) FindCurrentBalance(uuid string) float64 {
	// log.Println(acct)

	account, _ := b.Models.BankAccounts.Get(uuid)

	return account.CurrentBalance
}
