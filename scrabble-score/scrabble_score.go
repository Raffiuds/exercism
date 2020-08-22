package scrabble

import (
	"strings"
)

var scores = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

func Score(word string) int {

	word = strings.ToUpper(word)
	score := 0

	for key, value := range scores {
		count := 0
		for _, s := range key {
			if strings.Count(word, string(s)) <= len(word) {
				count = strings.Count(word, string(s))
				score += (value * count)
			}
		}

	}

	return score
}
