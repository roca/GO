package euler

import (
	"fmt"
	"math"
	"strings"

	"udemy.com/aml/dcm"
	"udemy.com/aml/matrix"
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
	Convert(angles IAngles, sequence seq)
}

type Angles struct {
	Sequence seq `default:"XYZ"`
	Phi      float64
	Theta    float64
	Si       float64
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
		// XZX: dcmToAnglesXZX,
		// YXY: dcmToAnglesYXY,
		// YZX: dcmToAnglesYZX,
		// ZXY: dcmToAnglesZXY,
		// XZY: dcmToAnglesXZY,
		// ZYX: dcmToAnglesZYX,
		// YXZ: dcmToAnglesYXZ,
	}

	angles, e := anglesMap[sequence](m)

	return angles, e
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
func (a *Angles) Convert(sequence seq) {
	a.Sequence = sequence
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

// func dcmToAnglesXZX(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
// func dcmToAnglesYXY(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
// func dcmToAnglesYZX(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
// func dcmToAnglesZXY(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
// func dcmToAnglesXZY(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
// func dcmToAnglesZYX(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
// func dcmToAnglesYXZ(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
