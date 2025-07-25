package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/wert2all/ai-commit/project"
)

type ClaudeProvider struct {
	apiKey string
	model  string
}

// GetProviderInfo implements Provider.
func (p *ClaudeProvider) GetProviderInfo() ProviderInfo {
	return ProviderInfo{Name: "Anthropic Claude", Model: p.model}
}

type claudeRequest struct {
	Model       string   `json:"model"`
	Prompt      string   `json:"prompt"`
	MaxTokens   int      `json:"max_tokens_to_sample"`
	Temperature float64  `json:"temperature"`
	Stop        []string `json:"stop_sequences"`
}

type claudeResponse struct {
	Completion string `json:"completion"`
	Stop       string `json:"stop"`
	Model      string `json:"model"`
}

func (p *ClaudeProvider) GenerateCommitMessage(projectContext project.ProjectContext) (string, error) {
	req := claudeRequest{
		Model:       p.model,
		Prompt:      fmt.Sprintf("\n\nHuman: Project Context:\n\n%s\n\nAssistant: %s", projectContext.Context, projectContext.SystemPrompt),
		MaxTokens:   50,
		Temperature: 0.7,
		Stop:        []string{"\n", "Human:"},
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	httpReq, err := http.NewRequest("POST", "https://api.anthropic.com/v1/complete", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-API-Key", p.apiKey)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	// nolint
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("error from Claude API (status %d): %s", resp.StatusCode, string(body))
	}

	var result claudeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	return result.Completion, nil
}

func NewClaudeProvider(apiKey string, model string) *ClaudeProvider {
	if model == "" {
		model = "claude-3-7-sonnet-latest" // default model
	}
	return &ClaudeProvider{
		apiKey: apiKey,
		model:  model,
	}
}
