package llm

import (
	"context"
	"fmt"

	"github.com/openai/openai-go/v3"
)

type Embedder interface {
	Embed(ctx context.Context, texts []string) ([][]float32, error)
}

func (c *Client) Embed(ctx context.Context, texts []string) ([][]float32, error) {
	if len(texts) == 0 {
		return nil, nil
	}

	resp, err := c.sdk.Embeddings.New(ctx, openai.EmbeddingNewParams{
		Model: c.cfg.EmbeddingModel,
		Input: openai.EmbeddingNewParamsInputUnion{OfArrayOfStrings: texts},
	})
	if err != nil {
		return nil, err
	}

	if len(resp.Data) != len(texts) {
		return nil, fmt.Errorf("embeddings: expected %d vectors; git %d", len(texts), len(resp.Data))
	}

	vecs := make([][]float32, len(texts))
	for _, d := range resp.Data {
		idx := int(d.Index)
		if idx < 0 || idx >= len(vecs) {
			return nil, fmt.Errorf("embeddings: index %d out of range", idx)
		}
		vec := make([]float32, len(d.Embedding))
		for i, f := range d.Embedding {
			vec[i] = float32(f)
		}
		vecs[idx] = vec
	}

	return vecs, nil
}
