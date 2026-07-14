package dcm

import (
	"math"

	"udemy.com/aml/matrix"
	"udemy.com/aml/vector"
)

// RotationX returns the [R]ᵀ rotation matrix about the X axis.
func RotationX(theta float64) matrix.Matrix {
	return matrix.New([3][3]float64{
		{1, 0, 0},
		{0, math.Cos(theta), math.Sin(theta)},
		{0, -math.Sin(theta), math.Cos(theta)},
	})
}

// RotationY returns the [R]ᵀ rotation matrix about the Y axis.
func RotationY(theta float64) matrix.Matrix {
	return matrix.New([3][3]float64{
		{math.Cos(theta), 0, -math.Sin(theta)},
		{0, 1, 0},
		{math.Sin(theta), 0, math.Cos(theta)},
	})
}

// RotationZ returns the [R]ᵀ rotation matrix about the Z axis.
func RotationZ(theta float64) matrix.Matrix {
	return matrix.New([3][3]float64{
		{math.Cos(theta), math.Sin(theta), 0},
		{-math.Sin(theta), math.Cos(theta), 0},
		{0, 0, 1},
	})
}

// IsOrthogonal reports whether m is orthogonal within an optional tolerance.
func IsOrthogonal(m matrix.Matrix, tol ...float64) bool {
	tolerance := 1e-10
	if len(tol) != 0 {
		tolerance = tol[0]
	}
	if math.Abs(m.Determinant()) > (1 + tolerance) {
		return false
	}
	inv, err := m.Inverse()
	if err != nil {
		return false
	}
	trans := m.Transpose()
	di := inv.Data()
	dt := trans.Data()
	for i := range 3 {
		for j := range 3 {
			if math.Abs(di[i][j]-dt[i][j]) > tolerance {
				return false
			}
		}
	}
	return true
}

// Normalize orthonormalizes m in place using a first-order correction of the
// first two rows and a cross product for the third.
func Normalize(m *matrix.Matrix) {
	x := m.Row(0)
	y := m.Row(1)
	vErr := vector.Dot(x, y)

	xOrth := x.Sub(y.Scale(0.5 * vErr))
	yOrth := y.Sub(x.Scale(0.5 * vErr))
	zOrth := vector.Cross(xOrth, yOrth)

	xNorm := xOrth.Scale(0.5 * (3.0 - vector.Dot(xOrth, xOrth)))
	yNorm := yOrth.Scale(0.5 * (3.0 - vector.Dot(yOrth, yOrth)))
	zNorm := zOrth.Scale(0.5 * (3.0 - vector.Dot(zOrth, zOrth)))

	*m = matrix.FromRows(xNorm, yNorm, zNorm)
}

// Normalize2 orthonormalizes m in place using Gram-Schmidt on its rows.
func Normalize2(m *matrix.Matrix) {
	r0, err := m.Row(0).Normalize()
	if err != nil {
		return
	}
	row1 := m.Row(1)
	r1, err := row1.Sub(r0.Scale(vector.Dot(r0, row1))).Normalize()
	if err != nil {
		return
	}
	r2 := vector.Cross(r0, r1)
	*m = matrix.FromRows(r0, r1, r2)
}

// Integrate advances dcm by dcmRates over dt, renormalizing the result.
func Integrate(dcm, dcmRates matrix.Matrix, dt float64) matrix.Matrix {
	out := dcm.Add(dcmRates.Scale(dt))
	Normalize(&out)
	return out
}

// KinematicRatesFromBodyRates returns Ṙ = -[ω×] R for body angular rates.
func KinematicRatesFromBodyRates(dcm matrix.Matrix, bodyRates vector.Vector) matrix.Matrix {
	skew := skewSymmetric(bodyRates)
	return skew.Scale(-1).Mul(dcm)
}

// KinematicRatesFromWorldRates returns Ṙ = -R [ω×] for world angular rates.
func KinematicRatesFromWorldRates(dcm matrix.Matrix, worldRates vector.Vector) matrix.Matrix {
	skew := skewSymmetric(worldRates)
	return dcm.Mul(skew.Scale(-1))
}

// skewSymmetric returns the skew-symmetric cross-product matrix [v×].
func skewSymmetric(v vector.Vector) matrix.Matrix {
	return matrix.New([3][3]float64{
		{0, -v.Z, v.Y},
		{v.Z, 0, -v.X},
		{-v.Y, v.X, 0},
	})
}

func RadiansToDegrees(radians float64) float64 { return radians * 180.0 / math.Pi }
func DegreesToRadians(degrees float64) float64 { return degrees * math.Pi / 180.0 }

func EulerAnglesFromRxyz(Rxyz matrix.Matrix) (phi, theta, si float64) {
	phi = math.Atan2(Rxyz.M23, Rxyz.M33)
	theta = -math.Asin(Rxyz.M13)
	si = math.Atan2(Rxyz.M12, Rxyz.M11)
	return
}

func EulerAnglesFromRzxz(Rzxz matrix.Matrix) (phi, theta, si float64) {
	phi = math.Atan2(Rzxz.M13, Rzxz.M23)
	theta = math.Acos(Rzxz.M33)
	si = math.Atan2(Rzxz.M31, -Rzxz.M32)
	return
}

// XYZEulerAngleRates maps body rates to XYZ Euler-angle rates.
// Singularity: at theta = ±90° the rates go to infinity.
func XYZEulerAngleRates(phi, theta, si float64, omegaBody vector.Vector) vector.Vector {
	tanTheta := math.Tan(theta)
	secTheta := 1.0 / math.Cos(theta)
	e := matrix.New([3][3]float64{
		{1, tanTheta * math.Sin(phi), tanTheta * math.Cos(phi)},
		{0, math.Cos(phi), -math.Sin(phi)},
		{0, math.Sin(phi) * secTheta, math.Cos(phi) * secTheta},
	})
	return e.MulVec(omegaBody)
}

// ZXZEulerAngleRates maps body rates to ZXZ Euler-angle rates.
// Singularity: at theta = 0 the rates go to infinity.
func ZXZEulerAngleRates(phi, theta, si float64, omegaBody vector.Vector) vector.Vector {
	cscTheta := 1.0 / math.Sin(theta)
	cotTheta := math.Cos(theta) * cscTheta
	e := matrix.New([3][3]float64{
		{-math.Sin(phi) * cotTheta, -math.Cos(phi) * cotTheta, 1},
		{math.Cos(phi), -math.Sin(phi), 0},
		{math.Sin(phi) * cscTheta, math.Cos(phi) * cscTheta, 0},
	})
	return e.MulVec(omegaBody)
}

// EulerIntegration advances x by xDot over dt.
func EulerIntegration(x, xDot vector.Vector, dt float64) vector.Vector {
	return x.Add(xDot.Scale(dt))
}
