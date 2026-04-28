package main

type AIPlayer struct {
	board          Board
	heatMap        [boardSize][boardSize]int
	hits           []Position
	huntMode       bool
	potentialShips []struct {
		size     int
		sunk     bool
		hits     int
		shipPos  []Position
		ships    []Ship
		opponent *HumanPlayer
	}
}
