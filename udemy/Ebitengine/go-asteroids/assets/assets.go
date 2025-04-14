package assets

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	_ "image/png"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed *
var assets embed.FS

var PlayerSprite = mustLoadImage("images/player.png")
var TiTleFont = mustLoadFontFace("fonts/title.ttf")
var MeteorSprites = mustLoadImages("images/meteors/*.png")
var MeteorSpritesSmall = mustLoadImages("images/meteors-small/*.png")
var LaserSprite = mustLoadImage("images/laser.png")
var ExplosionSprite = mustLoadImage("images/explosion.png")
var ExplosionSmallSprite = mustLoadImage("images/explosion-small.png")
var Explosion = createExplosion()
var ExhaustSprite = mustLoadImage("images/fire.png")

var ThrustSound = mustLoadOggVorbis("thrust")
var LaserOneSound = mustLoadOggVorbis("fire")
var LaserTwoSound = mustLoadOggVorbis("fire")
var LaserThreeSound = mustLoadOggVorbis("fire")

func mustLoadOggVorbis(name string) *vorbis.Stream {
	f, err := assets.ReadFile("audio/" + name + ".ogg")
	if err != nil {
		panic(err)
	}

	stream, err := vorbis.DecodeWithoutResampling(bytes.NewReader(f))
	if err != nil {
		panic(err)
	}
	return stream
}

func createExplosion() []*ebiten.Image {
	var frames []*ebiten.Image
	for i := 0; i < 12; i++ {
		frame := mustLoadImage(fmt.Sprintf("images/explosion/%d.png", i+1))
		frames = append(frames, frame)
	}
	return frames
}

func mustLoadImages(pattern string) []*ebiten.Image {
	matches, err := fs.Glob(assets, pattern)
	if err != nil {
		panic(err)
	}

	images := make([]*ebiten.Image, len(matches))
	for i, match := range matches {
		images[i] = mustLoadImage(match)
	}

	return images
}

func mustLoadFontFace(name string) *text.GoTextFaceSource {
	f, err := assets.ReadFile(name)
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(f)

	ts, err := text.NewGoTextFaceSource(r)
	if err != nil {
		panic(err)
	}

	return ts
}

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
