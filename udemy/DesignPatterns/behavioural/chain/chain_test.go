package chain

import (
	"fmt"
	"strings"
	"testing"
)

type myTestWriter struct {
	receiveMessage *string
}

func (m *myTestWriter) Write(p []byte) (int, error) {
	if m.receiveMessage == nil {
		m.receiveMessage = new(string)
	}
	tempMessage := fmt.Sprintf("%s%s", string(*m.receiveMessage), string(p))
	m.receiveMessage = &tempMessage
	return len(p), nil
}

func (m *myTestWriter) Next(s string) {
	m.Write([]byte(s))
}

func TestCreateDefaultChain(t *testing.T) {
	myWriter := myTestWriter{}
	writerLogger := WriterLogger{Writer: &myWriter}
	second := SecondLogger{NextChain: &writerLogger}
	chain := FirstLogger{NextChain: &second}
	t.Run("3 loggers 2 of them writes to console, second only if it finds"+
		" the word 'Hello', third writes to some variable if second found 'Hello'",
		func(t *testing.T) {
			chain.Next("message that breaks the chain")

			if myWriter.receiveMessage != nil {
				t.Error("Last link should not receive any message")
			}

			chain.Next("Hello")

			if myWriter.receiveMessage == nil || !strings.Contains(*myWriter.receiveMessage, "Hello") {
				t.Fatal("Last link didn't receive expected message")
			}
		})

	t.Run("2 loggers, second uses the closure implementation", func(t *testing.T) {
		myWriter = myTestWriter{}
		closureLogger := ClosureChain{
			Closure: func(s string) {
				fmt.Printf("My closure logger! Message: %s\n", s)
				myWriter.receiveMessage = &s
			},
		}
		writerLogger.NextChain = &closureLogger

		chain.Next("Hello closure logger")

		if *myWriter.receiveMessage != "Hello closure logger" {
			t.Fatal("Expected message wasn't received in myWriter")
		}
	})
}
