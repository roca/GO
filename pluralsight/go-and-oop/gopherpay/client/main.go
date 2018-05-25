package main

import "github.com/gopherpay/paybroker"

type PaymentOption interface {
	ProcessPayment(float32) bool
}

func main() {
	var option PaymentOption

	option = &paybroker.PaymentBrokerAccount{}

	option.ProcessPayment(500)
}
