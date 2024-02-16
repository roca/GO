package main

import (
	"fmt"
	"testing"
)

func Test_getBlockSize(t *testing.T) {
	type args struct {
		keyLen     int
		cipherType int
	}
	tests := []struct {
		name      string
		args      args
		want      int
		errorText string
	}{
		{"AES 16", args{16, typeAES}, 16, ""},
		{"AES 24", args{24, typeAES}, 16, ""},
		{"AES 32", args{32, typeAES},16, ""},
		{"DES 8", args{8, typeDES}, 8, ""},
		{"DES 16", args{16, typeDES}, 0, "crypto/des: invalid key size 16"},
		{"DES 24", args{24, typeDES}, 0, "crypto/des: invalid key size 24"},
		{"Unknown", args{1, -1}, 0, "invalid cipher type"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getBlockSize(tt.args.keyLen, tt.args.cipherType)
			cipherTypeName := getCipherTypeName(tt.args.cipherType)
			if (err != nil) && (fmt.Sprintf("%s", err) != tt.errorText) {
				t.Errorf("%s getBlockSize() error message incorrect. got: %v, want: %v", cipherTypeName, err, tt.errorText)
				return
			}
			if got != tt.want {
				t.Errorf("%s getBlockSize() = %v, want %v", cipherTypeName , got, tt.want)
			}
		})
	}
}
