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
	Succeeded() bool
	SetSucceeded(value bool)
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

func (b *BankAccountCommand) Succeeded() bool {
	return b.succeeded
}

func (b *BankAccountCommand) SetSucceeded(value bool) {
	b.succeeded = value
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

type CompositeBankAccountCommand struct {
	commands []Command
}

func (c *CompositeBankAccountCommand) Do(){
	for _, cmd := range c.commands {
		cmd.Do()
	}
}
func (c *CompositeBankAccountCommand) Undo(){
	for i := len(c.commands)-1; i >= 0; i-- {
		c.commands[i].Undo()
	}
}
func (c *CompositeBankAccountCommand) Succeeded() bool{
	for _, cmd := range c.commands {
		if !cmd.Succeeded() {
			return false
		}
	}
	return true
}
func (c *CompositeBankAccountCommand) SetSucceeded(value bool){
	for _, cmd := range c.commands {
		cmd.SetSucceeded(value)
	}
}

type MoneyTransferCommand struct {
	CompositeBankAccountCommand
	from, to *BankAccount
	amount int
}

func NewMoneyTransferCommand(from, to *BankAccount, amount int) *MoneyTransferCommand {
	c := &MoneyTransferCommand{from: from, to: to, amount: amount}
	c.commands = append(c.commands, NewBankAccountCommand(from, Withdraw, amount))
	c.commands = append(c.commands, NewBankAccountCommand(to, Deposit, amount))
	return c
}

func (m *MoneyTransferCommand) Do() {
	ok := true
	for _, cmd := range m.commands {
		if ok {
			cmd.Do()
			ok = cmd.Succeeded()
		} else {
			cmd.SetSucceeded(false)
		}
	}
}

func main() {
	from := &BankAccount{100}
	to := &BankAccount{0}
	mtc := NewMoneyTransferCommand(from, to, 25)
	mtc.Do()
	fmt.Println(from, to)
	mtc.Undo()
	fmt.Println(from, to)
}
