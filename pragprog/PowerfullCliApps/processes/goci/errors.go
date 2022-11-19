package main

import (
	"errors"
	"fmt"
)

var (
	ErrValidation = errors.New("Validation failed")
)

type stepError struct {
	step  string
	msg   string
	cause error
}

func (s *stepError) Error() string {
	return fmt.Sprintf("Step: %q: %s: Cause: %v", s.step, s.msg, s.cause)
}

func (s *stepError) Is(target error) bool {
	t, ok := target.(*stepError)
	if !ok {
		return false
	}
	return t.step == s.step
}

func (s *stepError) Unwrap() error {
	return s.cause
}
