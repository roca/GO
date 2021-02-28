package vector

import (
	"fmt"
	"math"
)

// type Point struct{ X, Y, Z int }

// type Ray struct {
// 	V Vector
// 	P Point
// 	T float64 // theta
// }

type Vector struct{ X, Y, Z float64 }

func New(values ...interface{}) (Vector, error) {

	switch v := values[0].(type) {
	case float64:
		switch l := len(values); l {
		case 1:
			return Vector{v, v, v}, nil
		case 2:
			return Vector{}, fmt.Errorf("Could not create Vector type")
		case 3:
			return Vector{values[0].(float64), values[1].(float64), values[2].(float64)}, nil
		default:
			return Vector{}, fmt.Errorf("Could not create Vector type")
		}
	case []float64:
		switch l := len(v); l {
		case 3:
			return Vector{v[0], v[1], v[2]}, nil
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
func (v *Vector) Negative() {
	v.Sop("*=", -1.0)
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
func Cross(l Vector, r Vector) *Vector {
	x := (l.Y * r.Z) - (l.Z * r.Y)
	y := (l.Z * r.X) - (l.X * r.Z)
	z := (l.X * r.Y) - (l.Y * r.X)
	return &Vector{x, y, z}
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

// Vector Operations

// func (u Vector) Divide(v Vector) Vector   { return Vector{u.X / v.X, u.Y / v.Y, u.Z / v.Z} }
// func (u Vector) MultMag(n float64) Vector { return Vector{u.X * n, u.Y * n, u.Z * n} }

// func (u Point) Add(v Vector) Vector {
// 	return Vector{float64(u.X) + v.X, float64(u.Y) + v.Y, float64(u.Z) + v.Z}
// }
// func (u Point) Subtract(v Point) Vector {
// 	return Vector{float64(u.X - v.X), float64(u.Y - v.Y), float64(u.Z - v.Z)}
// }

// func triangulate(input string) {
// 	var r1, r2 Ray
// 	fmt.Sscanf(input, "%d %d %f\n%d %d %f", &r1.P.X, &r1.P.Y, &r1.T, &r2.P.X, &r2.P.Y, &r2.T)

// 	// get unit vectors v.x and v.y using the provided angle theta
// 	r1.V = Vector{math.Sin(r1.T * math.Pi / 180), math.Cos(r1.T * math.Pi / 180), 0}
// 	r2.V = Vector{math.Sin(r2.T * math.Pi / 180), math.Cos(r2.T * math.Pi / 180), 0}

// 	// p3 = p1 + a*v1
// 	// p3 = p2 + b*v2
// 	// setting equations to equal each other and solving for a:
// 	// a = mag[ (p2-p1) cross v2 ] / mag[ v1 cross v2 ]
// 	// where a is magnitude of distance between r1.P and p3 (intersection point)

// 	a := r2.P.Subtract(r1.P).Cross(r2.V).Mag() / r1.V.Cross(r2.V).Mag()
// 	p3 := r1.P.Add(r1.V.MultMag(a))

// 	fmt.Printf("(%4.1f, %4.1f)\n", p3.X, p3.Y)
// }
