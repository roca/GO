// File name: ...\s14\02_error_handling2\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	// Also try produceError2() and 'false'
	_, _, err := produceError1(true)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("No error!")
	}
}

func produceError1(b bool) (int, int, error) {
	if !b {
		return 0, 0, nil
	}

	errMsg1 := errors.New("I'm just a value like any other values in Go! (Rob Pike told me)")
	fmt.Printf("**errors.New() generates type: %T\n", errMsg1)

	return 0, 0, errMsg1
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
func produceError2(b bool) (int, int, error) {

	if !b {
		return 0, 0, nil
	}

	errMsg2 := fmt.Errorf("it's %v that I'm just an error value", b)
	fmt.Printf("**errors.New() generates type: %T\n", errMsg2)

	return 0, 0, errMsg2
}
