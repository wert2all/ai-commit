package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/wert2all/ai-commit/changes"
)

type LocalProvider struct {
	model    string
	endpoint string
}

func NewLocalProvider(options map[string]interface{}) (*LocalProvider, error) {
	model, ok := options["model"].(string)
	if !ok {
		model = "llama2" // default model
	}

	endpoint, ok := options["endpoint"].(string)
	if !ok {
		endpoint = "http://localhost:11434/api/generate" // default Ollama endpoint
	}

	return &LocalProvider{
		model:    model,
		endpoint: endpoint,
	}, nil
}

func (p *LocalProvider) GenerateCommitMessage(projectContext string, changes changes.Changes) (string, error) {
	panic("not implemented")
	// Prepare the prompt using GenerateCommitMessagePrompt
	//nolint
	prompt := GenerateCommitMessagePrompt(projectContext, changes.ToString())

	// Prepare request body
	requestBody, err := json.Marshal(map[string]interface{}{
		"model":  p.model,
		"prompt": prompt,
	})
	if err != nil {
		return "", err
	}

	// Send request to local AI
	resp, err := http.Post(p.endpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	// nolint
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse response
	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	// Extract generated text
	responseText, ok := response["response"].(string)
	if !ok {
		return "", fmt.Errorf("invalid response from local AI: %s", string(body))
	}

	return strings.TrimSpace(responseText), nil
}
