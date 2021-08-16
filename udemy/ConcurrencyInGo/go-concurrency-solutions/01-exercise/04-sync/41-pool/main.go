package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

//create pool of bytes.Buffers which can be reused.

var bufPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("allocating new bytes.Buffer")
		return new(bytes.Buffer)
	},
}

func log(w io.Writer, val string) {
	b := bufPool.Get().(*bytes.Buffer)

	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(" : ")
	b.WriteString(val)
	b.WriteString("\n")

	n, err := w.Write(b.Bytes())
	if err != nil {
		panic(err)
	}

	bufPool.Put(b)
	fmt.Printf("%d Bytes written \n", n)
}

func main() {
	log(os.Stdout, "debug-string1")
	log(os.Stdout, "debug-string2")
}
