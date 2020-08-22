package isogram

import (
	"regexp"
	"strings"
)

func IsIsogram(word string) bool {

	word = strings.ToUpper(word)
	re, err := regexp.Compile(`(?m)[^\w]`)
	if err != nil {
		return false
	}

	word = re.ReplaceAllString(word, "")

	isogram := true

	for _, s := range word {
		if strings.Count(word, string(s)) != 1 {
			isogram = false
		}
	}

	return isogram
}
