// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

func Hey(remark string) string {

	var removeAllWhiteSpace, _ = regexp.Compile(`\s`)
	var sanitizedRemark = removeAllWhiteSpace.ReplaceAllString(remark, "")
	// fmt.Println(sanitizedRemark)

	if IsSilence(sanitizedRemark) {
		return "Fine. Be that way!"
	}

	if IsYelling(sanitizedRemark) {
		if strings.HasSuffix(sanitizedRemark, "?") {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	}

	if IsQuestion(sanitizedRemark) {
		return "Sure."
	}

	return "Whatever."
}

func IsYelling(remark string) bool {
	containsLetters, _ := regexp.Compile("[[:alpha:]]")
	return containsLetters.MatchString(remark) && strings.ToUpper(remark) == remark
}

func IsQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func IsSilence(remark string) bool {
	onlySilence, _ := regexp.Compile(`^\s*$`)
	return onlySilence.MatchString(remark)
}
