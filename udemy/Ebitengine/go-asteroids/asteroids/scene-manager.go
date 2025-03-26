package asteroids

import "github.com/hajimehoshi/ebiten/v2"

var (
	transitionFrom = ebiten.NewImage(ScreenWidth, ScreenHeight)
	transitionTo   = ebiten.NewImage(ScreenWidth, ScreenHeight)
)

const transitionMaxCount = 25

type Scene interface {
	Update(state *State) error
	Draw(screen *ebiten.Image)
}

type State struct {
	SceneManager *SceneManager
	Input        *Input
}

type SceneManager struct {
	current         Scene
	next            Scene
	transitionCount int
}

func (s *SceneManager) Draw(r *ebiten.Image) {
	if s.transitionCount == 0 {
		s.current.Draw(r)
		return
	}

	transitionFrom.Clear()
	s.current.Draw(transitionFrom)

	transitionTo.Clear()
	s.next.Draw(transitionTo)

	r.DrawImage(transitionFrom, nil)

	alpha := 1 - float32(s.transitionCount)/float32(transitionMaxCount)
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(alpha)

	r.DrawImage(transitionTo, op)
}

func (s *SceneManager) Update(_ *Input) error {
	if s.transitionCount == 0 {
		return s.current.Update(&State{
			SceneManager: s,
		})
	}

	s.transitionCount--
	if s.transitionCount > 0 {
		return nil
	}

	s.current = s.next
	s.next = nil
	return nil
}

func (s *SceneManager) GoToScene(scene Scene) {
	if s.current == nil {
		s.current = scene
	} else {
		s.next = scene
		s.transitionCount = transitionMaxCount
	}
}