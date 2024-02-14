package main

import "testing"

func Test_xor(t *testing.T) {
	type args struct {
		lhs bool
		rhs bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"truetrue", args{true, true}, false},
		{"truefalse", args{true, false}, true},
		{"falsetrue", args{false, true}, true},
		{"falsefalse", args{false, false}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xor(tt.args.lhs, tt.args.rhs); got != tt.want {
				t.Errorf("xor() = %v, want %v", got, tt.want)
			}
		})
	}
}
