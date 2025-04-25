package asteroids

import (
	"go-asteroids/assets"

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
