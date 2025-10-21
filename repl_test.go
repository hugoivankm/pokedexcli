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
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "first   second  third fourth",
			expected: []string{"first", "second", "third", "fourth"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "Happy Path",
			expected: []string{"happy", "path"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("FAIL: A slice of length %v is expected but received a slice of length %v", len(actual), len(c.expected))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("FAIL: Received %v but was expecting %v ", actual, expectedWord)
			}

		}
	}
}
