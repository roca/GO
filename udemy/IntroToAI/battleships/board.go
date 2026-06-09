package main

import (
	"fmt"
	"strconv"
)

type shipSpec struct {
	name string
	size int
}

var shipTypes = []shipSpec{
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

// newBoard returns a pointer to a board with every cell set to empty.
func newBoard() *Board {
	var b Board
	for i := range boardSize {
		for j := range boardSize {
			b[i][j] = empty
		}
	}
	return &b
}

// shipCells returns the positions a ship of the given size occupies starting
// at start, extending right when horizontal or down when vertical. It does not
// validate bounds; callers check each returned position.
func shipCells(start Position, size int, horizontal bool) []Position {
	cells := make([]Position, 0, size)
	for i := range size {
		if horizontal {
			cells = append(cells, Position{row: start.row, col: start.col + i})
		} else {
			cells = append(cells, Position{row: start.row + i, col: start.col})
		}
	}
	return cells
}

// parseCoord parses a board coordinate like "A0" into a Position. The first
// character is the column letter (A-J) and the remainder is the row number
// (0-9). Input is expected to be already trimmed and upper-cased. It returns an
// error describing the first validation failure.
func parseCoord(s string) (Position, error) {
	if len(s) < 2 {
		return Position{}, fmt.Errorf("invalid input format. Please enter a letter (A-J) followed by a number (0-9)")
	}

	if s[0] < 'A' || s[0] > 'J' {
		return Position{}, fmt.Errorf("invalid column. Please enter a letter between A and J")
	}
	col := int(s[0] - 'A')

	row, err := strconv.Atoi(s[1:])
	if err != nil || row < 0 || row >= boardSize {
		return Position{}, fmt.Errorf("invalid row. Please enter a number between 0 and 9")
	}

	return Position{row: row, col: col}, nil
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
	for i := range boardSize {
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
