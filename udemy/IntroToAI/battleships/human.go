package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Ship struct {
	ShipName      string
	StartPosition Position
	EndPosition   Position
}

type HumanPlayer struct {
	board    Board
	ships    []Ship
	opponent *AIPlayer
}

func NewHumanPlayer() *HumanPlayer {
	p := &HumanPlayer{}
	for i := range boardSize {
		for j := range boardSize {
			p.board[i][j] = empty
		}
	}
	return p
}

func (p *HumanPlayer) PlaceShips() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n== SHIP PLACEMENT ==")
	fmt.Println("Place your ships on the board.")
	fmt.Println("Format: A0 H (A0 starting position, H=horizontal or V=vertical)")
	fmt.Println("Positions are given as letter (A-J) for column and number (0-9) for row.")

	for _, shipType := range shipTypes {
		for {
			// Display current board state
			printBoards(&p.board, &Board{})

			fmt.Printf("\nPlace your %s (length %d): ", shipType.name, shipType.size)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(strings.ToUpper(input))
			parts := strings.Fields(input)

			if len(parts) != 2 {
				fmt.Println("Invalid input format. Use format like 'A0 H'. Please try again.")
				time.Sleep(2 * time.Second)
				continue
			}

			pos := parts[0]
			dir := parts[1]

			if len(pos) < 2 || (dir != "H" && dir != "V") {
				fmt.Println("Invalid input format. Use format like 'A0 H'. Please try again.")
				time.Sleep(2 * time.Second)
				continue
			}

		}
	}
}
