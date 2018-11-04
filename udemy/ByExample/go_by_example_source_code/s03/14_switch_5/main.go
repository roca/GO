// File name: ...\s03\14_switch_5\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"os"
)

func main() {
	var season string

	//for {	// infinite loop
	for season != "end" {
		fmt.Print("Enter your favorite season (type end to exit): ")
		fmt.Scanf("%s ", &season)

		// if season == "end" {
		// 	break
		// }

		switch season {
		case "spring", "SPRING":
			fmt.Fprintf(os.Stdout, "You Entered %s\n", season)
		case "summer":
			fmt.Println("You Entered", season)
			break
			fmt.Println("This line won't be reached", season) // Note
		case "fall":
			fmt.Println("You Entered", season)
		case "winter":
			fmt.Println("You Entered", season)
		default:
			fmt.Println("a new season - ", season)
		}
	}
}
