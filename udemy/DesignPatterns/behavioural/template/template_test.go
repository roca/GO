package template

import (
	"strings"
	"testing"
)

type TestStruct struct {
	ITemplate
}

func (m *TestStruct) Message() string {
	return "world"
}

func TestTemplate_ExecuteAlgorithm(t *testing.T) {
	t.Run("Using interfaces", func(t *testing.T) {
		s := &TestStruct{}
		template := &TemplateImpl{}
		s.ITemplate = template
		res := s.ExecuteAlgorithm(s)
		expected_words := []string{"world", "hello", "template"}
		expectedWordsOrErrors(res, expected_words, t)
	})

	t.Run("Using anonymous functions", func(t *testing.T) {
		m := new(AnonymousTemplate)
		expected_words := []string{"world", "hello", "template"}
		res := m.ExecuteAlgorithm(func() string {
			return "world"
		})
		expectedWordsOrErrors(res, expected_words, t)
	})
	t.Run("Using anonymous functions adapted to interface", func(t *testing.T) {
		messageRetriever := MessageRetrieverAdater(func() string {
			return "world"
		})

		if messageRetriever == nil {
			t.Fatal("Can not continue with a nil MessageRetriever")
		}

		template := &TemplateImpl{}
		res := template.ExecuteAlgorithm(messageRetriever)
		expected_words := []string{"world", "hello", "template"}
		expectedWordsOrErrors(res, expected_words, t)
	})
}

func expectedWordsOrErrors(res string, expected_words []string, t *testing.T) {
	for _, expected := range expected_words {
		if !strings.Contains(res, expected) {
			t.Errorf("Expected string '%s' wasn't found on returned string: '%s'", expected, res)
		}
	}
}
