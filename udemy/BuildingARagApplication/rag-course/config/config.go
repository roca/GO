package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BaseURL          string
	APIKey           string
	Model            string
	SystemPromptFile string
}

func Load() Config {
	_ = godotenv.Load() // Load environment variables from .env file if it exists

	cfg := Config{
		BaseURL:          os.Getenv("OPENAI_BASE_URL"),
		APIKey:           os.Getenv("OPENAI_API_KEY"),
		Model:            os.Getenv("OPENAI_MODEL"),
		SystemPromptFile: os.Getenv("SYSTEM_PROMPT_FILE"),
	}

	if cfg.BaseURL == "" {
		cfg.BaseURL = "https://api.openai.com/v1" // Default base URL
	}

	if cfg.Model == "" {
		cfg.Model = "gpt-4o-mini" // Default model
	}

	return cfg
}
