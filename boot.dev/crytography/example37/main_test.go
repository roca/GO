package main

import "testing"

func Test_hashPassword(t *testing.T) {
	type args struct {
		password1 string
		password2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test 1", args{"thisIsAPassword", "thisIsAPassword"}, true},
		{"Test 2", args{"thisIsAPassword", "thisIsAnotherPassword"}, false},
		{"Test 3", args{"corr3ct h0rse", "corr3ct h0rse"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashed, err := hashPassword(tt.args.password1)
			if err != nil {
				t.Errorf("Error hashing password: %v\n", err)
				return
			}

			match := checkPasswordHash(tt.args.password2, hashed)
			if match != tt.want {
				t.Errorf("hashPassword() = %v, want %v", match, tt.want)
			}
		})
	}
}
