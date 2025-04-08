package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MistralProvider struct {
	apiKey string
	model  string
}

type mistralRequest struct {
	Model       string   `json:"model"`
	Messages    []message `json:"messages"`
	Temperature float64  `json:"temperature"`
	MaxTokens   int      `json:"max_tokens"`
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

func (p *MistralProvider) GenerateCommitMessage(projectContext, changes string) (string, error) {
	req := mistralRequest{
		Model: p.model,
		Messages: []message{
			{
				Role: "system",
				Content: "You are a commit message generator. Generate a concise and descriptive commit message " +
					"following the Conventional Commits specification. The message should be in the format: " +
					"type(scope): description where type is one of: feat, fix, docs, style, refactor, test, or chore. " +
					"Consider the project structure, dependencies, and current branch when determining the scope. " +
					"Return only the commit message, nothing else.",
			},
			{
				Role:    "user",
				Content: fmt.Sprintf("Project Context:\n\n%s\n\nChanges:\n\n%s", projectContext, changes),
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
