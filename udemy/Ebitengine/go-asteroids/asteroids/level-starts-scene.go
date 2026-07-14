package asteroids

import (
	"fmt"
	"go-asteroids/assets"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type LevelStartsScene struct {
	game           *GameScene
	nextLevelTimer *Timer
	stars          []*Star
}

func (l *LevelStartsScene) Draw(screen *ebiten.Image) {
	// Draw the stars
	for _, s := range l.stars {
		s.Draw(screen)
	}

	textToDraw := fmt.Sprintf("LEVEL %d", l.game.currentLevel)
	op := &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign: text.AlignCenter,
		},
	}
	op.ColorScale.ScaleWithColor(color.White)
	op.GeoM.Translate(ScreenWidth/2, ScreenHeight/2)
	text.Draw(screen, textToDraw, &text.GoTextFace{
		Source: assets.TiTleFont,
		Size:   48,
	}, op)
}

func (l *LevelStartsScene) Update(state *State) error {
	l.nextLevelTimer.Update()
	if l.nextLevelTimer.IsReady() {
		l.game.meteorsForLevel += 2
		l.game.meteorCount = 0
		for k, v := range l.game.lasers {
			delete(l.game.lasers, k)
			l.game.space.Remove(v.lasterObj)
		}
		state.SceneManager.GoToScene(l.game)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		l.game.meteorsForLevel += 2
		l.game.meteorCount = 0
		for k, v := range l.game.lasers {
			delete(l.game.lasers, k)
			l.game.space.Remove(v.lasterObj)
		}
		state.SceneManager.GoToScene(l.game)
	}

	return nil
}
