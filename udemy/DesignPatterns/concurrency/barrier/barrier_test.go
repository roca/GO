package barrier

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestBarrier(t *testing.T) {
	t.Run("Correct endpoints", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/user-agent"}
		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "Accept-Encoding") || !strings.Contains(result, "Accept-Encoding") {
			t.Fail()
		}
		t.Log(result)
	})
	t.Run("One endpoint incorrect", func(t *testing.T) {
		endpoints := []string{"http://malformed-url", "http://httpbin.org/user-agent"}
		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "ERROR") {
			t.Fail()
		}
		t.Log(result)
	})
	t.Run("Very short timeout", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/user-agent"}
		TimeoutMilliseconds = 1

		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "Timeout") {
			t.Fail()
		}
		t.Log(result)
	})
}
func captureBarrierOutput(endpoints ...string) string {
	reader, writer, _ := os.Pipe()

	os.Stdout = writer

	out := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	Barrier(endpoints...)

	writer.Close()
	temp := <-out

	return temp
}
