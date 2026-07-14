package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 4
	res := count(b, false, false)
	if exp != res {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1\n")
	exp := 3
	res := count(b, true, false)
	if exp != res {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountBytes(t *testing.T) {
	s := "word1 word2 word3 word4\n"
	b := bytes.NewBufferString(s)
	exp := len([]byte(s))
	res := count(b, true, true)
	if exp != res {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
