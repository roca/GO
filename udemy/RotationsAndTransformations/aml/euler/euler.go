package euler

import (
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

func New(phi, theta, si float64, sequence seq) (Angles, error) {
	return Angles{Phi: phi, Theta: theta, Si: si, Sequence: sequence}, nil
}

// func NewFromDcm(m matrix.Matrix, sequence seq) (Angles, error) {
// 	if !dcm.IsOrthogonal(m) {
// 		return Angles{}, fmt.Errorf("This Matrix is not orthoganal %g", m)
// 	}

// 	anglesMap := map[seq]func(matrix.Matrix) (Angles, error){
// 		ZXZ: dcmToAnglesZXZ,
// 		XYX: dcmToAnglesXYX,
// 		YZY: dcmToAnglesYZY,
// 		ZYZ: dcmToAnglesZYZ,
// 		XZX: dcmToAnglesXZX,
// 		YXY: dcmToAnglesYXY,
// 		XYZ: dcmToAnglesXYZ,
// 		YZX: dcmToAnglesYZX,
// 		ZXY: dcmToAnglesZXY,
// 		XZY: dcmToAnglesXZY,
// 		ZYX: dcmToAnglesZYX,
// 		YXZ: dcmToAnglesYXZ,
// 	}

// 	angles, e := anglesMap[sequence](m)

// 	return angles, e
// }

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

// func dcmToAnglesZXZ(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
// func dcmToAnglesXYX(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
// func dcmToAnglesYZY(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
// func dcmToAnglesZYZ(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
// func dcmToAnglesXZX(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
// func dcmToAnglesYXY(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
// func dcmToAnglesXYZ(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
// func dcmToAnglesYZX(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
// func dcmToAnglesZXY(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
// func dcmToAnglesXZY(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
// func dcmToAnglesZYX(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
// func dcmToAnglesYXZ(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
