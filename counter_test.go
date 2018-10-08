package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitSentences(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "finds sentences",
			input:    "This is a test. WE PASSED! Or did we?",
			expected: []string{"This is a test.", "WE PASSED!", "Or did we?"},
		},
		{
			name:  "empty input",
			input: "",
		},
		{
			name:     "missing final punctuation",
			input:    "This is a test. WE PASSED! Or did we",
			expected: []string{"This is a test.", "WE PASSED!"},
		},
	}

	for _, tc := range testCases {
		output := SplitSentences(tc.input)
		assert.Equal(t, tc.expected, output)
	}
}

func TestCountWordsPerSentence(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected map[string]int
	}{
		{
			name:  "count words",
			input: []string{"This is a test.", "WE PASSED!", "Or did we?"},
			expected: map[string]int{
				"This is a test.": 4,
				"WE PASSED!":      2,
				"Or did we?":      3,
			},
		},
		{
			name:  "empty input",
			input: []string{""},
			expected: map[string]int{
				"": 0,
			},
		},
	}

	for _, tc := range testCases {
		output := CountWordsPerSentence(tc.input)
		assert.Equal(t, tc.expected, output)
	}
}
