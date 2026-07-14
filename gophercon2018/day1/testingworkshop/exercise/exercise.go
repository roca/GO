package exercise

import "strings"

func MyHasSuffix(instring string, sufffix string) bool {
	return strings.HasSuffix(instring, sufffix)
}

func MyIndex(instring string, substring string) int {
	return strings.Index(instring, substring)
}
