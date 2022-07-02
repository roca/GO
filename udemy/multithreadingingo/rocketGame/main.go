package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Our game constants
const (
	screenWidth, screenHeight = 640, 480
)

// Create our empty vars
var (
	err        error
	background *ebiten.Image
	spaceShip  *ebiten.Image
	playerOne  player
)

// Create the player class
type player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

// Run this code once at startup
func init() {
	background, _, err = ebitenutil.NewImageFromFile("assets/space.png")
	if err != nil {
		log.Fatal(err)
	}

	spaceShip, _, err = ebitenutil.NewImageFromFile("assets/rocket.png")
	if err != nil {
		log.Fatal(err)
	}

	playerOne = player{spaceShip, screenWidth / 2.0, screenHeight / 2.0, 4}
}

// Move the player depending on which key is pressed
func movePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		playerOne.yPos -= playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		playerOne.yPos += playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		playerOne.xPos -= playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		playerOne.xPos += playerOne.speed
	}
}

type Game struct{}

func (g *Game) Update() error {
	movePlayer()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(background, op)

	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(playerOne.xPos, playerOne.yPos)
	screen.DrawImage(playerOne.image, playerOp)
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Rocket in space")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
