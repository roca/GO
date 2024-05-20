package main

import "fmt"

func main() {
	myAddress := CreateAddress().
		SetStreet("Main St.").
		SetNumber(11).
		SetCity("New York").
		SetCountry("USA")

		fmt.Println("My Address is:",myAddress)
}

type Address struct {
	street  string
	number  int32
	city    string
	country string
}

func CreateAddress() *Address {
	return &Address{}
}

func (a *Address) SetStreet(street string) *Address {
	a.street = street
	return a
}

func (a *Address) SetNumber(number int32) *Address {
	a.number = number
	return a
}

func (a *Address) SetCity(city string) *Address {
	a.city = city
	return a
}

func (a *Address) SetCountry(country string) *Address {
	a.country = country
	return a
}
