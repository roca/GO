package main

import "testing"

func Test_genKeys(t *testing.T) {
	for i := 1; i < 4; i++ {
		pub, priv, err := genKeys()
		arePaired := keysArePaired(pub, priv)
		if err != nil || !arePaired {
			t.Errorf("Test %v failed: %v", i, err)
		}
	}
}
