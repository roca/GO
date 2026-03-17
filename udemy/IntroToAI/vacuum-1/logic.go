package main

import (
	"fmt"
	"time"
)

type PersonStatus struct {
	Name       string
	IsHome     bool
	Room       string // Person's room name
	DoorClosed bool
}

type LogicalWorld struct {
	Jack      PersonStatus
	Sarah     PersonStatus
	Johnny    PersonStatus
	IsWeekday bool
	Objects   map[string]bool
}

func NewLogicalWorld() *LogicalWorld {
	// Get current day to determine if it's a weekday or weekend.
	today := time.Now()
	weekday := today.Weekday()

	isWeekday := weekday >= time.Monday && weekday <= time.Friday

	return &LogicalWorld{
		Jack:      PersonStatus{Name: "Jack", IsHome: false, Room: "Jack's Room"},
		Sarah:     PersonStatus{Name: "Sarah", IsHome: false, Room: "Sarah's Room"},
		Johnny:    PersonStatus{Name: "Johnny", IsHome: false, Room: "Johnny's Room", DoorClosed: false},
		IsWeekday: isWeekday,
		Objects:   make(map[string]bool),
	}
}

func (world *LogicalWorld) UpdateObjectFound(objectName string) {
	world.Objects[objectName] = true

	// Apply object to person identification rules.

	// Rule 1: Ifs a backpack is found, then Jack is home.
	if objectName == "backpack" {
		world.Jack.IsHome = true
		fmt.Println("Logic: backpack found, deducing Jack is home.")
	}

	// Rule 2: If a bicycle is found, then Sarah is home.
	if objectName == "bicycle" {
		world.Sarah.IsHome = true
		fmt.Println("Logic: bicycle found, deducing Sarah is home.")
	}

	// Rule 3: If a skateboard is found, then Johnny is home.
	if objectName == "skateboard" {
		world.Johnny.IsHome = true
		fmt.Println("Logic: skateboard found, deducing Johnny is home.")
	}
}

// UpdateDoorStatus updates whether Johnny's door is open or closed.
func (world *LogicalWorld) UpdateDoorStatus(doorName string, isClosed bool) {
	if doorName == "Johnny's Door" {
		world.Johnny.DoorClosed = isClosed
		fmt.Printf("Logic: %s is now %s.\n", doorName, map[bool]string{true: "closed", false: "open"}[isClosed])
	}
}

// DetermineCleaningPriority decide on the order we clean rooms.
func (world *LogicalWorld) DetermineCleaningPriority() []string {
	availableRooms := []string{
		"Kitchen",
		"Living Room",
		"Jack's Room",
		"Sarah's Room",
		"Johnny's Room",
	}

	// Rule: If no one is home, then vacuum all rooms starting with the kitchen.
	if !world.Jack.IsHome && !world.Sarah.IsHome && !world.Johnny.IsHome {
		fmt.Println("Logic: No one is home, vacuuming all rooms starting with the kitchen.")
		return availableRooms
	}

	// Initialize a priority list with all available rooms.
	priorityList := make([]string, 0)
	skipRooms := make(map[string]bool)

	// Rule: If Sarah is home, then don't vacuum the living room.
	if world.Sarah.IsHome {
		fmt.Println("Logic: Sarah is home, skipping the living room.")
		skipRooms["Living Room"] = true
	}

	// Rule: If Johnny is home and his door is closed, then don't vacuum Johnny's room.
	if world.Johnny.IsHome && world.Johnny.DoorClosed {
		fmt.Println("Logic: Johnny is home and his door is closed, skipping Johnny's room.")
		skipRooms["Johnny's Room"] = true
	}

	// Rule: If Jack is home and it's a weekday, then vacuum Jack's room last.
	jackRoomLast := world.Jack.IsHome && world.IsWeekday

	// Build our priority list based on the rules. Add kitchen first, and filter out skipped rooms.
	priorityList = append(priorityList, "Kitchen")

	// Add all other rooms except the ones we need to skip and Jack's room if it should be last.
	for _, room := range availableRooms {
		if room == "Kitchen" {
			continue // Kitchen is already added as the highest priority.
		}
		if room == "Jack's Room" && jackRoomLast {
			continue // Skip Jack's room for now if it should be last.
		}
		if skipRooms[room] {
			continue // Skip rooms based on the rules.
		}
		priorityList = append(priorityList, room)
	}

	// Finally, add Jack's room at the end if it should be last.
	if jackRoomLast && !skipRooms["Jack's Room"] {
		priorityList = append(priorityList, "Jack's Room")
	}

	return priorityList
}

// Type for a robot with logic.

// Factory function to create a new robot with logic.
