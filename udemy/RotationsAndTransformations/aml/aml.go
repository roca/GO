package aml

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)
type Vector3 struct {
     X float64
     Y float64
     Z float64
}

func (v *Vector3) Plus(v1 *Vector3) (*Vector3) {
    v.X += v1.X
    v.Y += v1.Y
    v.Z += v1.Z
    return v
}

func MatPrint(X mat.Matrix) string {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	formatedString := fmt.Sprintf("%v\n", fa)

	print(formatedString)
	return formatedString
}
