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
