package main

import "fmt"

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

func isShipSunk(board *Board, row, col int, opponentShips []Ship) (bool, string) {
	// find the ship that was hit
	var hitShip *Ship
	for i := range opponentShips {
		ship := &opponentShips[i]
		// Check if the hit coordinates are within this ship's boundaries
		if (ship.StartPosition.row == ship.EndPosition.row &&
			row == ship.StartPosition.row &&
			col >= ship.StartPosition.col && col <= ship.EndPosition.col) ||
			(ship.StartPosition.col == ship.EndPosition.col &&
				col == ship.StartPosition.col &&
				row >= ship.StartPosition.row && row <= ship.EndPosition.row) {
			hitShip = ship // Found the ship that was hit
			break
		}
	}

	// If no ship was found at this hit location, return false
	if hitShip == nil {
		fmt.Println("No ship found at the hit location. This should not happen.")
		return false, ""
	}

	// Check if all parts of the ship are marked as hit
	hits := 0
	shipLength := 0

	if hitShip.StartPosition.row == hitShip.EndPosition.row { // Horizontal ship
		shipLength = abs(hitShip.EndPosition.col-hitShip.StartPosition.col) + 1
		for c := hitShip.StartPosition.col; c <= hitShip.EndPosition.col; c++ {
			if board[hitShip.StartPosition.row][c] == hit {
				hits++
			}
		}
	} else { // Vertical ship
		shipLength = abs(hitShip.EndPosition.row-hitShip.StartPosition.row) + 1
		for r := hitShip.StartPosition.row; r <= hitShip.EndPosition.row; r++ {
			if board[r][hitShip.StartPosition.col] == hit {
				hits++
			}
		}
	}

	// Check if the number of hits == the ship length
	// return true and ship name if ship is sunk
	if hits == shipLength {
		return true, hitShip.ShipName
	}

	// return false and no ship name if ship is not sunk
	return false, ""
}
