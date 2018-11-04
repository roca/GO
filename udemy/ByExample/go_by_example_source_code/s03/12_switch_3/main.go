// File name: ...\s03\12_switch_3\main.go
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
	case 3:
	case 4:
		fmt.Println("summer - ", seasonNo)
		fmt.Println("fall - ", seasonNo)
		fmt.Println("winter - ", seasonNo)
	default:
		fmt.Println("a new season - ", seasonNo)
	}

}
