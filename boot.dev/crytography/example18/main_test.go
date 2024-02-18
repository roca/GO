package main

import (
	"encoding/binary"
	"math/bits"
	"testing"
)

func Test_feistel(t *testing.T) {
	tests := []struct {
		name    string
		message []byte
		key     []byte
		rounds  int
	}{
		{"test1", []byte("General Kenobi!!!!"), []byte("thesecret"), 8},
		{"test2", []byte("Hello there!"), []byte("@n@kiN"), 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			roundKeys := makeRoundKeys(tt.key, tt.rounds)
			encrypted := feistel(tt.message, roundKeys)
			decrypted := feistel(encrypted, reverse(roundKeys))
			if string(decrypted) != string(tt.message) {
				t.Errorf("Decrypted: '%v', want '%v'", string(decrypted), string(tt.message))
			}
		})
	}
}

func makeRoundKeys(key []byte, rounds int) [][]byte {
	roundKeys := [][]byte{}
	for i := 0; i < rounds; i++ {
		ui := binary.BigEndian.Uint32(key)
		rotated := bits.RotateLeft32(uint32(ui), i)
		finalRound := make([]byte, len(key))
		binary.LittleEndian.PutUint32(finalRound, uint32(rotated))
		roundKeys = append(roundKeys, finalRound)
	}
	return roundKeys
}
