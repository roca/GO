package main

import (
	"fmt"
)

func main() {
	a := make(map[string]string)
	a["test"] = "value"
	testPonter(a)
	fmt.Printf("a: %v\n", a)
}

func testPonter(a map[string]string) {
	a["test2"] = "new value2"
	a["test3"] = "new value3"
}
