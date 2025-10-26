package indexdb

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"
)

// Document represents a document with content and metadata
type Document map[string]interface{}

// TokenizerFunc defines the signature for tokenizer functions
type TokenizerFunc func(string) []string

// BM25Result represents a search result with document and normalized score
type BM25Result struct {
	Document Document
	Score    float64
}

// BM25Index implements BM25 ranking function for document retrieval
type BM25Index struct {
	documents    []Document
	corpusTokens [][]string
	docLen       []int
	docFreqs     map[string]int
	avgDocLen    float64
	idf          map[string]float64
	indexBuilt   bool
	k1           float64
	b            float64
	tokenizer    TokenizerFunc
}

// NewBM25Index creates a new BM25Index with default parameters
func NewBM25Index(k1, b float64, tokenizer TokenizerFunc) *BM25Index {
	index := &BM25Index{
		documents:    make([]Document, 0),
		corpusTokens: make([][]string, 0),
		docLen:       make([]int, 0),
		docFreqs:     make(map[string]int),
		avgDocLen:    0.0,
		idf:          make(map[string]float64),
		indexBuilt:   false,
		k1:           k1,
		b:            b,
	}

	if tokenizer != nil {
		index.tokenizer = tokenizer
	} else {
		index.tokenizer = index.defaultTokenizer
	}

	return index
}

// defaultTokenizer splits text into tokens using regex
func (idx *BM25Index) defaultTokenizer(text string) []string {
	text = strings.ToLower(text)
	re := regexp.MustCompile(`\W+`)
	tokens := re.Split(text, -1)

	var result []string
	for _, token := range tokens {
		if token != "" {
			result = append(result, token)
		}
	}
	return result
}

// updateStatsAdd updates internal statistics when adding a document
func (idx *BM25Index) updateStatsAdd(docTokens []string) {
	idx.docLen = append(idx.docLen, len(docTokens))

	seenInDoc := make(map[string]bool)
	for _, token := range docTokens {
		if !seenInDoc[token] {
			idx.docFreqs[token]++
			seenInDoc[token] = true
		}
	}

	idx.indexBuilt = false
}

// calculateIDF computes inverse document frequency for all terms
func (idx *BM25Index) calculateIDF() {
	N := float64(len(idx.documents))
	idx.idf = make(map[string]float64)

	for term, freq := range idx.docFreqs {
		freqFloat := float64(freq)
		idfScore := math.Log(((N - freqFloat + 0.5) / (freqFloat + 0.5)) + 1)
		idx.idf[term] = idfScore
	}
}

// buildIndex builds the search index
func (idx *BM25Index) buildIndex() {
	if len(idx.documents) == 0 {
		idx.avgDocLen = 0.0
		idx.idf = make(map[string]float64)
		idx.indexBuilt = true
		return
	}

	totalDocLen := 0
	for _, length := range idx.docLen {
		totalDocLen += length
	}
	idx.avgDocLen = float64(totalDocLen) / float64(len(idx.documents))

	idx.calculateIDF()
	idx.indexBuilt = true
}

// AddDocument adds a document to the index
func (idx *BM25Index) AddDocument(key string, text any) error {
	document := Document{key: text.(string)}

	if document == nil {
		return fmt.Errorf("document must not be nil")
	}

	contentInterface, exists := document["content"]
	if !exists {
		return fmt.Errorf("document dictionary must contain a 'content' key")
	}

	content, ok := contentInterface.(string)
	if !ok {
		return fmt.Errorf("document 'content' must be a string")
	}

	docTokens := idx.tokenizer(content)

	idx.documents = append(idx.documents, document)
	idx.corpusTokens = append(idx.corpusTokens, docTokens)
	idx.updateStatsAdd(docTokens)

	return nil
}

// computeBM25Score computes BM25 score for a document given query tokens
func (idx *BM25Index) computeBM25Score(queryTokens []string, docIndex int) float64 {
	score := 0.0
	docTokens := idx.corpusTokens[docIndex]
	docLength := float64(idx.docLen[docIndex])

	// Count term frequencies in document
	docTermCounts := make(map[string]int)
	for _, token := range docTokens {
		docTermCounts[token]++
	}

	for _, token := range queryTokens {
		idf, exists := idx.idf[token]
		if !exists {
			continue
		}

		termFreq := float64(docTermCounts[token])

		numerator := idf * termFreq * (idx.k1 + 1)
		denominator := termFreq + idx.k1*(1-idx.b+idx.b*(docLength/idx.avgDocLen))
		score += numerator / (denominator + 1e-9)
	}

	return score
}

// Search performs BM25 search and returns top k results
func (idx *BM25Index) Search(queryText string, k int, scoreNormalizationFactor float64) ([]BM25Result, error) {
	if len(idx.documents) == 0 {
		return []BM25Result{}, nil
	}

	if queryText == "" {
		return nil, fmt.Errorf("query text must not be empty")
	}

	if k <= 0 {
		return nil, fmt.Errorf("k must be a positive integer")
	}

	if !idx.indexBuilt {
		idx.buildIndex()
	}

	if idx.avgDocLen == 0 {
		return []BM25Result{}, nil
	}

	queryTokens := idx.tokenizer(queryText)
	if len(queryTokens) == 0 {
		return []BM25Result{}, nil
	}

	type scoreDoc struct {
		score float64
		doc   Document
	}

	var rawScores []scoreDoc
	for i := 0; i < len(idx.documents); i++ {
		rawScore := idx.computeBM25Score(queryTokens, i)
		if rawScore > 1e-9 {
			rawScores = append(rawScores, scoreDoc{rawScore, idx.documents[i]})
		}
	}

	// Sort by score descending
	sort.Slice(rawScores, func(i, j int) bool {
		return rawScores[i].score > rawScores[j].score
	})

	// Take top k and normalize scores
	var normalizedResults []BM25Result
	limit := k
	if limit > len(rawScores) {
		limit = len(rawScores)
	}

	for i := 0; i < limit; i++ {
		rawScore := rawScores[i].score
		normalizedScore := math.Exp(-scoreNormalizationFactor * rawScore)
		normalizedResults = append(normalizedResults, BM25Result{
			Document: rawScores[i].doc,
			Score:    normalizedScore,
		})
	}

	// Sort by normalized score ascending (lower is better)
	sort.Slice(normalizedResults, func(i, j int) bool {
		return normalizedResults[i].Score < normalizedResults[j].Score
	})

	return normalizedResults, nil
}

// Len returns the number of documents in the index
func (idx *BM25Index) Len() int {
	return len(idx.documents)
}

// String returns a string representation of the index
func (idx *BM25Index) String() string {
	return fmt.Sprintf("BM25Index(count=%d, k1=%.2f, b=%.2f, index_built=%t)",
		len(idx.documents), idx.k1, idx.b, idx.indexBuilt)
}
