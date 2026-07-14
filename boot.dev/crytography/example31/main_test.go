package main

import (
	"encoding"
	"strings"
	"testing"
)

func Test_hasher(t *testing.T) {
	const (
		input1 = "The tunneling gopher digs downwards, "
		input2 = "unaware of what he will find."
	)
	first := newHasher()
	first.Write(input1)

	marshaler, ok := first.hash.(encoding.BinaryMarshaler)
	if !ok {
		t.Errorf("hash.Hash does not implement encoding.BinaryMarshaler")
	}

	state, err := marshaler.MarshalBinary()
	if err != nil {
		t.Errorf("MarshalBinary() failed: %v", err)
	}

	second := newHasher()

	unmarshaler, ok := second.hash.(encoding.BinaryUnmarshaler)
	if !ok {
		t.Errorf("hasher does not implement encoding.BinaryUnmarshaler")
	}

	if err := unmarshaler.UnmarshalBinary(state); err != nil {
		t.Errorf("UnmarshalBinary() failed: %v", err)
	}

	if !strings.Contains(string(state), input1) {
		t.Errorf("original state does not contains input1")
	}
}
