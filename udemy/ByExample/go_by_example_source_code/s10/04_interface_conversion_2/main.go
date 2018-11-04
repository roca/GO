// File name: ...\s10\04_interface_conversion_2\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"strconv"
)

func main() {

	// ref: https://golang.org/pkg/strconv/

	iResult, _ := strconv.Atoi("24")
	fmt.Printf("%d %T \n", iResult, iResult)

	sResult := strconv.Itoa(12)
	fmt.Printf("%s %T \n", sResult, sResult)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	i, _ := strconv.ParseInt("-42", 10, 64)
	u, _ := strconv.ParseUint("42", 10, 64)

	fmt.Println(b, f, i, u)
	fmt.Printf("%T %T %T %T\n", b, f, i, u)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	sb := strconv.FormatBool(true)
	sf := strconv.FormatFloat(3.1415, 'E', -1, 64)
	si := strconv.FormatInt(-42, 10)
	su := strconv.FormatUint(42, 10)

	fmt.Println(sb, sf, si, su)
	fmt.Printf("%T %T %T %T\n", sb, sf, si, su)
}
