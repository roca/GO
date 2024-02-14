package main

import "fmt"

func xor(lhs, rhs bool) bool {
	// ?
	return  (lhs || rhs) && !(lhs && rhs)
}

// don't touch below this line

func test(lhs, rhs bool) {
	res := xor(lhs, rhs)
	fmt.Printf("%v XOR %v = %v\n", lhs, rhs, res)
}

func main() {
	test(true, true)
	test(true, false)
	test(false, true)
	test(false, false)
}
