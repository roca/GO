package app

import (
	"context"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/donut"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
	"github.com/mum4k/termdash/widgets/text"
)

type widgets struct {
	donTimer       *donut.Donut
	disType        *segmentdisplay.SegmentDisplay
	textInfo       *text.Text
	textTimer      *text.Text
	updateDonTimer chan []int
	updateTxtInfo  chan string
	updateTxtTimer chan string
	updateTxtType  chan string
}

func (w *widgets) update(timer []int, txtType, txtInfo, txtTimer string, redrawCh chan<- bool) {
	if txtInfo != "" {
		w.updateTxtInfo <- txtInfo
	}
	if txtType != "" {
		w.updateTxtType <- txtType
	}
	if txtTimer != "" {
		w.updateTxtTimer <- txtTimer
	}
	if len(timer) > 0 {
		w.updateDonTimer <- timer
	}
	redrawCh <- true
}

func newWidgets(ctx context.Context, errorCh chan<- error) (*widgets, error) {
	w := &widgets{}
	var err error

	w.updateDonTimer = make(chan []int)
	w.updateTxtType = make(chan string)
	w.updateTxtInfo = make(chan string)
	w.updateTxtTimer = make(chan string)

	w.donTimer, err = newDonut(ctx, w.updateDonTimer, errorCh)
	if err != nil {
		return nil, err
	}

	w.disType, err = newSegmentDisplay(ctx, w.updateTxtType, errorCh)
	if err != nil {
		return nil, err
	}

	w.textInfo, err = newText(ctx, w.updateTxtInfo, errorCh)
	if err != nil {
		return nil, err
	}

	w.textTimer, err = newText(ctx, w.updateTxtTimer, errorCh)
	if err != nil {
		return nil, err
	}

	return w, nil
}

func newText(ctx context.Context, updateText <-chan string, errorCh chan<- error) (*text.Text, error) {

	txt, err := text.New()
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case t := <-updateText:
				txt.Reset()
				errorCh <- txt.Write(t)
			case <-ctx.Done():
				return
			}
		}
	}()

	return txt, nil
}

func newDonut(ctx context.Context, donutUpdater <-chan []int, errorCh chan<- error) (*donut.Donut, error) {

	don, err := donut.New(
		donut.Clockwise(),
		donut.CellOpts(cell.FgColor(cell.ColorBlue)),
	)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case d := <-donutUpdater:
				if d[0] <= d[1] {
					errorCh <- don.Absolute(d[0], d[1])
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return don, nil
}

func newSegmentDisplay(ctx context.Context, updateText <-chan string, errorCh chan<- error) (*segmentdisplay.SegmentDisplay, error) {

	sd, err := segmentdisplay.New()
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case t := <-updateText:
				if t == "" {
					t = " "
				}
				errorCh <- sd.Write([]*segmentdisplay.TextChunk{
					segmentdisplay.NewChunk(t),
				})
			case <-ctx.Done():
				return
			}
		}
	}()

	return sd, nil
}
