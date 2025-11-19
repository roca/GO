# Copilot Instructions for RAG Search System

## Project Architecture

This is a retrieval-augmented generation (RAG) system implementing both **vector similarity search** and **BM25 lexical search** for document retrieval. The system processes documents through a chunking → embedding → indexing → search pipeline.

### Core Components

- **`main.go`**: Entry point implementing the full RAG pipeline with three chunking strategies
- **`embedder/`**: VoyageAI embeddings integration (requires `VOYAGE_API_KEY` env var)
- **`indexdb/`**: Dual search backends - vector similarity (`vectordb.go`) and BM25 (`bm25db.go`)
- **Jupyter notebooks** (`001-004_*.ipynb`): Progressive Python implementations of each component

### Dual Implementation Pattern

The project maintains **parallel Go and Python implementations**:
- Go code in packages provides production-ready implementations
- Jupyter notebooks provide interactive exploration and prototyping
- Both share similar interfaces but use different external dependencies (VoyageAI Go SDK vs Python implementations)

## Key Patterns & Conventions

### Document Processing Pipeline
```go
// Standard workflow in main.go
chunks := ChunkBySection(text)              // 1. Chunk document
embeddings := embedder.GetEmbeddings(chunks) // 2. Generate embeddings  
vectordb.Add(chunks[i], embeddings[i])       // 3. Index in vector store
bm25db.Add(chunks[i])                        // 4. Index in BM25 store
results := vectordb.Search(query, k, metric) // 5. Search both stores
```

### Chunking Strategies
Three chunking methods implemented in `main.go`:
- `ChunkBySection()`: Split on `\n## ` (markdown headers) - **default strategy**
- `ChunkByChar()`: Fixed character windows with overlap
- `ChunkBySentence()`: Sentence-based chunking with overlap

### Search Interface Consistency
Both search backends follow the same pattern:
```go
type Result struct { Content string; Distance float64 }
Search(query string, topK int, params...) ([]Result, error)
```

### Error Handling & Environment
- VoyageAI API requires `VOYAGE_API_KEY` environment variable
- Go modules use `github.com/vendasta/langchaingo` (not `tmc/langchaingo`)
- Python notebooks implement from scratch (no external ML libraries)

## Development Workflow

### Running the System
```bash
export VOYAGE_API_KEY="your-key-here"
go run main.go  # Processes report.md and demonstrates search
```

### Testing Changes
- Modify chunking strategy by swapping function calls in `main.go`
- Test embeddings independently via `embedder.PrintEmbeddings()`
- Use notebooks for interactive experimentation before implementing in Go

### Adding New Search Methods
1. Implement in `indexdb/` package following the `VectorDB`/`Bm25DB` interface pattern
2. Add to `main.go` pipeline alongside existing search backends
3. Consider creating corresponding notebook for algorithm development

## File Dependencies

- **`report.md`**: Test document (interdisciplinary research report with structured sections)
- **VoyageAI SDK**: Critical dependency for embeddings generation
- **No external vector databases**: Custom in-memory implementations for both vector and BM25 search

## Search Behavior Notes

- **Vector search**: Returns similarity scores (lower = more similar)
- **BM25 search**: Returns relevance scores (higher = more relevant)  
- Both support configurable result limits (`k` parameter)
- Vector search supports multiple distance metrics (`cosine`, `euclidean`)