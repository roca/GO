package main

import (
	"fmt"
	"testing"
)

func Test_crypt(t *testing.T) {
	input := []byte("0110100001100101011011000110110001101111")
	key := []byte("0111001101101010011001100111010101100100")
	want := []byte("0001101100001111000010100001100100001011")

	bytes := crypt(input, key)
	s := toString(bytes)

	got := []byte(s)
	if len(got) != len(want) {
		t.Error("Wrong lengths")
		return
	}
	if string(got) != string(want) {
		t.Errorf("\ndata:\t%v\nkey:\t%v\noutput:\t%v\nwant:\t%v\n", string(input), string(key), string(got), string(want))
	}
}

func toString(bytes []byte) string {
	s := ""
	for _, v := range bytes {
		s += fmt.Sprintf("%d", v)
	}
	return s
}