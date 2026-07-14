// Package twofer implements ShareWith for manipulating a twofer phrase.
// `Two-fer` or `2-fer` is short for two for one. One for you and one for me.
package twofer

import "fmt"

// ShareWith returns "One for [NAME], one for me." if NAME is not "".
// Otherwise it returns "One for you, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
