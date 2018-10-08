package main

import (
	"strings"
	"text/scanner"
)

// CountWordsPerSentence given a list of sentences, return a map whose keys are
// the sentences and the value is the word count in the sentence
func CountWordsPerSentence(sentences []string) map[string]int {
	result := make(map[string]int)
	for _, sentence := range sentences {
		words := strings.Split(sentence, " ")
		result[sentence] = len(words)
	}

	return result
}

// SplitSentences returns a list of sentences split from a single string. A
// sentence is defined as a set of words that end with a period (.), question
// mark (?), or exclamation point (!).
func SplitSentences(sents string) []string {
	var b strings.Builder
	var s scanner.Scanner
	var sentences []string

	s.Init(strings.NewReader(sents))
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		switch {
		case s.TokenText() == ".", s.TokenText() == "?", s.TokenText() == "!":
			b.WriteString(s.TokenText())
			sentences = append(sentences, b.String())
			b.Reset()
		default:
			b.WriteString(s.TokenText())
			peek := s.Peek()
			if peek != '.' && peek != '?' && peek != '!' {
				b.WriteString(" ")
			}
		}
	}

	return sentences
}

func CountWords(sentence string) int {
	return 0
}
