package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"boot.dev/linko/internal/store"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	httpPort := flag.Int("port", 8899, "port to listen on")
	dataDir := flag.String("data", "./data", "directory to store data")
	flag.Parse()

	status := run(ctx, cancel, *httpPort, *dataDir)
	cancel()
	os.Exit(status)
}

func run(ctx context.Context, cancel context.CancelFunc, httpPort int, dataDir string) int {

	logger, closeF, err := initializeLogger(os.Getenv("LINKO_LOG_FILE"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize logger: %v\n", err)
		return 1
	}
	defer func() {
		if err := closeF(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to close logger: %v\n", err)
		}
	}()

	st, err := store.New(dataDir, logger)
	if err != nil {
		logger.Info(fmt.Sprintf("failed to create store: %v\n", err))
		return 1
	}

	s := newServer(*st, httpPort, cancel, logger)
	var serverErr error
	go func() {
		serverErr = s.start()
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger.Info("Linko is shutting down")
	if err := s.shutdown(shutdownCtx); err != nil {
		logger.Info(fmt.Sprintf("failed to shutdown server: %v\n", err))
		return 1
	}
	if serverErr != nil {
		logger.Info(fmt.Sprintf("server error: %v\n", serverErr))
		return 1
	}
	return 0
}

type closeFunc func() error

func initializeLogger(logFile string) (*slog.Logger, closeFunc, error) {
	if logFile != "" {
		logFile, err := os.OpenFile(os.Getenv("LINKO_LOG_FILE"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to open log file: %w", err)
		}

		bufferedFile := bufio.NewWriterSize(logFile, 8192)
		multiWriter := io.MultiWriter(os.Stderr, bufferedFile)

		closeF := func() error {
			if err := bufferedFile.Flush(); err != nil {
				return fmt.Errorf("failed to flush log file: %w", err)
			}
			if err := logFile.Close(); err != nil {
				return fmt.Errorf("failed to close log file: %w", err)
			}
			return nil
		}

		return slog.New(slog.NewTextHandler(multiWriter, nil)), closeF, nil

	} else {
		return slog.New(slog.NewTextHandler(os.Stderr, nil)), func() error { return nil }, nil
	}
}
