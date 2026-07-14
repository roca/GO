// File name: ...\s09\07_struct_anonymous_fields\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

type player struct {
	string
	int
}

func main() {

	player1 := player{"Muhammad Ali", 99}
	fmt.Println("Player 1:", player1)
	fmt.Printf("player1.int=%d player1.string=%s\n", player1.int, player1.string)

	player2 := player{ // Not to be used
		int:    36,
		string: "Federer",
	}
	fmt.Println("Player 2:", player2)
}
