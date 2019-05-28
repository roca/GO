package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func run(prog string, args ...string) ([]byte, error) {
	cmd := exec.Command(prog, args...)
	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		return []byte{}, err
	}

	errPipe, err := cmd.StderrPipe()
	if err != nil {
		return []byte{}, err
	}

	err = cmd.Start()
	if err != nil {
		return []byte{}, err
	}

	var outBuf bytes.Buffer
	var errBuf bytes.Buffer

	io.Copy(&outBuf, outPipe)
	io.Copy(&errBuf, errPipe)

	err = cmd.Wait()
	if err != nil {
		return []byte{}, err
	}

	if len(errBuf.Bytes()) != 0 {
		return outBuf.Bytes(), fmt.Errorf("%s", errBuf.Bytes())
	}

	return outBuf.Bytes(), nil
}

func main() {

}
