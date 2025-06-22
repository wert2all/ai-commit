package ai

func NewOpenRouterProvider(apiKey string, model string) *OpenAIProvider {
	if model == "" {
		model = "openrouter/optimus-alpha"
	}
	return NewOpenAiProvider("https://openrouter.ai/api/v1", apiKey, model)
}
