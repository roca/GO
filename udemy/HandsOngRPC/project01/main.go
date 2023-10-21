package main

import (
	"fmt"
	"log"
	"project01/basic"
	"time"
)

type logWriter struct{}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Printf(time.Now().Format("15:04:04") + " " + string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
	// basic.BasicHello()
	basic.BasicUser()
}
