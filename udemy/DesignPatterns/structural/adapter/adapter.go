package adapter

import "fmt"

//--------------------------------------------------------------------

type LegacyPrinter interface {
	Print(s string) string
}

//--------------------------------------------------------------------

type MyLegacyPrinter struct {}

func (l *MyLegacyPrinter) Print(s string) (newMsg string){
	newMsg = fmt.Sprintf("Legacy Printer: Adapter: %s!\n", s)
	println(newMsg)
	return
}

//--------------------------------------------------------------------

type NewPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg string
}

func (p *PrinterAdapter) PrintStored() (newMsg string) {
	newMsg = fmt.Sprintf("%s!\n",p.Msg)
	if p.OldPrinter != nil {
		newMsg =  p.OldPrinter.Print(p.Msg)
	}

	return
}