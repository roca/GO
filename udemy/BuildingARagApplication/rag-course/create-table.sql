CREATE EXTENSION IF NOT EXISTS vector;

CREATE TABLE documents (
    id text PRIMARY KEY,
    content text NOT NULL,
    metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
    embedding vector(768) NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now()
);

CREATE INDEX documents_embedding_idx ON documents USING hnsw (embedding vector_cosine_ops);
