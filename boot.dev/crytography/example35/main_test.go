package main

import "testing"

func Test_splitInHalf(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name       string
		args       args
		wantFirst  string
		wantSecond string
	}{
		{
			name:       "Test 1",
			args:       args{s: "123456"},
			wantFirst:  "123",
			wantSecond: "456",
		},
		{
			name:       "Test 2",
			args:       args{s: "1234567"},
			wantFirst:  "123",
			wantSecond: "4567",
		},
		{
			name:       "Test 3",
			args:       args{s: "123 4567"},
			wantFirst:  "123 ",
			wantSecond: "4567",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirst, gotSecond := splitInHalf(tt.args.s)
			if gotFirst != tt.wantFirst {
				t.Errorf("splitInHalf() gotFirst = '%v', want '%v'", gotFirst, tt.wantFirst)
			}
			if gotSecond != tt.wantSecond {
				t.Errorf("splitInHalf() gotSecond = '%v', want '%v'", gotSecond, tt.wantSecond)
			}
		})
	}

}
