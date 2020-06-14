package visitor

import (
	"fmt"
	"io"
	"os"
)

type MessageA struct {
	Msg    string
	Output io.Writer
}

func (m *MessageA) Accept(v IVisitor) {
	v.VisitA(m)
}
func (m *MessageA) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}
	fmt.Fprintf(m.Output, "A: %s", m.Msg)
}

type MessageB struct {
	Msg    string
	Output io.Writer
}

func (m *MessageB) Accept(v IVisitor) {
	v.VisitB(m)
}
func (m *MessageB) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}
	fmt.Fprintf(m.Output, "B: %s", m.Msg)
}

type IVisitor interface {
	VisitA(*MessageA)
	VisitB(*MessageB)
}

type IVisitable interface {
	Accept(IVisitor)
}

type MessageVisitor struct{}

func (mf *MessageVisitor) VisitA(m *MessageA) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited A)")
}
func (mf *MessageVisitor) VisitB(m *MessageB) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited B)")
}

func CreateMessageA(msg string, writer io.Writer) IVisitable {
	return &MessageA{
		Msg:    msg,
		Output: writer,
	}
}
