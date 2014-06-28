package main

import (
	"sort"
)

func formatDict(dict []string) []string {
	dict = validateWords(dict)
	dict = sortWords(dict)
	dict = uniqWords(dict)
	return dict
}

func validateWords(words []string) []string {
	validated := []string{}
	for _, word := range words {
		if len(word) < 1 {
			continue
		}

		// FIXME: sort.StringSlice.Sort() inserts unexpected words which start with '>'
		if word[0] == '>' {
			continue
		}

		validated = append(validated, word)
	}
	return validated
}

func sortWords(words []string) []string {
	ss := sort.StringSlice(words)
	ss.Sort()
	return []string(ss)
}

func uniqWords(words []string) []string {
	uniqed := []string{}

	prev := ""
	for _, word := range words {
		if prev != word {
			uniqed = append(uniqed, word)
		}
		prev = word
	}
	return uniqed
}
