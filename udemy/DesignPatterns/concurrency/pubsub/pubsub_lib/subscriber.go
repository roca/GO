package pubsub_lib

import (
	"fmt"
	"io"
	"os"
	"time"
)

type ISubscriber interface {
	Notify(interface{}) error
	Close()
}

type writerSubscriber struct {
	in     chan interface{}
	id     int
	Writer io.Writer
}

func (s *writerSubscriber) Notify(msg interface{}) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%#v", rec)
		}
	}()

	select {
	case s.in <- msg:
	case <-time.After(time.Second):
		err = fmt.Errorf("Timeout\n")
	}
	return
}

func (s *writerSubscriber) Close() {
	close(s.in)
}

func NewWriterSubscriber(id int, out io.Writer) ISubscriber {

	if out == nil {
		out = os.Stdout
	}

	s := &writerSubscriber{
		id:     id,
		in:     make(chan interface{}),
		Writer: out,
	}

	go func() {
		for msg := range s.in {
			fmt.Fprintf(s.Writer, "(W%d): %v\n", s.id, msg)
		}
	}()
	return s
}
