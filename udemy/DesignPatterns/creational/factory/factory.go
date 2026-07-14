package factory

import (
	"errors"
	"fmt"
)

type IPaymentMethod interface {
	Pay(amount float32) string
}

//Our current implemented Payment methods ar described here
const (
	Cash      = 1
	DebitCard = 2
)

//CreditPaymentMethod returns a pointer to a PaymentMethod object or an error
//if the method is not registered.
func GetPaymentMethod(m int) (IPaymentMethod, error) {
	paymentMethod := map[int]IPaymentMethod{
		1: new(CashPM),
		2: new(DebitCardPM),
	}
	if m, ok := paymentMethod[m]; ok {
		return m, nil
	}
	return nil, errors.New(fmt.Sprintf("Payment method %d not recognized", m))
}

type CashPM struct{}
type DebitCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("payed using cash. amount: %#0.2f", amount)
}
func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("payed using debit card. amount: %#0.2f", amount)
}

type NewDebitCardPM struct{}
func (c *NewDebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("payed using debit card (new). amount: %#0.2f", amount)
}
type CreditCardPM struct{}
func (c *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("payed using debit card (cash). amount: %#0.2f", amount)
}