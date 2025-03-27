package asteroids

import (
	"go-asteroids/assets"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type TitleScene struct{}

func (t *TitleScene) Draw(screen *ebiten.Image) {
	// Draw the title screen here
	textToDraw := "1 coin 1 play"
	tw := withOfText(assets.TiTleFont, textToDraw)
	text.Draw(screen, textToDraw, assets.TiTleFont, (ScreenWidth-tw)/2, ScreenHeight/2, color.White)
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
