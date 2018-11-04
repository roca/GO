// File name: ...\s15\06_reflection1\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(getType(10))
	fmt.Println(getType(true))
	fmt.Println(getType("hello"))
	fmt.Println(getType(10.31))

}
func getType(t interface{}) string {

	result := ""

	switch t := t.(type) {

	case int: // ...similar cases for int16, ...
		result = strconv.Itoa(t) + "/int"

	case bool:
		if t {
			result = "true/bool"
		} else {
			result = "false/bool"
		}
	case string:
		result = t + "/string"
	default:
		result = fmt.Sprintf("%v/unknown", t)
	}

	return result
}
