// File name: ...\s05\05_maps_employee\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {

	employeeSalary := map[string]float64{
		"Blake":  60000.00,
		"Parker": 120000.50,
		"Dakota": 93000.00,
	}

	fmt.Printf("Parker's Salary = %.2f\n", employeeSalary["Parker"])

	salary1, ok := employeeSalary["Dakota"]
	if ok {
		fmt.Printf("Value %.2f was found.\n", salary1)
	} else {
		fmt.Printf("The specified key was not found.\n")
	}

	if salary2, ok := employeeSalary["Jordan"]; ok {
		fmt.Printf("Value %.2f was found.\n", salary2)
		delete(employeeSalary, "Jordan")
	} else {
		fmt.Printf("The specified key was not found.\n")
	}

	_, ok = employeeSalary["Dakota"]
	if ok {
		fmt.Printf("Key exists.\n")
	} else {
		fmt.Printf("Key doesn't exist.\n")
	}

	for name, salary := range employeeSalary { // key:name, value:salary
		fmt.Println(name, salary)
	}
}
