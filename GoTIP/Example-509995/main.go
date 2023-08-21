package main

type x struct {
	a, b int
}

func (a x) myzero() (zero, error) {
	return 0, nil
}

func main() {

	// All the examples below will print 0
	// under the current proposal

	// Standard way
	i := 0
	println("ZERO", i)

	// From the new proposal (509995)
	// https://go-review.googlesource.com/c/go/+/509995

	// Example
	a := x{}

	l, _ := a.myzero()
	println("ZERO", l)
}
