package main

var directions = [][]int{
	{0, -1}, // North
	{1, 0},  // East
	{0, 1},  // South
	{-1, 0}, // West
}

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
	}
}

func Clean(robot *Robot, room *Room) {
	x, y := robot.Position.X, robot.Position.Y

	if !room.Grid[x][y].Cleaned && !room.Grid[x][y].Obstacle {
		room.Grid[x][y].Cleaned = true
		room.Grid[x][y].Type = "clean"
		room.CleanedCellCount++
	}

	CheckAdjacentObstacles(robot, room)
}

func CheckAdjacentObstacles(robot *Robot, room *Room) {
	x, y := robot.Position.X, robot.Position.Y

	for _, dir := range directions {
		newX := x + dir[0]
		newY := y + dir[1]
		RecordObstacle(robot, room, newX, newY)
	}
}

func RecordObstacle(robot *Robot, room *Room, x, y int) {
	if x >= 0 && x < room.Width && y >= 0 && y <= room.Height && room.Grid[x][y].Obstacle {
		if room.Grid[x][y].Type == "furniture" && room.Grid[x][y].ObstacleName != "" {
			robot.ObstaclesEncountered[room.Grid[x][y].ObstacleName] = true
		}
	}
}