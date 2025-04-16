package asteroids

import (
	"go-asteroids/assets"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TitleScene struct {
	meteors     map[int]*Meteor
	meteorCount int
	stars       []*Star
}

func (t *TitleScene) Draw(screen *ebiten.Image) {
	for _, s := range t.stars {
		s.Draw(screen)
	}
	
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

	for _, m := range t.meteors {
		m.Draw(screen)
	}

	
}

func (t *TitleScene) Update(state *State) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.SceneManager.GoToScene(NewGameScene())
	}

	if len(t.meteors) < 10 {
		m := NewMeteor(0.25, &GameScene{}, len(t.meteors)-1)
		t.meteorCount++
		t.meteors[t.meteorCount] = m
	}

	for _, m := range t.meteors {
		m.Update()
	}

	return nil
}
