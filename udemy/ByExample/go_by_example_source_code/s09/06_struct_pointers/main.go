// File name: ...\s09\06_struct_pointers\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

type player struct {
	name, sport string
	age         int
}

func main() {

	p1 := &player{"Michael Phelps", "swimming", 32}
	fmt.Printf("(*p1).name=%s p1.name=%s\n", (*p1).name, p1.name)

	(*p1).age = 20
	fmt.Println("Player 1:", (*p1))

	player2 := player{"Diego Maradona", "Soccer", 57}
	p2 := &player2
	
	fmt.Printf("(*p2).name=%s p2.name=%s\n", (*p2).name, p2.name)
	fmt.Println("Player 2:", (*p2))
}
