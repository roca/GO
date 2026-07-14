package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "developer", 60000}
	case Manager:
		return &Employee{"", "manager", 90000}
	default:
		panic("No role found")
	}
	return nil
}

func main() {
 m := NewEmployee(Manager)
 m.Name = "Sam"
 fmt.Println(m)
}