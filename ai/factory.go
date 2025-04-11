package ai

import (
	"fmt"

	"github.com/wert2all/ai-commit/ai/openai"
	"github.com/wert2all/ai-commit/project"
)

type (
	ProviderType string
	Provider     interface {
		GenerateCommitMessage(projectContext project.ProjectContext) (string, error)
	}
)

const (
	ProviderOpenAI     ProviderType = "openai"
	ProviderClaude     ProviderType = "claude"
	ProviderMistral    ProviderType = "mistral"
	ProviderGemini     ProviderType = "gemini"
	ProviderOpenRouter ProviderType = "openrouter"
	ProviderLocal      ProviderType = "local"
)

func NewProvider(config Config) (Provider, error) {
	switch config.Type {
	case ProviderMistral:
		return NewMistralProvider(config.APIKey, config.Model), nil
	case ProviderOpenAI:
		return openai.NewGPTProvider(config.APIKey, config.Model), nil
	case ProviderClaude:
		return NewClaudeProvider(config.APIKey, config.Model), nil
	case ProviderGemini:
		return NewGeminiProvider(config.APIKey, config.Model), nil
	case ProviderOpenRouter:
		return NewOpenRouterProvider(config.APIKey, config.Model), nil
	case ProviderLocal:
		return NewLocalProvider(config.Endpoint, config.Model), nil
	default:
		return nil, fmt.Errorf("unknown provider type: %s", config.Type)
	}
}
