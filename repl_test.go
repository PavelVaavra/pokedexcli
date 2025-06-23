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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "    ",
			expected: []string{},
		},
		{
			input:    "  heLLO	WOrld  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Hello  world ",
			expected: []string{"hello", "world"},
		},
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("len(actual) == %d != len(expected) == %d", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("%s != %s", word, expectedWord)
			}
		}
}
}