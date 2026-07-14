package asteroids

import (
	"go-asteroids/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

type LifeIndicator struct {
	position Vector
	rotation float64
	sprite   *ebiten.Image
}

func NewLifeIndicator(position Vector) *LifeIndicator {
	sprite := assets.LifeIndicator
	return &LifeIndicator{
		position: position,
		sprite:   sprite,
	}
}
func (l *LifeIndicator) Update() {}

func (l *LifeIndicator) Draw(screen *ebiten.Image) {
	bounds := l.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &colorm.DrawImageOptions{}
	op.GeoM.Translate(halfW, halfH)
	cm := colorm.ColorM{}
	cm.Scale(1.0, 1.0, 1.0, 0.2)

	op.GeoM.Translate(l.position.X, l.position.Y)

	colorm.DrawImage(screen, l.sprite, cm, op)
}
