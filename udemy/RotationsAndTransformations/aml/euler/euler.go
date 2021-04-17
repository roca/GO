package euler

import (
	"fmt"
	"math"
	"strings"

	"udemy.com/aml/dcm"
	"udemy.com/aml/matrix"
	"udemy.com/aml/vector"
)

type seq string

const (
	ZXZ seq = "ZXZ"
	XYX     = "XYX"
	YZY     = "YZY"
	ZYZ     = "ZYZ"
	XZX     = "XZX"
	YXY     = "YXY"
	XYZ     = "XYZ"
	YZX     = "YZX"
	ZXY     = "ZXY"
	XZY     = "XZY"
	ZYX     = "ZYX"
	YXZ     = "YXZ"
)

type IAngles interface {
	ToDCM() (matrix.Matrix, error)
	Convert(Angles, seq) (Angles, error)
	KinematicRates(vector.Vector) (Angles, error)
}
type Angles struct {
	Sequence seq `default:"XYZ"`
	Phi      float64
	Theta    float64
	Si       float64
}

func New(phi, theta, si float64, sequence ...seq) Angles {
	if len(sequence) != 0 {
		return Angles{Phi: phi, Theta: theta, Si: si, Sequence: sequence[0]}
	}
	return Angles{Phi: phi, Theta: theta, Si: si, Sequence: XYZ}
}
func (a Angles) ToDCM() (matrix.Matrix, error) {
	rotations := map[string]func(float64) (matrix.Matrix, error){
		"X": dcm.RotationX,
		"Y": dcm.RotationY,
		"Z": dcm.RotationZ,
	}
	axises := strings.Split(string(a.Sequence), "")

	R1, _ := rotations[axises[0]](a.Phi)
	R2, _ := rotations[axises[1]](a.Theta)
	R3, _ := rotations[axises[2]](a.Si)

	R23, _ := R2.Mop("*", R3)
	R123, _ := R1.Mop("*", *R23)

	return *R123, nil
}
func (a *Angles) Convert(angles Angles, sequence seq) (Angles, error) {
	dcm, _ := angles.ToDCM()

	newAngles, _ := DcmToAngles(dcm, sequence)

	return newAngles, nil
}
func (angles Angles) KinematicRates(bodyRates vector.Vector) (Angles, error) {
	ratesMap := map[seq]func(vector.Vector) (Angles, error){
		XYZ: angles.ratesMatrixXYZ,
		ZXZ: angles.ratesMatrixZXZ,
		XYX: angles.ratesMatrixXYX,
		YZY: angles.ratesMatrixYZY,
		ZYZ: angles.ratesMatrixZYZ,
		XZX: angles.ratesMatrixXZX,
		YXY: angles.ratesMatrixYXY,
		YZX: angles.ratesMatrixYZX,
		ZXY: angles.ratesMatrixZXY,
		XZY: angles.ratesMatrixXZY,
		ZYX: angles.ratesMatrixZYX,
		YXZ: angles.ratesMatrixYXZ,
	}
	rates, _ := ratesMap[angles.Sequence](bodyRates)

	return rates, nil
}
func DcmToAngles(m matrix.Matrix, sequence seq) (Angles, error) {
	if !dcm.IsOrthogonal(m) {
		return Angles{}, fmt.Errorf("This Matrix is not orthoganal %g", m)
	}

	anglesMap := map[seq]func(matrix.Matrix) (Angles, error){
		XYZ: dcmToAnglesXYZ,
		ZXZ: dcmToAnglesZXZ,
		XYX: dcmToAnglesXYX,
		YZY: dcmToAnglesYZY,
		ZYZ: dcmToAnglesZYZ,
		XZX: dcmToAnglesXZX,
		YXY: dcmToAnglesYXY,
		YZX: dcmToAnglesYZX,
		ZXY: dcmToAnglesZXY,
		XZY: dcmToAnglesXZY,
		ZYX: dcmToAnglesZYX,
		YXZ: dcmToAnglesYXZ,
	}

	angles, e := anglesMap[sequence](m)

	return angles, e
}
func Integrate(angles Angles, angleRates Angles, dt float64) (Angles, error) {
	if angles.Sequence != angleRates.Sequence {
		return Angles{}, fmt.Errorf("Can not integrate: %s != %s", angles.Sequence, angleRates.Sequence)
	}
	phiNew := angles.Phi + (angleRates.Phi * dt)
	thetaNew := angles.Theta + (angleRates.Theta * dt)
	siNew := angles.Si + (angleRates.Si * dt)

	return Angles{Phi: phiNew, Theta: thetaNew, Si: siNew, Sequence: angles.Sequence}, nil
}
func LinearInterpolate(startAngles, endAngles Angles, t float64) (Angles, error) {
	if startAngles.Sequence != endAngles.Sequence {
		return Angles{}, fmt.Errorf("Sequence %s != %s", startAngles.Sequence, endAngles.Sequence)
	}
	if t < 0.0 {
		return startAngles, nil
	}
	if t > 1.0 {
		return endAngles, nil
	}
	phiNew := (1-t)*startAngles.Phi + t*endAngles.Phi
	thetaNew := (1-t)*startAngles.Theta + t*endAngles.Theta
	siNew := (1-t)*startAngles.Si + t*endAngles.Si

	return Angles{Phi: phiNew, Theta: thetaNew, Si: siNew, Sequence: startAngles.Sequence}, nil
}

