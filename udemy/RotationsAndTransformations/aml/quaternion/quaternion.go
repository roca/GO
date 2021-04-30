package quaternion

import (
	"fmt"
	"math"

	"strings"

	"udemy.com/aml/dcm"
	"udemy.com/aml/euler"
	"udemy.com/aml/matrix"
	"udemy.com/aml/vector"
)

type IQuaternion interface {
	Norm() float64
}

type Quaternion struct {
	S, X, Y, Z float64
}

func New(s, x, y, z float64) (Quaternion, error) {
	return Quaternion{S: s, X: x, Y: y}, nil
}

func (m *Quaternion) Copy() (Quaternion, error) {
	u := m
	return *u, nil
}
func (m *Quaternion) CopyPointer() (*Quaternion, error) {
	var new *Quaternion
	x, _ := m.Copy()
	new = &x
	return new, nil
}

// Quaternion Operations
func (q Quaternion) Conjugate() (Quaternion, error) {
	return Quaternion{S: q.S, X: -q.X, Y: -q.Y, Z: -q.Z}, nil
}
func (q Quaternion) Norm() float64 {
	return math.Sqrt(q.S*q.S + q.X*q.X + q.Y*q.Y + q.Z*q.Z)
}
func (q Quaternion) Inverse() (Quaternion, error) {
	qI, _ := q.Conjugate()
	mag := q.Norm()
	return Quaternion{S: qI.S / mag, X: qI.X / mag, Y: qI.Y / mag, Z: qI.Z / mag}, nil
}
func (q Quaternion) Unit() (Quaternion, error) {
	mag := q.Norm()
	if mag > 0.0 {
		return Quaternion{S: q.S / mag, X: q.X / mag, Y: q.Y / mag, Z: q.Z / mag}, nil
	}
	return Quaternion{S: q.S, X: q.X, Y: q.Y, Z: q.Z}, nil
}
func (q *Quaternion) Normalise() {
	mag := q.Norm()
	if mag > 0.0 {
		q.S /= mag
		q.X /= mag
		q.Y /= mag
		q.Z /= mag
	}
}
func (q Quaternion) Dot(q2 Quaternion) float64 {
	return (q.S*q2.S + q.X*q2.X + q.Y*q2.Y + q.Z*q2.Z)
}
func IsUnitQuat(q Quaternion, tol ...float64) bool {
	tolerance := .0000000001
	if len(tol) != 0 {
		tolerance = tol[0]
	}
	return (q.Norm() - 1.0) < (2.0 * tolerance)
}
func Angles2Quat(angle euler.Angles) (Quaternion, error) {
	dcm, _ := angle.ToDCM()
	return Dcm2Quat(dcm)
}

func (q Quaternion) ToAngles(sequence euler.Seq) (euler.Angles, error) {
	dcm, _ := Quat2DCM(q)
	return euler.DcmToAngles(dcm, sequence)
}

