package asteroids

import (
	"go-asteroids/assets"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type Alien struct {
	game          *GameScene
	sprite        *ebiten.Image
	alienObj      *resolv.Circle
	position      Vector
	angle         float64
	movement      Vector
	isIntelligent bool
}

func NewAlien(baseVelocity float64, g *GameScene) *Alien {
	var alien Alien

	alienType := rand.Intn(3)

	sprite := assets.AlienSprites[rand.Intn(len(assets.AlienSprites))]

	switch alienType {
	case 0:
		// Stupid alien that comes in from the right and shoots in random directions
		x := float64(ScreenWidth + 100)
		y := float64(rand.Intn(ScreenHeight-100) + 100)

		target := Vector{X: 0, Y: y}

		pos := Vector{X: x, Y: y}

		velocity := baseVelocity + rand.Float64()*2.5

		movement := Vector{
			X: target.X - velocity,
			Y: 0,
		}

		alien := Alien{
			game:          g,
			sprite:        sprite,
			position:      pos,
			alienObj:      resolv.NewCircle(pos.X, pos.Y, float64(sprite.Bounds().Dx()/2)),
			movement:      movement,
			isIntelligent: false,
		}

		alien.alienObj.SetPosition(pos.X, pos.Y)
	case 1:
		// Stupid alien that comes in from the left and shoots in random directions
		x := -100.0
		y := float64(rand.Intn(ScreenHeight-100) + 100)

		target := Vector{X: 0, Y: y}

		pos := Vector{X: x, Y: y}

		velocity := baseVelocity + rand.Float64()*2.5

		movement := Vector{
			X: target.X + velocity,
			Y: 0,
		}

		alien := Alien{
			game:          g,
			sprite:        sprite,
			position:      pos,
			alienObj:      resolv.NewCircle(pos.X, pos.Y, float64(sprite.Bounds().Dx()/2)),
			movement:      movement,
			isIntelligent: false,
		}

		alien.alienObj.SetPosition(pos.X, pos.Y)
	case 2:
		middle := Vector{X: ScreenWidth / 2, Y: ScreenHeight / 2}

		angle := rand.Float64() * 2 * math.Pi
		r := ScreenWidth / 2.0

		pos := Vector{
			X: middle.X + r*math.Cos(angle),
			Y: middle.Y + r*math.Sin(angle),
		}

		velocity := baseVelocity + rand.Float64()*1.5
		target :=  g.player.position

		direction := Vector{
			X: target.X - pos.Y,
			Y: target.Y - pos.X,
		}

		normalizedDirection := direction.Normalize()

		movement := Vector{	
			X: normalizedDirection.X * velocity,
			Y: normalizedDirection.Y * velocity,
		}

		alien := Alien{
			game:          g,
			sprite:        sprite,
			position:      pos,
			alienObj:      resolv.NewCircle(pos.X, pos.Y, float64(sprite.Bounds().Dx()/2)),
			angle:        angle,
			movement:      movement,
			isIntelligent: true,
		}

		alien.alienObj.SetPosition(pos.X, pos.Y)
	}

	alien.alienObj.Tags().Set(TagAlien)
	return &alien
}

func (a *Alien) Update() {
	dx := a.movement.X
	dy := a.movement.Y
	
	a.position.X += dx
	a.position.Y += dy

	a.alienObj.SetPosition(a.position.X, a.position.Y)
}

func (a *Alien) Draw(screen *ebiten.Image) {
	bounds := a.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW,-halfH)
	op.GeoM.Translate(a.position.X, a.position.Y)
	screen.DrawImage(a.sprite, op)
}
