package main

import (
	"go-asteroids/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	rotationPerSecond = math.Pi
	maxVelocity   = 8.0
)

var curVelocity float64

type Player struct {
	game           *Game
	sprite         *ebiten.Image
	rotation       float64
	position       Vector
	playerVelocity float64
}

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite

	p := &Player{
		sprite: sprite,
		game:   game,
	}

	return p
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, op)
}

func (p *Player) Update() {
	speed := rotationPerSecond / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		// fmt.Println("Left pressed")
		p.rotation -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		// fmt.Println("Right pressed")
		p.rotation += speed
	}

	p.accelerate()
}

func (p *Player) accelerate() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if curVelocity < maxVelocity {
			curVelocity = p.playerVelocity + 4
		}

		if curVelocity >= maxVelocity {
			curVelocity = maxVelocity
		}

		p.playerVelocity = curVelocity

		// Move in the direction we are pointing
		dx := math.Sin(p.rotation) * p.playerVelocity
		dy := math.Cos(p.rotation) * -p.playerVelocity

		// Move the player on the screen
		p.position.X += dx
		p.position.Y += dy
	}
}
