package main

import "testing"

func Test_findKet(t *testing.T) {
	tests := []struct {
		name      string
		encrypted []byte
		decrypted string
	}{
		{"yes", []byte{0x1b, 0x2c, 0x3d}, "yes"},
		{"car", []byte{0x2a, 0xff, 0xea}, "car"},
		{"she", []byte{0x7d, 0x31, 0x32}, "she"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := findKey(tt.encrypted, tt.decrypted)
			if err != nil && err.Error() != "key not found" {
				t.Errorf("findKey error message not equal to 'key not found ' got '%v'", err)
				return
			}
			if err == nil {
				got := string(crypt(tt.encrypted, key))
				if got != tt.decrypted {
					t.Errorf("Decrypting with key = %v gives %s . want %s", key, got, tt.decrypted)
				}
				t.Logf("Key found: %x\n", key)
			}
		})
	}
}
