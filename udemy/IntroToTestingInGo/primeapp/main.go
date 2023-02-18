package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// print a welcome message
	intro()

	//  create a channel to indicate when the user wants to quit
	doneChan := make(chan bool)

	// start a goroutine to read user input and run program
	go readUserInput(doneChan)

	// block util the doneChan get a value
	<-doneChan

	// close the channel
	close(doneChan)

	// say goodbye
	fmt.Println("Goodbye!")

}

// readUserInput reads user input and runs the prime checker
func readUserInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

// checkNumbers checks if the user input is a number and if it is, checks if it is prime
func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()
	input := scanner.Text()

	if input == "q" {
		return "Goodbye!", true
	}

	num, err := strconv.Atoi(input)
	if err != nil {
		return "Please enter a whole number or 'q' to quit", false
	}

	_, msg := isPrime(num)

	return msg, false
}

func intro() {
	fmt.Println("Welcome to the Prime Number Checker!")
	fmt.Println("Enter a number to check if it is prime.")
	fmt.Println("Enter 'q' to quit.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
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
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not prime, it is divisible by %d", n, i)
		}
	}
	return true, fmt.Sprintf("%d is a prime number!", n)
}
