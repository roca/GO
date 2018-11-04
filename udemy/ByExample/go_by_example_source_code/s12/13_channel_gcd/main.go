// File name: ...\s12\13_channel_gcd\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

// ASSIGNMENT: Write two functions to compute gcd (greatest common divisor)
// of two integers using recursive and channel.
// Ref: https://en.wikipedia.org/wiki/Greatest_common_divisor

func main() {
	fmt.Println("==> gcd using recursive")
	fmt.Printf("gcd1(2, 4)\t%d \n", gcd1(2, 4))
	fmt.Printf("gcd1(27, 9)\t%d \n", gcd1(27, 9))
	fmt.Printf("gcd1(27, 4)\t%d \n", gcd1(27, 4))
	fmt.Printf("gcd1(111, 259)\t%d \n", gcd1(111, 259))

	fmt.Println("\n==> gcd using channels")
	g := gcd2(2, 4)
	fmt.Printf("gcd2(2, 4)\t%d \n", <-g)

	fmt.Printf("gcd2(27, 9)\t%d \n", <-gcd2(27, 9))
	fmt.Printf("gcd2(27, 4)\t%d \n", <-gcd2(27, 4))
	fmt.Printf("gcd2(111, 259)\t%d \n", <-gcd2(111, 259))
}

// gcd of two integers using recursive
func gcd1(x int, y int) int {
	if y == 0 {
		return x
	}

	return gcd1(y, (x % y))
}

// gcd of two integers using channels
func gcd2(x int, y int) chan int {
	c := make(chan int)
	go func() {
		for y != 0 {
			x, y = y, x%y
		}
		c <- x
		close(c)
	}()
	return c
}
