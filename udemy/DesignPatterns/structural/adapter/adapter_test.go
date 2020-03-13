package adapter

import "testing"

func TestAdapter(t *testing.T) {
	msg := "Hello World"

	adapter := PrinterAdapter{OldPrinter: &MyLegacyPrinter{}, Msg: msg}
	returnedMsg := adapter.PrintStored()
	expected := "Legacy Printer: Adapter: Hello World!\n"
	if returnedMsg != expected {
		t.Errorf("Message did'nt match: %s\n", expected)
	}
}

func TestNoAdapter(t *testing.T) {
	msg := "Hello World"

	adapter := PrinterAdapter{OldPrinter: nil, Msg: msg}
	returnedMsg := adapter.PrintStored()
	expected := "Hello World!\n"
	if returnedMsg != expected {
		t.Errorf("Message did'nt match: %s\n", expected)
	}
}
