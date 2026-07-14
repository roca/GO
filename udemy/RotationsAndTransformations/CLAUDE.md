# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

A Go implementation of an Attitude Math Library (AML) for 3D rotations and transformations, following the Udemy course "Complete Guide to Rotations and Transformations". The `AttitudeMathLibrary/` directory holds the original C++ reference implementation; the Go port under `aml/` is the primary work.

All commands below run from this directory (where `go.mod` lives). The module path is `udemy.com` and requires Go 1.25.

Note: the `gonum.org/v1/gonum` and `gonum.org/v1/plot` dependencies must be fetchable to build `aml/aml.go` (the `MatPrint` helper) and the `*_plot`/`quat_example3` programs. The five core packages (`vector`, `matrix`, `dcm`, `euler`, `quaternion`) and the `test/` package have no gonum dependency and build/test on their own.

## Commands

```bash
# Run the full test suite (tests live in the separate `test` package)
go test ./test/

# Run a single test by name
go test ./test/ -run TestCase08 -v

# Build/run an example or plot program (each is its own package main)
go run ./example/
go run ./quat_plot/      # writes points.png into the program's directory

go build ./...           # compile everything
go vet ./...
```

## Architecture

The library is layered — each package builds on the ones below it, so a change to a lower layer ripples upward:

- `aml/vector` — `Vector{X,Y,Z}`. Foundation type with no internal deps. `Cross`, `Dot`, `Unit`, `Normalize`.
- `aml/matrix` — `Matrix` (3×3, flat `M11..M33` fields). Depends on `vector`. `Transpose`, `Determinant` (via Levi-Civita `epsilon`), `Inverse`, `Diag`.
- `aml/dcm` — Direction Cosine Matrices. Depends on `matrix` + `vector`. Axis rotations `RotationX/Y/Z`, orthogonality checks/repair (`IsOrthogonal`, `Normalize`), kinematic rate matrices, integration.
- `aml/euler` — `Angles{Sequence, Phi, Theta, Si}`. Depends on `dcm`, `matrix`, `vector`. Supports all 12 rotation sequences (`XYZ`, `ZXZ`, ...) via per-sequence `dcmToAngles*` and `ratesMatrix*` functions dispatched through maps keyed by `Seq`.
- `aml/quaternion` — `Quaternion{S,X,Y,Z}`. Top layer, depends on all of the above. Conversions round-trip through DCMs (`Dcm2Quat`/`Quat2DCM`, `Angles2Quat`/`ToAngles`), plus `SlerpInterpolate`/`linearInterpolate`.

`aml/aml.go` holds only the `MatPrint` gonum helper.

### Cross-cutting conventions

- **Value semantics.** `Vector`, `Matrix`, and `Quaternion` are immutable value types. Arithmetic methods (`Add`, `Sub`, `Mul`, `Scale`, `Neg`, `Transpose`, `MulVec`, ...) take value receivers and return new values; they never mutate the receiver. Copy is free — just assign.
- **Typed constructors.** `vector.New(x, y, z)`, `vector.Splat(s)`, `vector.FromSlice([]float64)`; `matrix.New([3][3]float64)`, `matrix.FromRows(v0, v1, v2)`, `matrix.Splat`, `matrix.Identity`. `matrix.Data()` returns `[3][3]float64`.
- **Errors only where operations can fail.** Infallible operations (add, scale, transpose, rotation-matrix construction) return a single value. Errors are returned only for genuinely fallible operations: `Vector.Normalize`/`Matrix.Inverse` (degenerate input), and the conversions (`Dcm2Quat`, `Quat2DCM`, `ToDCM`, `DcmToAngles`, `Angles2Quat`).
- **Rotation matrices are `[R]ᵀ` form** (transpose of the rotation, i.e. the `-sin` sits in the upper triangle for `RotationX/Y/Z`).

### Tests and examples

- All tests are in the `test/` package (package `test`, not per-module `_test.go` files) and use `testify/assert`.
- The `example*/`, `*_plot/`, and `quat_*/` directories are each a standalone `package main` demonstrating a feature; the plot programs render PNGs using `gonum.org/v1/plot`.

## History

This Go port was modernized from Go 1.14 to Go 1.25. The original course code used stringly-typed operator methods (`Sop`/`Vop`/`Mop`/`Qop`, where `"*="` mutated and `"*"` copied) and `interface{}` variadic `New` constructors that type-switched at runtime; both were replaced with the typed, value-semantics API above. Several bugs preserved from the course exercises were fixed in the process: `quaternion.New` dropped its `z` argument, `Qop("*")` swapped the X/Y outputs of the Hamilton product, and `dcm.KinematicRatesFromWorldRates`/`Normalize2`/`ZXZEulerAngleRates` had incorrect math. The C++ reference in `AttitudeMathLibrary/AML/` remains the source of truth when precision matters.
