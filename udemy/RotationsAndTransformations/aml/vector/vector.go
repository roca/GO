package vector

import (
	"fmt"
	"math"
)

type Vector struct {
	X, Y, Z float64
}

func New(values ...interface{}) (Vector, error) {

	switch values[0].(type) {
	case float64:
		switch l := len(values); l {
		case 1:
			return Vector{values[0].(float64), values[0].(float64), values[0].(float64)}, nil
		case 3:
			return Vector{values[0].(float64), values[1].(float64), values[2].(float64)}, nil
		default:
			return Vector{}, fmt.Errorf("Could not create Vector type")
		}
	case []float64:
		a := values[0].([]float64)
		switch l := len(a); l {
		case 3:
			return Vector{a[0], a[1], a[2]}, nil
		default:
			return Vector{}, fmt.Errorf("Could not create Vector type")
		}
	default:
		return Vector{}, nil
	}
}
func NewX() Vector { return Vector{1.0, 0.0, 0.0} }
func NewY() Vector { return Vector{0.0, 1.0, 0.0} }
func NewZ() Vector { return Vector{0.0, 0.0, 1.0} }
func (v *Vector) Data() []float64 {
	return []float64{v.X, v.Y, v.Z}
}
func (v *Vector) Negative() (Vector, error) {
	u, _ := v.Copy()
	u.Sop("*=", -1.0)
	return u, nil
}

func (m *Vector) Copy() (Vector, error) {
	dataV := m.Data()
	u, _ := New(dataV)
	return u, nil
}

func (v *Vector) Sop(operation string, value float64) (*Vector, error) {
	switch o := operation; {
	case o == "+=" || o == "+":
		v.X += value
		v.Y += value
		v.Z += value
		return v, nil
	case o == "-=" || o == "-":
		v.X -= value
		v.Y -= value
		v.Z -= value
		return v, nil
	case o == "*=" || o == "*":
		v.X *= value
		v.Y *= value
		v.Z *= value
		return v, nil
	case o == "/=" || o == "/":
		v.X /= value
		v.Y /= value
		v.Z /= value
		return v, nil
	default:
		return &Vector{}, fmt.Errorf("Vector has no such operation '%s'", o)
	}
}
func (v *Vector) Vop(operation string, u Vector) (*Vector, error) {
	switch o := operation; {
	case o == "+=" || o == "+":
		v.X += u.X
		v.Y += u.Y
		v.Z += u.Z
		return v, nil
	case o == "-=" || o == "-":
		v.X -= u.X
		v.Y -= u.Y
		v.Z -= u.Z
		return v, nil
	case o == "*=" || o == "*":
		v.X *= u.X
		v.Y *= u.Y
		v.Z *= u.Z
		return v, nil
	case o == "/=" || o == "/":
		v.X /= u.X
		v.Y /= u.Y
		v.Z /= u.Z
		return v, nil
	default:
		return &Vector{}, fmt.Errorf("Vector has no such operation '%s'", o)
	}
}
func (u Vector) Mag() float64 {
	return math.Sqrt((u.X * u.X) + (u.Y * u.Y) + (u.Z * u.Z))
}
func (u Vector) Norm() float64 { return u.Mag() }
func (v *Vector) Normalize() (*Vector, error) {
	mag := v.Mag()
	if mag == 0 {
		return &Vector{}, fmt.Errorf("Can't normalize this vector the magnitdue was 0")
	}
	v.X /= mag
	v.Y /= mag
	v.Z /= mag
	return v, nil
}
func Cross(l Vector, r Vector) Vector {
	x := (l.Y * r.Z) - (l.Z * r.Y)
	y := (l.Z * r.X) - (l.X * r.Z)
	z := (l.X * r.Y) - (l.Y * r.X)
	return Vector{x, y, z}
}
func Dot(u Vector, v Vector) float64 { return (u.X * v.X) + (u.Y * v.Y) + (u.Z * v.Z) }
func Unit(u Vector) *Vector {
	mag := u.Norm()
	if mag > 0.0 {
		v, _ := u.Sop("/=", mag)
		return v
	}

	return &u
}
