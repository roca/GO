package main

import (
	"fmt"
	"time"
)

type logWriter struct{}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Printf(time.Now().Format("15:04:04") + " " + string(bytes))
}
