package asteroids

import "github.com/hajimehoshi/ebiten/v2"

const (
	rotationSpeedMin                    = -0.02
	rotationSpeedMax                    = 0.02
	numberOfSmallMeteorsFromLargeMeteor = 4
)

type Meteor struct {
	game          *GameScene
	position      Vector
	rotation      float64
	movement      Vector
	angle         float64
	rotationSpeed float64
	sprite        *ebiten.Image
}

func NewMeteor(baseVelocity float64, game *GameScene, index int) *Meteor {
	// Target the center of the screen

	// Pick a random angle

	// The distance from the center that meteor should spawn at. Half the width, and some arbitrary distance

	// Create the position vector, using the angle and simple math

	// Keep the meteor moving toward the center of the screen

	// Assign a sprite to the meteor

	// Create the meteor object and return it

	return nil
}
