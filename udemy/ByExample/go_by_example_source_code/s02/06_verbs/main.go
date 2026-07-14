// File name: ...\s02\06_verbs\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {
	var i uint8 = 20
	var f float32 = 214.437
	var b bool = true
	var msg string = "hello"

	type myStringT string
	var myMsg myStringT = "special message"

	fmt.Printf("%d %v %T \n", i, i, i)
	fmt.Printf("%5.1f %.4f %g %T\n", f, f, f, f)
	fmt.Printf("%t %v %T\n", b, b, b)
	fmt.Printf("%s %T\n", msg, msg)
	fmt.Printf("%v %T\n", myMsg, myMsg)
}
