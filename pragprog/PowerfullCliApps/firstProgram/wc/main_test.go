package main

import (
	"bytes"
	"testing"
)

func TestCount(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 4
	res := count(b)
	if exp != res {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
