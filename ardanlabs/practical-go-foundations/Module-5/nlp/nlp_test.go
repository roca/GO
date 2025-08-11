package nlp

import (
	"os"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
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

type tokenizedCase struct {
	Text   string   `toml:"text"`
	Tokens []string `toml:"tokens"`
	Name   string   `toml:"name"`
}

func loadTokenizedCases(t *testing.T) []tokenizedCase {
	file, err := os.Open("testdata/tokenize_cases.toml")
	require.NoError(t, err)
	defer file.Close()

	decoder := toml.NewDecoder(file)

	var data struct {
		Cases []tokenizedCase `toml:"cases"`
	}

	_, err = decoder.Decode(&data)
	require.NoError(t, err)

	return data.Cases

}

func TestCasesFromToml(t *testing.T) {
	cases := loadTokenizedCases(t)

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
