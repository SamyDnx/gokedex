package main

import (
	"testing"
)

func TestCleanInput (t *testing.T) {
	cases := []struct {
		input		string
		expected	[]string
	}{
		{
			input:		"  hello world   ",
			expected:	[]string{"hello", "world"},
		},
		{
			input:		"\n   caca   \tpipi   ",
			expected:	[]string{"caca", "pipi"},
		},
		{
			input:		"PROUTE FLAtuLENCE   ",
			expected:	[]string{"proute", "flatulence"},
		},
		{
			input:		"",
			expected:	[]string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("length do not match: got %d and %d", len(actual), len(c.expected))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if expectedWord != word {
				t.Error("words do not match")
			}
		}
	}
}
