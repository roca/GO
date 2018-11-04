// File name: ...\s09\09_struct_embedded_anonymous\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

// ASSIGNMENT: embedded anonymous structs
// change the previous example to use an anonymous field of type of struct.

type generalInfo struct {
	country, hairColor string
}

type player struct {
	name, sport string
	age         int
	generalInfo
}

func main() {

	var player1 player
	player1.name = "Wayne Gretzky"
	player1.age = 57
	player1.sport = "Hockey"

	player1.country = "Canada"
	player1.hairColor = "Brown"

	fmt.Println(player1)
}
