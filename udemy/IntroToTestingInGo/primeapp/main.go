package main

import "fmt"

func main() {
	
	for i := 0; i < 10; i++ {
		prime, msg := isPrime(i)
		if prime {
			fmt.Println(msg)
		}
	}

}

// isPrime returns true if the number is prime and a message if it is not
func isPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	//negative numbers are not prime
	if n < 0 {
		return false, "Negative numbers are not prime, by definition!"
	}
	
	// use the modulus operator to check if the number is prime
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not prime, it is divisible by %d", n, i)
		}
	}
	return true, fmt.Sprintf("%d is a prime number!", n)
}
