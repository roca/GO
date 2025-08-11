package nlp

import (
	"os"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestTokenize(t *testing.T) {
	text := "Who's on first?"
	tokens := Tokenize(text)

	expected := []string{"who", "s", "on", "first"}
	/*
		if !slices.Equal(expected, tokens) {
			t.Fatalf("expected %#v, got %#v", expected, tokens)
		}
	*/
	require.Equal(t, expected, tokens)
}

func TestTokenizeTable(t *testing.T) {
	var cases = []struct {
		text     string
		expected []string
	}{
		{"Who's on first?", []string{"who", "s", "on", "first"}},
		{"What's on second?", []string{"what", "s", "on", "second"}},
		{"", nil},
	}

	for _, tc := range cases {
		t.Run(tc.text, func(t *testing.T) {
			tokens := Tokenize(tc.text)
			/*
				if !slices.Equal(tc.expected, tokens) {
					t.Fatalf("expected %#v, got %#v", tc.expected, tokens)
				}
			*/
			require.Equalf(t, tc.expected, tokens, "expected %#v, got %#v", tc.expected, tokens)
		})
	}
}

var inCI = os.Getenv("CI") != ""

func TestInCI(t *testing.T) {
	if !inCI {
		t.Skip("no in CI")
	}
}

type tokenizedTomlCase struct {
	Text   string   `toml:"text"`
	Tokens []string `toml:"tokens"`
	Name   string   `toml:"name"`
}

type tokenizedYamlCase struct {
	Text   string   `yaml:"text"`
	Tokens []string `yaml:"tokens"`
	Name   string   `yaml:"name"`
}

func loadTokenizedTomlCases(t *testing.T) []tokenizedTomlCase {
	file, err := os.Open("testdata/tokenize_cases.toml")
	require.NoError(t, err)
	defer file.Close()

	decoder := toml.NewDecoder(file)

	var data struct {
		Cases []tokenizedTomlCase `toml:"cases"`
	}

	_, err = decoder.Decode(&data)
	require.NoError(t, err)

	return data.Cases

}

func loadTokenizedYamlCases(t *testing.T) []tokenizedYamlCase {
	file, err := os.Open("testdata/tokenize_cases.yml")
	require.NoError(t, err)
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	var data []tokenizedYamlCase

	err = decoder.Decode(&data)
	require.NoError(t, err)

	return data

}

func TestCasesFromToml(t *testing.T) {
	cases := loadTokenizedTomlCases(t)

	for _, tc := range cases {
		name := tc.Name
		if name == "" {
			name = tc.Text
		}

		t.Run(name, func(t *testing.T) {
			tokens := Tokenize(tc.Text)

			//TOML does not have nil
			if tokens == nil {
				tokens = []string{}
			}
			require.Equalf(t, tc.Tokens, tokens, "expected %#v, got %#v", tc.Tokens, tokens)
		})
	}

}

func TestCasesFromYaml(t *testing.T) {
	cases := loadTokenizedYamlCases(t)

	for _, tc := range cases {
		name := tc.Name
		if name == "" {
			name = tc.Text
		}

		t.Run(name, func(t *testing.T) {
			tokens := Tokenize(tc.Text)
			require.Equalf(t, tc.Tokens, tokens, "expected %#v, got %#v", tc.Tokens, tokens)
		})
	}

}

func FuzzTokenizer(f *testing.F) {
	fn := func(t *testing.T, text string) {
		tokens := Tokenize(text)
		ltext := strings.ToLower(text)
		for _, tok := range tokens {
			require.Contains(t, ltext, tok)
		}
	}
	f.Fuzz(fn)
}
