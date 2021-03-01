package matrix

import (
	"fmt"

	"udemy.com/aml/vector"
)

type Matrix struct {
	M11, M12, M13 float64
	M21, M22, M23 float64
	M31, M32, M33 float64
}

func New(values ...interface{}) (Matrix, error) {
	switch values[0].(type) {
	case float64:
		switch l := len(values); l {
		case 1:
			return Matrix{
				values[0].(float64),
				values[0].(float64),
				values[0].(float64),
				values[0].(float64),
				values[0].(float64),
				values[0].(float64),
				values[0].(float64),
				values[0].(float64),
				values[0].(float64),
			}, nil
		case 9:
			return Matrix{
				values[0].(float64), values[1].(float64), values[2].(float64),
				values[3].(float64), values[4].(float64), values[5].(float64),
				values[6].(float64), values[7].(float64), values[8].(float64),
			}, nil
		default:
			return Matrix{}, fmt.Errorf("Could not create Vector type")
		}
	case []float64:
		a := values[0].([]float64)
		switch l := len(a); l {
		case 9:
			return Matrix{
				a[0], a[1], a[2],
				a[3], a[4], a[5],
				a[6], a[7], a[8],
			}, nil
		default:
			return Matrix{}, fmt.Errorf("Could not create Matrixtype")
		}
	case [][]float64:
		a := values[0].([][]float64)
		switch l := len(a[0]) + len(a[1]) + len(a[2]); l {
		case 9:
			return Matrix{
				a[0][0], a[0][1], a[0][2],
				a[1][0], a[1][1], a[1][2],
				a[2][0], a[2][1], a[2][2],
			}, nil
		default:
			return Matrix{}, fmt.Errorf("Could not create Matrixtype")
		}
	case vector.Vector:
		switch l := len(values); l {
		case 3:
			u1 := values[0].(vector.Vector)
			u2 := values[1].(vector.Vector)
			u3 := values[2].(vector.Vector)
			return Matrix{
				u1.X, u1.Y, u1.Z,
				u2.X, u2.Y, u2.Z,
				u3.X, u3.Y, u3.Z,
			}, nil
		default:
			return Matrix{}, fmt.Errorf("Could not create Vector type")
		}
	case []vector.Vector:
		a := values[0].([]vector.Vector)
		switch l := len(a); l {
		case 1:
			return Matrix{
				a[0].X, a[0].Y, a[0].Z,
				a[1].X, a[1].Y, a[1].Z,
				a[2].X, a[2].Y, a[2].Z,
			}, nil
		default:
			return Matrix{}, fmt.Errorf("Could not create Matrix type")
		}
	default:
		return Matrix{}, nil
	}
}
func (m *Matrix) Data() []float64 {
	return []float64{
		m.M11, m.M12, m.M13,
		m.M21, m.M22, m.M23,
		m.M31, m.M32, m.M33,
	}
}

// Operator Assignments (Matrix)
// Operator Assignments (Scalar)
// Special Type creation

// Matrix / Matrix Operations
// Matrix / Vector Operations
// Matrix / Scalar Operations
// Matrix Operations
