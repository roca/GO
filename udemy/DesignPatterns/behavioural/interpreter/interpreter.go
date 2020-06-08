package interpreter

import (
	"strconv"
	"strings"
)

const (
	SUM = "sum"
	SUB = "sub"
)

//-------------------------------------------------------------------------------
type Interpreter interface {
	Read() int
}

type value int

func (v *value) Read() int {
	return int(*v)
}

type operationSum struct {
	Left  Interpreter
	Right Interpreter
}

func (a *operationSum) Read() int {
	return a.Left.Read() + a.Right.Read()
}

type operationSubtract struct {
	Left  Interpreter
	Right Interpreter
}

func (a *operationSubtract) Read() int {
	return a.Left.Read() - a.Right.Read()
}

func operatorFactory(o string, left, right Interpreter) Interpreter {
	switch o {
	case SUM:
		return &operationSum{
			Left:  left,
			Right: right,
		}
	case SUB:
		return &operationSubtract{
			Left:  left,
			Right: right,
		}
	}
	return nil
}

type polishNotationStack []Interpreter

func (p *polishNotationStack) Push(s Interpreter) {
	*p = append(*p, s)
}
func (p *polishNotationStack) Pop() Interpreter {
	length := len(*p)
	if length > 0 {
		temp := (*p)[length-1]
		*p = (*p)[:length-1]
		return temp
	}
	return nil
}

func Calculate(o string) (int, error) {
	stack := polishNotationStack{}
	operators := strings.Split(o, " ")
	for _, v := range operators {
		if v == SUM || v == SUB {
			right := stack.Pop()
			left := stack.Pop()
			mathFunc := operatorFactory(v, left, right)
			res := value(mathFunc.Read())
			stack.Push(&res)
		} else {
			val, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			temp := value(val)
			stack.Push(&temp)
		}

	}
	return int(stack.Pop().Read()), nil
}
