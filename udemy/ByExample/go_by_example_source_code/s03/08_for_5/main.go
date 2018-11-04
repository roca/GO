// File name: ...\s03\08_for_5\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	fmt.Println("Press CTRL+C to stop the program.")

	i := 2
	for {
		fmt.Print(i, " ")
		i++
	}
}
