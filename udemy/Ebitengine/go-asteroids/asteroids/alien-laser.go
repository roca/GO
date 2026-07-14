package asteroids

import (
	"go-asteroids/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

const (
	alienLaserSpeedPerSecond = 1000.0
)

type AlienLaser struct {
	position  Vector
	rotation  float64
	sprite    *ebiten.Image
	lasterObj *resolv.ConvexPolygon
}

func NewAlienLaser(pos Vector, rotation float64) *AlienLaser {
	sprite := assets.AlienLaserSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos.X -= halfW
	pos.Y -= halfH

	al := &AlienLaser{
		position:  pos,
		rotation:  rotation,
		sprite:    sprite,
		lasterObj: resolv.NewRectangle(pos.X, pos.Y, float64(bounds.Dx()), float64(bounds.Dy())),
	}

	al.lasterObj.SetPosition(pos.X, pos.Y)
	al.lasterObj.Tags().Set(TagLaser)

	return al
}

func (al *AlienLaser) Update() {
	speed := alienLaserSpeedPerSecond / float64(ebiten.TPS())

	dx := math.Sin(al.rotation) * speed
	dy := math.Cos(al.rotation) * -speed

	al.position.X += dx
	al.position.Y += dy

	al.lasterObj.SetPosition(al.position.X, al.position.Y)
}

func (al *AlienLaser) Draw(screen *ebiten.Image) {
	bounds := al.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(al.rotation)
	op.GeoM.Translate(al.position.X, al.position.Y)

	screen.DrawImage(al.sprite, op)
}
