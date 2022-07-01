package main

import (
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	. "github.com/roca/GO/tree/staging/udemy/multithreadingingo/deadlocks_train/common"
	. "github.com/roca/GO/tree/staging/udemy/multithreadingingo/deadlocks_train/deadlock"
)

var (
	trains        [4]*Train
	intersections [4]*Intersection
)

const (
	screenWidth, screenHeight = 320, 320
	trainLength               = 70
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawTracks(screen)
	DrawIntersections(screen)
	DrawTrains(screen)
	ebitenutil.DebugPrint(screen, "Trains in a box")
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func main() {
	for i := 0; i < 4; i++ {
		trains[i] = &Train{Id: i, Front: 0, TrainLength: trainLength}
	}
	for i := 0; i < 4; i++ {
		intersections[i] = &Intersection{Id: i, Mutex: sync.Mutex{}, LockedBy: -1}
	}

	go MoveTrain(trains[0], 300, []*Crossing{
		{Position: 125, Intersections: intersections[0]},
		{Position: 175, Intersections: intersections[1]},
	})
	go MoveTrain(trains[1], 300, []*Crossing{
		{Position: 125, Intersections: intersections[1]},
		{Position: 175, Intersections: intersections[2]},
	})
	go MoveTrain(trains[2], 300, []*Crossing{
		{Position: 125, Intersections: intersections[2]},
		{Position: 175, Intersections: intersections[3]},
	})
	go MoveTrain(trains[3], 300, []*Crossing{
		{Position: 125, Intersections: intersections[3]},
		{Position: 175, Intersections: intersections[0]},
	})

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Trains in a box")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
