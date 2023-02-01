package main

import "fmt"


type BankAccount struct {
	Balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.Balance += amount
	fmt.Println("Deposited", amount, ", balance is now", b.Balance)
}

func (b *BankAccount) Withdraw(amount int)  {
	if b.Balance >= amount  {
		b.Balance -= amount
		fmt.Println("Withdrew", amount, ", balance is now", b.Balance)
	}
}



func main() {
	ba := BankAccount{0}
	var commands []func()
	commands = append(commands, func() {
		ba.Deposit(100)
	})
	commands = append(commands, func() {
		ba.Withdraw(25)
	})

	for _, cmd := range commands {
		cmd()
	}
}
