package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/wert2all/ai-commit/changes"
)

type GeminiProvider struct {
	apiKey string
	model  string
}

type geminiRequest struct {
	Contents []content `json:"contents"`
	GenerationConfig generationConfig `json:"generationConfig"`
}

type content struct {
	Parts []part `json:"parts"`
}

type part struct {
	Text string `json:"text"`
}

type generationConfig struct {
	Temperature     float64 `json:"temperature"`
	MaxOutputTokens int     `json:"maxOutputTokens"`
}

type geminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func NewGeminiProvider(apiKey string, model string) *GeminiProvider {
	if model == "" {
		model = "gemini-pro" // default model
	}
	return &GeminiProvider{
		apiKey: apiKey,
		model:  model,
	}
}

func (p *GeminiProvider) GenerateCommitMessage(projectContext string, changes changes.Changes) (string, error) {
	req := geminiRequest{
		Contents: []content{
			{
				Parts: []part{
					{Text: SystemPrompt},
					{Text: fmt.Sprintf("Project Context:\n\n%s\n\nChanges:\n\n%s", projectContext, changes.ToString())},
				},
			},
		},
		GenerationConfig: generationConfig{
			Temperature:     0.7,
			MaxOutputTokens: 50,
		},
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1/models/%s:generateContent?key=%s", p.model, p.apiKey)
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	// nolint
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("error from Gemini API (status %d): %s", resp.StatusCode, string(body))
	}

	var result geminiResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	if len(result.Candidates) == 0 || len(result.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no response from Gemini API")
	}

	return result.Candidates[0].Content.Parts[0].Text, nil
}
