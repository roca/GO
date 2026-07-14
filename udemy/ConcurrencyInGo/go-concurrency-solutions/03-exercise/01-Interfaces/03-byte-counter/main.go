package main

import "fmt"

// ByteCounter type
type ByteCounter int

func (bc *ByteCounter) Write(b []byte) (n int, err error) {
	*bc += ByteCounter(len(b))
	return len(b), nil

}

// TODO: Implement Write method for ByteCounter
// to count the number of bytes written.

func main() {
	var b ByteCounter
	fmt.Fprintf(&b, "hello world")
	fmt.Println(b)
}
