package main

import (
	"flag"
)

func main() {
	var configFile, algorithm string
	var animate, cat bool
	flag.StringVar(&configFile, "file", "empty.json", "configuration file")
	flag.StringVar(&algorithm, "algorithm", "snake", "cleaning algorithm")
	flag.BoolVar(&animate, "animate", true, "animate while cleaning")
	flag.BoolVar(&cat, "cat", false, "include a cat in the room")
	flag.Parse()

	room := NewRoom(configFile, animate)
	if cat {
		room.Cat = NewCat(room)
	}

	// Get a robot.
	robot := NewRobot(1, 1)

	// Assign a cleaning algorithm.
	switch algorithm {
	case "random":
		robot.CleanRoom = CleanRoomRandomWalk
	case "slam":
		robot.CleanRoom = CleanRoomSlam
	case "spiral":
		robot.CleanRoom = CleanSpiralPattern
	default:
		// Default to snaking pattern.
		robot.CleanRoom = CleanRoomSnake
	}

	// Clean the room.
	robot.CleanRoom(room, robot)
}
