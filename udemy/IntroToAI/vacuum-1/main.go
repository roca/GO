package main

import (
	"flag"
	"fmt"
)

func main() {
	var configFile, algorithm string
	var animate bool
	flag.StringVar(&configFile, "file", "empty.json", "configuration file")
	flag.StringVar(&algorithm, "algorithm", "random", "cleaning algorithm")
	flag.BoolVar(&animate, "animate", true, "animate while cleaning")
	flag.Parse()

	room := NewRoom(configFile, animate)
	fmt.Println(room.CleanableCellCount)

	// Get a robot.
	robot := NewRobot(1, 1)

	// Assign a cleaning algorithm.
	switch algorithm {
	case "random":
		robot.CleanRoom = CleanRoomRandomWalk
	case "slam":
		robot.CleanRoom = CleanRoomSlam
	default:
		// Do nothing.
	}

	// Clean the room.
	robot.CleanRoom(room, robot)
}

