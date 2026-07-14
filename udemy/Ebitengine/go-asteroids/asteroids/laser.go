package asteroids

import (
	"go-asteroids/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

const (
	laserSpeedPerSecond = 1000.0
)

type Laser struct {
	game      *GameScene
	rotation  float64
	position  Vector
	sprite    *ebiten.Image
	lasterObj *resolv.ConvexPolygon
}

func NewLaser(pos Vector, rotation float64, index int, g *GameScene) *Laser {
	// Set the sprite
	sprite := assets.LaserSprite

	// Position X & Y coordinates from the center of the sprite.
	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos.X -= halfW
	pos.Y -= halfH

	// Create a laser object for collision detection
	l := &Laser{
		game:      g,
		position:  pos,
		rotation:  rotation,
		sprite:    sprite,
		lasterObj: resolv.NewRectangle(pos.X, pos.Y, float64(bounds.Dx()), float64(bounds.Dy())),
	}

	// Set the position of the collision object.
	l.lasterObj.SetPosition(pos.X, pos.Y)
	l.lasterObj.SetData(&ObjectData{index: index})
	l.lasterObj.Tags().Set(TagLaser)

	return l
}

func (l *Laser) Update() {
	// How fats should the laser move?
	speed := laserSpeedPerSecond / float64(ebiten.TPS())

	dx := math.Sin(l.rotation) * speed
	dy := math.Cos(l.rotation) * -speed

	// Update the position of the laser
	l.position.X += dx
	l.position.Y += dy

	l.lasterObj.SetPosition(l.position.X, l.position.Y)
}

func (l *Laser) Draw(screen *ebiten.Image) {
	bounds := l.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(l.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(l.position.X, l.position.Y)

	screen.DrawImage(l.sprite, op)
}
