package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(isMatch("abc", "a***abc"))
}

func isMatch(s string, p string) bool {
	re, _ := regexp.Compile("[*]+")
	p = string(re.ReplaceAll([]byte(p), []byte("*")))

	match, _ := regexp.MatchString("^"+p+"$", s)
	return match
}
