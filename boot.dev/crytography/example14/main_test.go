package main

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func Test_crypt(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	textCh := make(chan byte)
	keyCh := make(chan byte)
	result := make(chan byte)

	input := []byte("0110100001100101011011000110110001101111")
	key := []byte("0111001101101010011001100111010101100100")
	want := []byte("0001101100001111000010100001100100001011")

	go func() {
		defer close(textCh)
		for _, v := range input {
			textCh <- v
		}
	}()

	go func() {
		defer close(keyCh)
		for _, v := range key {
			keyCh <- v
		}
	}()

	go crypt(textCh, keyCh, result)

	res := []byte{}
	for v := range result {
		res = append(res, v)
	}

	if toString(res) != string(want) {
		t.Errorf("crypt() = %s; want %s", toString(res), string(want))
	}

	_ = w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldOut

	expected := ""
	for i := range res {
		expected += fmt.Sprintf("Crypted byte: %d\n", i+1)
	}

	if string(out) != expected {
		t.Errorf("Expected %s, got %s", expected, string(out))
	}
}

func toString(bytes []byte) string {
	s := ""
	for _, v := range bytes {
		s += fmt.Sprintf("%d", v)
	}
	return s
}
