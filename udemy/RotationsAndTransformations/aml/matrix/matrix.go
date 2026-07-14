package matrix

import (
	"fmt"
	"math"

	"udemy.com/aml/vector"
)

// Matrix is a 3x3 matrix with value semantics: all methods return new
// Matrices and never mutate the receiver.
type Matrix struct {
	M11, M12, M13 float64
	M21, M22, M23 float64
	M31, M32, M33 float64
}

// New returns a Matrix from a 3x3 array of rows.
func New(rows [3][3]float64) Matrix {
	return Matrix{
		rows[0][0], rows[0][1], rows[0][2],
		rows[1][0], rows[1][1], rows[1][2],
		rows[2][0], rows[2][1], rows[2][2],
	}
}

// FromRows returns a Matrix whose rows are the given vectors.
func FromRows(r0, r1, r2 vector.Vector) Matrix {
	return Matrix{
		r0.X, r0.Y, r0.Z,
		r1.X, r1.Y, r1.Z,
		r2.X, r2.Y, r2.Z,
	}
}

// Splat returns a Matrix with every element set to s.
func Splat(s float64) Matrix {
	return Matrix{s, s, s, s, s, s, s, s, s}
}

// Identity returns the 3x3 identity matrix.
func Identity() Matrix {
	return Matrix{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
}

// Data returns the matrix as rows of a 3x3 array.
func (m Matrix) Data() [3][3]float64 {
	return [3][3]float64{
		{m.M11, m.M12, m.M13},
		{m.M21, m.M22, m.M23},
		{m.M31, m.M32, m.M33},
	}
}

// Row returns the i-th row (0-based) as a vector.
func (m Matrix) Row(i int) vector.Vector {
	switch i {
	case 0:
		return vector.New(m.M11, m.M12, m.M13)
	case 1:
		return vector.New(m.M21, m.M22, m.M23)
	default:
		return vector.New(m.M31, m.M32, m.M33)
	}
}

// Add returns m + n.
func (m Matrix) Add(n Matrix) Matrix {
	return Matrix{
		m.M11 + n.M11, m.M12 + n.M12, m.M13 + n.M13,
		m.M21 + n.M21, m.M22 + n.M22, m.M23 + n.M23,
		m.M31 + n.M31, m.M32 + n.M32, m.M33 + n.M33,
	}
}

// Sub returns m - n.
func (m Matrix) Sub(n Matrix) Matrix {
	return Matrix{
		m.M11 - n.M11, m.M12 - n.M12, m.M13 - n.M13,
		m.M21 - n.M21, m.M22 - n.M22, m.M23 - n.M23,
		m.M31 - n.M31, m.M32 - n.M32, m.M33 - n.M33,
	}
}

// Scale returns m scaled by s.
func (m Matrix) Scale(s float64) Matrix {
	return Matrix{
		m.M11 * s, m.M12 * s, m.M13 * s,
		m.M21 * s, m.M22 * s, m.M23 * s,
		m.M31 * s, m.M32 * s, m.M33 * s,
	}
}

// Neg returns -m.
func (m Matrix) Neg() Matrix { return m.Scale(-1) }

// Mul returns the matrix product m * n.
func (m Matrix) Mul(n Matrix) Matrix {
	a := m.Data()
	b := n.Data()
	var c [3][3]float64
	for i := range 3 {
		for j := range 3 {
			for k := range 3 {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return New(c)
}

// MulVec returns the matrix-vector product m * v.
func (m Matrix) MulVec(v vector.Vector) vector.Vector {
	return vector.New(
		m.M11*v.X+m.M12*v.Y+m.M13*v.Z,
		m.M21*v.X+m.M22*v.Y+m.M23*v.Z,
		m.M31*v.X+m.M32*v.Y+m.M33*v.Z,
	)
}

// Transpose returns the transpose of m.
func (m Matrix) Transpose() Matrix {
	return Matrix{
		m.M11, m.M21, m.M31,
		m.M12, m.M22, m.M32,
		m.M13, m.M23, m.M33,
	}
}

// Determinant returns the determinant of m via the Levi-Civita symbol.
func (m Matrix) Determinant() float64 {
	data := m.Data()
	det := 0.0
	for i := range 3 {
		for j := range 3 {
			for k := range 3 {
				q := [3]int{i, j, k}
				s := 1.0
				for l := range 3 {
					s *= data[l][q[l]]
				}
				det += float64(epsilon(i+1, j+1, k+1)) * s
			}
		}
	}
	return det
}

// Inverse returns the inverse of m. It returns an error if m is singular.
func (m Matrix) Inverse() (Matrix, error) {
	det := m.Determinant()
	if math.Abs(det) == 0.0 {
		return Matrix{}, fmt.Errorf("matrix: cannot invert singular matrix")
	}
	return Matrix{
		(m.M22*m.M33 - m.M32*m.M23) / det,
		(m.M13*m.M32 - m.M33*m.M12) / det,
		(m.M12*m.M23 - m.M22*m.M13) / det,

		(m.M23*m.M31 - m.M33*m.M21) / det,
		(m.M11*m.M33 - m.M31*m.M13) / det,
		(m.M13*m.M21 - m.M23*m.M11) / det,

		(m.M21*m.M32 - m.M31*m.M22) / det,
		(m.M12*m.M31 - m.M32*m.M11) / det,
		(m.M11*m.M22 - m.M21*m.M12) / det,
	}, nil
}

// Diag returns a diagonal matrix from the components of v.
func Diag(v vector.Vector) Matrix {
	return Matrix{
		v.X, 0, 0,
		0, v.Y, 0,
		0, 0, v.Z,
	}
}

// DiagV returns the diagonal of m as a vector.
func (m Matrix) DiagV() vector.Vector {
	return vector.New(m.M11, m.M22, m.M33)
}

// epsilon is the Levi-Civita permutation symbol for the given indices.
func epsilon(values ...int) int {
	seen := make(map[int]bool)
	for _, v := range values {
		if seen[v] {
			return 0
		}
		seen[v] = true
	}
	ep := 1
	for l, v1 := range values {
		for n, v2 := range values {
			if v1 != v2 && l > n {
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
