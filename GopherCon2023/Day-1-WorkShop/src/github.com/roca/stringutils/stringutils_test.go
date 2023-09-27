package stringutils

import (
	"os"
	"testing"
)

var tests map[string]struct {
	input string
	want  string
}

func setup() {
	tests = map[string]struct {
		input string
		want  string
	}{
		"hello":       {"hello", "HELLO"},
		"HELLO":       {"HELLO", "HELLO"},
		"hello world": {"hello world", "HELLO WORLD"},
		"one":         {input: "Johnny", want: "JOHNNY"},
		"two":         {input: "世界", want: "世界"},
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	// teardown
	os.Exit(code)
}

func TestUpper(t *testing.T) {
	s := "hello"
	got := Upper(s)
	want := "HELLO"
	if got != want {
		t.Errorf("Upper(%q) = %q, want %q", s, got, want)
	}
}

func TestUpperMultipleCases(t *testing.T) {

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := Upper(test.input)
			if got != test.want {
				t.Errorf("Upper(%q) = %q, want %q", test.input, got, test.want)
			}
		})
	}
}

func BenchmarkUpper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Upper("hello")
	}
}
