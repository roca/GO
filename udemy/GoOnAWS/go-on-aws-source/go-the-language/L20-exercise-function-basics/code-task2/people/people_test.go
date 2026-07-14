package people_test

import (
	"testing"
	"walkintoabar/people"
)

func TestMoodString(t *testing.T) {
	//begin expected
	tests := []struct {
		name        string
		mood        people.Mood
		expectedStr string
	}{
		{"neutral", people.Neutral, "neutral"},
		{"unknown", people.Mood(100), "unknown"}, // Test case for unknown mood
	}
	//end expected

	//begin test
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.mood.String()
			if result != test.expectedStr {
				t.Errorf("Mood %d: Expected %s, but got %s", test.mood, test.expectedStr, result)
			}
		})
	}
	//end test
}
