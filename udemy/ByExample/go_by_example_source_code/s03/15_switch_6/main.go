// File name: ...\s03\15_switch_6\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {
	sportID := 1
	isYoung := true

	switch {
	case sportID == 1 && isYoung:
		fmt.Println("playing soccer ...")
	case sportID == 1 && !isYoung:
		fmt.Println("Watching soccer ...")
	case sportID == 2 && isYoung:
		fmt.Println("playing chess ...")
	case sportID == 2 && !isYoung:
		fmt.Println("still can play chess ...")
	default:
		fmt.Println("We'll discuss later ...")
	}
}
