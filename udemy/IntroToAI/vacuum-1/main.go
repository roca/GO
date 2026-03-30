package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var configFile, algorithm string
	var animate, cat, isHouse, useLogic bool
	flag.StringVar(&configFile, "file", "empty.json", "configuration file")
	flag.StringVar(&algorithm, "algorithm", "snake", "cleaning algorithm")
	flag.BoolVar(&animate, "animate", true, "animate while cleaning")
	flag.BoolVar(&cat, "cat", false, "include a cat in the room")
	flag.BoolVar(&isHouse, "house", false, "config file has multiple rooms (house)")
	flag.BoolVar(&useLogic, "logic", false, "use propositional logic for cleaning decisions")
	flag.Parse()

	var house *House

	if !isHouse {
		// If not a house , we just have one room. Create a house, and assign one room to it.
		// This way, we can use the same loop for houses and for individual rooms.
		var rooms []*Room

		// get a room from JSON config.
		room := NewRoom(configFile, animate)
		rooms = append(rooms, room)
		var h House
		h.Rooms = rooms
		house = &h
	} else {
		// We are doing a complete house. Just get a house from JSON config.
		house = NewHouse(configFile, animate)
	}

	// Add cats to rooms if nessary.
	if cat {
		for _, room := range house.Rooms {
			room.Cat = NewCat(room)
		}
	}

	roomCount := 0

	if useLogic {
		// Use propositional logic to make cleaning decisions.
		fmt.Println("Using propositional logic for cleaning decisions.")
		robot := NewRobotWithLogic(1, 1)

		// Assign a cleaning algorithm.
		setAlgorithm(algorithm, robot.Robot)

		// Scan the house
		roomNameToIndex := robot.ScanHouseWithLogic(house)
		fmt.Println("\nLogical state after scanning:")
		fmt.Printf("Today is %s (Weekday: %t)\n", time.Now().Weekday(), robot.World.IsWeekday)
		fmt.Printf("Jack is home: %t\n", robot.World.Jack.IsHome)
		fmt.Printf("Sarah is home: %t\n", robot.World.Sarah.IsHome)
		fmt.Printf("Johnny is home: %t\n", robot.World.Johnny.IsHome)
		fmt.Printf("Johnny's Door is closed: %t\n", robot.World.Johnny.DoorClosed)
		if robot.World.Johnny.DoorClosed {
			fmt.Println("Logic: Will not vacuum Johnny's room because his door is closed.")
		} else {
			fmt.Println("Logic: Will vacuum Johnny's room because his door is open.")
		}

		// Determine cleaning priority based on logical rules
		cleaningPriority := robot.World.DetermineCleaningPriority()
		fmt.Println("\nDetermined cleaning priority based on propositional logic:")
		for i, roomName := range cleaningPriority {
			fmt.Printf("%d: %s\n", i+1, roomName)
		}

		for k,  v := range roomNameToIndex {
			fmt.Println(k, "->", v)
		}

		fmt.Println("\nPress enter to start cleaning...")
		fmt.Scanln() // Wait for user input before starting cleaning.

	} else {
		// Use the original cleaning approach without propositional logic, and for multiple rooms.
		for _, room := range house.Rooms {

			// Get a robot.
			robot := NewRobot(1, 1)

			// Assign a cleaning algorithm.
			setAlgorithm(algorithm, robot)

			// Clean the room.
			robot.CleanRoom(room, robot)
			roomCount++
		}
	}

	fmt.Printf("All done. Cleaned a total of %d room(s)\n", roomCount)
}

// SetAlgorithm is a helper function to assign a cleaning algorithm to a robot based on the algorithm name.
func setAlgorithm(algorithm string, robot *Robot) {
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
}
