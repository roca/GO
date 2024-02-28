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
			args:       args{s: "hello"},
			wantFirst:  "he",
			wantSecond: "llo",
		},
		{
			name:       "Test 2",
			args:       args{s: "hello world"},
			wantFirst:  "hello",
			wantSecond: " world",
		},
		{
			name:       "Test 3",
			args:       args{s: "hello world!"},
			wantFirst:  "hello ",
			wantSecond: "world!",
		},
		{
			name:       "Test 4",
			args:       args{s: "hello world! 123"},
			wantFirst:  "hello wo",
			wantSecond: "rld! 123",
		},
		{
			name:       "Test 5",
			args:       args{s: "hello world! 1234"},
			wantFirst:  "hello wo",
			wantSecond: "rld! 1234",
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
