package logs

import (
	"testing"
	"time"
)

const (
	n = 10_000
)

var (
	initialTime    = time.Date(2022, time.July, 4, 0, 0, 0, 0, time.UTC)
	initialTimeFmt = initialTime.Format(timeFmt)
	logs           []Log
)

func init() {
	logs = make([]Log, n)
	t := initialTime
	for i := 0; i < n; i++ {
		t = t.Add(173 * time.Millisecond)
		logs[i].Time = t.Format(timeFmt)
	}
}

func Benchmark_sortByTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortByTime(logs)
		if logs[0].Time != initialTimeFmt {
			b.Fatal(logs[:10])
		}
	}
}
