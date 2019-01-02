package main

import (
	"fmt"
	"testing"
)

func TestFound(t *testing.T) {

	var xtemp int
	x1 := 0
	x2 := 1
	for x := 0; x < 5; x++ {
		xtemp = x2
		x2 = x2 + x1
		x1 = xtemp
	}
	fmt.Println(x2)

	matchingStrings := []string{"ian", "Ian", "iuiygaygn", "I d skd a efju N"}
	for _, s := range matchingStrings {
		if !Found(s, "(?i)^i.*a.*n$") {
			t.Errorf("Failed regex match on string %s", s)
		}

	}

	nonMatchingStrings := []string{"ihhhhhn", "ina", "xian", "qwerty", "ixxxxxn", "xihhhaxxxn", "ixxhhnx"}
	for _, s := range nonMatchingStrings {
		if Found(s, "(?i)^i.*a.*n$") {
			t.Errorf("Failed regex match on string %s", s)
		}

	}

}
