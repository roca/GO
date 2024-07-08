package nlp

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

// var tokenizetCases = []struct { // anonymous struct
// 	text   string
// 	tokens []string
// }{
// 	{"What's on first?", []string{"what", "s", "on", "first"}},
// 	{"", nil},
// }

// Exercise: Read test cases from tokennize_cases.toml

type tokenizeCase struct {
	Text   string
	Tokens []string
}

func loadTokenizeCases(t *testing.T) []tokenizeCase {
	/*
		data, err := ioutil.ReadFile("tokenize_cases.toml")
		require.NoError(t, err, "Read file")
	*/

	var testCases struct {
		Cases []tokenizeCase
	}

	// err = toml.Unmarshal(data, &testCases)
	_, err := toml.DecodeFile("testdata/tokenize_cases.toml", &testCases)
	require.NoError(t, err, "Unmarshal TOML")
	return testCases.Cases
}

func TestTokenizeTable(t *testing.T) {
	for _, tc := range loadTokenizeCases(t) {
		t.Run(tc.Text, func(t *testing.T) {
			tokens := Tokenize(tc.Text)
			require.Equal(t, tc.Tokens, tokens)
		})
	}
}

func TestTokenize(t *testing.T) {
	text := "What's on second?"
	expected := []string{"what", "s", "on", "second"}
	tokens := Tokenize(text)
	require.Equal(t, expected, tokens)
}
