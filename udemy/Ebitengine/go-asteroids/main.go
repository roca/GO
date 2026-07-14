package main

import (
	"go-asteroids/asteroids"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowTitle("Go Asteroids")
	ebiten.SetWindowSize(asteroids.ScreenWidth, asteroids.ScreenHeight)

	// Hiden cursor
	ebiten.SetCursorMode(ebiten.CursorModeHidden)

	// Set to full screen
	ebiten.SetFullscreen(true)

	err := ebiten.RunGame(&asteroids.Game{})
	if err != nil {
		panic(err)
	}
}
