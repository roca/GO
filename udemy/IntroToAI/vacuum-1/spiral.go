package main

import "time"

func CleanSpiralPattern(room *Room, robot *Robot) {
	// Set start time and movecount
	startTime := time.Now()
	moveCount := 0

	// Find to the center of the room

	// Find a valid point near the center

	// Find path to center (using A*).

	// Move to the center point.

	// Create a spiral pattern.

	// Follow the spiral pattern (for loop)
	for {
		// Skip if cell is already cleaned, or an obstacle.

		// Find path to the next point (using A*).

		// Move along path

	}

	// Final cleanup

	// Calulate cleaning time
	cleaningTime := time.Since(startTime)

	// Dispaly final statistics
	displaySummary(room, robot, moveCount, cleaningTime)
}
