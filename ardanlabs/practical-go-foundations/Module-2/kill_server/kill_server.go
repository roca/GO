package main

import (
	"fmt"
	"log/slog"
	"os"
)

func main() {
	fmt.Println(KillServer("server.pid"))
}

func KillServer(pidFile string) error {
	file, err := os.Open(pidFile)
	if err != nil {
		return err
	}
	defer file.Close()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return fmt.Errorf("%q -bad pid: %s", pidFile, err)
	}

	slog.Info("Killing", "pid", pid)

	return nil
}
