package main

import "fmt"

type CreditAccount struct{}

func (c *CreditAccount) AvailableFunds() float32 {
	fmt.Println("Getting credit funds")
	return 250
}

type CheckingAccount struct{}

func (c *CheckingAccount) AvailableFunds() float32 {
	fmt.Println("Getting checking funds")
	return 125
}

type HybridAccount struct {
	creditAccount   CreditAccount
	checkingAccount CheckingAccount
}

func (h *HybridAccount) AvailableFunds() float32 {
	return h.creditAccount.AvailableFunds()
}

func main() {
	ha := &HybridAccount{}
	fmt.Println(ha.AvailableFunds())
	fmt.Println(ha.checkingAccount.AvailableFunds())
}
