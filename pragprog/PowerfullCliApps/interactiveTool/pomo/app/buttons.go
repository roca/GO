package app

import (
	"context"

	"github.com/mum4k/termdash/widgets/button"
	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/interactiveTool/pomo/pomodoro"
)

type buttonSet struct {
	btStart *button.Button
	btPause *button.Button
}

func newButtonSet(ctx context.Context, config *pomodoro.IntervalConfig, w *widgets, redrawCh chan<-bool, errorCh chan <- error) (*buttonSet, error) {
	return &buttonSet{}, nil
}
