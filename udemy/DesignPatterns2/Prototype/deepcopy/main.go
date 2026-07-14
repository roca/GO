package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

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

func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

func (p *Person) DeepCopy2() *Person { //This version will uses serialization method
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	//fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)

	return &result

}

func main() {
	john := Person{"John", &Address{"123 London Road", "London", "UK"}, []string{"Alice", "Bob", "Charlie"}}

	// Example 1: There is a problem here
	// jane := john // copy by value
	// jane.Name = "Jane" // Ok, not changing john's Name
	// jane.Address.StreetAddress = "321 Baker Street"

	// Example 2: This is better
	// jane := john // copy by reference
	// jane.Name = "Jane" // Ok, not changing john's Name
	// jane.Address = &Address{"321 Baker Street", "London", "UK"} // Ok, changing jane's Address`s StreetAddress`

	// Example 3: This is even better!
	jane := *john.DeepCopy2()
	jane.Name = "Jane"                                          // Ok, not changing john's Name
	jane.Address = &Address{"321 Baker Street", "London", "UK"} // Ok, changing jane's Address`s StreetAddress`
	jane.Friends = append(jane.Friends, "Jack")

	fmt.Printf("%p %v\n", &john, john)
	fmt.Printf("%p %v\n", &jane, jane)

}
