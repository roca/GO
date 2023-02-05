package main

import "fmt"

// Momento Design Pattern

type Memento struct {
	Balance int
}

type BankAcount struct {
	balance int
	changes []*Memento
	current int
}

// Stringer interface
func (b *BankAcount) String() string {
	return fmt.Sprintf("Balance = %d, current = %d", b.balance, b.current)
}

func NewBankAcount(balance int) *BankAcount {
	ba := &BankAcount{balance: balance}
	ba.changes = append(ba.changes, &Memento{balance})
	return ba
}

// Deposit adds amount to balance and returns a memento
func (b *BankAcount) Deposit(amount int) *Memento {
	b.balance += amount
	m := &Memento{b.balance}
	b.changes = append(b.changes, m)
	b.current++
	fmt.Printf("Deposited %d, balance is now %d\n", amount, b.balance)
	return m
}

// Restore restores the balance to that of the memento
func (b *BankAcount) Restore(m *Memento) {
	if m != nil {
		b.balance = m.Balance
		b.changes = append(b.changes, m)
		b.current = len(b.changes) - 1
	}
}

// Undo reverts to previous balance
func (b *BankAcount) Undo() *Memento {
	if b.current > 0 {
		b.current--
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}
	return nil
}

// Redo reverts to next balance
func (b *BankAcount) Redo() *Memento {
	if b.current+1 < len(b.changes) {
		b.current++
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}
	return nil
}

func main() {
	ba := NewBankAcount(100)
	ba.Deposit(50)
	ba.Deposit(25)
	fmt.Println(ba)
	ba.Undo()
	fmt.Println("Undo 1:", ba)
	ba.Undo()
	fmt.Println("Undo 2:", ba)
	ba.Redo()
	fmt.Println("Redo 1:", ba)
}
