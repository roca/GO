package bridge

import (
	"errors"
	"fmt"
	"io"
)

type PrintAPI interface {
	PrintMessage(string) error
}

type PrinterImpl1 struct{}

func (p *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Printf("%s\n",msg)
	return nil
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (d *PrinterImpl2) PrintMessage(msg string) error {
	if d.Writer == nil {
		return errors.New("You need to pass an io.Writer to PrinterImpl2")
	}
	fmt.Fprintf(d.Writer, "%s", msg)
	return nil
}

type PrinterAbstraction interface {
	Print() error
}
type NormalPrinter struct {
	Msg     string
	Printer PrintAPI
}

func (c *NormalPrinter) Print() error {
	c.Printer.PrintMessage(c.Msg)
	return nil
}

type PacktPrinter struct {
	Msg     string
	Printer PrintAPI
}

func (c *PacktPrinter) Print() error {
	c.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s",c.Msg))
	return nil
}
