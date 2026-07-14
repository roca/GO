package shapes

import (
	"bytes"
	"io"

	"example.com/strategy"
)

type ConsoleSquare struct {
	strategy.PrintOutPut
}

func (c *ConsoleSquare) Print() error {
	r := bytes.NewReader([]byte("Square\n"))
	io.Copy(c.Writer, r)
	return nil
}
