package main

import (
	"fmt"
	"strings"
)

// Double dispatch Visitor pattern

type ExpressionVisitor interface {
	VisitDoubleExpression(e *DoubleExpression)
	VisitAdditionExpression(e *AdditionExpression)
}

type Expression interface {
	Accept(ev ExpressionVisitor)
}

type DoubleExpression struct {
	value float64
}

func (d *DoubleExpression) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression(d)
}

type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression(a)
}

type ExpressionPrinter struct {
	sb strings.Builder
}

func (ep *ExpressionPrinter) VisitDoubleExpression(e *DoubleExpression) {
	ep.sb.WriteString(fmt.Sprintf("%g", e.value))
}

func (ep *ExpressionPrinter) VisitAdditionExpression(e *AdditionExpression) {
	ep.sb.WriteString("(")
	e.left.Accept(ep)
	ep.sb.WriteString("+")
	e.right.Accept(ep)
	ep.sb.WriteString(")")
}

func (ep *ExpressionPrinter) String() string {
	return ep.sb.String()
}

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{strings.Builder{}}
}

type ExpressionEvaluator struct {
	result float64
}

func (ee *ExpressionEvaluator) VisitDoubleExpression(e *DoubleExpression) {
	ee.result = e.value
}

func (ee *ExpressionEvaluator) VisitAdditionExpression(e *AdditionExpression) {
	e.left.Accept(ee)
	a := ee.result
	e.right.Accept(ee)
	b := ee.result
	ee.result = a + b
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
	ep := NewExpressionPrinter()
	e.Accept(ep)
	// fmt.Println(ep.String())

	ee := &ExpressionEvaluator{}
	e.Accept(ee)
	fmt.Printf("%s = %g\n", ep.String(), ee.result)
}
