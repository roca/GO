package arithmetic

import "errors"

func Sum(args ...int) (res int) {
	for _, v := range args {
		res += v
	}
	return
}

func Subtract(args ...int) int {
	if len(args) < 2 {
		return 0
	}
	res := args[0]
	for i := 1; i < len(args); i++ {
		res -= args[i]
	}
	return res
}

func Multiply(args ...int) int {
	if len(args) < 2 {
		return 0
	}
	res := args[0]
	for i := 1; i < len(args); i++ {
		res *= args[i]
	}
	return res
}

func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("You cannot divide by zero")
	}
	return float64(a) / float64(b), nil
}
