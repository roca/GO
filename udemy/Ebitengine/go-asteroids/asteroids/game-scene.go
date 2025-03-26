package asteroids

import "github.com/hajimehoshi/ebiten/v2"

type GameScene struct {
	player *Player
}

func NewGameScene() *GameScene {
	g := &GameScene{}
	g.player = NewPlayer(g)

	return g
}

func (g *GameScene) Update(state *State) error {
	g.player.Update()
	return nil
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *GameScene) Layout(outsideWidth, outsideHeight int) (ScreenWidth, ScreenHeight int) {
	return outsideWidth, outsideHeight
}