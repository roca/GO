package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		StreetAddress: a.StreetAddress,
		City:          a.City,
		Country:       a.Country,
	}
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person{
	q := p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return q
}

func main() {
	john := Person{"John", &Address{"123 London Road", "London", "UK"}}

	// jane := john // copy by value
	// jane.Address.StreetAddress = "321 Baker Street"

	//n deep copying

	jane := john // copy by reference
	jane.Name = "Jane" // Ok, not changing john's Name
	jane.Address = &Address{"321 Baker Street", "London", "UK"} // Ok, changing jane's Address`s StreetAddress`

	fmt.Println(john,john.Address)
	fmt.Println(jane,jane.Address)
}
