// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	s = clearString(s)

	stgs := strings.Split(s, " ")

	acron := ""

	for _, stg := range stgs {
		acron = acron + prefix(stg)
	}

	return strings.ToUpper(acron)
}

func prefix(s string) string {

	if len(s) == 0 {
		return ""
	}
	return s[0:1]
}

func clearString(s string) string {

	re, err := regexp.Compile(`[^a-zA-Z0-9 -']+`)

	if err != nil {
		return ""
	}

	s = re.ReplaceAllString(s, " ")

	return strings.TrimSpace(s)
}
