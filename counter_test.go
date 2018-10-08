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

func TestCountCharacters(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected map[rune]int
	}{
		{
			name:  "count characters",
			input: "This is a test. WE PASSED! Or did we?",
			expected: map[rune]int{
				'A': 2,
				'D': 3,
				'E': 4,
				'H': 1,
				'I': 3,
				'O': 1,
				'P': 1,
				'R': 1,
				'S': 5,
				'T': 3,
				'W': 2,
			},
		},
		{
			name:     "empty input",
			input:    "",
			expected: map[rune]int{},
		},
	}

	for _, tc := range testCases {
		output := CountCharacters(tc.input)
		assert.EqualValues(t, tc.expected, output)
	}
}
