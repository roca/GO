// File name: ...\s03\10_switch_1\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	seasonNo := 3

	switch seasonNo {

	case 1:
		fmt.Println("spring - ", seasonNo)
		// break
	case 2:
		fmt.Println("summer - ", seasonNo)
	case 3:
		fmt.Println("fall - ", seasonNo)
	case 4:
		fmt.Println("winter - ", seasonNo)
	default:
		fmt.Println("a new season - ", seasonNo)
	}
}
