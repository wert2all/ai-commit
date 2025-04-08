package ai

import "fmt"

type ProviderType string

const (
	ProviderOpenAI ProviderType = "openai"
	// Add more providers here
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
	default:
		return nil, fmt.Errorf("unknown provider type: %s", config.Type)
	}
}
