package main

import (
	"bytes"
	"errors"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name   string
		col    int
		op     string
		exp    string
		files  []string
		expErr error
	}{
		{name: "RunAvgFile", col: 3, op: "avg", exp: "227.6\n",
			files:  []string{"./testdata/example.csv"},
			expErr: nil},
		{name: "RunAvgMultiFiles", col: 3, op: "avg", exp: "233.84\n",
			files:  []string{"./testdata/example.csv", "./testdata/example2.csv"},
			expErr: nil},
		{name: "RunFailedRead", col: 2, op: "avg", exp: "",
			files:  []string{"./testdata/example.csv", "./testdata/fakefile.csv"},
			expErr: os.ErrNotExist},
		{name: "RunFailColumn", col: 0, op: "avg", exp: "",
			files:  []string{"./testdata/example.csv"},
			expErr: ErrInvalidColumn},
		{name: "RunFailNoFiles", col: 2, op: "avg", exp: "",
			files:  []string{},
			expErr: ErrNoFiles},
		{name: "RunFailOperation", col: 2, op: "invalid", exp: "",
			files:  []string{"./testdata/example.csv"},
			expErr: ErrInvalidOperation},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var res bytes.Buffer
			err := run(tc.files, tc.op, tc.col, &res)

			if tc.expErr != nil {
				if err == nil {
					t.Errorf("Expected error %v, Got nil instead", tc.expErr)
				}
				if !errors.Is(err, tc.expErr) {
					t.Errorf("Expected error %q, got %q instead", tc.expErr, err)
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}

			if res.String() != tc.exp {
				t.Errorf("Expected %q, got %q instead", tc.exp, res.String())
			}
		})
	}
}
