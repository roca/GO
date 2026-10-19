package pgvector

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"rag-course/vector"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/pgvector/pgvector-go"
	pgxvec "github.com/pgvector/pgvector-go/pgx"
)

type Options struct {
	DSN          string
	EmbeddingDim int
}

type Store struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, opts Options) (*Store, error) {
	if opts.DSN == "" {
		return nil, errors.New("pgvector: DSN is required")
	}
	if opts.EmbeddingDim <= 0 {
		return nil, errors.New("pgvector: EmbeddingDim must be greater than 0")
	}

	cfg, err := pgxpool.ParseConfig(opts.DSN)
	if err != nil {
		return nil, fmt.Errorf("pgvector: failed to parse DSN: %w", err)
	}

	if err := ensureExtension(ctx, opts.DSN); err != nil {
		return nil, fmt.Errorf("pgvector: install extension: %w", err)
	}

	cfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		return pgxvec.RegisterTypes(ctx, conn)
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("pgvector: failed to create connection pool: %w", err)
	}

	s := &Store{
		pool: pool,
	}

	if err := s.migrate(ctx, opts.EmbeddingDim); err != nil {
		pool.Close()
		return nil, fmt.Errorf("pgvector: failed to migrate: %w", err)
	}

	return s, nil
}

func ensureExtension(ctx context.Context, dsn string) error {
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	_, err = conn.Exec(ctx, "CREATE EXTENSION IF NOT EXISTS vector")
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) migrate(ctx context.Context, embeddingDim int) error {
	stmts := []string{
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS documents (
			id TEXT PRIMARY KEY,
			content TEXT NOT NULL,
			metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
			embedding vector(%d) NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT now())
			`, embeddingDim),
		`CREATE INDEX IF NOT EXISTS documents_embedding_idx 
			ON documents USING hnsw (embedding vector_cosine_ops)`,
	}

	for _, q := range stmts {
		if _, err := s.pool.Exec(ctx, q); err != nil {
			return fmt.Errorf("pgvector: failed to execute migration statement: %q -  %w", firstLine(q), err)
		}
	}

	return nil
}

func firstLine(s string) string {
	for i, c := range s {
		if c == '\n' {
			return s[:i]
		}
	}
	return s
}

func (s *Store) Upsert(ctx context.Context, docs []vector.Document) error {
	if len(docs) == 0 {
		return nil
	}

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("pgvector: failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	const stmt = `INSERT INTO documents (id, content, metadata, embedding)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (id) DO UPDATE SET
	content = EXCLUDED.content,
	metadata = EXCLUDED.metadata,
	embedding = EXCLUDED.embedding`

	for _, doc := range docs {
		meta, err := marshalMetadata(doc.Metadata)
		if err != nil {
			return fmt.Errorf("pgvector: failed to marshal metadata for %s: %w", doc.ID, err)
		}

		_, err = tx.Exec(ctx, stmt, doc.ID, doc.Content, meta, pgvector.NewVector(doc.Embedding))
		if err != nil {
			return fmt.Errorf("pgvector: failed to upsert document %s: %w", doc.ID, err)
		}
	}

	return tx.Commit(ctx)
}

func marshalMetadata(metadata map[string]string) ([]byte, error) {
	// Implement the logic to marshal metadata to JSONB format
	// For example, you can use json.Marshal(metadata)
	if len(metadata) == 0 {
		return []byte("{}"), nil
	}

	return json.Marshal(metadata)
}

func unmarshalMetadata(data []byte, dst *map[string]string) error {
	// Implement the logic to unmarshal JSONB data to metadata map
	// For example, you can use json.Unmarshal(data, &metadata)
	if len(data) == 0 {
		*dst = nil
		return nil
	}
	return json.Unmarshal(data, dst)
}

func (s *Store) Query(ctx context.Context, embedding []float32, topK int) ([]vector.Result, error) {
	if topK <= 0 {
		return nil, errors.New("pgvector: topK must be greater than 0")
	}

	const stmt = `SELECT id, content, metadata, embedding <=> $1 AS distance
	FROM documents
	ORDER BY embedding <=> $1
	LIMIT $2`

	rows, err := s.pool.Query(ctx, stmt, pgvector.NewVector(embedding), topK)
	if err != nil {
		return nil, fmt.Errorf("pgvector: failed to query documents: %w", err)
	}
	defer rows.Close()

	var results []vector.Result
	for rows.Next() {
		var (
			r        vector.Result
			metaRaw  []byte
			distance float64
		)

		if err := rows.Scan(&r.ID, &r.Content, &metaRaw, &distance); err != nil {
			return nil, fmt.Errorf("pgvector: failed to scan row: %w", err)
		}

		if err := unmarshalMetadata(metaRaw, &r.Metadata); err != nil {
			return nil, fmt.Errorf("pgvector: failed to unmarshal metadata for %s: %w", r.ID, err)
		}

		r.Score = float32(1 - distance) // Convert distance to similarity score
		results = append(results, r)
	}

	return results, rows.Err()
}
func (s *Store) Delete(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return nil
	}

	const stmt = `DELETE FROM documents WHERE id = ANY($1)`
	_, err := s.pool.Exec(ctx, stmt, ids)
	if err != nil {
		return fmt.Errorf("pgvector: failed to delete documents: %w", err)
	}

	return nil
}

func (s *Store) DeleteBySource(ctx context.Context, source string) error {
	if source == "" {
		return errors.New("pgvector: source is required")
	}

	const stmt = `DELETE FROM documents WHERE metadata->>'source' = $1`
	_, err := s.pool.Exec(ctx, stmt, source)
	if err != nil {
		return fmt.Errorf("pgvector: failed to delete documents by source: %w", err)
	}
	return nil
}

func (s *Store) Close() error {
	s.pool.Close()
	return nil
}
