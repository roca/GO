package main

import "fmt"

type Payment struct {
	From   string
	To     string
	Amount float64 // USD
}

func (p *Payment) Process() {
	fmt.Printf("%s - > $%.2f -> %s\n", p.From, p.Amount, p.To)
}
