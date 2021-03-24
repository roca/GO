package matrix

import (
	"fmt"
	"math"

	"udemy.com/aml/vector"
)

type IMatrix interface {
	New(values ...interface{}) (Matrix, error)
	Negate() (Matrix, error)
	Copy() (Matrix, error)
	change(values ...interface{}) error
	Data() [][]float64
	Mop(operation string, u Matrix) (*Matrix, error)
	Sop(operation string, value float64) (*Matrix, error)
	Inverse() (Matrix, error)
}

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
	m, _ := New([][]float64{
		{1., 0., 0.},
		{0., 1., 0.},
		{0., 0., 1.},
	})
	return m
}

func (m *Matrix) Negate() (Matrix, error) {
	u, _ := m.Copy()
	u.Sop("*=", -1.0)
	return u, nil
}

func (m *Matrix) Copy() (Matrix, error) {
	dataM := m.Data()
	u, _ := New(dataM)
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
		z := Matrix{}
		data := z.Data()

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				for k := 0; k < 3; k++ {
					data[i][j] += dataM[i][k] * dataU[k][j]
				}
			}
		}
		_ = m.change(data)
		return m, nil
	case o == "/=" || o == "/":
		inverse, _ := u.Inverse()
		_, _ = m.Mop("*=", inverse)
		return m, nil
	default:
		return &Matrix{}, fmt.Errorf("Matrix has no such operation '%s'", o)
	}
}

func (m *Matrix) Vop(operation string, v vector.Vector) (vector.Vector, error) {
	switch o := operation; {
	case o == "*":
		u, _ := vector.New([]float64{
			m.M11*v.X + m.M12*v.Y + m.M13*v.Z,
			m.M21*v.X + m.M22*v.Y + m.M23*v.Z,
			m.M31*v.X + m.M32*v.Y + m.M33*v.Z,
		})
		return u, nil
	default:
		return vector.Vector{}, fmt.Errorf("Matrix has no such operation '%s'", o)
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
				dataM[i][j] += value
			}
		}
		m.change(dataM)
		return m, nil
	case o == "-=" || o == "-":
		for i, row := range dataM {
			for j, _ := range row {
				dataM[i][j] -= value
			}
		}
		m.change(dataM)
		return m, nil
	case o == "*=" || o == "*":
		for i, row := range dataM {
			for j, _ := range row {
				dataM[i][j] *= value
			}
		}
		m.change(dataM)
		return m, nil
	case o == "/=" || o == "/":
		for i, row := range dataM {
			for j, _ := range row {
				dataM[i][j] /= value
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

func Transpose(m Matrix) (Matrix, error) {
	dataM := m.Data()
	u, _ := m.Copy()
	dataU := u.Data()
	for i, row := range dataM {
		for j, _ := range row {
			for k := 0; k < 3; k++ {
				dataU[i][j] = dataM[j][i]
			}
		}
	}
	u.change(dataU)
	return u, nil
}
func Determinant(m Matrix) (float64, error) {
	// det1 := m.M11 * ((m.M22 * m.M33) - (m.M32 * m.M23))
	// det2 := m.M12 * ((m.M21 * m.M33) - (m.M23 * m.M31))
	// det3 := m.M13 * ((m.M21 * m.M32) - (m.M22 * m.M31))
	// det := det1 - det2 + det3
	det := 0.0
	data := m.Data()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				q := []int{i, j, k}
				s := 1.0
				for l := 0; l < 3; l++ {
					s *= data[l][q[l]]
				}
				det += float64(Epsilon(i+1, j+1, k+1)) * s
			}
		}
	}

	return det, nil
}

// fmt.Printf("%s %d\n", p, e[p])
func Diag(v interface{}) (Matrix, error) {
	switch v.(type) {
	case vector.Vector:
		u := v.(vector.Vector)
		m := Matrix{}
		m.M11 = u.X
		m.M22 = u.Y
		m.M33 = u.Z
		return m, nil
	case Matrix:
		m := v.(Matrix)
		u := Matrix{}
		u.M11 = m.M11
		u.M22 = m.M22
		u.M33 = m.M33
		return u, nil
	default:
		return Matrix{}, fmt.Errorf("Matrix has no such operation")

	}
}
func DiagV(v interface{}) (vector.Vector, error) {
	m, e := Diag(v)
	u, _ := vector.New([]float64{m.M11, m.M22, m.M33})
	return u, e
}
func (m *Matrix) Inverse() (Matrix, error) {
	det, _ := Determinant(*m)
	dataM := m.Data()
	if math.Abs(det) > 0.0 {
		dataM[0][0] = (m.M22*m.M33 - m.M32*m.M23) / det
		dataM[0][1] = (m.M13*m.M32 - m.M33*m.M12) / det
		dataM[0][2] = (m.M12*m.M23 - m.M22*m.M13) / det

		dataM[1][0] = (m.M23*m.M31 - m.M33*m.M21) / det
		dataM[1][1] = (m.M11*m.M33 - m.M31*m.M13) / det
		dataM[1][2] = (m.M13*m.M21 - m.M23*m.M11) / det

		dataM[2][0] = (m.M21*m.M32 - m.M31*m.M22) / det
		dataM[2][1] = (m.M12*m.M31 - m.M32*m.M11) / det
		dataM[2][2] = (m.M11*m.M22 - m.M21*m.M12) / det

	}
	u, _ := New(dataM)
	return u, nil
}

func Epsilon(values ...int) int {
	m := make(map[int]int)
	for _, v := range values {
		m[v] = v
	}
	if len(m) < len(values) {
		return 0.0
	}

	ep := 1
	for l, v1 := range values {
		for n, v2 := range values {

			if (v1 != v2) && (l > n) {
				ep *= sign(v1 - v2)
			}
		}
	}

	return ep
}
func sign(i int) int {
	if i < 0 {
		return -1
	}
	return 1
}
