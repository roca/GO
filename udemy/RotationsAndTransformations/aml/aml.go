package aml

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func MatPrint(X mat.Matrix) string {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	formatedString := fmt.Sprintf("%v\n", fa)

	print(formatedString)
	return formatedString
}
