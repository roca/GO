package euler

import (
	"fmt"

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
func NewFromDcm(m matrix.Matrix, sequence seq) (Angles, error) {
	if !dcm.IsOrthogonal(m) {
		return Angles{}, fmt.Errorf("This Matrix is not orthoganal %g", m)
	}

	anglesMap := map[seq]func(matrix.Matrix) (Angles, error){
		ZXZ: dcmToAnglesZXZ,
		XYX: dcmToAnglesXYX,
		YZY: dcmToAnglesYZY,
		ZYZ: dcmToAnglesZYZ,
		XZX: dcmToAnglesXZX,
		YXY: dcmToAnglesYXY,
		XYZ: dcmToAnglesXYZ,
		YZX: dcmToAnglesYZX,
		ZXY: dcmToAnglesZXY,
		XZY: dcmToAnglesXZY,
		ZYX: dcmToAnglesZYX,
		YXZ: dcmToAnglesYXZ,
	}

	angles, e := anglesMap[sequence](m)

	return angles, e
}

func (a Angles) ToDCM() (matrix.Matrix, error) {
	dcmMap := map[seq]func() (matrix.Matrix, error){
		ZXZ: a.toDCMZXZ,
		XYX: a.toDCMXYX,
		YZY: a.toDCMYZY,
		ZYZ: a.toDCMZYZ,
		XZX: a.toDCMXZX,
		YXY: a.toDCMYXY,
		XYZ: a.toDCMXYZ,
		YZX: a.toDCMYZX,
		ZXY: a.toDCMZXY,
		XZY: a.toDCMXZY,
		ZYX: a.toDCMZYX,
		YXZ: a.toDCMYXZ,
	}

	dcm, e := dcmMap[a.Sequence]()

	return dcm, e
}
func (a *Angles) Convert(sequence seq) {
	a.Sequence = sequence
}

func (a Angles) toDCMZXZ() (matrix.Matrix, error) { return matrix.Matrix{}, nil }
func (a Angles) toDCMXYX() (matrix.Matrix, error) { return matrix.Matrix{}, nil }
func (a Angles) toDCMYZY() (matrix.Matrix, error) { return matrix.Matrix{}, nil }
func (a Angles) toDCMZYZ() (matrix.Matrix, error) { return matrix.Matrix{}, nil }
func (a Angles) toDCMXZX() (matrix.Matrix, error) { return matrix.Matrix{}, nil }
func (a Angles) toDCMYXY() (matrix.Matrix, error) { return matrix.Matrix{}, nil }
func (a Angles) toDCMXYZ() (matrix.Matrix, error) {
	Rx, _ := dcm.RotationX(a.Phi)
	Ry, _ := dcm.RotationY(a.Theta)
	Rz, _ := dcm.RotationZ(a.Si)

	Ryz, _ := Ry.Mop("*", Rz)
	Rxyz, _ := Rx.Mop("*", *Ryz)

	return *Rxyz, nil
}
func (a Angles) toDCMYZX() (matrix.Matrix, error) { return matrix.Matrix{}, nil }
func (a Angles) toDCMZXY() (matrix.Matrix, error) { return matrix.Matrix{}, nil }
func (a Angles) toDCMXZY() (matrix.Matrix, error) { return matrix.Matrix{}, nil }
func (a Angles) toDCMZYX() (matrix.Matrix, error) { return matrix.Matrix{}, nil }
func (a Angles) toDCMYXZ() (matrix.Matrix, error) { 
	Ry, _ := dcm.RotationX(a.Phi)
	Rx, _ := dcm.RotationY(a.Theta)
	Rz, _ := dcm.RotationZ(a.Si)

	Rxz, _ := Rx.Mop("*", Rz)
	Ryxz, _ := Ry.Mop("*", *Rxz)

	return *Ryxz, nil
}

func dcmToAnglesZXZ(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
func dcmToAnglesXYX(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
func dcmToAnglesYZY(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
func dcmToAnglesZYZ(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
func dcmToAnglesXZX(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
func dcmToAnglesYXY(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
func dcmToAnglesXYZ(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
func dcmToAnglesYZX(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
func dcmToAnglesZXY(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
func dcmToAnglesXZY(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
func dcmToAnglesZYX(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
func dcmToAnglesYXZ(dcm matrix.Matrix) (Angles, error) { return Angles{}, nil }
