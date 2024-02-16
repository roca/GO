package main

import (
	"errors"
	"testing"
)

func Test_getBlockSize(t *testing.T) {
	type args struct {
		keyLen     int
		cipherType int
	}
	tests := []struct {
		name string
		args args
		want int
		err  error
	}{
		{"AES 16", args{16, typeAES}, 16, nil},
		{"AES 24", args{24, typeAES}, 24, nil},
		{"AES 32", args{32, typeAES}, 32, nil},
		{"AES 64", args{64, typeAES}, 64, nil},
		{"DES 8", args{8, typeDES}, 8, nil},
		{"DES 16", args{16, typeDES}, 16, nil},
		{"DES 24", args{24, typeDES}, 24, nil},
		{"Unknown", args{1, -1}, 0, errors.New("unknown cipher type")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getBlockSize(tt.args.keyLen, tt.args.cipherType)
			if (err != nil) != (tt.err != nil) {
				t.Errorf("getBlockSize() error = %v, wantErr %v", err, tt.err)
				return
			}
			if got != tt.want {
				t.Errorf("getBlockSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
