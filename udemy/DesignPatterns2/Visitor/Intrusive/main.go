package main

import (
	"fmt"
	"strings"
)

// Intrusive Visitor pattern

type Expression interface {
	Print(sb *strings.Builder)
}

type DoubleExpression struct {
	value float64
}

func (d *DoubleExpression) Print(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("%g", d.value))
}

type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Print(sb *strings.Builder) {
	sb.WriteString("(")
	if a.left != nil {
		a.left.Print(sb)
	}
	sb.WriteString("+")
	if a.right != nil {
		a.right.Print(sb)
	}
	sb.WriteString(")")
}

func main() {
	// 1 + (2+3)
	e := &AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}

	sb := strings.Builder{}
	e.Print(&sb)
	fmt.Println(sb.String())
}
