package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
)

func main() {
	err := KillServer("server.pid")
	if err != nil {
		fmt.Println("EROR:", err)
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("not found")
		}
		for e := err; e != nil; e = errors.Unwrap(e) {
			fmt.Printf("> %s\n", e)
		}
	}
}

func KillServer(pidFile string) error {
	file, err := os.Open(pidFile)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			slog.Warn("close", "file", pidFile, "error", err)
		}
	}()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return fmt.Errorf("%q -bad pid: %w", pidFile, err)
	}

	slog.Info("Killing", "pid", pid)

	if err := os.Remove(pidFile); err != nil {
		slog.Warn("delete", "file", pidFile, "error", err)
	}

	return nil
}
