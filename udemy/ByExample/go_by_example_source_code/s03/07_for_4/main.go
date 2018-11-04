// File name: ...\s03\07_for_4\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)
package main

import "fmt"

func main() {
	for i := -2; i < 10; i++ {

		if i == 4 {
			break
		}

		if i == 2 {
			continue
		}

		fmt.Println(i)
	}
}
