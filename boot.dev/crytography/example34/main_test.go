package main

import "testing"

func Test_checksumMatches(t *testing.T) {
	type args struct {
		message  string
		checksum string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "pa$$w0rd",
			args: args{
				message:  "pa$$w0rd",
				checksum: "4b358ed84b7940619235a22328c584c7bc4508d4524e75231d6f450521d16a17",
			},
			want: true,
		},
		{
			name: "buil4WithB1ologee",
			args: args{
				message:  "buil4WithB1ologee",
				checksum: "1c489a153271aaf3b234aa154b1a2eef5248eb9ab402e4d3c8b7bc3d81fed1a8",
			},
			want: false,
		},
		{
			name: "br3ak1ngB@d1sB3st",
			args: args{
				message:  "br3ak1ngB@d1sB3st",
				checksum: "5d178e1c6fd5d76415e1632f84e5192fb50ef244d42a02148fedbf991d914546",
			},
			want: false,
		},
		{
			name: "b3ttterC@llS@ulI$B3tter",
			args: args{
				message:  "b3ttterC@llS@ulI$B3tter",
				checksum: "8d42f2dc81476123974619969a42b27b8d8a4fa507be99c9623f614ad2d859f7",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checksumMatches(tt.args.message, tt.args.checksum); got != tt.want {
				t.Errorf("checksumMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}
