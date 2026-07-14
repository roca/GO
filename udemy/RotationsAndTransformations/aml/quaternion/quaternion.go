package quaternion

import (
	"fmt"
	"math"

	"udemy.com/aml/dcm"
	"udemy.com/aml/euler"
	"udemy.com/aml/matrix"
	"udemy.com/aml/vector"
)

// Quaternion is a rotation quaternion with value semantics: all methods
// return new Quaternions and never mutate the receiver.
type Quaternion struct {
	S, X, Y, Z float64
}

// New returns a Quaternion with the given components.
func New(s, x, y, z float64) Quaternion {
	return Quaternion{S: s, X: x, Y: y, Z: z}
}

// Conjugate returns the conjugate of q.
func (q Quaternion) Conjugate() Quaternion {
	return Quaternion{S: q.S, X: -q.X, Y: -q.Y, Z: -q.Z}
}

// Norm returns the magnitude of q.
func (q Quaternion) Norm() float64 {
	return math.Sqrt(q.S*q.S + q.X*q.X + q.Y*q.Y + q.Z*q.Z)
}

// Inverse returns the inverse of q.
func (q Quaternion) Inverse() Quaternion {
	c := q.Conjugate()
	mag := q.Norm()
	return Quaternion{S: c.S / mag, X: c.X / mag, Y: c.Y / mag, Z: c.Z / mag}
}

// Unit returns the unit quaternion in the direction of q, or q unchanged if
// it has zero magnitude.
func (q Quaternion) Unit() Quaternion {
	mag := q.Norm()
	if mag > 0 {
		return q.Scale(1 / mag)
	}
	return q
}

// Normalize returns the unit quaternion in the direction of q.
func (q Quaternion) Normalize() Quaternion { return q.Unit() }

// Dot returns the dot product q · r.
func (q Quaternion) Dot(r Quaternion) float64 {
	return q.S*r.S + q.X*r.X + q.Y*r.Y + q.Z*r.Z
}

// IsUnitQuat reports whether q has unit norm within an optional tolerance.
func IsUnitQuat(q Quaternion, tol ...float64) bool {
	tolerance := 1e-10
	if len(tol) != 0 {
		tolerance = tol[0]
	}
	return math.Abs(q.Norm()-1.0) < (2.0 * tolerance)
}

// Add returns q + r.
func (q Quaternion) Add(r Quaternion) Quaternion {
	return Quaternion{S: q.S + r.S, X: q.X + r.X, Y: q.Y + r.Y, Z: q.Z + r.Z}
}

// Sub returns q - r.
func (q Quaternion) Sub(r Quaternion) Quaternion {
	return Quaternion{S: q.S - r.S, X: q.X - r.X, Y: q.Y - r.Y, Z: q.Z - r.Z}
}

// Mul returns the Hamilton product r ⊗ q.
func (q Quaternion) Mul(r Quaternion) Quaternion {
	return Quaternion{
		S: r.S*q.S - r.X*q.X - r.Y*q.Y - r.Z*q.Z,
		X: r.S*q.X + r.X*q.S - r.Y*q.Z + r.Z*q.Y,
		Y: r.S*q.Y + r.X*q.Z + r.Y*q.S - r.Z*q.X,
		Z: r.S*q.Z - r.X*q.Y + r.Y*q.X + r.Z*q.S,
	}
}

// Scale returns q scaled by s.
func (q Quaternion) Scale(s float64) Quaternion {
	return Quaternion{S: q.S * s, X: q.X * s, Y: q.Y * s, Z: q.Z * s}
}

// RotateVec rotates v by the rotation represented by q.
func (q Quaternion) RotateVec(v vector.Vector) (vector.Vector, error) {
	m, err := Quat2DCM(q)
	if err != nil {
		return vector.Vector{}, err
	}
	return m.MulVec(v), nil
}

// Angles2Quat converts Euler angles to a quaternion via their DCM.
func Angles2Quat(angle euler.Angles) (Quaternion, error) {
	m, err := angle.ToDCM()
	if err != nil {
		return Quaternion{}, err
	}
	return Dcm2Quat(m)
}

// ToAngles converts q to Euler angles for the given sequence via its DCM.
func (q Quaternion) ToAngles(sequence euler.Seq) (euler.Angles, error) {
	m, err := Quat2DCM(q)
	if err != nil {
		return euler.Angles{}, err
	}
	return euler.DcmToAngles(m, sequence)
}

