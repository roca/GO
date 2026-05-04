package main

import "math/rand"

type AIPlayer struct {
	board          Board
	heatMap        [boardSize][boardSize]int
	hits           []Position
	shipsSunk      int
	huntMode       bool
	potentialShips []struct {
		size    int
		sunk    bool
		hits    int
		shipPos []Position
	}
	ships    []Ship
	opponent *HumanPlayer
}

func NewAIPlayer() *AIPlayer {
	p := &AIPlayer{
		shipsSunk: 0,
		huntMode:  false,
	}

	for i := range boardSize {
		for j := range boardSize {
			p.board[i][j] = empty
		}
	}

	p.intializeHeatMap()

	// Initialize potential ships tracking
	p.potentialShips = make([]struct {
		size    int
		sunk    bool
		hits    int
		shipPos []Position
	}, len(shipTypes))

	for i, shipTyper := range shipTypes {
		p.potentialShips[i].size = shipTyper.size
		p.potentialShips[i].sunk = false
		p.potentialShips[i].hits = 0
		p.potentialShips[i].shipPos = make([]Position, 0)
	}

	return p
}

func (p *AIPlayer) intializeHeatMap() {
	for i := range boardSize {
		for j := range boardSize {
			//Start with base probability
			p.heatMap[i][j] = 1

			// Increase probability in a checkkerboard pattern
			if (i+j)%2 == 0 {
				p.heatMap[i][j] += 1
			}

			// Higher probability in the center of the board
			centerX, centerY := boardSize/2, boardSize/2
			centerDistance := abs(i-centerX) + abs(j-centerY)
			if centerDistance <= 3 {
				p.heatMap[i][j] += 2
			}
		}
	}
}

func (p *AIPlayer) GetBoard() *Board {
	return &p.board
}

func (p *AIPlayer) PlaceShips() {
	// Place ships in a mix of edge and center clusters.

	// Attempt to place larger ships near the edges.
	for i, shipType := range shipTypes {
		placed := false
		attempts := 0

		for !placed && attempts < 100 {

			// Decide on placement strategy based on ship size
			var row, col int
			horizontal := rand.Intn(2) == 0

			if shipType.size >= 4 {
				if horizontal {
					row = rand.Intn(boardSize)
					if rand.Intn(2) == 0 {
						col = rand.Intn(3) // Place near left edge
					} else {
						// Place near right edge
						col = boardSize - shipType.size - rand.Intn(3)
					}
				} else {
					col = rand.Intn(boardSize)
					if rand.Intn(2) == 0 {
						row = rand.Intn(3) // Place near top edge
					} else {
						// Place near bottom edge
						row = boardSize - shipType.size - rand.Intn(3)
					}
				}
			} else {
				// Place smaller ships in a more distributed pattern
				if horizontal {
					row = rand.Intn(boardSize)
					col = rand.Intn(boardSize - shipType.size + 1)
				} else {
					row = rand.Intn(boardSize - shipType.size + 1)
					col = rand.Intn(boardSize)
				}
			}

			// Check if placement is valid
			positions := []Position{}
			valid := true

			for j := range shipType.size {
				var r, c int
				if horizontal {
					r, c = row, col+j
				} else {
					r, c = row+j, col
				}

				// Check validity of position
				if r < 0 || r >= boardSize || c < 0 || c >= boardSize || p.board[r][c] == ship {
					valid = false
					break
				}

				// Check surrounding cells to avoid placing ships adjacent to each other
				for dr := -1; dr <= 1; dr++ {
					for dc := -1; dc <= 1; dc++ {
						nr, nc := r+dr, c+dc
						if nr >= 0 && nr < boardSize && nc >= 0 && nc < boardSize &&
							p.board[nr][nc] == ship && (dr != 0 || dc != 0) {
							// Avoid placing ships diagonally or directly adjacent to another ship
							if i < 2 { // Only for larger ships
								valid = false
								break
							}
						}
					}
					if !valid {
						break
					}
				}

				if !valid {
					break
				}

				positions = append(positions, Position{r, c})
			}

			if valid {
				// Place the ship
				for _, pos := range positions {
					p.board[pos.row][pos.col] = ship
				}

				// Append this ship to the slice of ai ships
				newShip := Ship{
					StartPosition: positions[0],
					EndPosition:   positions[len(positions)-1],
					ShipName:      shipType.name,
				}

				p.ships = append(p.ships, newShip)
				placed = true
			}
		}

		// If we couldn't place the ship, fall back to random placement
		if !placed {
			for {
				horizontal := rand.Intn(2) == 0
				var row, col int
				if horizontal {
					row = rand.Intn(boardSize)
					col = rand.Intn(boardSize - shipType.size + 1)
				} else {
					row = rand.Intn(boardSize - shipType.size + 1)
					col = rand.Intn(boardSize)
				}

				// Check if placement is valid
				valid := true
				positions := []Position{}

				for i := range shipType.size {
					var r, c int
					if horizontal {
						r, c = row, col+i
					} else {
						r, c = row+i, col
					}

					if p.board[r][c] == ship {
						valid = false
						break
					}

					positions = append(positions, Position{r, c})
				}

				if valid {
					// Place the ship
					for _, pos := range positions {
						p.board[pos.row][pos.col] = ship
					}

					// Append this ship to the slice of ai ships
					newShip := Ship{
						StartPosition: positions[0],
						EndPosition:   positions[len(positions)-1],
						ShipName:      shipType.name,
					}

					p.ships = append(p.ships, newShip)
					break
				}
			}
		}
	}
}
