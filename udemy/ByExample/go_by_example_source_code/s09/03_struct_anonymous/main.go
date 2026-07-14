// File name: ...\s09\03_struct_anonymous\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {

	player1 := struct {
		name, sport string
		age         int
	}{"Sidney Crosby", "Hockey", 30}

	fmt.Println("Player 1=", player1)

	player2 := struct {
		name, sport string
		age         int
	}{
		age:   21,
		sport: "Swimming",
		name:  "Katie Ledecky",
	}
	fmt.Println("Player 2=", player2)
}
