package main

import "fmt"

func PlayGame(start, takeTurn func(), haveWinner func() bool, winningPlayer func() int) {
	start()
	for !haveWinner() {
		takeTurn()
	}
	fmt.Printf("Player %d wins.\n", winningPlayer())
}

func main() {
	turn, maxTurns, currentPlayer := 1, 10, 0

	start := func() {
		turn = 1
		currentPlayer = 0
		fmt.Printf("Starting a game of chess with %d turns\n", maxTurns)
	}

	takeTurn := func() {
		turn++
		fmt.Printf("Turn %d taken by player %d\n", turn, currentPlayer)
		currentPlayer = (currentPlayer + 1) % 2
	}

	haveWinner := func() bool {
		return turn == maxTurns
	}

	winningPlayer := func() int {
		return currentPlayer
	}

	PlayGame(start, takeTurn, haveWinner, winningPlayer)
}
