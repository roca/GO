package bridge

import "testing"

func TestPrintAPI1(t *testing.T) {
	api1 := PrinterImpl1{}
	err := api1.PrintMessage("Hello")
	if err != nil {
		t.Errorf("Error trying to use the API1 implementation: Message: %s\n", err.Error())
	}
}
