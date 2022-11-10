package main

import "errors"

var (
	ErrNotNumber        = errors.New("ata is not numeric")
	ErrInvalidColumn    = errors.New("Invalid column number")
	ErrNoFiles          = errors.New("No input files")
	ErrInvalidOperation = errors.New("Invalid operation")
)
