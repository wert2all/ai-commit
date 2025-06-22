package ai

import (
	"github.com/sashabaranov/go-openai"
	oai "github.com/wert2all/ai-commit/ai/openai"
)

func NewGPTProvider(apiKey string, model string) *oai.OpenAIProvider {
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}
	return &oai.OpenAIProvider{
		Client: openai.NewClient(apiKey),
		Model:  model,
	}
}
