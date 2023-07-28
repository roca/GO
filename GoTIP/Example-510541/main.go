package main

// Run this code using `gotip run main.go`

func main() {

	// All the examples below will print 0 to 9
	// under the current proposal

	// Standard for loop
	for i := 0; i < 10; i++ {
		print(i)
	}
	println()

	// From the new proposal (510541)
	// https://go-review.googlesource.com/c/go/+/510541

	// Eaxmple 1
	for i := range 10 {
		print(i)
	}
	println()

	// Eaxmple 2
	i := 0
	for range 10 {
		print(i)
		i++
	}
	println()

	// Eaxmple 3
	i = 0
	n := 10
	for range n {
		print(i)
		i++
	}
	println()
}
