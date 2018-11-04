// File name: ...\s09\02_struct_player\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

type player struct {
	name, sport string
	age         int
}

func main() {

	player1 := player{"Leo Messi", "Soccer", 30}
	fmt.Println("Player 1=", player1)
	fmt.Printf("(1) name=%s age=%d\n", player1.name, player1.age)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	player2 := player{
		age:   36,
		sport: "Tennis",
		name:  "Roger Federer",
	}

	fmt.Println("Player 2=", player2)
	fmt.Printf("(2) name=%s age=%d\n", player2.name, player2.age)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	var player3 player
	fmt.Println("Player 3=", player3)

	player3.name = "Usain Bolt"
	player3.sport = "Sprinter"
	player3.age = 31
	fmt.Println("Player 3 updated=", player3)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	player4 := player{
		name:  "Michael Jordan",
		sport: "basketball",
	}
	fmt.Println("Player 4=", player4)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	player5 := new(player)
	player5.name = "Tiger Woods"
	player5.sport = "Golf"
	player5.age = 42

	fmt.Printf("*player5=%v \n", *player5)
	fmt.Printf("*player5=%+v \n", *player5)
}
