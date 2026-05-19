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
	hiddenShip = "E"
	headerRow  = "  A B C D E F G H I J"
	headerCol  = "0123456789"
)

func main() {
	// Create players
	human := NewHumanPlayer()
	ai := NewAIPlayer()

	human.opponent = ai
	ai.opponent = human

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
	ai.PlaceShips()

	// Human player places ships
	human.PlaceShips()

	// Main game loop
	gameOver := false
	playerTurn := true

	for !gameOver {
		// Display boards
		printBoards(human.GetBoard(), ai.GetBoard())

		// Players take turn
		if playerTurn {
			fmt.Println("\n=== YOUR TURN ===")
			// let player take turn
			_, _ = human.TakeTurn(ai.GetBoard())

			// Check for win condition
			if checkWinCondition(ai.GetBoard()) {
				gameOver = true
				printBoards(human.GetBoard(), ai.GetBoard())
				fmt.Println("\nCongratulations! You win! You sank all enemy ships.")
			}
		} else {
			// Print heatmap
			//	fmt.Println("Heat map:")
			//	for i := range boardSize {
			//		for j := range boardSize {
			//			fmt.Printf("%3d ", ai.heatMap[i][j])
			//		}
			//		fmt.Println()
			//	}

			fmt.Println("\n=== AI'S TURN ===")
			// let AI take turn
			_, _ = ai.TakeTurn(human.GetBoard())

			// fmt.Println("\nAI has taken its turn. Press enter to continue...")
			// reader.ReadString('\n')

			// Check for win condition
			if checkWinCondition(human.GetBoard()) {
				gameOver = true
				printBoards(human.GetBoard(), ai.GetBoard())
				fmt.Println("\nGame Over! The AI wins! Your fleet has been sunk.")
			}
		}

		// Switch turns
		playerTurn = !playerTurn

		// Check win condition
	}

	fmt.Println("\nThanks for playing Battleships! Press enter to exit...")
	reader.ReadString('\n')
}
