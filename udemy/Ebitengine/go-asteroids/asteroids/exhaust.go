package asteroids

import (
	"go-asteroids/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const exhaustSpawnOffset = -50.0

type Exhaust struct {
	position Vector
	rotation float64
	sprite   *ebiten.Image
}

func NewExhaust(pos Vector, rotation float64) *Exhaust {
	sprite := assets.ExhaustSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos.X -= halfW
	pos.Y -= halfH

	return &Exhaust{
		position: pos,
		rotation: rotation,
		sprite:   sprite,
	}
}

func (e *Exhaust) Draw(screen *ebiten.Image) {
	 bounds := e.sprite.Bounds()
	 halfW := float64(bounds.Dx()) / 2
	 halfH := float64(bounds.Dy()) / 2

	 op := &ebiten.DrawImageOptions{}
	 op.GeoM.Translate(-halfW, -halfH)
	 op.GeoM.Rotate(e.rotation)
	 op.GeoM.Translate(halfW, halfH)
	 op.GeoM.Translate(e.position.X, e.position.Y)

	screen.DrawImage(e.sprite, op)
}

func (e *Exhaust) Update() {
	speed := maxVelocity / float64(ebiten.TPS())
	e.position.X += math.Sin(e.rotation) * speed
	e.position.Y += math.Cos(e.rotation) * speed
}	
