package main

import "fmt"

var shipTypes = []struct {
	name string
	size int
}{
	{"Carrier", 5},
	{"Battleship", 4},
	{"Cruiser", 3},
	{"Submarine", 3},
	{"Destroyer", 2},
}

type Board [boardSize][boardSize]string

type Position struct {
	row, col int
}

func printBoards(playerBoard, opponentBoard *Board) {
	// Clear the screen
	// Fow Windows users there is a pacakage called "github.com/inancgumus/screen" that can be used to clear the console in a cross-platform way.
	fmt.Print("\033[H\033[2J")

	fmt.Println("\n=== BATTLESHIP ===")
	fmt.Println()

	fmt.Println("  OPPONENT'S BOARD:")
	fmt.Println(headerRow)

	// Print opponewnt's board
	// Player shouild not see the enemy's ships, only hits and misses
	for i := 0; i < boardSize; i++ {
		fmt.Printf("%d ", i)
		for j := range boardSize {
			cell := opponentBoard[i][j]
			switch cell {
			case ship:
				// Don't show ships on opponent's board
				fmt.Printf("%s ", hiddenShip)
			default:
				fmt.Printf("%s ", cell)
			}
		}
		fmt.Println()
	}

	fmt.Println("\n  YOUR BOARD:")
	fmt.Println(headerRow)

	// Print player's board - all information is visible to the player
	for i := range boardSize {
		fmt.Printf("%d ", i)
		for j := range boardSize {
			fmt.Printf("%s ", playerBoard[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}
