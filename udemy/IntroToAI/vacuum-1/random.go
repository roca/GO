package main

import (
	"math"
	"time"
)

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

func bresenhamLine(x0, y0, x1, y1 int) []Point {
	// https://www.tutorialspoint.com/computer_graphics/bresenhams_line_generation_algorithm.htm

	// Intialize a slice to store all points on the line
	var points []Point

	// Calulate the absolute difference between the endpoints
	dx := Abs(x1 - x0)
	dy := Abs(y1 - y0)

	// Determine the direction of movement along each axis
	sx := -1
	if x0 < x1 {
		sx = 1
	}
	sy := -1
	if y0 < y1 {
		sy = 1
	}

	// Calculate the initial error value
	err := dx - dy

	// for loop until we reach the end point
	for {
		// Add the current point to our result
		points = append(points, Point{X: x0, Y: y0})

		// Check to see if we've reach the end point
		if x0 == x1 && y0 == y1 {
			break
		}

		// Calculate the error for the next step
		e2 := err * 2

		// If moving in the x direction would keep us closer to the ideal line
		if e2 > -dy {
			// If we've reached the endpoint, stop
			if x0 == x1 {
				break
			}
			// Update the error and move in the x direction
			err -= dy
			x0 += sx
		}

		// If moving in the y direction would keep us closer to the ideal line
		if e2 < dx {
			// If we've reached the endpoint, stop
			if y0 == y1 {
				break
			}
			// Update the error and move in the y direction
			err += dx
			y0 += sy
		}

	}

	// return the list of points
	return points
}

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func findNearestDirtyCell(room *Room, position Point) Point {
	var nearestCell Point = Point{X: -1, Y: -1}
	minDistance := math.MaxFloat64

	for i := 1; i <= room.Width-1; i++ {
		for j := 1; j <= room.Height-1; j++ {
			distance := heuristic(position, Point{X: i, Y: j})
			if distance < minDistance {
				minDistance = distance
				nearestCell = Point{X: i, Y: j}
			}
		}
	}

	return nearestCell
}
