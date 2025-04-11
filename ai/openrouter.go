package ai

import (
	"github.com/sashabaranov/go-openai"
)

func NewOpenRouterProvider(apiKey string, model string) *OpenAIProvider {
	if model == "" {
		model = "openrouter/optimus-alpha"
	}

	config := openai.DefaultConfig(apiKey)
	config.BaseURL = "https://openrouter.ai/api/v1"

	return &OpenAIProvider{
		client: openai.NewClientWithConfig(config),
		model:  model,
	}
}
