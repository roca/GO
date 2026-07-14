// File name: ...\s08\14_func_panic_recover\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {

	defer func() {
		errMsg := recover()
		fmt.Println(errMsg)
	}()

	// Only use one of the following at the same time.

	var x map[string]int
	x["key"] = 10
	fmt.Println(x)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	// panic("BOO!!!")
	fmt.Println("This line won't be reached!")
}
