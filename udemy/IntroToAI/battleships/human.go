package main

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

}
