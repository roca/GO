// File name: ...\s03\06_for_3\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	i := 2
	for {
		if i == 8 {
			break
		}

		if i == 4 || i == 6 {
			i++
			continue
		}

		// fmt.Print(++i)	//compiler error
		// fmt.Print(i++)	//compiler error

		fmt.Print(i, " ")
		i++
	}

}
