package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

func checkDestination(targetURL string) error {
	resp, err := http.DefaultClient.Get(targetURL)
	if err != nil {
		return fmt.Errorf("destination unreachable: %w", err)
	}
	defer resp.Body.Close()
	_, _ = io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return fmt.Errorf("destination returned status %d", resp.StatusCode)
	}
	return nil
}

func requestLogger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			logger.Info(fmt.Sprintf("Served request: %s %s", r.Method, r.URL.Path))
		})
	}
}
