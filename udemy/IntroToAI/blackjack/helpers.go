package main

import (
	"fmt"
	"runtime"

	"github.com/inancgumus/screen"
)

func clearScreen() {
	if runtime.GOOS == "windows" {
		screen.Clear()
		screen.MoveTopLeft()
	} else {
		fmt.Print("\033[H\033[2J")
	}

}
