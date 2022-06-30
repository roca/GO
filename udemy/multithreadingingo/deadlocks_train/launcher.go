package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	. "github.com/roca/GO/tree/staging/udemy/multithreadingingo/deadlocks_train/common"
	. "github.com/roca/GO/tree/staging/udemy/multithreadingingo/deadlocks_train/deadlock"
)

var (
	trains        [4]*Train
	intersections [4]*Intersection
)

const trainLength = 70

func update(screen *ebiten.Image) error {
	if !ebiten.IsDrawingSkipped() {
		DrawTracks(screen)
		DrawIntersections(screen)
		DrawTrains(screen)
	}
	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 320, 3, "Trains in a box"); err != nil {
		log.Fatal(err)
	}
}
