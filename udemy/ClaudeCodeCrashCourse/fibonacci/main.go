package main

import "fmt"

func fibonacci(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("negative input not allowed: %d", n)
	}
	if n <= 1 {
		return n, nil
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b, nil
}

func main() {
	fmt.Println("Fibonacci Sequence:")
	for i := 0; i <= 20; i++ {
		result, err := fibonacci(i)
		if err != nil {
			fmt.Printf("F(%d): %s\n", i, err)
			continue
		}
		fmt.Printf("F(%d) = %d\n", i, result)
	}
}
