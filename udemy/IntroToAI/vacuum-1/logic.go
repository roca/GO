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

	// Rule 1: If a backpack is found, then Jack is home.
	if objectName == "backpack" {
		if !world.Jack.IsHome {
			fmt.Println("Logic: backpack found, deducing Jack is home.")
		}
		world.Jack.IsHome = true
	}

	// Rule 2: If a bicycle is found, then Sarah is home.
	if objectName == "bicycle" {
		if !world.Sarah.IsHome {
			fmt.Println("Logic: bicycle found, deducing Sarah is home.")
		}
		world.Sarah.IsHome = true
	}

	// Rule 3: If a skateboard is found, then Johnny is home.
	if objectName == "skateboard" {
		if !world.Johnny.IsHome {
			fmt.Println("Logic: skateboard found, deducing Johnny is home.")
		}
		world.Johnny.IsHome = true
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

// RobotWithLogic is a specialized robot that has access to the logical world and can make decisions based on it.
type RobotWithLogic struct {
	*Robot // Embed the basic Robot struct to inherit its properties and methods.
	World  *LogicalWorld
}

// NewRobotWithLogic creates a new RobotWithLogic instance with an initialized logical world.
func NewRobotWithLogic(startX, startY int) *RobotWithLogic {
	return &RobotWithLogic{
		Robot: NewRobot(startX, startY),
		World: NewLogicalWorld(),
	}
}

func (robot *RobotWithLogic) ScanHouseWithLogic(house *House) map[string]int {
	// Create a map which maps room indices to room names.
	roomNameToIndex := make(map[string]int)

	// Identify all rooms and generate our mapping.
	for i, room := range house.Rooms {
		roomName := ""

		for x := range room.Width {
			for y := range room.Height {

				if room.Grid[x][y].Type == "furniture" {
					if roomName == "" {
						switch room.Grid[x][y].ObstacleName {
						case "bed":
							if roomName == "" && roomNameToIndex["Jack's Room"] == 0 {
								roomName = "Jack's Room"
							} else if roomNameToIndex["Sarah's Room"] == 0 && roomNameToIndex["Johnny's Room"] == 0 {
								roomName = "Sarah's Room"
							} else {
								roomName = "Johnny's Room"
							}
						case "desk":
							if roomName == "" {
								roomName = "study"
							}
						case "sofa", "tv":
							roomName = "Living Room"
						case "fridge", "stove", "sink":
							roomName = "Kitchen"
						}
					}

					// Johnny's door
					if room.Grid[x][y].ObstacleName == "johnny's door" {
						// change to true to close the door, false to open it.
						robot.World.UpdateDoorStatus("Johnny's Door", false) // Assume door is open when we see it.
					}
				}
			}
		}

		// If we can't determine the name of the room, give it a default name.
		if roomName == "" {
			roomName = fmt.Sprintf("Room %d", i)
		}

		roomNameToIndex[roomName] = i
		fmt.Printf("Identified room %s (index %d)\n", roomName, i)
	}

	// Scan the house for objects to build out logical world.
	fmt.Println("Robot is scanning the house for objects....")

	for _, room := range house.Rooms {
		for x := range room.Width {
			for y := range room.Height {
				if room.Grid[x][y].Type == "furniture" && room.Grid[x][y].ObstacleName != "" {
					robot.World.UpdateObjectFound(room.Grid[x][y].ObstacleName)
				}
			}
		}
	}

	return roomNameToIndex
}
