package vector

import (
	"fmt"
	"math"
	"strings"
)

type IVector interface {
	New(values ...interface{}) (Vector, error)
	Data() []float64
	Negative() (Vector, error)
	Copy() (Vector, error)
	CopyPointer() (*Vector, error)
	Sop(operation string, value float64) (*Vector, error)
	Vop(operation string, u Vector) (*Vector, error)
	Mag() float64
	Norm() float64
	Normalize() (*Vector, error)
	Cross(l Vector, r Vector) Vector
	Dot(u Vector, v Vector) float64
	Unit(u Vector) *Vector
	Mutate()
	change(values ...interface{}) error
}

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
	_, _ = u.Sop("*=", -1.0)
	return u, nil
}

func (m *Vector) Copy() (Vector, error) {
	u := m
	return *u, nil
}
func (m *Vector) CopyPointer() (*Vector, error) {
	var new *Vector
	x, _ := m.Copy()
	new = &x
	return new, nil
}

func (v *Vector) Mutate(values ...interface{}) {
	_ = v.change(values)
}

func (v *Vector) change(values ...interface{}) error {
	switch values[0].(type) {
	case float64:
		switch l := len(values); l {
		case 1:
			v.X = values[0].(float64)
			v.Y = values[0].(float64)
			v.Z = values[0].(float64)
			return nil
		case 3:
			v.X = values[0].(float64)
			v.Y = values[1].(float64)
			v.Z = values[2].(float64)
			return nil
		default:
			return fmt.Errorf("Could not create Vector type")
		}
	case []float64:
		a := values[0].([]float64)
		switch l := len(a); l {
		case 3:
			v.X = a[0]
			v.Y = a[1]
			v.Z = a[2]
			return nil
		default:
			return fmt.Errorf("Could not create Vector type")
		}
	default:
		return nil
	}
}

func (v *Vector) Sop(operation string, value float64) (*Vector, error) {
	var new *Vector
	if !strings.Contains(operation, "=") {
		new, _ = v.CopyPointer()
	} else {
		new = v
	}

	switch o := operation; {
	case o == "+=" || o == "+":
		new.X += value
		new.Y += value
		new.Z += value
		return new, nil
	case o == "-=" || o == "-":
		new.X -= value
		new.Y -= value
		new.Z -= value
		return new, nil
	case o == "*=" || o == "*":
		new.X *= value
		new.Y *= value
		new.Z *= value
		return new, nil
	case o == "/=" || o == "/":
		new.X /= value
		new.Y /= value
		new.Z /= value
		return new, nil
	default:
		return &Vector{}, fmt.Errorf("Vector has no such operation '%s'", o)
	}
}
func (v *Vector) Vop(operation string, u Vector) (*Vector, error) {
	var new *Vector
	if !strings.Contains(operation, "=") {
		new, _ = v.CopyPointer()
	} else {
		new = v
	}

	switch o := operation; {
	case o == "+=" || o == "+":
		new.X += u.X
		new.Y += u.Y
		new.Z += u.Z
		return new, nil
	case o == "-=" || o == "-":
		new.X -= u.X
		new.Y -= u.Y
		new.Z -= u.Z
		return new, nil
	case o == "*=" || o == "*":
		new.X *= u.X
		new.Y *= u.Y
		new.Z *= u.Z
		return new, nil
	case o == "/=" || o == "/":
		new.X /= u.X
		new.Y /= u.Y
		new.Z /= u.Z
		return new, nil
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
