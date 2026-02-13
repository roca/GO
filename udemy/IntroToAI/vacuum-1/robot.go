package main

type Robot struct {
	Position             Point
	Path                 []Point
	CleanRoom            func(*Room, *Robot)
	Direction            float64
	ObstaclesEncountered map[string]bool
}

func NewRobot(startX, startY int) *Robot {
	return &Robot{
		Position:             Point{X: startX, Y: startY},
		Path:                 []Point{{X: startX, Y: startY}},
		ObstaclesEncountered: make(map[string]bool),
		CleanRoom:            CleanRoomRandomWalk,
	}
}
