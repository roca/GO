package llm

import (
	"rag-course/config"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Client struct {
	cfg config.Config
	sdk openai.Client
}

func New(cfg config.Config) *Client {
	opts := []option.RequestOption{}

	if cfg.BaseURL != "" {
		opts = append(opts, option.WithBaseURL(cfg.BaseURL))
	}

	if cfg.APIKey != "" {
		opts = append(opts, option.WithAPIKey(cfg.APIKey))
	}

	sdk := openai.NewClient(opts...)

	return &Client{
		cfg: cfg,
		sdk: sdk,
	}
}
