package ai

import "fmt"

type ProviderType string

const (
	ProviderOpenAI   ProviderType = "openai"
	ProviderClaude  ProviderType = "claude"
	ProviderMistral ProviderType = "mistral"
	ProviderGemini  ProviderType = "gemini"
)

type Config struct {
	Type    ProviderType
	APIKey  string
	Model   string
	Options map[string]interface{}
}

func NewProvider(config Config) (Provider, error) {
	switch config.Type {
	case ProviderOpenAI:
		return NewOpenAIProvider(config.APIKey, config.Model), nil
	case ProviderClaude:
		return NewClaudeProvider(config.APIKey, config.Model), nil
	case ProviderMistral:
		return NewMistralProvider(config.APIKey, config.Model), nil
	case ProviderGemini:
		return NewGeminiProvider(config.APIKey, config.Model), nil
	default:
		return nil, fmt.Errorf("unknown provider type: %s", config.Type)
	}
}
