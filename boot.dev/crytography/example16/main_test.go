package main

import (
	"reflect"
	"testing"
)

func Test_padWithZeros(t *testing.T) {
	type args struct {
		block       []byte
		desiredSize int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"1 byte", args{[]byte{0xFF}, 4}, []byte{0xFF, 0x00, 0x00, 0x00}},
		{"2 bytes", args{[]byte{0xFA, 0xBC}, 8}, []byte{0xFA, 0xBC, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"3 bytes", args{[]byte{0x12, 0x34, 0x56}, 12}, []byte{0x12, 0x34, 0x56, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := padWithZeros(tt.args.block, tt.args.desiredSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("padWithZeros() = %v, want %v", got, tt.want)
			}
		})
	}
}
