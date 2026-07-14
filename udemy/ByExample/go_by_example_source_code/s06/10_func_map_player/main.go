// File name: ...\s06\10_func_map_player\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

// ASSIGNMENT - using 'map of maps' implement the following functionality:
// 1. addPlayer("firstName", "lastName")
//	  This function adds a new player to our map of maps;
//    'first name' must be unique (for simplicity and only in this program).
// 2. hasPlayer("firstName", "lastName") - returns bool
//    to check if a pair of firstName/lastName exist in our map.
//

var players = make(map[string]map[string]bool)

func main() {
	addPlayer("Leo", "Messi")
	addPlayer("Roger", "Federer")
	addPlayer("Michael", "Jordan")

	fmt.Println()
	fmt.Println(hasPlayer("Roger", "Federer"))
	fmt.Println(hasPlayer("Michael", "Jordan"))
	fmt.Println(hasPlayer("Michael", "Messi"))
	fmt.Println(hasPlayer("Leo", "Jordan"))

	fmt.Println()
	fmt.Println(players)

	fmt.Println()
	for k1, v1 := range players {
		for k2, v2 := range v1 {
			fmt.Println(k1, k2, v2)
		}
	}
}

func addPlayer(fName, lName string) {
	n := players[fName]
	if n == nil {
		n = make(map[string]bool)
		players[fName] = n
	}
	n[lName] = true
}

func hasPlayer(fName, lName string) bool {
	return players[fName][lName]
}
