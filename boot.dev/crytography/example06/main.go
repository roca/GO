package main

import (
	"fmt"
	"strings"
)

func getHexString(b []byte) string {
	var text []string
	for _, v := range b {
		text = append(text, fmt.Sprintf("%02x", v))
	}
	return strings.Join(text, ":")
}

func getBinaryString(b []byte) string {
	var text []string
	for _, v := range b {
		text = append(text, fmt.Sprintf("%08b", v))
	}
	return strings.Join(text, ":")
}

// don't touch below this line

func testHex(s string) {
	myBytes := []byte(s)
	fmt.Printf("String: '%s', Hex: %v\n", s, getHexString(myBytes))
	fmt.Println("========")
}

func testBinary(s string) {
	myBytes := []byte(s)
	fmt.Printf("String: '%s', Bin: %v\n", s, getBinaryString(myBytes))
	fmt.Println("========")
}

func main() {
	testHex("Hello")
	testHex("World")
	testBinary("Hello")
	testBinary("World")
}
