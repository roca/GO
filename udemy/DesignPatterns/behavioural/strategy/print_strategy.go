package strategy

import "io"

type IPrintStrategy interface {
	Print() error
	SetLog(io.Writer)
	SetWriter(io.Writer)
}

type PrintOutPut struct {
	Writer    io.Writer
	LogWriter io.Writer
}

func (d *PrintOutPut) SetLog(w io.Writer) {
	d.LogWriter = w
}

func (d *PrintOutPut) SetWriter(w io.Writer) {
	d.Writer = w
}
