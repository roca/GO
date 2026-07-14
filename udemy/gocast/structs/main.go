package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jim.updateFirstName("jimmy")
	jim.print()
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	pointerToPerson.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
