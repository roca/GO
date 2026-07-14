package main

import (
	"fmt"
	"log"
)

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Println("Hello, I am", p.name, "and I am", p.age, "years old")
}

type tiredPerson struct {
	name string
	age  int
}

func (p *tiredPerson) SayHello() {
	fmt.Println("Sorry, I'm too tired to say hello")
}

func NewPerson(name string, age int) (Person, error) {
	if age <= 0 {
		err := fmt.Errorf("Age %d is invalid\n", age)
		return nil, err
	}
	if age > 100 {
		return &tiredPerson{name, age}, nil
	}
	return &person{name, age}, nil
}

func main() {
	p, err := NewPerson("John", 30)
	if err != nil {
		log.Fatal(err)
	}
	p.SayHello()
	p2, err := NewPerson("Edward", 101)
	if err != nil {
		log.Fatal(err)
	}
	p2.SayHello()
}
