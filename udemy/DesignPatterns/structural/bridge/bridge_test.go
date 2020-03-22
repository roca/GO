package bridge

import (
	"errors"
	"strings"
	"testing"
)

func TestPrintAPI1(t *testing.T) {
	api1 := PrinterImpl1{}
	err := api1.PrintMessage("Hello")
	if err != nil {
		t.Errorf("Error trying to use the API1 implementation: Message: %s\n", err.Error())
	}
}

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
	}
	err = errors.New("Content received on Writer was empty")
	return
}

func TestPrintAPI2(t *testing.T) {
	api2 := PrinterImpl2{}
	err := api2.PrintMessage("Hello")
	if err != nil {
		expectedErrorMessage := "You need to pass an io.Writer to PrinterImpl2"
		if !strings.Contains(err.Error(), expectedErrorMessage) {
			t.Errorf("Error message was not correct.\n Actual %s\n Expected:: %s\n", err.Error(), expectedErrorMessage)
		}
	}
	testWriter := TestWriter{}
	api2 = PrinterImpl2{
		Writer: &testWriter,
	}
	expectedMessage := "Hello"
	err = api2.PrintMessage(expectedMessage)
	if err != nil {
		t.Errorf("Error trying to use the API2 implementation: %s\n", err.Error())
	}
	if testWriter.Msg != expectedMessage {
		t.Fatalf("API2 did not write correctly on the io.Writer. \n Actual: %s\n Expected: %s\n", testWriter.Msg, expectedMessage)
	}
}

func TestNormalPrinter_Print(t *testing.T) {
	expectedMessage := "Hello io.Writer"
	normal := NormalPrinter{
		Msg:     expectedMessage,
		Printer: &PrinterImpl1{},
	}
	err := normal.Print()
	if err != nil {
		t.Error(err.Error())
	}
	testWriter := TestWriter{}
	normal = NormalPrinter{
		Msg: expectedMessage,
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}
	err = normal.Print()
	if err != nil {
		t.Error(err.Error())
	}
	if testWriter.Msg != expectedMessage {
		t.Errorf("The expected message on the io.Writer doesn't match actual.\n Actual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}

func TestPacktPrinter_Print(t *testing.T) {
	passedMessage := "Hello io.Writer"
	expectedMessage := "Message from Packt: Hello io.Writer"
	packt := PacktPrinter{
		Msg:     passedMessage,
		Printer: &PrinterImpl1{},
	}
	err := packt.Print()
	if err != nil {
		t.Error(err.Error())
	}
	testWriter := TestWriter{}
	packt = PacktPrinter{
		Msg: passedMessage,
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}
	err = packt.Print()
	if err != nil {
		t.Error(err.Error())
	}
	if testWriter.Msg != expectedMessage {
		t.Errorf("The expected message on the io.Writer doesn't match actual.\n Actual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}
