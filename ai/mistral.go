package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/wert2all/ai-commit/changes"
)

type MistralProvider struct {
	apiKey string
	model  string
}

type mistralRequest struct {
	Model       string    `json:"model"`
	Messages    []message `json:"messages"`
	Temperature float64   `json:"temperature"`
	MaxTokens   int       `json:"max_tokens"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type mistralResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func NewMistralProvider(apiKey string, model string) *MistralProvider {
	if model == "" {
		model = "mistral-medium" // default model
	}
	return &MistralProvider{
		apiKey: apiKey,
		model:  model,
	}
}

func (p *MistralProvider) GenerateCommitMessage(projectContext string, changes changes.Changes) (string, error) {
	req := mistralRequest{
		Model: p.model,
		Messages: []message{
			{
				Role:    "system",
				Content: SystemPrompt,
			},
			{
				Role:    "user",
				Content: fmt.Sprintf("Project Context:\n\n%s\n\nChanges:\n\n%s", projectContext, changes.ToString()),
			},
		},
		Temperature: 0.7,
		MaxTokens:   50,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	httpReq, err := http.NewRequest("POST", "https://api.mistral.ai/v1/chat/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+p.apiKey)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("error from Mistral API (status %d): %s", resp.StatusCode, string(body))
	}

	var result mistralResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("no response from Mistral API")
	}

	return result.Choices[0].Message.Content, nil
}
