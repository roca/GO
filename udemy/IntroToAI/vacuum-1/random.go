package main

import "time"

func CleanRoomRandomWalk(room *Room, robot *Robot) {
	startTime := time.Now()
	moveCount := 0

	// Set variables
	maxMoves := room.Width * room.Height * 5
	stuckCount := 0
	maxStuckCount := 5 // Max number of consecutive failed moves before changing strategy

	// Clean current position
	Clean(robot, room)
	if room.Animate {
		room.Display(robot, false)
		time.Sleep(moveDelay)
	}

	for moveCount < maxMoves && room.CleanableCellCount < room.CleanedCellCount {
		// Generate a random angle in radians

		// Calulate a direction vector based on the random angle

		// Use Bressenham's line algorithm to determine the path from the robot's current position to the target position

		// If we didn't move very much, increment stuck counter and possibly change strategy

		// If stuck too many times, use A* to find a path to nearest dirty cell

		// Add some adaptive behavior, Scan for dirty cells every once in awhile

		// end for

		// Final sweep to ensure complete coverage
	}

	// Calulate cleaning time
	cleaningTime := time.Since(startTime)

	displaySummary(room, robot, moveCount, cleaningTime)
}
