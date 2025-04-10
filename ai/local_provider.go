package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/wert2all/ai-commit/project"
)

type LocalProvider struct {
	model    string
	endpoint string
}

func NewLocalProvider(endpoint string, model string) (*LocalProvider, error) {
	return &LocalProvider{
		model:    model,
		endpoint: endpoint + "/api/generate",
	}, nil
}

func (p *LocalProvider) GenerateCommitMessage(projectContext project.ProjectContext) (string, error) {
	// Prepare request body
	requestBody, err := json.Marshal(map[string]any{
		"model":  p.model,
		"prompt": generatePrompt(projectContext),
	})
	if err != nil {
		return "", err
	}

	// Send request to local AI
	resp, err := http.Post(p.endpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("error responce from ollama: %s", resp.Status)
	}

	// nolint
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return parseOllamaResponse(body)
}

func generatePrompt(projectContext project.ProjectContext) string {
	return "\n" + projectContext.SystemPrompt + "\n\n" + projectContext.Context
}

func parseOllamaResponse(body []byte) (string, error) {
	// Split the response by newlines to get individual JSON objects
	lines := strings.Split(string(body), "\n")

	var fullResponse strings.Builder

	for _, line := range lines {
		if line == "" {
			continue
		}

		// Parse each JSON object
		var respObj map[string]any
		if err := json.Unmarshal([]byte(line), &respObj); err != nil {
			return "", fmt.Errorf("failed to parse JSON: %w", err)
		}

		// Extract the token from the "response" field
		if token, ok := respObj["response"].(string); ok {
			fullResponse.WriteString(token)
		}

	}
	return strings.TrimSpace(fullResponse.String()), nil
}
