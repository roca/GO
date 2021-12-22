package main

import (
	"fmt"
	"log"
)

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

func NewPerson(name string, age int) (*Person, error) {
	if age <= 0 {
		err := fmt.Errorf("Age %d is invalid\n", age)
		return nil, err
	}
	return &Person{
		Name:     name,
		Age:      age,
		EyeCount: 2,
	}, nil
}

func main() {
	p, err := NewPerson("John", 0)
	if err != nil {
		log.Fatal(err)
	}
	println(p.Name)
	println(p.Age)
	p.EyeCount = 1
	println(p.EyeCount)
}
