package vector

import (
	"fmt"
	"math"
)

// Vector is a 3D vector with value semantics: all methods return new
// Vectors and never mutate the receiver.
type Vector struct {
	X, Y, Z float64
}

// New returns a Vector with the given components.
func New(x, y, z float64) Vector {
	return Vector{x, y, z}
}

// Splat returns a Vector with every component set to s.
func Splat(s float64) Vector {
	return Vector{s, s, s}
}

// FromSlice returns a Vector from a length-3 slice.
func FromSlice(a []float64) (Vector, error) {
	if len(a) != 3 {
		return Vector{}, fmt.Errorf("vector: need 3 values, got %d", len(a))
	}
	return Vector{a[0], a[1], a[2]}, nil
}

// Axis unit vectors.
func UnitX() Vector { return Vector{1, 0, 0} }
func UnitY() Vector { return Vector{0, 1, 0} }
func UnitZ() Vector { return Vector{0, 0, 1} }

// Data returns the components as a slice.
func (v Vector) Data() []float64 { return []float64{v.X, v.Y, v.Z} }

// Add returns v + u.
func (v Vector) Add(u Vector) Vector {
	return Vector{v.X + u.X, v.Y + u.Y, v.Z + u.Z}
}

// Sub returns v - u.
func (v Vector) Sub(u Vector) Vector {
	return Vector{v.X - u.X, v.Y - u.Y, v.Z - u.Z}
}

// Mul returns the elementwise product v * u.
func (v Vector) Mul(u Vector) Vector {
	return Vector{v.X * u.X, v.Y * u.Y, v.Z * u.Z}
}

// Div returns the elementwise quotient v / u.
func (v Vector) Div(u Vector) Vector {
	return Vector{v.X / u.X, v.Y / u.Y, v.Z / u.Z}
}

// Scale returns v scaled by s.
func (v Vector) Scale(s float64) Vector {
	return Vector{v.X * s, v.Y * s, v.Z * s}
}

// Neg returns -v.
func (v Vector) Neg() Vector {
	return Vector{-v.X, -v.Y, -v.Z}
}

// Mag returns the Euclidean magnitude of v.
func (v Vector) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Norm is an alias for Mag.
func (v Vector) Norm() float64 { return v.Mag() }

// Normalize returns the unit vector in the direction of v. It returns an
// error if v has zero magnitude.
func (v Vector) Normalize() (Vector, error) {
	mag := v.Mag()
	if mag == 0 {
		return Vector{}, fmt.Errorf("vector: cannot normalize zero-magnitude vector")
	}
	return v.Scale(1 / mag), nil
}

// Cross returns the cross product l × r.
func Cross(l, r Vector) Vector {
	return Vector{
		l.Y*r.Z - l.Z*r.Y,
		l.Z*r.X - l.X*r.Z,
		l.X*r.Y - l.Y*r.X,
	}
}

// Dot returns the dot product u · v.
func Dot(u, v Vector) float64 { return u.X*v.X + u.Y*v.Y + u.Z*v.Z }

// Unit returns the unit vector in the direction of u, or u unchanged if it
// has zero magnitude.
func Unit(u Vector) Vector {
	mag := u.Norm()
	if mag > 0 {
		return u.Scale(1 / mag)
	}
	return u
}
