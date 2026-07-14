package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const lucidBaseURL = "https://api.lucid.co"

type LucidClient struct {
	oauth      *OAuthClient
	httpClient *http.Client
	baseURL    string
}

func NewLucidClient(oauth *OAuthClient) *LucidClient {
	return &LucidClient{
		oauth:      oauth,
		httpClient: &http.Client{Timeout: 30 * time.Second},
		baseURL:    lucidBaseURL,
	}
}

func (c *LucidClient) doRequest(ctx context.Context, method, path string, body interface{}, extraHeaders map[string]string) (*http.Response, error) {
	token, err := c.oauth.GetValidToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get valid token: %w", err)
	}

	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Lucid-Api-Version", "1")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	for k, v := range extraHeaders {
		req.Header.Set(k, v)
	}

	return c.httpClient.Do(req)
}

type DocumentSummary struct {
	DocumentID   string `json:"documentId"`
	Title        string `json:"title"`
	EditURL      string `json:"editUrl"`
	ViewURL      string `json:"viewUrl"`
	Product      string `json:"product"`
	LastModified string `json:"lastModified"`
}

type SearchResult struct {
	Documents  []DocumentSummary `json:"documents"`
	PageSize   int               `json:"pageSize"`
	PageNumber int               `json:"pageNumber"`
	TotalHits  int               `json:"totalHits"`
}

type Document struct {
	DocumentID   string `json:"documentId"`
	Title        string `json:"title"`
	EditURL      string `json:"editUrl"`
	ViewURL      string `json:"viewUrl"`
	Product      string `json:"product"`
	LastModified string `json:"lastModified"`
	Owner        string `json:"owner"`
	PageCount    int    `json:"pageCount"`
}

func (c *LucidClient) SearchDocuments(ctx context.Context, keywords string, pageSize, pageNumber int) (*SearchResult, error) {
	body := map[string]interface{}{}
	if keywords != "" {
		body["keywords"] = keywords
	}
	if pageSize > 0 {
		body["pageSize"] = pageSize
	}
	if pageNumber > 0 {
		body["pageNumber"] = pageNumber
	}

	resp, err := c.doRequest(ctx, "POST", "/documents/search", body, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search documents returned %d: %s", resp.StatusCode, string(respBody))
	}

	var result SearchResult
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return &result, nil
}

func (c *LucidClient) GetDocument(ctx context.Context, documentID string) (*Document, error) {
	resp, err := c.doRequest(ctx, "GET", "/documents/"+documentID, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get document returned %d: %s", resp.StatusCode, string(respBody))
	}

	var doc Document
	if err := json.Unmarshal(respBody, &doc); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return &doc, nil
}

func (c *LucidClient) ExportDocument(ctx context.Context, documentID string, format string) ([]byte, string, error) {
	mimeType := "image/png"
	if format == "pdf" {
		mimeType = "application/pdf"
	}

	resp, err := c.doRequest(ctx, "GET", "/documents/"+documentID, nil, map[string]string{
		"Accept": mimeType,
	})
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errBody, _ := io.ReadAll(resp.Body)
		return nil, "", fmt.Errorf("export document returned %d: %s", resp.StatusCode, string(errBody))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("failed to read export data: %w", err)
	}

	return data, mimeType, nil
}
