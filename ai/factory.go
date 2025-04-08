package ai

import (
	"fmt"
	"os"
)

type ProviderType string

const (
	ProviderOpenAI  ProviderType = "openai"
	ProviderClaude  ProviderType = "claude"
	ProviderMistral ProviderType = "mistral"
	ProviderGemini  ProviderType = "gemini"
	ProviderLocal   ProviderType = "local"
)

type Config struct {
	Type    ProviderType
	APIKey  string
	Model   string
	Options map[string]any
}

func NewProvider(providerName string, model string) (Provider, error) {
	// Get API key based on provider
	apiKey, err := getAPIKey(providerName)
	if err != nil {
		return nil, err
	}

	config := Config{
		Type:   ProviderType(providerName),
		APIKey: apiKey,
		Model:  model,
	}

	switch config.Type {
	case ProviderOpenAI:
		return NewOpenAIProvider(config.APIKey, config.Model), nil
	case ProviderClaude:
		return NewClaudeProvider(config.APIKey, config.Model), nil
	case ProviderMistral:
		return NewMistralProvider(config.APIKey, config.Model), nil
	case ProviderGemini:
		return NewGeminiProvider(config.APIKey, config.Model), nil
	case ProviderLocal:
		provider, err := NewLocalProvider(config.Options)
		if err != nil {
			return nil, err
		}
		return provider, nil
	default:
		return nil, fmt.Errorf("unknown provider type: %s", config.Type)
	}
}

func getAPIKey(providerName string) (string, error) {
	switch providerName {
	case "openai":
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			return "", fmt.Errorf("OPENAI_API_KEY environment variable is not set")
		}
		return apiKey, nil

	case "claude":
		apiKey := os.Getenv("CLAUDE_API_KEY")
		if apiKey == "" {
			return "", fmt.Errorf("CLAUDE_API_KEY environment variable is not set")
		}
		return apiKey, nil

	case "mistral":
		apiKey := os.Getenv("MISTRAL_API_KEY")
		if apiKey == "" {
			return "", fmt.Errorf("MISTRAL_API_KEY environment variable is not set")
		}
		return apiKey, nil

	case "gemini":
		apiKey := os.Getenv("GEMINI_API_KEY")
		if apiKey == "" {
			return "", fmt.Errorf("GEMINI_API_KEY environment variable is not set")
		}
		return apiKey, nil

	case "local":
		// No API key needed for local provider
		return "", nil

	default:
		return "", fmt.Errorf("Unknown provider: %s", providerName)
	}
}
