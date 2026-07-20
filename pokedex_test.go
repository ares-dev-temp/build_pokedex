package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},

		{
			input:    "   heyyo what  is this it working  ",
			expected: []string{"heyyo", "what", "is", "this", "it", "working"},
		},

		{
			input:    "tests it  can   work for  the best",
			expected: []string{"tests", "it", "can", "work", "for", "the", "best"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			//error and continue here
			t.Errorf("Actual array output doesn't match the expected array length")
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("actual word: %s doesnt match the expected word %s", word, expectedWord)
			}
		}
	}
}
