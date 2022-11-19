package main

import (
	"bytes"
	"errors"
	"testing"
)

func TestRun(t *testing.T) {
	var testCases = []struct {
		name     string
		proj     string
		out      string
		expErr error
	}{
		{name: "success", proj: "./testdata/tool/", out: "Go Build: SUCCESS\n", expErr: nil},
		{name: "fail", proj: "./testdata/toolErr/", out: "", expErr: &stepErr{step: "go build"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var out bytes.Buffer
			err := run(tc.proj, &out)
			if tc.expErr != nil {
				if err == nil {
					t.Errorf("Expected error: %q, Got 'nil' instead.", tc.expErr)
				}
				if !errors.Is(err, tc.expErr) {
					t.Errorf("Expected error: %q, Got: %q", tc.expErr, err)
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %q", err)
			}

			if out.String() != tc.out {
				t.Errorf("Expected output: %q, Got: %q", tc.out, out.String())
			}
		})
	}
}
