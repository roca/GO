// File name: ...\s03\02_if_else_2\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {
	seasonNo := 1

	if seasonNo == 1 {
		fmt.Println("spring - ", seasonNo)
	} else if seasonNo == 2 {
		fmt.Println("summer - ", seasonNo)
	} else if seasonNo == 3 {
		fmt.Println("fall - ", seasonNo)
	} else if seasonNo == 4 {
		fmt.Println("winter - ", seasonNo)
	} else {
		fmt.Println("a brand new season - ", seasonNo)
	}
}
