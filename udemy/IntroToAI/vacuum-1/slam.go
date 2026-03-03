package main

import (
	"math"
	"time"
)

func CleanRoomSlam(room *Room, robot *Robot) {
	// Set start time and movecount
	startTime := time.Now()
	moveCount := 0

	// Initialize the robot's internal map
	robotMap := initializeRobotMap(room.Width, room.Height)

	// Initialize visited cells (tracking)
	visited := make(map[Point]bool)

	// Initialize a frontier
	frontier := make(map[Point]bool)

	// Mark starting position as visited and update map for the first time
	visited[robot.Position] = true
	updateRobotMap(robot.Position, robotMap, room)

	// Clean the current position
	Clean(robot, room)

	// Add neighbors to the frontier
	addNeighborsToFrontier(robot.Position, robotMap, frontier, visited, room)

	// Display the initial state
	if room.Animate {
		room.Display(robot, false)
		time.Sleep(moveDelay)
	}

	// for - if the frontier is not empty and the room is not fully cleaned
	for len(frontier) > 0 && room.CleanedCellCount < room.CleanableCellCount {
		// Get closest frontier point
		target := getCloesestFrontierPoint(robot.Position, frontier)

		// If not valid target, break
		if target.X == -1 && target.Y == -1 {
			break
		}

		// remove target from frontier
		delete(frontier, target)

		// find path to target using A*
		path := Astar(room, robot.Position, target)

		// If not path found, go to next frontier point (continue)
		if len(path) <= 1 {
			continue
		}

		// Move along the path
		for i := 1; i < len(path); i++ {
			// Update robot position
			robot.Position = path[i]
			robot.Path = append(robot.Path, robot.Position)

			// Clean the current position
			Clean(robot, room)

			// Mark as visited and update map
			visited[robot.Position] = true
			// Update map (internal) based on what we can see
			updateRobotMap(robot.Position, robotMap, room)

			// Updatre frontier with newly discovered cells
			addNeighborsToFrontier(robot.Position, robotMap, frontier, visited, room)

			// Display the room
			if room.Animate {
				room.Display(robot, false)
				time.Sleep(moveDelay)
			}

			moveCount++
		}

		// every 10 minutes, do a more thorough frontier check
		if moveCount%10 == 0 {
			updateAllFrontiers(robotMap, frontier, visited, room)
		}

		// check if we have sufficient coverage -- break
		if float64(room.CleanedCellCount)/float64(room.CleanableCellCount) > 0.95 {
			break
		}
	}

	// final cleanup phase
	cleanRemainingCells(room, robot, &moveCount)

	// Calculate cleaning time
	cleaningTime := time.Since(startTime)

	// Display final statistics
	displaySummary(room, robot, moveCount, cleaningTime)

}

func cleanRemainingCells(room *Room, robot *Robot, moveCount *int) {
	// Find all cells that should be ceanable but haven't been cleaned
	for i := 1; i < room.Width-1; i++ {
		for j := 1; j < room.Height-1; j++ {
			// If the cell is not an obstacle, not cleaned, and known to the robot
			if !room.Grid[i][j].Obstacle && !room.Grid[i][j].Cleaned {
				path := Astar(room, robot.Position, Point{X: i, Y: j})
				if len(path) <= 1 {
					continue
				}

				// Move along path
				for k := 1; k < len(path); k++ {
					// Update robot position
					robot.Position = path[k]
					robot.Path = append(robot.Path, robot.Position)

					// Clean the current position
					Clean(robot, room)

					// Display the room
					if room.Animate {
						room.Display(robot, false)
						time.Sleep(moveDelay)
					}
					*moveCount++
				}
			}
		}
	}
}

func updateAllFrontiers(robotMap [][]int, frontier map[Point]bool, visited map[Point]bool, room *Room) {
	for x := 0; x < room.Width-1; x++ {
		for y := 0; y < room.Height-1; y++ {
			// If a cell is free but not visited, add to frontier
			point := Point{X: x, Y: y}
			if robotMap[x][y] == 1 && !visited[point] && !frontier[point] && room.Grid[x][y].Obstacle {
				// Check to see if it is accessible (has at least one visited neighbor)
				for _, dir := range directions {
					nx, ny := x+dir[0], y+dir[1]
					neighborPoint := Point{X: nx, Y: ny}
					if nx >= 0 && nx < room.Width && ny >= 0 && ny < room.Height && visited[neighborPoint] {
						frontier[point] = true
						break
					}
				}

			}
		}
	}
}

func getCloesestFrontierPoint(position Point, frontier map[Point]bool) Point {
	closestPoint := Point{X: -1, Y: -1}
	minDistance := math.MaxFloat64

	for point := range frontier {
		distance := heuristic(position, point)
		if distance < minDistance {
			minDistance = distance
			closestPoint = point
		}
	}

	return closestPoint
}

func initializeRobotMap(width, height int) [][]int {
	//  = unknown, 1 = free, 2 = obstacle, 3 = cleaned
	robotMap := make([][]int, width)
	for i := range robotMap {
		robotMap[i] = make([]int, height)
	}

	return robotMap
}

func updateRobotMap(position Point, robotMap [][]int, room *Room) {
	if room.Grid[position.X][position.Y].Cleaned {
		robotMap[position.X][position.Y] = 3 // cleaned
	} else {
		robotMap[position.X][position.Y] = 1 // free
	}

	// Scan surroundings
	for _, dir := range directions {
		newX, newY := position.X+dir[0], position.Y+dir[1]

		// Check if position is within bounds
		if newX >= 0 && newX < len(robotMap) && newY >= 0 && newY < len(robotMap[0]) {
			if room.Grid[newX][newY].Obstacle {
				robotMap[newX][newY] = 2 // obstacle
			} else if robotMap[newX][newY] == 0 {
				robotMap[newX][newY] = 1 // free
			} else if room.Grid[newX][newY].Cleaned {
				robotMap[newX][newY] = 3 // cleaned
			}
		}
	}
}

func addNeighborsToFrontier(position Point, robotMap [][]int, frontier map[Point]bool, visited map[Point]bool, room *Room) {
	// Check adjacent cells
	for _, dir := range directions {
		newX, newY := position.X+dir[0], position.Y+dir[1]
		newPoint := Point{X: newX, Y: newY}

		// Check if position is valid, not visited, not an obstacle, and not already in frontier
		if newX >= 0 && newX <= len(robotMap) && newY >= 0 && newY < len(robotMap[0]) &&
			!visited[newPoint] && !frontier[newPoint] && room.IsValid(newX, newY) {
			// Add to frontier
			frontier[newPoint] = true
		}
	}
}
