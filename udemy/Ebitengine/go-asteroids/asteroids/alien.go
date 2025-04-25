package asteroids

import (
	"go-asteroids/assets"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type Alien struct {
	game          *GameScene
	sprite        *ebiten.Image
	alienIbject   *resolv.Circle
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
	}
}
