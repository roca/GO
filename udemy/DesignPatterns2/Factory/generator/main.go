package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// functional
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{
			Name:         name,
			Position:     position,
			AnnualIncome: annualIncome,
		}
	}
}

type EmployeeFactory struct {
	Position string
	AnnualIncome int
}
func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{
		Name:         name,
		Position:     f.Position,
		AnnualIncome: f.AnnualIncome,
	}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{
		Position: position,
		AnnualIncome: annualIncome,
	}
}

func main() {
	developerFactory := NewEmployeeFactory("developer", 100000)
	managerFactory := NewEmployeeFactory("manager", 200000)

	dmitri := developerFactory("Dmitri")
	frank := managerFactory("Frank")

	fmt.Printf("%+v\n", *dmitri)
	fmt.Printf("%+v\n", *frank)

	developerFactory2 := NewEmployeeFactory2("developer", 130000)
	bossFactory2 := NewEmployeeFactory2("CEO", 260000)

	ahmed := developerFactory2.Create("Ahmed")
	bob := bossFactory2.Create("Bob")
	bob.AnnualIncome = 100000

	fmt.Printf("%+v\n", *ahmed)
	fmt.Printf("%+v\n", *bob)

}
