package ingest

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"rag-course/llm"
	"rag-course/vector"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

const debounceDelay = 500 * time.Millisecond

func Watch(ctx context.Context, opts Options, embedder llm.Embedder, store vector.Store, logger *log.Logger) error {
	if filepath.Clean(opts.SourceDir) == filepath.Clean(opts.ProcessedDir) {
		return errors.New("source and processed directories must differ")
	}

	if err := os.MkdirAll(opts.SourceDir, 0o755); err != nil {
		return fmt.Errorf("create source dir: %w", err)
	}

	if err := os.MkdirAll(opts.ProcessedDir, 0o755); err != nil {
		return fmt.Errorf("create processed dir: %w", err)
	}

	w, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("create watcher: %w", err)
	}
	defer w.Close()

	if err := w.Add(opts.SourceDir); err != nil {
		return fmt.Errorf("watch source dir: %w", err)
	}

	processedDir, err := filepath.Abs(opts.ProcessedDir)
	if err != nil {
		return fmt.Errorf("resolve processed dir: %w", err)
	}

	handle := func(path string) {
		if err := processOne(ctx, path, opts, embedder, store); err != nil {
			logger.Printf("process %s: %v", filepath.Base(path), err)
			return
		}
		dst := filepath.Join(opts.ProcessedDir, filepath.Base(path))
		if err := os.Rename(path, dst); err != nil {
			logger.Printf("move %s to processed %s: %v", path, dst, err)
			return
		}

		logger.Printf("ingested %s", filepath.Base(path))
	}

	entries, err := os.ReadDir(opts.SourceDir)
	if err != nil {
		return fmt.Errorf("read source dir: %w", err)
	}

	for _, e := range entries {
		if e.IsDir() || strings.HasPrefix(e.Name(), ".") {
			continue
		}
		go handle(filepath.Join(opts.SourceDir, e.Name()))
	}

	var (
		timersMu sync.Mutex
		timers   = make(map[string]*time.Timer)
	)

	schedual := func(path string) {
		timersMu.Lock()
		defer timersMu.Unlock()
		if t, ok := timers(path); ok {
			t.Reset(debounceDelay)
			return
		}

		timers[path] = time.AfterFunc(debounceDelay, func() {
			timersMu.Lock()
			delete(timers, path)
			timersMu.Unlock()
			handle(path)
		})
	}

}
