package dcm

import (
	"errors"
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

func Rotation(thetaX, thetaY, thetaZ float64) (matrix.Matrix, error) {
	rX, _ := RotationX(thetaX)
	rY, _ := RotationY(thetaY)
	rZ, _ := RotationZ(thetaZ)

	r, _ := rX.Mop("*=", rY)
	_, _ = r.Mop("*=", rZ)

	return *r, nil
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

func Normalize(m *matrix.Matrix) error {
	X, _ := vector.New(m.M11, m.M12, m.M13)
	Y, _ := vector.New(m.M21, m.M22, m.M23)
	vError := vector.Dot(X, Y)

	xError, _ := X.Sop("*", .5*vError)
	yError, _ := Y.Sop("*", .5*vError)

	xOrth, _ := X.Vop("-", *yError)
	yOrth, _ := Y.Vop("-", *xError)
	zOrth := vector.Cross(X, Y)

	xNorm, _ := xOrth.Sop("*", 0.5*(3.0-vector.Dot(*xOrth, *xOrth)))
	yNorm, _ := yOrth.Sop("*", 0.5*(3.0-vector.Dot(*yOrth, *yOrth)))
	zNorm, _ := zOrth.Sop("*", 0.5*(3.0-vector.Dot(zOrth, zOrth)))

	dcmT, _ := matrix.New(*xNorm, *yNorm, *zNorm)
	dcm, _ := matrix.Transpose(dcmT)

	*m = dcm

	return nil
}

func Intergrate(dcm, dcmRates *matrix.Matrix, dt float64) (matrix.Matrix, error) {
	dcmDt, _ := dcmRates.Sop("*", dt)
	dcmNew, _ := dcm.Mop("+", *dcmDt)
	_ = Normalize(dcmNew)
	if !IsOrthogonal(*dcmNew) {
		return matrix.Matrix{}, errors.New("DCM is not Orthoganal")
	}

	return *dcmNew, nil
}

func KinematicRatesFromBodyRates(dcm *matrix.Matrix, bodyRates *vector.Vector) (matrix.Matrix, error) {
	return matrix.Matrix{}, nil
}
func KinematicRatesFromWorldRates(dcm *matrix.Matrix, worldRates *vector.Vector) (matrix.Matrix, error) {
	return matrix.Matrix{}, nil
}
