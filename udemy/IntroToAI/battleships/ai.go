package main

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
