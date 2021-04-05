package euler

import "udemy.com/aml/matrix"

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
	New(phi, theta, si float64, sequence seq) (Angles, error)
	ToDCM(angles *Angles) matrix.Matrix
	ToAngles(dcm *matrix.Matrix, sequence seq) (Angles, error)
	Convert(angles *Angles, sequence seq) (Angles, error)
}

type Angles struct {
	Sequence seq
}
