package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office *Address
}

func (e *Employee) String() string {
	return fmt.Sprintf("%s works at Suite: %d, %s, %s", e.Name, e.Office.Suite, e.Office.StreetAddress, e.Office.City)
}

func (emp *Employee) DeepCopy() *Employee { //This version will uses serialization method
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(emp)

	// peek into the structure
	// fmt.Println(b.String())

	d := gob.NewDecoder(&b)
	result := Employee{}
	_ = d.Decode(&result)

	return &result

}

var mainOffice = Employee{"", &Address{0, "123 East Dr.", "London"}}
var auxOffice = Employee{"", &Address{0, "66 West Dr.", "London"}}

func newEmployee(proto *Employee, name string, suite int) *Employee{
	emp := proto.DeepCopy()
	emp.Name = name
	emp.Office.Suite = suite
	return emp
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}

func main() {
	john := NewMainOfficeEmployee("John", 100)
	jane := NewAuxOfficeEmployee("Jane", 80)

	fmt.Println(john.String())
	fmt.Println(jane.String())
}