// DCM Conversion Functions
func Dcm2Quat(r matrix.Matrix) (Quaternion, error) {
	if !dcm.IsOrthogonal(r) {
		return Quaternion{}, nil
	}

	q0_den := (1.0 + r.M11 + r.M22 + r.M33)
	s0 := math.Sqrt(q0_den)

	q1_den := (1.0 + r.M11 - r.M22 - r.M33)
	s1 := math.Sqrt(q1_den)

	q2_den := (1.0 - r.M11 + r.M22 - r.M33)
	s2 := math.Sqrt(q2_den)

	q3_den := (1.0 - r.M11 - r.M22 + r.M33)
	s3 := math.Sqrt(q3_den)

	q2q3 := r.M23 + r.M32
	q1q3 := r.M31 + r.M13
	q1q2 := r.M12 + r.M21
	q0q1 := r.M23 - r.M32
	q0q2 := r.M31 - r.M13
	q0q3 := r.M12 - r.M21

	qs := make(map[float64]Quaternion)
	qs[q0_den] = Quaternion{
		S: 0.5 * s0,
		X: 0.5 * q0q1 / s0,
		Y: 0.5 * q0q2 / s0,
		Z: 0.5 * q0q3 / s0,
	}
	qs[q1_den] = Quaternion{
		S: 0.5 * q0q1 / s1,
		X: 0.5 * s1,
		Y: 0.5 * q1q2 / s1,
		Z: 0.5 * q1q3 / s1,
	}
	qs[q2_den] = Quaternion{
		S: 0.5 * q0q2 / s2,
		X: 0.5 * q1q2 / s2,
		Y: 0.5 * s2,
		Z: 0.5 * q2q3 / s2,
	}
	qs[q3_den] = Quaternion{
		S: 0.5 * q0q3 / s3,
		X: 0.5 * q1q3 / s3,
		Y: 0.5 * q2q3 / s3,
		Z: 0.5 * s3,
	}

	q_dens := []float64{q0_den, q1_den, q2_den, q3_den}
	max_q_den := q_dens[0]
	for _, q_den := range q_dens {
		max_q_den = math.Max(max_q_den, q_den)
	}
	return qs[max_q_den], nil
}
func Quat2DCM(rhs Quaternion) (matrix.Matrix, error) {
	TOL := 0.0001

	// Check if valid rotation quaternion
	if IsUnitQuat(rhs, TOL) {
		q0 := rhs.S
		q1 := rhs.X
		q2 := rhs.Y
		q3 := rhs.Z
		q0_2 := q0 * q0
		q1_2 := q1 * q1
		q2_2 := q2 * q2
		q3_2 := q3 * q3
		q1q2 := q1 * q2
		q0q3 := q0 * q3
		q1q3 := q1 * q3
		q0q2 := q0 * q2
		q2q3 := q2 * q3
		q0q1 := q0 * q1
		data := [][]float64{
			{
				q0_2 + q1_2 - q2_2 - q3_2,
				2.0 * (q1q2 + q0q3),
				2.0 * (q1q3 - q0q2),
			},
			{
				2.0 * (q1q2 - q0q3),
				q0_2 - q1_2 + q2_2 - q3_2,
				2.0 * (q2q3 + q0q1),
			},
			{2.0 * (q1q3 + q0q2), 2.0 * (q2q3 - q0q1), q0_2 - q1_2 - q2_2 + q3_2},
		}
		dcm, _ := matrix.New(data)
		return dcm, nil
	}
	return matrix.Matrix{}, fmt.Errorf("Quaternion Norm %f != 1.0", rhs.Norm())
}

// Quaternion / Quaternion Operations
func (q *Quaternion) Qop(operation string, u Quaternion) (*Quaternion, error) {
	var new *Quaternion
	if !strings.Contains(operation, "=") {
		new, _ = q.CopyPointer()
	} else {
		new = q
	}

	switch o := operation; {
	case o == "+=" || o == "+":
		new.S += u.S
		new.X += u.X
		new.Y += u.Y
		new.Z += u.Z
		return new, nil
	case o == "-=" || o == "-":
		new.S -= u.S
		new.X -= u.X
		new.Y -= u.Y
		new.Z -= u.Z
	case o == "*=" || o == "*":
		s_new := (u.S * q.S) - (u.X * q.X) - (u.Y * q.Y) - (u.Z * q.Z)
		x_new := (u.S * q.X) + (u.X * q.S) - (u.Y * q.Z) + (u.Z * q.Y)
		y_new := (u.S * q.Y) + (u.X * q.Z) + (u.Y * q.S) - (u.Z * q.X)
		z_new := (u.S * q.Z) - (u.X * q.Y) + (u.Y * q.X) + (u.Z * q.S)
		new.S = s_new
		new.Y = x_new
		new.X = y_new
		new.Z = z_new
		return new, nil
	default:
		return &Quaternion{}, fmt.Errorf("Matrix has no such operation '%s'", o)
	}
	return new, nil
}

// Quaternion / Scalar Operations
func (q *Quaternion) Sop(operation string, s float64) (*Quaternion, error) {
	var new *Quaternion
	if !strings.Contains(operation, "=") {
		new, _ = q.CopyPointer()
	} else {
		new = q
	}

	switch o := operation; {
	case o == "+=" || o == "+":
		new.S += s
		new.X += s
		new.Y += s
		new.Z += s
		return new, nil
	case o == "-=" || o == "-":
		new.S -= s
		new.X -= s
		new.Y -= s
		new.Z -= s
	case o == "*=" || o == "*":
		new.S *= s
		new.X *= s
		new.Y *= s
		new.Z *= s
	case o == "/=" || o == "/":
		new.S /= s
		new.X /= s
		new.Y /= s
		new.Z /= s
		return new, nil
	default:
		return &Quaternion{}, fmt.Errorf("Matrix has no such operation '%s'", o)
	}
	return new, nil
}

// Quaternion / Vector Operations
func (q Quaternion) Vop(operation string, v vector.Vector) (vector.Vector, error) {

	switch o := operation; {
	case o == "*":
		dcm, _ := Quat2DCM(q)
		u, _ := dcm.Vop("*", v)
		return u, nil
	default:
		return vector.Vector{}, fmt.Errorf("Matrix has no such operation '%s'", o)
	}
}
