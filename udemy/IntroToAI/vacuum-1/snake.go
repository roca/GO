package main

import "time"

func CleanRoomSnake(room *Room, robot *Robot) {
	// Initialize start time and moveCount
	startTime := time.Now()
	moveCount := 0

	// Generate the snaking pattern.
	coveragePoints := generateSnakingPattern(room)

	// Clean cell
	Clean(robot, room)

	if room.Animate {
		room.Display(robot, false)
		time.Sleep(moveDelay)
	}

	// Visit each point in the coverage pattern (for)
	for _, point := range coveragePoints {
		// Skip cell if already clean.
		if room.Grid[point.X][point.Y].Cleaned {
			continue
		}

		// Find path to the next point
		path := Astar(room, robot.Position, point)

		// If no path found, try the next point
		if len(path) == 0 {
			continue
		}

		// Move along the path (for)
		for i := 1; i < len(path); i++ {
			// Update robot position
			robot.Position = path[i]
			robot.Path = append(robot.Path, robot.Position)

			// Clean the cell
			Clean(robot, room)

			// Display the room if Animate set to true
			if room.Animate {
				room.Display(robot, false)
				time.Sleep(moveDelay)
			}

			// Increment moveCount
			moveCount++

		}
	}

	// Do final sweep
	finalCleanup(room, robot, &moveCount)

	cleaningTime := time.Since(startTime)

	// Display final statistics
	displaySummary(room, robot, moveCount, cleaningTime)

}

func generateSnakingPattern(room *Room) []Point {
	var points []Point

	var directiionX = 1

	for y := 1; y < room.Height-1; y++ {
		if directiionX == 1 {
			// Moving left to right
			for x := 1; x < room.Width-1; x++ {
				if !room.Grid[x][y].Obstacle {
					points = append(points, Point{X: x, Y: y})
				}
			}
		} else {
			// Moving right to left
			for x := room.Width - 2; x >= 1; x-- {
				if !room.Grid[x][y].Obstacle {
					points = append(points, Point{X: x, Y: y})
				}
			}
		}

		directiionX *= -1 // Change direction for the next row
	}

	return points
}
