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

func (b *BankAccount) Withdraw(amount int) bool {
	if b.balance-amount >= overdraftlimit {
		b.balance -= amount
		fmt.Println("Withdrew", amount, ", balance is now", b.balance)
		return true
	}
	return false
}

type Command interface {
	Do()
	Undo()
}

type Action int

const (
	Deposit Action = iota
	Withdraw
)

type BankAccountCommand struct {
	account   *BankAccount
	action    Action
	amount    int
	succeeded bool
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account: account, action: action, amount: amount}
}

func (b *BankAccountCommand) Do() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
		b.succeeded = true
	case Withdraw:
		b.succeeded = b.account.Withdraw(b.amount)
	}
}

func (b *BankAccountCommand) Undo() {
	if !b.succeeded {
		return
	}
	switch b.action {
	case Deposit:
		b.account.Withdraw(b.amount)
	case Withdraw:
		b.account.Deposit(b.amount)
	}
}

func main() {
	ba := BankAccount{}
	cmd := NewBankAccountCommand(&ba, Deposit, 100)
	cmd.Do()
	cmd2 := NewBankAccountCommand(&ba, Withdraw, 25)
	cmd2.Do()
	cmd2.Undo()
}
