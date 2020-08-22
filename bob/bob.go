// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	remark = removeSpace(remark)

	if isQuestion(remark) && !isUpper(remark) {
		return "Sure."
	}

	if isUpper(remark) && !isQuestion(remark) {
		return "Whoa, chill out!"
	}

	if isQuestion(remark) && isUpper(remark) {
		return "Calm down, I know what I'm doing!"
	}

	if remark == "" {
		return "Fine. Be that way!"
	}

	return "Whatever."
}

func isUpper(s string) bool {

	return strings.ToUpper(s) == s && text(s)
}

func isQuestion(s string) bool {

	return strings.HasSuffix(s, "?")
}

func text(s string) bool {

	is, err := regexp.Match(`(?m)[a-zA-Z]`, []byte(s))
	if err != nil {
		return false
	}

	return is
}

func removeSpace(s string) string {

	return strings.TrimSpace(s)
}
