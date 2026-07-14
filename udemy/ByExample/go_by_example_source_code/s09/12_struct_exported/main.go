// File name: ...\s09\12_struct_exported\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"strings"

	"./types"
)

func changeAthleteName1(p athletes.Player) {
	p.Name = "Anderson Silva-Spider"
}

func changeAthleteName2(p *athletes.Player) {
	p.Name = "Anderson Silva-Spider"
	p.Country = strings.ToUpper(p.Country)
}

func main() {

	player1 := athletes.Player{"Anderson Silva", "MMA", 43, athletes.Info{"Brazil", "Black"}}
	// info1 := athletes.Info{Country: "Brazil", HairColor: "Black"}
	// player1 := athletes.Player{Name: "Anderson Silva", Sport: "MMA", Age: 43, Info: info1}
	fmt.Println("(1) player1:", player1)

	changeAthleteName1(player1)
	fmt.Println("(2) player1:", player1)

	changeAthleteName2(&player1)
	fmt.Println("(3) player1:", player1)

	fmt.Println("(4) player1:", *player1.ToLowercase())
}
