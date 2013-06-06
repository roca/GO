// bio project bio.go
package bio

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func CountBaseOccurences(sequence, base string) int {
	return strings.Count(strings.ToUpper(sequence), strings.ToUpper(base))
}

func SequenceFromRosalindFile(filePath string) string {

	// read whole the file
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// write whole the body
	err = ioutil.WriteFile("output.txt", b, 0644)
	if err != nil {
		panic(err)
	}

	return BytesToString(b, "")

}

func BytesToString(c []byte, acc string) string {

	if len(c) == 0 {
		return acc
	} else {
		head := c[0]
		tail := c[1:]
		return BytesToString(tail, acc+fmt.Sprintf("%c", head))
	}
}

func Reverse(s string) string {
	b := []rune(s)
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
