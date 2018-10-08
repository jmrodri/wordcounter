package main

import (
	"io"
	"strings"
	"text/scanner"
	"unicode"
)

// CountWordsPerSentence given a list of sentences, return a map whose keys are
// the sentences and the value is the word count in the sentence
func CountWordsPerSentence(sentences []string) map[string]int {
	result := make(map[string]int)
	for _, sentence := range sentences {
		words := strings.Split(sentence, " ")
		if len(words) == 1 && words[0] == "" {
			result[sentence] = 0
			continue
		}
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

// CountCharacters will count the number of characters in the given string.
// Unlike the CountWordsPerSentence which splits the sentences and keeps track
// of them, CountCharacters treats it as a big blob of data. CountCharacters
// returns a map with the given rune (character) and the total count.
func CountCharacters(sents string) map[rune]int {
	result := make(map[rune]int)
	rdr := strings.NewReader(strings.ToUpper(sents))
	for {
		if c, _, err := rdr.ReadRune(); err != nil {
			if err == io.EOF {
				break
			}
		} else if unicode.IsLetter(c) {
			result[c] = result[c] + 1
		}
	}
	return result
}
