package main

import (
	"os"
	"testing"
)

func TestFilterOut(t *testing.T) {
	testCases := []struct {
		name     string
		file     string
		ext      string
		minSize  int64
		expected bool
	}{
		{"FilterNoExtension", "testdata/dir.log", "", 0, false},
		{"FilterNoExtensionMatch", "testdata/dir.log", ".log", 0, false},
		{"FilterNoExtensionNoMatch", "testdata/dir.log", ".sh", 0, true},
		{"FilterNoExtensionSizeMatch", "testdata/dir.log", ".log", 10, false},
		{"FilterNoExtensionSizeNoMatch", "testdata/dir.log", ".log", 20, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			info, err := os.Stat(tc.file)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			res := filterOut(tc.file, tc.ext, tc.minSize, info)
			if res != tc.expected {
				t.Errorf("Expected %t, got %t instead.\n", tc.expected, res)
			}
		})
	}
}
