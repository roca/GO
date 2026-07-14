package main

import (
	"fmt"
	"strings"
)

// Reflective Visitor pattern

type Expression interface {
}

type DoubleExpression struct {
	value float64
}

type AdditionExpression struct {
	left, right Expression
}

func Print(e Expression, sb *strings.Builder) {
	switch v := e.(type) {
	case *DoubleExpression:
		sb.WriteString(fmt.Sprintf("%g", v.value))
	case *AdditionExpression:
		sb.WriteString("(")
		Print(v.left, sb)
		sb.WriteString("+")
		Print(v.right, sb)
		sb.WriteString(")")
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	// 1 + (2+3)-
	e := &AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}

	sb := strings.Builder{}
	Print(e, &sb)
	fmt.Println(sb.String())
}
