package main


func main() {

	// All the examples below will print 0
	// under the current proposal

	// Standard way
	i := 0
	println("ZERO", i)


	// From the new proposal (509995)
	// https://go-review.googlesource.com/c/go/+/509995

	// Example
	var x struct {
		a, b int
	}

	x = zero
	println("ZERO", x)
}
