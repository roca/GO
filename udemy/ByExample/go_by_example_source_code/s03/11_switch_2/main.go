// File name: ...\s03\11_switch_2\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	var seasonNo int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &seasonNo)

	switch seasonNo {
	case 1:
		fmt.Println("spring - ", seasonNo)
	case 2:
		fmt.Println("summer - ", seasonNo)
		fallthrough
	case 3:
		fmt.Println("fall - ", seasonNo)
		fallthrough
	case 4:
		fmt.Println("winter - ", seasonNo)
	default:
		fmt.Println("a new season - ", seasonNo)
	}
}
