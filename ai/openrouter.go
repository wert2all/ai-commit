package ai

import "github.com/wert2all/ai-commit/ai/openai"

func NewOpenRouterProvider(apiKey string, model string) *openai.OpenAIProvider {
	if model == "" {
		model = "openrouter/optimus-alpha"
	}
	return openai.NewOpenAiProvider("https://openrouter.ai/api/v1", apiKey, model)
}
