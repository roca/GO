// File name: ...\s12\09_channel_direction\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ASSIGNMENT: Assuming your current bank balance is 200, write
// a program containing these functions:
// 1) credit() function with a 'bi-directional channel' parameter that sends random
// numbers (between 1 and 10) to the channel that will be added to the balance.
// 2) debit() function with a 'send-only channel' parameter that sends random
// numbers (between -10 and -1) to the channel that will be subtracted from the balance.
// 3) balance() function with a 'receive-only channel' parameter that adds/subtracts
// the random values (generated in steps 1 & 2) to the balance.
//
// Here's a sample output that indicates that the balance is randomly changing:
//		Press Enter to stop the program ...
// 		=> 200 + (6) = 206
// 		=> 206 + (-10) = 196
// 		=> 196 + (3) = 199
// 		=> 199 + (-2) = 197
//
var currentBalance = 200

func main() {
	fmt.Println("Press Enter to stop the program ...")
	rand.Seed(time.Now().UnixNano())

	c := make(chan int) // bidirectional

	go credit(c)
	go debit(c)
	go balance(c)

	var input string
	fmt.Scanln(&input) // press "Enter" to exit
}

func credit(c chan int) { // a bi-directional channel
	for i := 0; ; i++ {
		c <- rand.Intn(9) + 1 //random(1, 10)
	}
}

func debit(c chan<- int) { //a "send-only" channel

	for i := 0; ; i++ {
		c <- rand.Intn(9) - 10 //random(-10, -1)
	}
}

func balance(c <-chan int) { //a "receive-only" channel
	for {
		num, ok := <-c
		if ok == false {
			fmt.Println("Error:")
			break
		}

		oldBalance := currentBalance
		currentBalance += num

		fmt.Printf("=> %d + (%d) = %d\n", oldBalance, num, currentBalance)
		time.Sleep(1 * time.Second)
	}
}
