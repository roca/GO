package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_alpha_prompt(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = oldOut

	expected := `Encrypted password: e1a08e8c08c45a98a908ec8fd8bf72f14e09866e
Decrypted password: k33pThisPasswordSafe
========
Encrypted password: bba18ec869
Decrypted password: 12345
========
Encrypted password: fefbd8ac3ddf409c961bfbb3c19d79d9680f876a1d0a
Decrypted password: thePasswordOnMyLuggage
========
Encrypted password: fafac7863df347839c36d7a9db
Decrypted password: pizza_the_HUt
========
`
	if string(out) != expected {
		t.Errorf("Expected %s, got %s", expected, string(out))
	}
}
