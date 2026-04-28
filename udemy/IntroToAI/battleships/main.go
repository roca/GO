package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	boardSize  = 10
	empty      = "."
	ship       = "O"
	hit        = "X"
	miss       = "~"
	hiddenShip = "."
	headerRow  = "  A B C D E F G H I J"
	headerCol  = "0123456789"
)

func main() {
	// Create players
	human := NewHumanPlayer()

	// Welcome message
	fmt.Println("\n=== WELCOME TO BATTLESHIPS ===")
	fmt.Println("Legend:")
	fmt.Printf(" %s - Empty water\n", empty)
	fmt.Printf(" %s - Your ship\n", ship)
	fmt.Printf(" %s - Hit\n", hit)
	fmt.Printf(" %s - Miss\n", miss)
	fmt.Println("\nPress enter to start the game...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	// Place ships
	// AI place ships (player does not see this)

	// Human player places ships

	// Main game loop
	gameOver := false
	// playerTurn := true

	for !gameOver {
		// Display boards

		// Players take turn

		// Switch turns

		// Check win condition
	}

	fmt.Println("\nThanks for playing Battleships! Press enter to exit...")
	reader.ReadString('\n')
}
