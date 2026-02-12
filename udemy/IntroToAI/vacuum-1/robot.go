package main

type Robot struct {
	Position             Point
	Path                 []Point
	CleanRoom            func(*Room, *Robot)
	Direction            float64
	ObstaclesEncountered map[string]bool
}

func NewRobot() *Robot {
	return &Robot{}
}
