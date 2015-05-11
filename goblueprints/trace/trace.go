package trace

//Tracer is the interface thatnwe have described an object capablemof
// tracing events throughout code

import (
	"fmt"
	"io"
)

type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

func Off() Tracer {
	return &nilTracer{}
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))

}
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}
