package main

import (
	"fmt"
	"strings"
)

type person struct {
	name    string
	address string
	city    string
	state   string
	zip     string
}

type people []person

func (p people) String() string {
	lines := []string{}
	for _, v := range p {
		lines = append(lines, fmt.Sprint(v.name, " | ", v.address, " ", v.city, " ", v.state, " ", v.zip))
	}

	return strings.Join(lines, "\n")
}

func main() {
	peeps := people{
		{name: "John", address: "123 Main St", city: "Jamestown", state: "NY", zip: "14701"},
		{name: "Jane", address: "234 Fleet St", city: "Columbia", state: "MD", zip: "21150"},
		{name: "Terry", address: "345 Charles Ave", city: "Gergetown", state: "DC", zip: "20007"},
	}

	fmt.Println(peeps)
}
