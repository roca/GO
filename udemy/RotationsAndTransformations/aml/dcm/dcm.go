package dcm

import (
	"fmt"
	"math"

	"udemy.com/aml/matrix"
	"udemy.com/aml/vector"
)

func RotationX(theta float64) (matrix.Matrix, error) {
	m, e := matrix.New([][]float64{
		{1., 0., 0.},
		{0., math.Cos(theta), math.Sin(theta)},
		{0., -1.0 * math.Sin(theta), math.Cos(theta)},
	})
	if e != nil {
		return matrix.Matrix{}, e
	}
	return m, nil
}
func RotationY(theta float64) (matrix.Matrix, error) {
	m, e := matrix.New([][]float64{
		{math.Cos(theta), 0., -1.0 * math.Sin(theta)},
		{0., 1.0, 0.},
		{math.Sin(theta), 0., math.Cos(theta)},
	})
	if e != nil {
		return matrix.Matrix{}, e
	}
	return m, nil
}
func RotationZ(theta float64) (matrix.Matrix, error) {
	m, e := matrix.New([][]float64{
		{math.Cos(theta), math.Sin(theta), 0.},
		{-1.0 * math.Sin(theta), math.Cos(theta), 0.},
		{0., 0., 1.},
	})
	if e != nil {
		return matrix.Matrix{}, e
	}
	return m, nil
}

func XYZRotation(phi, theta, si float64) (matrix.Matrix, error) {
	Rx, _ := RotationX(phi)
	Ry, _ := RotationY(theta)
	Rz, _ := RotationZ(si)

	Ryz, _ := Ry.Mop("*", Rz)
	Rxyz, _ := Rx.Mop("*", *Ryz)

	return *Rxyz, nil
}

func ZXZRotation(phi, theta, si float64) (matrix.Matrix, error) {
	Rz, _ := RotationZ(phi)
	Rx, _ := RotationX(theta)
	Rz2, _ := RotationZ(si)

	Rxz, _ := Rx.Mop("*", Rz2)
	Rzxz, _ := Rz.Mop("*", *Rxz)

	return *Rzxz, nil
}

func IsOrthogonal(m matrix.Matrix) bool {
	det, _ := matrix.Determinant(m)
	if math.Abs(det) > 1.0000000001 {
		return false
	}
	mInverse, _ := m.Inverse()
	mTranspose, _ := matrix.Transpose(m)
	dataI := mInverse.Data()
	dataT := mTranspose.Data()

	for i, row := range dataI {
		for j, vI := range row {
			if math.Abs(vI-dataT[i][j]) > .0000000001 {
				return false
			}
		}
	}
	return true
}
func Normalize2(m *matrix.Matrix) error {
	V1, _ := vector.New(m.M11, m.M12, m.M13)
	V2, _ := vector.New(m.M21, m.M22, m.M23)
	V3, _ := vector.New(m.M31, m.M32, m.M33)

	U1 := vector.Cross(V2, V3)
	_, _ = U1.Sop("*=", vector.Dot(V1, V3))
	U2 := vector.Cross(V3, V1)
	_, _ = U2.Sop("*=", vector.Dot(V3, V1))
	U3 := vector.Cross(U1, U2)

	dcm, _ := matrix.New(U1, U2, U3)
	fmt.Printf("%f\n-------------------\n", dcm.Data())
	*m = dcm

	return nil
}

func Normalize(m *matrix.Matrix) error {
	X, _ := vector.New(m.M11, m.M12, m.M13)
	Y, _ := vector.New(m.M21, m.M22, m.M23)
	vError := vector.Dot(X, Y)

	xError, _ := X.Sop("*", .5*vError)
	yError, _ := Y.Sop("*", .5*vError)

	xOrth, _ := X.Vop("-", *yError)
	yOrth, _ := Y.Vop("-", *xError)
	zOrth := vector.Cross(X, Y)

	dotX := 0.5 * (3.0 - vector.Dot(*xOrth, *xOrth))
	dotY := 0.5 * (3.0 - vector.Dot(*yOrth, *yOrth))
	dotZ := 0.5 * (3.0 - vector.Dot(zOrth, zOrth))

	xNorm, _ := xOrth.Sop("*", dotX)
	yNorm, _ := yOrth.Sop("*", dotY)
	zNorm, _ := zOrth.Sop("*", dotZ)

	dcm, _ := matrix.New(*xNorm, *yNorm, *zNorm)
	dcmT, _ := matrix.Transpose(dcm)

	m.Mutate(dcmT.Data())

	return nil
}

