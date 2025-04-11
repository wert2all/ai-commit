package ai

import (
	"github.com/sashabaranov/go-openai"
)

type (
	OpenRouterProvider struct {
		client *openai.Client
		model  string
	}
)

func NewOpenRouterProvider(apiKey string, model string) *OpenRouterProvider {
	if model == "" {
		model = "optimus-alpha" // default model
	}

	config := openai.DefaultConfig(apiKey)
	config.BaseURL = "https://openrouter.ai/api/v1"

	openai.NewClientWithConfig(config)
	return &OpenRouterProvider{
		client: openai.NewClientWithConfig(config),
		model:  model,
	}
}
