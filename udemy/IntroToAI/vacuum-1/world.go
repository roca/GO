package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	// Display characters
	charRobot          = "🔴"
	charWall           = "🟦"
	charFurniture      = "🪑"
	charClean          = "🧽"
	charDirty          = "🟫"
	charPath           = "🟢"
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
	// Load from JSON config configFile
	roomConfig, err := LoadRoomConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to load room config: %v", err)
	}

	// Convert dimensions to grid cells.
	gridWidth := roomConfig.Width / cellSize
	gridHeight := roomConfig.Height / cellSize

	// Create grid
	grid := make([][]Cell, gridWidth)
	for i := range grid {
		grid[i] = make([]Cell, gridHeight)
		for j := range grid[i] {
			grid[i][j] = Cell{Type: "dirty", Cleaned: false, Obstacle: false}
		}
	}

	// Add walls
	for i := 0; i < gridWidth; i++ {
		grid[i][0] = Cell{Type: "wall", Cleaned: false, Obstacle: true, ObstacleName: "wall"}
		grid[i][gridHeight-1] = Cell{Type: "wall", Cleaned: false, Obstacle: true, ObstacleName: "wall"}
	}
	for j := 0; j < gridHeight; j++ {
		grid[0][j] = Cell{Type: "wall", Cleaned: false, Obstacle: true, ObstacleName: "wall"}
		grid[gridWidth-1][j] = Cell{Type: "wall", Cleaned: false, Obstacle: true, ObstacleName: "wall"}
	}

	// Add furniture
	for _, f := range roomConfig.Furniture {
		x := f.X / cellSize
		y := f.Y / cellSize
		width := f.Width / cellSize
		height := f.Height / cellSize

		for i := x; i < x+width; i++ {
			for j := y; j < y+height; j++ {
				grid[i][j] = Cell{Type: f.Type, Cleaned: false, Obstacle: true, ObstacleName: f.Name}
			}
		}
	}

	// Count cleanable cells
	cleanableCellCount := 0
	for i := 0; i < gridWidth; i++ {
		for j := 0; j < gridHeight; j++ {
			if !grid[i][j].Obstacle {
				cleanableCellCount++
			}
		}
	}

	return &Room{
		Grid:               grid,
		Width:              gridWidth,
		Height:             gridHeight,
		CleanableCellCount: cleanableCellCount,
		CleanedCellCount:   0,
		Animate:            animate,
	}
}

func LoadRoomConfig(filename string) (*RoomConfig, error) {
	// Implement JSON loading logic here

	// Read JSON file.
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config RoomConfig
	// Parse JSON data.
	if err := json.Unmarshal(jsonData, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (r *Room) Display(robot *Robot, showPath bool) {
	// Clear the screen
	fmt.Print("\033[H\033[2J")

	for j := range r.Height {
		for i := range r.Width {
			if robot.Position.X == i && robot.Position.Y == j {
				fmt.Print(charRobot)
			} else if showPath && isInPath(Point{X: i, Y: j}, robot.Path) {
				fmt.Print(charPath)
			} else {
				cell := r.Grid[i][j]
				switch cell.Type {
				case "wall":
					fmt.Print(charWall)
				case "furniture":
					fmt.Print(charFurniture)
				case "clean":
					fmt.Print(charClean)
				case "dirty":
					fmt.Print(charDirty)
				}
			}
		}
		fmt.Println()
	}

	// Dispaly cleaning progress
	percentCleaned := float64(r.CleanedCellCount) / float64(r.CleanableCellCount) * 100
	fmt.Printf("Cleaning Progress: %.2f%% (%d/%d) cells cleaned)\n", percentCleaned, r.CleanedCellCount, r.CleanableCellCount)
}

func isInPath(point Point, path []Point) bool {
	for _, p := range path {
		if p.X == point.X && p.Y == point.Y {
			return true
		}
	}
	return false
}

func displaySummary(room *Room, robot *Robot, moveCount int, cleaningTime time.Duration) {
	// Display the final room state with the robot's path
	fmt.Println("\nFinal room state with robot's path:")
	room.Display(robot, true)

	// Display summary
	fmt.Println("\n============= Cleaning Summary =============")
	fmt.Printf("Room size: %d x %d (%d cm x %d cm)\n", room.Width, room.Height, room.Width*cellSize, room.Height*cellSize)

	//Calulate coverage percentage
	percentCleaned := float64(room.CleanedCellCount) / float64(room.CleanableCellCount) * 100
	fmt.Printf("Coverage: %.2f%% (%d/%d cells cleaned)\n", percentCleaned, room.CleanedCellCount, room.CleanableCellCount)

	//Display time and moves
	fmt.Printf("Total moves: %d\n", moveCount)
	fmt.Printf("Cleaning time: %v\n", cleaningTime)

	// Calculate efficiency (cells cleaned per move)
	efficiency := float64(room.CleanedCellCount) / float64(moveCount)
	fmt.Printf("Efficiency: %.2f cells cleaned per move\n", efficiency)

	fmt.Println()
	fmt.Println("===========================================")
}

func (room *Room) IsValid(x, y int) bool {
	if x >= 0 && x < room.Width && y >= 0 && y < room.Height && !room.Grid[x][y].Obstacle {
		return true
	}
	return false
}
