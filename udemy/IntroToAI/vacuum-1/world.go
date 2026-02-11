package main

import "time"

const (
	// Display characters
	charRobot          = "🔴"
	charWall           = "🟦"
	charFurniture      = "🪑"
	charClean          = "🧽"
	charDirty          = "🟫"
	carPath            = "🟢"
	charCat            = "🐱" // Display character for the cat
	catStopProbability = 0.1 // Probability that the cat will stop moving at each step
	catStopDuration    = 5   // Duration (in seconds) that the cat will stop moving when it decides to stop
	moveDelay          = 50 * time.Millisecond
	cellSize           = 10
)

type Point struct {
	X, Y int
}

type Cell struct {
	Type         string // "clean", "dirty", "wall", "furniture", "bike"
	Cleaned      bool
	Obstacle     bool
	ObstacleName string
}

type Furniture struct {
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Name   string `json:"name"`
	Type   string `json:"type"`
}

type Room struct {
	Grid               [][]Cell
	Width              int
	Height             int
	CleanableCellCount int
	CleanedCellCount   int
	Animate            bool
}

type RoomConfig struct {
	Width     int         `json:"width"`
	Height    int         `json:"height"`
	Furniture []Furniture `json:"furniture"`
}

func NewRoom(configFile string, animate bool) *Room {
	return &Room{}
}
