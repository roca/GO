package application

import (
	"log"
)

type BankService struct{}

func (b *BankService) FindCurrentBalance(acct string) float64 {
	log.Println(acct)

	return 100.0
}
