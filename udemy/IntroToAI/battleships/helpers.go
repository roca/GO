package main

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func checkWinCondition(board *Board) bool {
	for i := range boardSize {
		for j := range boardSize {
			if board[i][j] == ship {
				return false
			}
		}
	}
	return true
}

func isShipSunk(board *Board, row, col int, player *HumanPlayer, ai *AIPlayer) (bool, string) {
	return false, ""
}
