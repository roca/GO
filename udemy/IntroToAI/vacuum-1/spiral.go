package main

import "time"

func CleanSpiralPattern(room *Room, robot *Robot) {
	// Set start time and movecount
	startTime := time.Now()
	moveCount := 0

	// Find to the center of the room
	centerX := room.Width / 2
	centerY := room.Height / 2

	// Find a valid point near the center
	centerPoint := findNearestCleanablePoint(room, Point{X: centerX, Y: centerY})

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

func findNearestCleanablePoint(room *Room, target Point) Point {
	if room.IsValid(target.X, target.Y) && !room.Grid[target.X][target.Y].Obstacle {
		return target
	}

	// Search for a valid point in expanding circles
	for radius := 1; radius < room.Width || radius < room.Height; radius++ {
		// Check all points at the current radius
		for dx := -radius; dx <= radius; dx++ {
			for dy := -radius; dy <= radius; dy++ {
				if abs(dx) != radius && abs(dy) != radius {
					continue
				}

				x, y := target.X+dx, target.Y+dy

				// Check to see if this point is valid and not an obstical
				if room.IsValid(x, y) && !room.Grid[x][y].Obstacle {
					return Point{X: x, Y: y}
				}
			}
		}
	}
	// If no valid point, return the starting point
	return Point{X: 1, Y: 1}
}
