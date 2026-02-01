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
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "   complex   TEST   case   ",
			expected: []string{"complex", "test", "case"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		// Check the length of the actual slice against the expected slice
		// If they don't match, we error out immediately for this case
		// because iterating over indices might cause a panic (index out of range)
		if len(actual) != len(c.expected) {
			t.Errorf("The lengths are not equal: %v vs %v",
				len(actual),
				len(c.expected),
			)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			// Check each word in the slice
			if word != expectedWord {
				t.Errorf("Error in case '%v': expected '%v' but got '%v'",
					c.input,
					expectedWord,
					word,
				)
			}
		}
	}
}
