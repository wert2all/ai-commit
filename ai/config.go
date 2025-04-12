package ai

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Options struct {
	WithCommit              bool
	WithChangedFilesContent bool
}

type Config struct {
	Directory string
	Type      ProviderType
	Endpoint  string
	APIKey    string
	Model     string
	Options   Options
}

func ReadConfig() (*Config, error) {
	providerName := flag.String("provider", "openai", "AI provider to use (openai, claude, mistral, gemini, openrouter, local)")
	model := flag.String("model", "", "Model to use (e.g., gpt-3.5-turbo, claude-2, mistral-medium, gemini-pro)")
	projectDir := flag.String("dir", ".", "Project directory path")
	endpoint := flag.String("endpoint", "", "Local provider endpoint1")
	withoutCommit := flag.Bool("without-commit", false, "commit a source after generate")
	withFilesContent := flag.Bool("with-files-content", false, "include content of changed files to context")

	flag.Parse()

	// Convert relative path to absolute
	absProjectDir, err := filepath.Abs(*projectDir)
	if err != nil {
		return nil, fmt.Errorf("error resolving project directory path: %v", err)
	}

	// Get API key based on provider
	apiKey, err := getAPIKey(*providerName)
	if err != nil {
		return nil, err
	}
	config := Config{
		Type:      ProviderType(*providerName),
		APIKey:    apiKey,
		Model:     *model,
		Endpoint:  *endpoint,
		Directory: absProjectDir,
		Options: Options{
			WithCommit:              !*withoutCommit,
			WithChangedFilesContent: *withFilesContent,
		},
	}
	return &config, nil
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

	case "openrouter":
		apiKey := os.Getenv("OPENROUTER_API_KEY")
		if apiKey == "" {
			return "", fmt.Errorf("OPENROUTER_API_KEY environment variable is not set")
		}
		return apiKey, nil
	case "local":
		// No API key needed for local provider
		return "", nil

	default:
		return "", fmt.Errorf("unknown provider: %s", providerName)
	}
}
