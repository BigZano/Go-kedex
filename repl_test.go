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
			input:    " hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    " smashing pumpkins",
			expected: []string{"smashing", "pumpkins"},
		},
		{
			input:    " alice in chains",
			expected: []string{"alice", "in", "chains"},
		},
		{
			input:    " Rage Against ThE MaChInE",
			expected: []string{"rage", "against", "the", "machine"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("expected %d words, got %d", len(c.expected), len(actual))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("at index %d: expected %q, got %q", i, expectedWord, word)
			}

		}
	}
}
