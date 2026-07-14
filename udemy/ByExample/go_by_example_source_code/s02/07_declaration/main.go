// File name: ...\s02\07_declaration\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	s1, s2 := "msg1", "msg2"
	var s3 = "msg3"
	var s4 string = "msg4"

	fmt.Println(s1 + " " + s2 + " " + s3 + " " + s4) //msg1 msg2 msg3 msg4
}
