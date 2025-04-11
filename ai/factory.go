package ai

import (
	"fmt"
)

type ProviderType string

const (
	ProviderOpenAI  ProviderType = "openai"
	ProviderClaude  ProviderType = "claude"
	ProviderMistral ProviderType = "mistral"
	ProviderGemini  ProviderType = "gemini"
	ProviderLocal   ProviderType = "local"
)

func NewProvider(config Config) (Provider, error) {
	switch config.Type {
	case ProviderMistral:
		return NewMistralProvider(config.APIKey, config.Model), nil
	case ProviderOpenAI:
		return NewOpenAIProvider(config.APIKey, config.Model), nil
	case ProviderClaude:
		return NewClaudeProvider(config.APIKey, config.Model), nil
	case ProviderGemini:
		return NewGeminiProvider(config.APIKey, config.Model), nil
	case ProviderLocal:
		provider, err := NewLocalProvider(config.Endpoint, config.Model)
		if err != nil {
			return nil, err
		}
		return provider, nil
	default:
		return nil, fmt.Errorf("unknown provider type: %s", config.Type)
	}
}