// Dcm2Quat converts an orthogonal DCM to a rotation quaternion, using the
// numerically most stable of the four extraction formulas.
func Dcm2Quat(r matrix.Matrix) (Quaternion, error) {
	if !dcm.IsOrthogonal(r) {
		return Quaternion{}, fmt.Errorf("quaternion: DCM is not orthogonal")
	}

	q0den := 1.0 + r.M11 + r.M22 + r.M33
	q1den := 1.0 + r.M11 - r.M22 - r.M33
	q2den := 1.0 - r.M11 + r.M22 - r.M33
	q3den := 1.0 - r.M11 - r.M22 + r.M33

	q2q3 := r.M23 + r.M32
	q1q3 := r.M31 + r.M13
	q1q2 := r.M12 + r.M21
	q0q1 := r.M23 - r.M32
	q0q2 := r.M31 - r.M13
	q0q3 := r.M12 - r.M21

	qs := map[float64]Quaternion{
		q0den: {S: 0.5 * math.Sqrt(q0den), X: 0.5 * q0q1 / math.Sqrt(q0den), Y: 0.5 * q0q2 / math.Sqrt(q0den), Z: 0.5 * q0q3 / math.Sqrt(q0den)},
		q1den: {S: 0.5 * q0q1 / math.Sqrt(q1den), X: 0.5 * math.Sqrt(q1den), Y: 0.5 * q1q2 / math.Sqrt(q1den), Z: 0.5 * q1q3 / math.Sqrt(q1den)},
		q2den: {S: 0.5 * q0q2 / math.Sqrt(q2den), X: 0.5 * q1q2 / math.Sqrt(q2den), Y: 0.5 * math.Sqrt(q2den), Z: 0.5 * q2q3 / math.Sqrt(q2den)},
		q3den: {S: 0.5 * q0q3 / math.Sqrt(q3den), X: 0.5 * q1q3 / math.Sqrt(q3den), Y: 0.5 * q2q3 / math.Sqrt(q3den), Z: 0.5 * math.Sqrt(q3den)},
	}

	maxDen := math.Max(math.Max(q0den, q1den), math.Max(q2den, q3den))
	return qs[maxDen], nil
}

// Quat2DCM converts a unit rotation quaternion to a DCM.
func Quat2DCM(q Quaternion) (matrix.Matrix, error) {
	const tol = 0.0001
	if !IsUnitQuat(q, tol) {
		return matrix.Matrix{}, fmt.Errorf("quaternion: norm %f != 1.0", q.Norm())
	}
	q0, q1, q2, q3 := q.S, q.X, q.Y, q.Z
	q0_2, q1_2, q2_2, q3_2 := q0*q0, q1*q1, q2*q2, q3*q3
	q1q2, q0q3 := q1*q2, q0*q3
	q1q3, q0q2 := q1*q3, q0*q2
	q2q3, q0q1 := q2*q3, q0*q1
	return matrix.New([3][3]float64{
		{q0_2 + q1_2 - q2_2 - q3_2, 2.0 * (q1q2 + q0q3), 2.0 * (q1q3 - q0q2)},
		{2.0 * (q1q2 - q0q3), q0_2 - q1_2 + q2_2 - q3_2, 2.0 * (q2q3 + q0q1)},
		{2.0 * (q1q3 + q0q2), 2.0 * (q2q3 - q0q1), q0_2 - q1_2 - q2_2 + q3_2},
	}), nil
}

// KinematicRatesBodyRates returns q̇ for body angular rates.
func KinematicRatesBodyRates(q Quaternion, bodyRates vector.Vector) Quaternion {
	p, r2, r := bodyRates.X, bodyRates.Y, bodyRates.Z
	return Quaternion{
		S: 0.5 * (-q.X*p - q.Y*r2 - q.Z*r),
		X: 0.5 * (q.S*p + q.Z*r2 - q.Y*r),
		Y: 0.5 * (-q.Z*p + q.S*r2 + q.X*r),
		Z: 0.5 * (q.Y*p - q.X*r2 + q.S*r),
	}
}

// KinematicRatesWorldRates returns q̇ for world angular rates.
func KinematicRatesWorldRates(q Quaternion, worldRates vector.Vector) Quaternion {
	p, r2, r := worldRates.X, worldRates.Y, worldRates.Z
	return Quaternion{
		S: 0.5 * (-q.X*p - q.Y*r2 - q.Z*r),
		X: 0.5 * (q.S*p - q.Z*r2 + q.Y*r),
		Y: 0.5 * (q.Z*p + q.S*r2 - q.X*r),
		Z: 0.5 * (-q.Y*p + q.X*r2 + q.S*r),
	}
}

// Integrate advances quat by quatRates over dt, renormalizing the result.
func Integrate(quat, quatRates Quaternion, dt float64) Quaternion {
	return quat.Add(quatRates.Scale(dt)).Normalize()
}

// linearInterpolate performs normalized linear interpolation (nlerp).
func linearInterpolate(startQuat, endQuat Quaternion, t float64) Quaternion {
	q0 := startQuat.Unit()
	q1 := endQuat.Unit()
	if t < 0.0 {
		return q0
	}
	if t > 1.0 {
		return q1
	}
	return q0.Scale(1.0 - t).Add(q1.Scale(t)).Unit()
}

// SlerpInterpolate performs spherical linear interpolation between two
// quaternions, falling back to nlerp for small angles.
func SlerpInterpolate(startQuat, endQuat Quaternion, t float64) Quaternion {
	q0 := startQuat.Unit()
	q1 := endQuat.Unit()
	if t < 0.0 {
		return q0
	}
	if t > 1.0 {
		return q1
	}

	quatDot := q0.Dot(q1)
	if quatDot < 0 {
		q1 = q1.Scale(-1.0)
		quatDot = -quatDot
	}

	theta := math.Acos(quatDot)
	if theta < 0.0001 {
		return linearInterpolate(startQuat, endQuat, t)
	}

	a := math.Sin((1.0-t)*theta) / math.Sin(theta)
	b := math.Sin(t*theta) / math.Sin(theta)
	return q0.Scale(a).Add(q1.Scale(b)).Unit()
}
