package main

import "fmt"

var overdraftlimit = -500

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Println("Deposited", amount, ", balance is now", b.balance)
}

func (b *BankAccount) With(amount int) {
	if b.balance-amount >= overdraftlimit {
		b.balance -= amount
		fmt.Println("Withdrew", amount, ", balance is now", b.balance)
	}
}

type Command interface {
	Call()
}

type Action int
const (
	Deposit Action = iota
	Withdraw
)

type BankAccountCommand struct {
	account *BankAccount
	action Action
	amount int
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account, action, amount}
}

func (b *BankAccountCommand) Call() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
	case Withdraw:
		b.account.With(b.amount)
	}
}

func main() {
	ba := BankAccount{}
	cmd := NewBankAccountCommand(&ba, Deposit, 100)
	cmd.Call()
	cmd2 := NewBankAccountCommand(&ba, Withdraw, 50)
	cmd2.Call()
}
