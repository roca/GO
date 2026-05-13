package main

import (
	"fmt"
	"math/rand"
)

const (
	// Heatmap Weights
	baseProbability      = 1   // Starting probability for valid cells
	checkerboardBonus    = 1   // Bonus for checkerboard pattern (less likeley adjacent placements)
	centerProximityBonus = 2   // Bonus for being closer to the center of the board (where players often place ships)
	maxCenterDistance    = 3   // How far from the center qualifies for the center proximity bonus
	huntModeBoost        = 100 // Sigificant boost for cells adjacent to hits in hunt mode
	shipFitBonus         = 2   // Base bonus multiplier for fitting a ship
)

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

	p.initializeHeatMap()

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

func (p *AIPlayer) initializeHeatMap() {
	for i := range boardSize {
		for j := range boardSize {
			//Start with base probability
			p.heatMap[i][j] = baseProbability

			// Increase probability in a checkkerboard pattern
			if (i+j)%2 == 0 {
				p.heatMap[i][j] += checkerboardBonus
			}

			// Higher probability in the center of the board
			centerX, centerY := boardSize/2, boardSize/2
			centerDistance := abs(i-centerX) + abs(j-centerY)
			if centerDistance <= maxCenterDistance {
				p.heatMap[i][j] += centerProximityBonus
			}
		}
	}
}

// updateHeatMap recalculates the heatmap probabilities based on the current game state.
// It considers potenial ship placements and prioritizes targets during hunt mode.
func (p *AIPlayer) updateHeatMap(opponentBoard *Board) {
	// 1. Reset heat map (clear prveious probabilities)
	for i := range boardSize {
		for j := range boardSize {
			p.heatMap[i][j] = 0
		}
	}

	// 2. Calculate base probabilities & ship fit probabilities
	for r := range boardSize {
		for c := range boardSize {
			// Skip cells that have already been targeted.
			if opponentBoard[r][c] == hit || opponentBoard[r][c] == miss {
				continue
			}

			// Assign base probability to valid untargeted cells
			p.heatMap[r][c] = baseProbability

			// Interate through opponents ships that have not been sunk yet.
			for _, shipData := range p.potentialShips {
				if shipData.sunk {
					continue
				}

				shipSize := shipData.size

				// Check horizontal fit: can and unsunk ship of this size fit horizontally starting here?
				if c+shipSize <= boardSize {
					canFitHorizontal := true
					for k := range shipSize {
						// Check if any cell needed for the ship is already a miss or a hit
						if opponentBoard[r][c+k] == miss || opponentBoard[r][c+k] == hit {
							canFitHorizontal = false
							break
						}
						if canFitHorizontal {
							// Increase probability based on ship size if it fits
							p.heatMap[r][c] += shipFitBonus * shipSize
						}
					}
				}

				// Check vertical fit
				if r+shipSize <= boardSize {
					canFitVertical := true
					for k := range shipSize {
						// Check if any cell needed for the ship is already a miss or a hit
						if opponentBoard[r+k][c] == miss || opponentBoard[r+k][c] == hit {
							canFitVertical = false
							break
						}
						if canFitVertical {
							// Increase probability based on ship size if it fits
							p.heatMap[r][c] += shipFitBonus * shipSize
						}
					}
				}

			}
		}
	}

	// 3. Apply hunt mode boost if applicable
	if p.huntMode && len(p.hits) > 0 {
		p.applyHuntModeBoosts(opponentBoard)
	}

}

func (p *AIPlayer) applyHuntModeBoosts(opponentBoard *Board) {

}

func (p *AIPlayer) TakeTurn(opponentBoard *Board) (Position, bool) {
	fmt.Println("\nAI is taking its turn...")
	if p.huntMode {
		fmt.Println("AI is in HUNT MODE: Prioritizing targets around hits!")
	} else {
		fmt.Println("AI is in probability target mode!")
	}

	// Update heat map based on game state

	// Select a target based on strategy (hunt mode vs probability mode)

	if p.huntMode {
		// If in hunt mode ...
		// find the hightest probability cell(s)
		// select a randowm target from highest probability cells
		// if can't find one, fallback to random targeting
	} else {
		// ... else, do something else
		// find the hightest probability cell(s)
		// select a random target from highest probability cells
	}

	// Perform the attack
	// check to see if we hit a ship
	// if we hit, enter hunt mode
	// check to see if ship was sunk

	return Position{}, false
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
			attempts++

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
