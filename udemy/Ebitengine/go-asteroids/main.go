package main

import (
	"go-asteroids/asteroids"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowTitle("Go Asteroids")
	ebiten.SetWindowSize(asteroids.ScreenWidth, asteroids.ScreenHeight)

	err := ebiten.RunGame(&asteroids.Game{})
	if err != nil {
		panic(err)
	}
}
