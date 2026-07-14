package main

// Momento Design Pattern

type BankAcount struct {
	balance int
}

type Memento struct {
	Balance int
}

func (b *BankAcount) Deposit(amount int) *Memento {
	b.balance += amount
	return &Memento{b.balance}
}

func (b *BankAcount) Restore(m *Memento) {
	b.balance = m.Balance
}

func NewBankAcount(balance int) (*BankAcount, *Memento) {
	return &BankAcount{balance}, &Memento{balance}
}

func main() {
	ba,m0  := NewBankAcount(100)
	m1 := ba.Deposit(50)
	m2 := ba.Deposit(25)
	println(ba.balance)
	ba.Restore(m1)
	println(ba.balance)
	ba.Restore(m2)
	println(ba.balance)

	ba.Restore(m0)
	println(ba.balance)
}
