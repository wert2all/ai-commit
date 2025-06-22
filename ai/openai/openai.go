package openai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
	"github.com/wert2all/ai-commit/project"
)

type (
	OpenAIProvider struct {
		Client *openai.Client
		Model  string
	}
)

func (p *OpenAIProvider) GenerateCommitMessage(projectContext project.ProjectContext) (string, error) {
	resp, err := p.Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: p.Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: projectContext.SystemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Project Context:\n\n%s\n\n", projectContext.Context),
				},
			},
			Temperature: 0.7,
			MaxTokens:   50,
		},
	)
	if err != nil {
		return "", fmt.Errorf("error calling OpenAI API: %v", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI API")
	}

	return resp.Choices[0].Message.Content, nil
}

func NewOpenAiProvider(baseURL string, apiKey string, model string) *OpenAIProvider {
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = baseURL

	return &OpenAIProvider{
		Client: openai.NewClientWithConfig(config),
		Model:  model,
	}
}
