package main

import "time"

func CleanRoomSlam(room *Room, robot *Robot) {
	// Set start time and movecount
	startTime := time.Now()
	moveCount := 0

	// Initialize the robot's internal map
	robotMap := initializeRobotMap(room.Width, room.Height)

	// Initialize visited cells (tracking)
	visited := make(map[Point]bool)

	// Initialize a frontier
	frontier := make([]Point, bool)

	// Mark starting position as visited and update map for the first time
	visited[robot.Position] = true
	updateRobotMap(robot.Position, robotMap, room)

	// Clean the current position
	Clean(robot, room)

	// Add neighbors to the frontier

	// Display the initial state

	// for - if the frontier is not empty and the room is not fully cleaned
	for {

		// Get closest frontier point

		// If not valid target, break

		// remove target from frontier

		// find path to target using A*

		// If not path found, go to next frontier point (continue)

		// Move along the path
		for {

			// Update map (internal)

		}

		// every 10 minutes, do a more thorough frontier check

		// check if we have sufficient coverage -- break
	}

	// final cleanup phase

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