func SmoothInterpolate(startAngles, endAngles Angles, t float64) (Angles, error) {
	if startAngles.Sequence != endAngles.Sequence {
		return Angles{}, fmt.Errorf("Sequence %s != %s", startAngles.Sequence, endAngles.Sequence)
	}
	if t < 0.0 {
		return startAngles, nil
	}
	if t > 1.0 {
		return endAngles, nil
	}
	t2 := t * t
	t3 := t2 * t
	t4 := t3 * t
	t5 := t4 * t

	deltaPhi := endAngles.Phi - startAngles.Phi
	deltaTheta := endAngles.Theta - startAngles.Theta
	deltaSi := endAngles.Si - startAngles.Si

	phiNew := 6*deltaPhi*t5 + -15*deltaPhi*t4 + 10*deltaPhi*t3 + startAngles.Phi
	thetaNew := 6*deltaTheta*t5 + -15*deltaTheta*t4 + 10*deltaTheta*t3 + startAngles.Theta
	siNew := 6*deltaSi*t5 + -15*deltaSi*t4 + 10*deltaSi*t3 + startAngles.Si

	return Angles{Phi: phiNew, Theta: thetaNew, Si: siNew, Sequence: startAngles.Sequence}, nil
}

func dcmToAnglesXYZ(dcm matrix.Matrix) (Angles, error) {
	phi := math.Atan2(dcm.M23, dcm.M33)
	theta := -math.Asin(dcm.M13)
	psi := math.Atan2(dcm.M12, dcm.M11)

	return Angles{Phi: phi, Theta: theta, Si: psi, Sequence: XYZ}, nil
}
func dcmToAnglesZXZ(dcm matrix.Matrix) (Angles, error) {
	phi := math.Atan2(dcm.M13, dcm.M23)
	theta := math.Acos(dcm.M33)
	psi := math.Atan2(dcm.M31, -dcm.M32)

	return Angles{Phi: phi, Theta: theta, Si: psi, Sequence: ZXZ}, nil
}
func dcmToAnglesXYX(dcm matrix.Matrix) (Angles, error) {
	phi := math.Atan2(dcm.M21, dcm.M31)
	theta := math.Acos(dcm.M11)
	psi := math.Atan2(dcm.M12, -dcm.M13)

	return Angles{Phi: phi, Theta: theta, Si: psi, Sequence: XYX}, nil
}
func dcmToAnglesYZY(dcm matrix.Matrix) (Angles, error) {
	phi := math.Atan2(dcm.M32, dcm.M12)
	theta := math.Acos(dcm.M22)
	psi := math.Atan2(dcm.M23, -dcm.M21)
	return Angles{Phi: phi, Theta: theta, Si: psi, Sequence: YZY}, nil
}
func dcmToAnglesZYZ(dcm matrix.Matrix) (Angles, error) {
	phi := math.Atan2(dcm.M23, -dcm.M13)
	theta := math.Acos(dcm.M33)
	psi := math.Atan2(dcm.M32, dcm.M31)
	return Angles{Phi: phi, Theta: theta, Si: psi, Sequence: ZYZ}, nil
}
func dcmToAnglesXZX(dcm matrix.Matrix) (Angles, error) {
	phi := math.Atan2(dcm.M31, -dcm.M21)
	theta := math.Acos(dcm.M11)
	psi := math.Atan2(dcm.M13, dcm.M12)

	return Angles{Phi: phi, Theta: theta, Si: psi, Sequence: XZX}, nil
}
func dcmToAnglesYXY(dcm matrix.Matrix) (Angles, error) {
	phi := math.Atan2(dcm.M12, -dcm.M32)
	theta := math.Acos(dcm.M22)
	psi := math.Atan2(dcm.M21, dcm.M23)

	return Angles{Phi: phi, Theta: theta, Si: psi, Sequence: YXY}, nil
}
func dcmToAnglesYZX(dcm matrix.Matrix) (Angles, error) {
	phi := math.Atan2(dcm.M31, dcm.M11)
	theta := -math.Asin(dcm.M21)
	psi := math.Atan2(dcm.M23, dcm.M22)

	return Angles{Phi: phi, Theta: theta, Si: psi, Sequence: YZX}, nil
}
func dcmToAnglesZXY(dcm matrix.Matrix) (Angles, error) {
	phi := math.Atan2(dcm.M12, dcm.M22)
	theta := -math.Asin(dcm.M32)
	psi := math.Atan2(dcm.M31, dcm.M33)

	return Angles{Phi: phi, Theta: theta, Si: psi, Sequence: ZXY}, nil
}
func dcmToAnglesXZY(dcm matrix.Matrix) (Angles, error) {
	phi := math.Atan2(-dcm.M32, dcm.M22)
	theta := math.Asin(dcm.M12)
	psi := math.Atan2(-dcm.M13, dcm.M11)

	return Angles{Phi: phi, Theta: theta, Si: psi, Sequence: XZY}, nil
}
func dcmToAnglesZYX(dcm matrix.Matrix) (Angles, error) {
	phi := math.Atan2(-dcm.M21, dcm.M11)
	theta := math.Asin(dcm.M31)
	psi := math.Atan2(-dcm.M32, dcm.M33)

	return Angles{Phi: phi, Theta: theta, Si: psi, Sequence: ZYX}, nil
}
func dcmToAnglesYXZ(dcm matrix.Matrix) (Angles, error) {
	phi := math.Atan2(-dcm.M13, dcm.M33)
	theta := math.Asin(dcm.M23)
	psi := math.Atan2(-dcm.M21, dcm.M22)

	return Angles{Phi: phi, Theta: theta, Si: psi, Sequence: YXZ}, nil
}
func (angles Angles) ratesMatrixXYZ(bodyRates vector.Vector) (Angles, error) {

	cosPhi := math.Cos(angles.Phi)
	sinPhi := math.Sin(angles.Phi)
	cosTheta := math.Cos(angles.Theta)
	tanTheta := math.Tan(angles.Theta)
	secTheta := 1.0 / cosTheta

	data := [][]float64{
		{1.0, sinPhi * tanTheta, cosPhi * tanTheta},
		{0.0, cosPhi, -sinPhi},
		{0.0, sinPhi * secTheta, cosPhi * secTheta},
	}
	m, _ := matrix.New(data)
	rates, _ := m.Vop("*", bodyRates)

	return Angles{Phi: rates.X, Theta: rates.Y, Si: rates.Z, Sequence: angles.Sequence}, nil
}
func (angles Angles) ratesMatrixZXZ(bodyRates vector.Vector) (Angles, error) {

	cosPhi := math.Cos(angles.Phi)
	sinPhi := math.Sin(angles.Phi)
	cosTheta := math.Cos(angles.Theta)
	sinTheta := math.Sin(angles.Theta)
	cscTheta := 1.0 / sinTheta

	data := [][]float64{
		{-sinPhi * cosTheta * cscTheta, -cosPhi * cosTheta * cscTheta, sinTheta * cscTheta},
		{cosPhi * sinTheta * cscTheta, -sinPhi * sinTheta * cscTheta, 0.0},
		{sinPhi * cscTheta, cosPhi * cscTheta, 0.0},
	}

	m, _ := matrix.New(data)
	rates, _ := m.Vop("*", bodyRates)

	return Angles{Phi: rates.X, Theta: rates.Y, Si: rates.Z, Sequence: angles.Sequence}, nil
}
func (angles Angles) ratesMatrixYZY(bodyRates vector.Vector) (Angles, error) {

	cosPhi := math.Cos(angles.Phi)
	sinPhi := math.Sin(angles.Phi)
	cosTheta := math.Cos(angles.Theta)
	sinTheta := math.Sin(angles.Theta)
	cscTheta := 1.0 / sinTheta

	data := [][]float64{
		{-cosPhi * cosTheta * cscTheta, sinTheta * cscTheta, -sinPhi * cosTheta * cscTheta},
		{-sinPhi * sinTheta * cscTheta, 0.0, cosPhi * sinTheta * cscTheta},
		{cosPhi * cscTheta, 0.0, sinPhi * cscTheta},
	}

	m, _ := matrix.New(data)
	rates, _ := m.Vop("*", bodyRates)

	return Angles{Phi: rates.X, Theta: rates.Y, Si: rates.Z, Sequence: angles.Sequence}, nil
}
func (angles Angles) ratesMatrixZYZ(bodyRates vector.Vector) (Angles, error) {

	cosPhi := math.Cos(angles.Phi)
	sinPhi := math.Sin(angles.Phi)
	cosTheta := math.Cos(angles.Theta)
	sinTheta := math.Sin(angles.Theta)
	cscTheta := 1.0 / sinTheta

	data := [][]float64{
		{cosPhi * cosTheta * cscTheta, -sinPhi * cosTheta * cscTheta, sinTheta},
		{sinPhi * sinTheta * cscTheta, cosPhi * sinTheta * cscTheta, 0.0},
		{-cosPhi * cscTheta, sinPhi * cscTheta, 0.0},
	}

	m, _ := matrix.New(data)
	rates, _ := m.Vop("*", bodyRates)

	return Angles{Phi: rates.X, Theta: rates.Y, Si: rates.Z, Sequence: angles.Sequence}, nil
}
func (angles Angles) ratesMatrixXYX(bodyRates vector.Vector) (Angles, error) {

	cosPhi := math.Cos(angles.Phi)
	sinPhi := math.Sin(angles.Phi)
	cosTheta := math.Cos(angles.Theta)
	sinTheta := math.Sin(angles.Theta)
	cscTheta := 1.0 / sinTheta

	data := [][]float64{
		{sinTheta * cscTheta, -sinPhi * cosTheta * cscTheta, -cosPhi * cosTheta * cscTheta},
		{0.0, cosPhi * sinTheta * cscTheta, -sinPhi * sinTheta * cscTheta},
		{0.0, sinPhi * cscTheta, cosPhi * cscTheta},
	}

	m, _ := matrix.New(data)
	rates, _ := m.Vop("*", bodyRates)

	return Angles{Phi: rates.X, Theta: rates.Y, Si: rates.Z, Sequence: angles.Sequence}, nil
}
func (angles Angles) ratesMatrixXZX(bodyRates vector.Vector) (Angles, error) {

	cosPhi := math.Cos(angles.Phi)
	sinPhi := math.Sin(angles.Phi)
	cosTheta := math.Cos(angles.Theta)
	sinTheta := math.Sin(angles.Theta)
	cscTheta := 1.0 / sinTheta

	data := [][]float64{
		{sinTheta * cscTheta, cosPhi * cosTheta * cscTheta, -sinPhi * cosTheta * cscTheta},
		{0.0, sinPhi * sinTheta * cscTheta, cosPhi * sinTheta * cscTheta},
		{0.0, -cosPhi * cscTheta, sinPhi * cscTheta},
	}

	m, _ := matrix.New(data)
	rates, _ := m.Vop("*", bodyRates)

	return Angles{Phi: rates.X, Theta: rates.Y, Si: rates.Z, Sequence: angles.Sequence}, nil
}
func (angles Angles) ratesMatrixYXY(bodyRates vector.Vector) (Angles, error) {

	cosPhi := math.Cos(angles.Phi)
	sinPhi := math.Sin(angles.Phi)
	cosTheta := math.Cos(angles.Theta)
	sinTheta := math.Sin(angles.Theta)
	cscTheta := 1.0 / sinTheta

	data := [][]float64{
		{-sinPhi * cosTheta * cscTheta, sinTheta * cscTheta, cosPhi * cosTheta * cscTheta},
		{sinTheta * cosPhi * cscTheta, 0.0, sinTheta * sinPhi * cscTheta},
		{sinPhi * cscTheta, 0.0, -cosPhi * cscTheta},
	}

	m, _ := matrix.New(data)
	rates, _ := m.Vop("*", bodyRates)

	return Angles{Phi: rates.X, Theta: rates.Y, Si: rates.Z, Sequence: angles.Sequence}, nil
}
func (angles Angles) ratesMatrixYZX(bodyRates vector.Vector) (Angles, error) {

	cosPhi := math.Cos(angles.Phi)
	sinPhi := math.Sin(angles.Phi)
	cosTheta := math.Cos(angles.Theta)
	sinTheta := math.Sin(angles.Theta)
	secTheta := 1.0 / cosTheta

	data := [][]float64{
		{cosPhi * sinTheta * secTheta, cosTheta * secTheta, sinPhi * sinTheta * secTheta},
		{-sinPhi * cosTheta * secTheta, 0.0, cosPhi * cosTheta * secTheta},
		{cosPhi * secTheta, 0.0, sinPhi * secTheta},
	}

	m, _ := matrix.New(data)
	rates, _ := m.Vop("*", bodyRates)

	return Angles{Phi: rates.X, Theta: rates.Y, Si: rates.Z, Sequence: angles.Sequence}, nil
}
func (angles Angles) ratesMatrixZXY(bodyRates vector.Vector) (Angles, error) {

	cosPhi := math.Cos(angles.Phi)
	sinPhi := math.Sin(angles.Phi)
	cosTheta := math.Cos(angles.Theta)
	sinTheta := math.Sin(angles.Theta)
	secTheta := 1.0 / cosTheta

	data := [][]float64{
		{sinPhi * sinTheta * secTheta, cosPhi * sinTheta * secTheta, cosTheta * secTheta},
		{cosTheta * cosPhi * secTheta, -sinPhi * cosTheta * secTheta, 0.0},
		{sinPhi * secTheta, cosPhi * secTheta, 0.0},
	}

	m, _ := matrix.New(data)
	rates, _ := m.Vop("*", bodyRates)

	return Angles{Phi: rates.X, Theta: rates.Y, Si: rates.Z, Sequence: angles.Sequence}, nil
}
func (angles Angles) ratesMatrixXZY(bodyRates vector.Vector) (Angles, error) {

	cosPhi := math.Cos(angles.Phi)
	sinPhi := math.Sin(angles.Phi)
	cosTheta := math.Cos(angles.Theta)
	sinTheta := math.Sin(angles.Theta)
	secTheta := 1.0 / cosTheta

	data := [][]float64{
		{cosTheta * secTheta, -cosPhi * sinTheta * secTheta, sinPhi * sinTheta * secTheta},
		{0.0, sinPhi * cosTheta * secTheta, cosPhi * cosTheta * secTheta},
		{0.0, cosPhi * secTheta, -sinPhi * secTheta},
	}

	m, _ := matrix.New(data)
	rates, _ := m.Vop("*", bodyRates)

	return Angles{Phi: rates.X, Theta: rates.Y, Si: rates.Z, Sequence: angles.Sequence}, nil
}
func (angles Angles) ratesMatrixZYX(bodyRates vector.Vector) (Angles, error) {

	cosPhi := math.Cos(angles.Phi)
	sinPhi := math.Sin(angles.Phi)
	cosTheta := math.Cos(angles.Theta)
	sinTheta := math.Sin(angles.Theta)
	secTheta := 1.0 / cosTheta

	data := [][]float64{
		{-cosPhi * sinTheta * secTheta, sinPhi * sinTheta * secTheta, cosTheta * secTheta},
		{sinPhi * cosTheta * secTheta, cosPhi * cosTheta * secTheta, 0.0},
		{cosPhi * secTheta, -sinPhi * secTheta, 0.0},
	}

	m, _ := matrix.New(data)
	rates, _ := m.Vop("*", bodyRates)

	return Angles{Phi: rates.X, Theta: rates.Y, Si: rates.Z, Sequence: angles.Sequence}, nil
}
func (angles Angles) ratesMatrixYXZ(bodyRates vector.Vector) (Angles, error) {

	cosPhi := math.Cos(angles.Phi)
	sinPhi := math.Sin(angles.Phi)
	cosTheta := math.Cos(angles.Theta)
	sinTheta := math.Sin(angles.Theta)
	secTheta := 1.0 / cosTheta

	data := [][]float64{
		{sinPhi * sinTheta * secTheta, cosTheta * secTheta, -cosPhi * sinTheta * secTheta},
		{cosPhi * cosTheta * secTheta, 0.0, sinPhi * cosTheta * secTheta},
		{-sinPhi * secTheta, 0.0, cosPhi * secTheta},
	}

	m, _ := matrix.New(data)
	rates, _ := m.Vop("*", bodyRates)

	return Angles{Phi: rates.X, Theta: rates.Y, Si: rates.Z, Sequence: angles.Sequence}, nil
}
