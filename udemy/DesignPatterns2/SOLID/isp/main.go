package main

// ISP: Interface Segregation Principle
// ISP is a design principle that states that you should separate the interface from the implementation.

type Document struct{}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct{}

func (m *MultiFunctionPrinter) Print(d Document) {}
func (m *MultiFunctionPrinter) Fax(d Document)   {}
func (m *MultiFunctionPrinter) Scan(d Document)  {}

type OldFashionPrinter struct{}

func (o *OldFashionPrinter) Print(d Document) {}

// Deprecated: ...
func (o *OldFashionPrinter) Fax(d Document) {
	panic("operation not supported")
}

// Deprecated: ...
func (o *OldFashionPrinter) Scan(d Document) {
	panic("operation not supported")
}

type Printer interface {
	Print(d Document)
}
type Scanner interface {
	Scan(d Document)
}
type Faxer interface {
	Fax(d Document)
}

type MyPrinter struct{}

func (m *MyPrinter) Print(d Document) {}

type Photocopier struct{}

func (m *Photocopier) Print(d Document) {}
func (m *Photocopier) Fax(d Document)   {}
func (m *Photocopier) Scan(d Document)  {}

type MultiFunctionDevice interface {
	Printer
	Scanner
	Faxer
}

type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
	fax     Faxer
}

func (m *MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}
func (m *MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}
func (m *MultiFunctionMachine) Fax(d Document) {
	m.fax.Fax(d)
}

func main() {

}
