package main

import "fmt"

type Person struct {
	// address: Separate builder 1
	StreetAddress, Postcode, City string

	// job: Separate builder 2
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{b}
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct {
	*PersonBuilder
}

func (b *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	b.person.StreetAddress = streetAddress
	return b
}

func (b *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	b.person.Postcode = postcode
	return b
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}

type PersonJobBuilder struct {
	*PersonBuilder
}

func (b *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	b.person.CompanyName = companyName
	return b
}

func (b *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	b.person.Position = position
	return b
}

func (b *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	b.person.AnnualIncome = annualIncome
	return b
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func main() {
	pb := NewPersonBuilder()
	pb.
		Lives().
		At("123 London Road").
		In("London").
		WithPostcode("SW12BC")
	pb.
		Works().
		At("Fabrikam").
		AsA("Engineer").
		Earning(123000)
	person := pb.Build()
	fmt.Println(person)
}
