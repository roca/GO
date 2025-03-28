package asteroids

import (
	"go-asteroids/assets"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
)

type TitleScene struct{}

func (t *TitleScene) Draw(screen *ebiten.Image) {
	// Draw the title screen here
	textToDraw := "1 coin 1 play"

	op := &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign: text.AlignCenter,
		},
	}
	op.ColorScale.ScaleWithColor(color.White)
	op.GeoM.Translate(float64(ScreenWidth)/2, ScreenHeight-200)

	text.Draw(screen, textToDraw, &text.GoTextFace{
		Source: assets.TiTleFont,
		Size:   48,
	}, op)
}

func (t *TitleScene) Update(state *State) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.SceneManager.GoToScene(NewGameScene())
	}
	return nil
}

func withOfText(f font.Face, t string) int {
	_, textWidth := font.BoundString(f, t)
	return textWidth.Round()
}
