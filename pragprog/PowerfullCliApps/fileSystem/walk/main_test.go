package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name     string
		root     string
		cfg      config
		expected string
	}{
		{name: "NoFilter", root: "testdata", cfg: config{ext: "", size: 0, list: true}, expected: "testdata/dir.log\ntestdata/dir2/script.sh\n"},
		{name: "FilterExtensionMatch", root: "testdata", cfg: config{ext: ".log", size: 0, list: true}, expected: "testdata/dir.log\n"},
		{name: "FilterExtensionSizeMatch", root: "testdata", cfg: config{ext: ".log", size: 10, list: true}, expected: "testdata/dir.log\n"},
		{name: "FilterExtensioneNoMatch", root: "testdata", cfg: config{ext: ".gz", size: 0, list: true}, expected: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer
			if err := run(tc.root, &buffer, tc.cfg); err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			res := buffer.String()
			if res != tc.expected {
				t.Errorf("Expected %s, got %s instead.\n", tc.expected, res)
			}

		})
	}
}

func TestRunDelExtension(t *testing.T) {
	testCases := []struct{
		name string
		cfg config
		extNoDelete string
		nDelete int
		nNoDelete int
		expected string
	}{
		{ name: "DeleteExtensionNoMatch", cfg: config{ext: ".log", del: true}, extNoDelete: ".gz", nDelete: 0, nNoDelete: 10, expected: ""},
		{ name: "DeleteExtensionMatch", cfg: config{ext: ".log", del: true}, extNoDelete: "", nDelete: 10, nNoDelete: 0, expected: ""},
		{ name: "DeleteExtensionMxed", cfg: config{ext: ".log", del: true}, extNoDelete: ".gz", nDelete: 5, nNoDelete: 5, expected: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var (
				buffer bytes.Buffer
				logBuffer bytes.Buffer
			)
			tc.cfg.wLog = &logBuffer

			tempDir, cleanup := createTempDir(t, map[string]int{
				tc.cfg.ext: tc.nDelete,
				tc.extNoDelete: tc.nNoDelete,
			})
			defer cleanup()

			if err := run(tempDir, &buffer, tc.cfg); err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			res := buffer.String()

			if res != tc.expected {
				t.Errorf("Expected %s, got %s instead.\n", tc.expected, res)
			}

			filesLeft, err := ioutil.ReadDir(tempDir)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if len(filesLeft) != tc.nNoDelete {
				t.Errorf("Expected %d files left, got %d instead.\n", tc.nNoDelete, len(filesLeft))
			}

			expLogLines := tc.nDelete + 1
			lines := bytes.Split(logBuffer.Bytes(), []byte{'\n'})
			if len(lines) != expLogLines {
				t.Errorf("Expected %d log lines, got %d instead.\n", expLogLines, len(lines))
			}
		})
	}
}

func createTempDir(t *testing.T, files map[string]int) (dirname string, cleanup func()) {
	t.Helper()

	tempDir, err := ioutil.TempDir("", "walk")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	for k, n := range files {
		for j := 1; j <= n; j++ {
			fname := fmt.Sprintf("file%d%s", j, k)
			fpath := filepath.Join(tempDir, fname)
			if err := ioutil.WriteFile(fpath, []byte("dummy"), 0644); err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
		}
	}
	return tempDir, func() { os.RemoveAll(tempDir) }
}
