package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

type exceptionStep struct {
	step
}

func newExceptionStep(name, exe, message, proj string, args []string) exceptionStep {
	s := exceptionStep{}
	s.step = newStep(name, exe, message, proj, args)
	return s
}

func (s exceptionStep) execute() (string, error) {
	cmd := exec.Command(s.exe, s.args...)
	cmd.Dir = s.proj
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", &stepErr{
			step:  s.name,
			msg:   fmt.Sprintf("Failed to execute '%s'", s.name),
			cause: err,
		}
	}

	if out.Len() > 0 {
		return "", &stepErr{
			step:  s.name,
			msg:   fmt.Sprintf("Invalid format: %s", out.String()),
			cause: nil,
		}
	}

	return s.message, nil
}
