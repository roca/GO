package application

import "log"

type BankService struct{}

func (b *BankService) GetCurrentBalance(accountNumber int) float32 {
	log.Println(accountNumber)
	return 100.0 * float32(accountNumber)
}