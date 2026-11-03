package web

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"rag-course/llm"
	"rag-course/rag"
	"rag-course/vector"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//go:embed templates/*.gohtml
var templateFS embed.FS

type Options struct {
	Addr             string
	SystemPromptFIle string
	Title            string
	Store            vector.Store
	ProcessedDir     string
	ImagesDir        string
}

type Server struct {
	client       *llm.Client
	embedder     *llm.Client
	retriever    *rag.Retriever
	store        vector.Store
	processedDir string
	imagesDir    string
	tpl          *template.Template
	system       string
	title        string
}

func New(client, embedder *llm.Client, retriever *rag.Retriever, opts Options) (*Server, error) {
	tpl, err := template.ParseFS(templateFS, "templates/*.gohtml")
	if err != nil {
		return nil, fmt.Errorf("parse templates: %w", err)
	}

	title := opts.Title
	if title == "" {
		title = "RAG Chat"
	}

	return &Server{
		client:       client,
		embedder:     embedder,
		retriever:    retriever,
		store:        opts.Store,
		processedDir: opts.ProcessedDir,
		imagesDir:    opts.ImagesDir,
		tpl:          tpl,
		system:       readSystemPrompt(opts.SystemPromptFIle),
		title:        title,
	}, nil
}

func (s *Server) Routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Get("/chat", s.handleChatPage)

	return r
}

func (s *Server) handleChatPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := s.tpl.ExecuteTemplate(w, "chat.gohtml", map[string]any{
		"Title": s.title,
		//"CaptionEnabled": s.client.HasVision(),
	}); err != nil {

		log.Printf("[web] template error: %v", err)
	}
}

func (s *Server) Run(ctx context.Context, addr string) error {
	srv := &http.Server{
		Addr:              addr,
		Handler:           s.Routes(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		err := srv.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
		errCh <- nil
	}()

	select {
	case <-ctx.Done():
		shutDownCtrx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = srv.Shutdown(shutDownCtrx)
		return nil
	case err := <-errCh:
		return err
	}
}

func readSystemPrompt(path string) string {
	if path == "" {
		return ""
	}

	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return ""
	}

	return strings.TrimSpace(string(data))
}
