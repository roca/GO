package asteroids

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Star struct {
	x, y, r    float32
	brightness float32
}

func NewStart() *Star {
	return &Star{
		x:          rand.Float32() * ScreenWidth,
		y:          rand.Float32() * ScreenHeight,
		r:          rand.Float32() * (3 - 1),
		brightness: rand.Float32() * 0xff,
	}
}

func (s *Star) Draw(screen *ebiten.Image) {
	c := color.RGBA{
		R: uint8(0xbb * s.brightness / 0xff),
		G: uint8(0xdd * s.brightness / 0xff),
		B: uint8(0xff * s.brightness / 0xff),
		A: 0xff,
	}

	vector.DrawFilledCircle(screen, s.x, s.y, s.r, c, true)
}

func (s *Star) Update() {}

func GenerateStars(n int) []*Star {
	var stars []*Star
	for i := 0; i < n; i++ {
		stars = append(stars, NewStart())
	}
	return stars
}