func Intergrate(dcm, dcmRates *matrix.Matrix, dt float64) (matrix.Matrix, error) {
	dcmDt, _ := dcmRates.Sop("*", dt)
	dcmNew, _ := dcm.Mop("+", *dcmDt)
	_ = Normalize(dcmNew)

	return *dcmNew, nil
}

func KinematicRatesFromBodyRates(dcm *matrix.Matrix, bodyRates *vector.Vector) (matrix.Matrix, error) {
	p := bodyRates.X
	q := bodyRates.Y
	r := bodyRates.Z
	skeMatrix, _ := matrix.New([][]float64{
		{0.0, -r, q},
		{r, 0.0, -p},
		{-q, p, 0.0},
	})
	_, _ = skeMatrix.Sop("*=", -1.0)
	_, _ = skeMatrix.Mop("*=", *dcm)
	// _ = Normalize(&skeMatrix)
	// if !IsOrthogonal(skeMatrix) {
	// 	return matrix.Matrix{}, errors.New("DCM is not Orthoganal")
	// }
	return skeMatrix, nil
}
func KinematicRatesFromWorldRates(dcm *matrix.Matrix, worldRates *vector.Vector) (matrix.Matrix, error) {
	p := worldRates.X
	q := worldRates.Y
	r := worldRates.Z
	skeMatrix, _ := matrix.New([][]float64{
		{0.0, -r, q},
		{r, 0.0, -p},
		{-q, p, 0.0},
	})
	finalSkeMatrix, _ := dcm.Mop("*", skeMatrix)
	// _ = Normalize(&finalSkeMatrix)
	// if !IsOrthogonal(finalSkeMatrix) {
	// 	return matrix.Matrix{}, errors.New("DCM is not Orthoganal")
	// }
	return *finalSkeMatrix, nil
}
func RadiansToDegrees(radians float64) (degrees float64) {
	degrees = radians * 180.0 / math.Pi
	return
}
func DegreesToRadians(degrees float64) (radians float64) {
	radians = degrees * math.Pi / 180.0
	return
}

func EulerAnglesFromRxyz(Rxyz matrix.Matrix) (phi, theta, si float64) {
	phi = math.Atan2(Rxyz.M23, Rxyz.M33)
	theta = -1.0 * math.Asin(Rxyz.M13)
	si = math.Atan2(Rxyz.M12, Rxyz.M11)
	return
}
func EulerAnglesFromRzxz(Rzxz matrix.Matrix) (phi, theta, si float64) {
	phi = math.Atan2(Rzxz.M13, Rzxz.M23)
	theta = math.Acos(Rzxz.M33)
	si = math.Atan2(Rzxz.M31, -1.0*Rzxz.M32)
	return
}

// Euler Angle Rates
// Singularity: For theta = 0 degrees, Rates go to infinity

func XYZEulerAngleRates(phi, theta, si float64, omega_body vector.Vector) (vector.Vector, error) {
	Exyz, _ := matrix.New(0)

	Exyz.M11 = 1.0
	Exyz.M12 = math.Tan(theta) * math.Sin(phi)
	Exyz.M13 = math.Tan(theta) * math.Cos(phi)

	Exyz.M21 = 0.0
	Exyz.M22 = math.Cos(phi)
	Exyz.M23 = -1.0 * math.Sin(phi)

	Exyz.M31 = 0.0
	Exyz.M32 = math.Sin(phi) / math.Cos(theta)
	Exyz.M33 = math.Cos(phi) / math.Cos(theta)

	w, _ := Exyz.Vop("*", omega_body)

	return w, nil
}

func ZXZEulerAngleRates(phi, theta, si float64, omega_body vector.Vector) (vector.Vector, error) {
	Ezxz, _ := matrix.New(0)

	Ezxz.M11 = -1.0 * math.Sin(phi) * math.Cos(theta) / math.Sin(theta)
	Ezxz.M11 = -1.0 * math.Cos(phi) * math.Cos(theta) / math.Sin(theta)
	Ezxz.M13 = 1.0

	Ezxz.M21 = math.Cos(phi)
	Ezxz.M22 = -1.0 * math.Sin(phi)
	Ezxz.M23 = 0.0

	Ezxz.M31 = math.Sin(phi) / math.Sin(theta)
	Ezxz.M32 = math.Cos(phi) / math.Sin(theta)
	Ezxz.M33 = 0.0

	w, _ := Ezxz.Vop("*", omega_body)

	return w, nil
}

func EulerIntergration(x, xDot vector.Vector, dt float64) (vector.Vector, error) {
	v, _ := xDot.Sop("*", dt)
	w, _ := x.Vop("+", *v)
	return *w, nil
}
