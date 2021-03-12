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
		case 3:
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

// Special Type creation
func Identity() Matrix {
	m,_ := New([][]float64{
		{1., 0., 0.},
		{0., 1., 0.},
		{0., 0., 1.},
	})
	return m
}

func (m *Matrix) Negate() (Matrix,error) {
	dataM := m.Data()
	for i, row := range dataM {
			for j, _ := range row {
				for k := 0; k < 3; k++ {
					dataM[i][j] *= -1.0
				}
			}
		}
		u,_ := New(dataM)
	return u, nil
}

func (m *Matrix) change(values ...interface{}) error {
	switch values[0].(type) {
	case float64:
		switch l := len(values); l {
		case 1:
			m.M11 = values[0].(float64)
			m.M12 = values[0].(float64)
			m.M13 = values[0].(float64)
			m.M21 = values[0].(float64)
			m.M22 = values[0].(float64)
			m.M23 = values[0].(float64)
			m.M31 = values[0].(float64)
			m.M32 = values[0].(float64)
			m.M33 = values[0].(float64)
			return nil
		case 9:
			m.M11 = values[0].(float64)
			m.M12 = values[1].(float64)
			m.M13 = values[2].(float64)
			m.M21 = values[3].(float64)
			m.M22 = values[4].(float64)
			m.M23 = values[5].(float64)
			m.M31 = values[6].(float64)
			m.M32 = values[7].(float64)
			m.M33 = values[8].(float64)
			return nil
		default:
			return fmt.Errorf("Could not create Vector type")
		}
	case []float64:
		a := values[0].([]float64)
		switch l := len(a); l {
		case 9:
			m.M11 = a[0]
			m.M12 = a[1]
			m.M13 = a[2]
			m.M21 = a[3]
			m.M22 = a[4]
			m.M23 = a[5]
			m.M31 = a[6]
			m.M32 = a[7]
			m.M33 = a[8]
			return nil
		default:
			return fmt.Errorf("Could not create Matrixtype")
		}
	case [][]float64:
		a := values[0].([][]float64)
		switch l := len(a[0]) + len(a[1]) + len(a[2]); l {
		case 9:
			m.M11 = a[0][0]
			m.M12 = a[0][1]
			m.M13 = a[0][2]
			m.M21 = a[1][0]
			m.M22 = a[1][1]
			m.M23 = a[1][2]
			m.M31 = a[2][0]
			m.M32 = a[2][1]
			m.M33 = a[2][2]
			return nil
		default:
			return fmt.Errorf("Could not create Matrixtype")
		}
	case vector.Vector:
		switch l := len(values); l {
		case 3:
			u1 := values[0].(vector.Vector)
			u2 := values[1].(vector.Vector)
			u3 := values[2].(vector.Vector)
			m.M11 = u1.X
			m.M12 = u1.Y
			m.M13 = u1.Z
			m.M21 = u2.X
			m.M22 = u2.Y
			m.M23 = u2.Z
			m.M31 = u3.X
			m.M32 = u3.Y
			m.M33 = u3.Z
			return nil
		default:
			return fmt.Errorf("Could not create Vector type")
		}
	case []vector.Vector:
		a := values[0].([]vector.Vector)
		switch l := len(a); l {
		case 3:
			m.M11 = a[0].X
			m.M12 = a[0].Y
			m.M13 = a[0].Z
			m.M21 = a[1].X
			m.M22 = a[1].Y
			m.M23 = a[1].Z
			m.M31 = a[2].X
			m.M32 = a[2].Y
			m.M33 = a[2].Z
			return nil
		default:
			return fmt.Errorf("Could not create Matrix type")
		}
	default:
		return nil
	}
}

func (m *Matrix) Data() [][]float64 {
	return [][]float64{
		{m.M11, m.M12, m.M13},
		{m.M21, m.M22, m.M23},
		{m.M31, m.M32, m.M33},
	}
}

// Operator Assignments (Matrix)
// +=, -=, *=, /=
func (m *Matrix) Mop(operation string, u Matrix) (*Matrix, error) {
	switch o := operation; {
	case o == "+=" || o == "+":
		m.M11 += u.M11
		m.M12 += u.M12
		m.M13 += u.M13
		m.M21 += u.M21
		m.M22 += u.M22
		m.M23 += u.M23
		m.M31 += u.M31
		m.M32 += u.M32
		m.M33 += u.M33
		return m, nil
	case o == "-=" || o == "-":
		m.M11 -= u.M11
		m.M12 -= u.M12
		m.M13 -= u.M13
		m.M21 -= u.M21
		m.M22 -= u.M22
		m.M23 -= u.M23
		m.M31 -= u.M31
		m.M32 -= u.M32
		m.M33 -= u.M33
		return m, nil
	case o == "*=" || o == "*":
		dataM := m.Data()
		dataU := u.Data()
		for i, row := range dataM {
			for j, _ := range row {
				for k := 0; k < 3; k++ {
					dataM[i][j] = dataM[i][k] * dataU[k][i]
				}
			}
		}
		m.change(dataM)
		return m, nil
	case o == "/=" || o == "/":
		inverse, _ := u.Inverse()
		m.Mop("/=", inverse)
		return m, nil
	default:
		return &Matrix{}, fmt.Errorf("Matrix has no such operation '%s'", o)
	}
}

// Operator Assignments (Scalar)
// +=, -=, *=, /=
func (m *Matrix) Sop(operation string, value float64) (*Matrix, error) {
	dataM := m.Data()
	switch o := operation; {
	case o == "+=" || o == "+":
		for i, row := range dataM {
			for j, _ := range row {
				for k := 0; k < 3; k++ {
					dataM[i][j] += value
				}
			}
		}
		m.change(dataM)
		return m, nil
	case o == "-=" || o == "-":
		for i, row := range dataM {
			for j, _ := range row {
				for k := 0; k < 3; k++ {
					dataM[i][j] -= value
				}
			}
		}
		m.change(dataM)
		return m, nil
	case o == "*=" || o == "*":
		for i, row := range dataM {
			for j, _ := range row {
				for k := 0; k < 3; k++ {
					dataM[i][j] *= value
				}
			}
		}
		m.change(dataM)
		return m, nil
	case o == "/=" || o == "/":
		for i, row := range dataM {
			for j, _ := range row {
				for k := 0; k < 3; k++ {
					dataM[i][j] /= value
				}
			}
		}
		m.change(dataM)
		return m, nil
	default:
		return &Matrix{}, fmt.Errorf("Vector has no such operation '%s'", o)
	}
}



// Matrix / Vector Operations
// M * V = V
// Matrix / Scalar Operations
// M * S, M + S, M - S , M / S
// S * M, S + M, S - M , S / M
// Matrix Operations

func (m *Matrix) DiagV(n Matrix) (vector.Vector, error) { return vector.Vector{}, nil }
func (m *Matrix) DiagM(v vector.Vector) (Matrix, error) { return Matrix{}, nil }
func (m *Matrix) Transpose() (Matrix, error)            { return Matrix{}, nil }
func (m *Matrix) Inverse() (Matrix, error)              { return Matrix{}, nil }
func (m *Matrix) Determinant() (float64, error)         { return 0.0, nil }
