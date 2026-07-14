package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const (
	// Heatmap Weights
	baseProbability      = 1  // Starting probability for valid cells
	checkerboardBonus    = 1  // Bonus for checkerboard pattern (less likeley adjacent placements)
	centerProximityBonus = 2  // Bonus for being closer to the center of the board (where players often place ships)
	maxCenterDistance    = 3  // How far from the center qualifies for the center proximity bonus
	huntModeBoost        = 15 // Sigificant boost for cells adjacent to hits in hunt mode
	shipFitBonus         = 2  // Base bonus multiplier for fitting a ship
)

type shipTracker struct {
	size    int
	sunk    bool
	hits    int
	shipPos []Position
}

type AIPlayer struct {
	board          Board
	heatMap        [boardSize][boardSize]int
	hits           []Position
	shipsSunk      int
	huntMode       bool
	potentialShips []shipTracker
	ships          []Ship
	opponent       *HumanPlayer
}

func NewAIPlayer() *AIPlayer {
	p := &AIPlayer{
		shipsSunk: 0,
		huntMode:  false,
		board:     *newBoard(),
	}

	p.initializeHeatMap()

	// Initialize potential ships tracking
	p.potentialShips = make([]shipTracker, len(shipTypes))

	for i, shipType := range shipTypes {
		p.potentialShips[i].size = shipType.size
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
	p.initializeHeatMap()

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
					}
					if canFitHorizontal {
						// Increase probability based on ship size if it fits
						for k := range shipSize {
							p.heatMap[r][c+k] += shipFitBonus
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
					}
					if canFitVertical {
						// Increase probability based on ship size if it fits
						for k := range shipSize {
							p.heatMap[r+k][c] += shipFitBonus
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
	// Determine the hit pattern: single hit, horizontal line, or vertical line
	isSingleHit := len(p.hits) == 1
	isHorizontal := false
	isVertical := false

	if !isSingleHit {
		firstHit := p.hits[0]
		isHorizontal = true
		isVertical = true

		for i := 1; i < len(p.hits); i++ {
			if p.hits[i].row != firstHit.row {
				isHorizontal = false
			}
			if p.hits[i].col != firstHit.col {
				isVertical = false
			}
		}

		// If hits are not aligned horizontally or vertically, treat as multiple single points
		// for adjacent checks.
		if !isHorizontal && !isVertical {
			isSingleHit = true // fall back to checking adjacent cells for all hits if not clearly aligned
		}
	}

	boostCell := func(r, c int) {
		if r >= 0 && r < boardSize && c >= 0 && c < boardSize &&
			opponentBoard[r][c] != hit && opponentBoard[r][c] != miss {
			p.heatMap[r][c] += huntModeBoost
		}
	}

	if isSingleHit {
		// Boost all valid neighbors if the hits(s)
		for _, hitPos := range p.hits {
			boostCell(hitPos.row-1, hitPos.col) // Up
			boostCell(hitPos.row+1, hitPos.col) // Down
			boostCell(hitPos.row, hitPos.col-1) // Left
			boostCell(hitPos.row, hitPos.col+1) // Right
		}
	} else if isHorizontal {
		// Boost cells to the left and right of the horizontal line of hits
		row := p.hits[0].row
		minCol, maxCol := p.hits[0].col, p.hits[0].col
		for _, hitPos := range p.hits {
			if hitPos.col < minCol {
				minCol = hitPos.col
			}
			if hitPos.col > maxCol {
				maxCol = hitPos.col
			}
		}
		boostCell(row, minCol-1) // Left of the line
		boostCell(row, maxCol+1) // Right of the line
	} else if isVertical {
		// Boost cells above and below the vertical line of hits
		col := p.hits[0].col
		minRow, maxRow := p.hits[0].row, p.hits[0].row
		for _, hitPos := range p.hits {
			if hitPos.row < minRow {
				minRow = hitPos.row
			}
			if hitPos.row > maxRow {
				maxRow = hitPos.row
			}
		}
		boostCell(minRow-1, col) // Above the line
		boostCell(maxRow+1, col) // Below the line
	}

}

// bestTarget returns the highest-probability untargeted cell on the heat map,
// breaking ties at random. If no untargeted cell carries a positive
// probability, it falls back to a random untargeted cell.
func (p *AIPlayer) bestTarget(opponentBoard *Board) Position {
	maxProb := 0
	candidates := []Position{}

	for i := range boardSize {
		for j := range boardSize {
			if opponentBoard[i][j] == hit || opponentBoard[i][j] == miss {
				continue
			}
			if p.heatMap[i][j] > maxProb {
				maxProb = p.heatMap[i][j]
				candidates = []Position{{i, j}}
			} else if p.heatMap[i][j] == maxProb {
				candidates = append(candidates, Position{i, j})
			}
		}
	}

	if len(candidates) > 0 {
		return candidates[rand.Intn(len(candidates))]
	}

	// Fallback: random untargeted cell.
	for {
		r, c := rand.Intn(boardSize), rand.Intn(boardSize)
		if opponentBoard[r][c] != hit && opponentBoard[r][c] != miss {
			return Position{r, c}
		}
	}
}

func (p *AIPlayer) TakeTurn(opponentBoard *Board) (Position, bool) {
	fmt.Println("\nAI is taking its turn...")
	if p.huntMode {
		fmt.Println("AI is in HUNT MODE: Prioritizing targets around hits!")
	} else {
		fmt.Println("AI is in probability target mode!")
	}

	// Update heat map based on game state
	p.updateHeatMap(opponentBoard)

	// Select a target. Hunt mode already influences the heat map upstream via
	// applyHuntModeBoosts, so the selection itself is identical in both modes:
	// pick the highest-probability untargeted cell.
	target := p.bestTarget(opponentBoard)
	targetRow, targetCol := target.row, target.col

	// Perform the attack
	isHit := opponentBoard[targetRow][targetCol] == ship

	// check to see if we hit a ship
	if isHit {
		opponentBoard[targetRow][targetCol] = hit
		fmt.Printf("AI targets %c%d... HIT!\n", 'A'+targetCol, targetRow)

		// Update hit tracking
		p.hits = append(p.hits, Position{targetRow, targetCol})

		// if we hit, enter hunt mode
		p.huntMode = true

		// check to see if ship was sunk
		sunk, shipName := isShipSunk(opponentBoard, targetRow, targetCol, p.opponent.ships)
		if sunk {
			fmt.Printf("AI sunk your %s!\n", shipName)
			p.shipsSunk++
			p.huntMode = false    // exit hunt mode when a ship is shipsSunk
			p.hits = []Position{} // reset hits tracking for next hunt
		}
	} else {
		opponentBoard[targetRow][targetCol] = miss
		fmt.Printf("AI targets %c%d... Miss.\n", 'A'+targetCol, targetRow)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nPress enter to continue...")
	reader.ReadString('\n')

	return Position{targetRow, targetCol}, isHit
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

			for _, pos := range shipCells(Position{row: row, col: col}, shipType.size, horizontal) {
				r, c := pos.row, pos.col

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

				positions = append(positions, pos)
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

				for _, pos := range shipCells(Position{row: row, col: col}, shipType.size, horizontal) {
					if p.board[pos.row][pos.col] == ship {
						valid = false
						break
					}

					positions = append(positions, pos)
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
