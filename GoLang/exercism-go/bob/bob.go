// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	var answer string
	if remark = strings.TrimSpace(remark); remark == "" {
		answer = "Fine. Be that way!"
	} else if remark = strings.TrimSpace(remark); string(remark[len(remark)-1]) == "?" && strings.ToUpper(remark) == remark && strings.ContainsAny(strings.ToLower(remark), "abcdefghijklmnopqrstuvwxyz") {
		answer = "Calm down, I know what I'm doing!"
	} else if remark = strings.TrimSpace(remark); strings.ToUpper(remark) == remark && strings.ContainsAny(strings.ToLower(remark), "abcdefghijklmnopqrstuvwxyz") {
		answer = "Whoa, chill out!"
	} else if remark = strings.TrimSpace(remark); string(remark[len(remark)-1]) == "?" {
		answer = "Sure."
	} else {
		answer = "Whatever."
	}
	return answer
}
