package vectordb

import (
	"example15/embedder"
	"fmt"
	"math"
	"sort"
)

type Embedding []float64

type VectorDB map[string]Embedding

func (db VectorDB) Add(content string, embedding Embedding) {
	db[content] = embedding
}

func (db VectorDB) Search(query string, topK int, distanceMetric string) ([]string, error) {
	queryEmbedding, err := embedder.GetEmbeddings([]string{query})
	if err != nil {
		return nil, fmt.Errorf("failed to get embedding for query: %w", err)
	}

	// Find the top K closest embeddings
	type result struct {
		content  string
		distance float64
	}

	var results []result

	for content, embedding := range db {
		var dist float64
		switch distanceMetric {
		case "euclidean":
			dist, err = euclideanDistance(embedding, queryEmbedding[0])
		case "cosine":
			dist, err = cosineDistance(embedding, queryEmbedding[0])
		default:
			return nil, fmt.Errorf("unsupported distance metric: %s", distanceMetric)
		}
		if err != nil {
			return nil, fmt.Errorf("failed to compute distance: %w", err)
		}
		results = append(results, result{content, dist})
	}

	// Sort results by distance
	sort.Slice(results, func(i, j int) bool {
		return results[i].distance < results[j].distance
	})

	// Return top K results
	var topResults []string
	for i := 0; i < topK && i < len(results); i++ {
		topResults = append(topResults, results[i].content)
	}

	return topResults, nil
}

func euclideanDistance(a, b Embedding) (float64, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("embeddings must be of the same length")
	}
	var sum float64
	for i := range a {
		diff := a[i] - b[i]
		sum += diff * diff
	}
	return math.Sqrt(sum), nil

}

func dotProduct(a, b Embedding) (float64, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("embeddings must be of the same length")
	}
	var sum float64
	for i := range a {
		sum += a[i] * b[i]
	}
	return sum, nil
}

func magnitude(a Embedding) float64 {
	var sum float64
	for _, v := range a {
		sum += v * v
	}
	return math.Sqrt(sum)
}

func cosineDistance(a, b Embedding) (float64, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("embeddings must be of the same length")
	}
	dot, err := dotProduct(a, b)
	if err != nil {
		return 0, err
	}
	magA := magnitude(a)
	magB := magnitude(b)
	if magA == 0 || magB == 0 {
		return 0, fmt.Errorf("magnitude of embedding cannot be zero")
	}

	cosineSim := dot / (magA * magB)
	cosineSim = math.Min(math.Max(cosineSim, -1), 1) // Clamp to [-1, 1] to avoid numerical issues

	return 1.0 - cosineSim, nil
}
