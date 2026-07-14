package indexdb

import (
	"example15/embedder"
	"fmt"
	"math"
	"sort"
)

type Embedding []float64

type VectorDB map[string]Embedding

func NewVectorDB() VectorDB {
	return make(VectorDB)
}

func (db VectorDB) AddDocument(content string, embedding Embedding) error {
	db[content] = embedding
	return nil
}

type Result struct {
	Content  string
	Distance float64
}

func (db VectorDB) Search(query string, topK int, distanceMetric string) ([]Result, error) {
	queryEmbedding, err := embedder.GetEmbeddings([]string{query})
	if err != nil {
		return nil, fmt.Errorf("failed to get embedding for query: %w", err)
	}

	// Find the top K closest embeddings
	var results []Result

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
		results = append(results, Result{content, dist})

		fmt.Println(dist, ":", content[0:min(100, len(content))])
	}

	// Sort results by distance
	sort.Slice(results, func(i, j int) bool {
		return results[i].Distance < results[j].Distance
	})

	// Return top K results
	var topResults []Result
	for i := 0; i < topK && i < len(results); i++ {
		topResults = append(topResults, results[i])
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
	cosineSim = max(-1.0, min(1.0, cosineSim)) // Clamp to [-1, 1] to avoid numerical issues

	return 1.0 - cosineSim, nil
}
