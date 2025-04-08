package ai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
	"github.com/wert2all/ai-commit/changes"
)

type OpenAIProvider struct {
	client *openai.Client
	model  string
}

func NewOpenAIProvider(apiKey string, model string) *OpenAIProvider {
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}
	return &OpenAIProvider{
		client: openai.NewClient(apiKey),
		model:  model,
	}
}

func (p *OpenAIProvider) GenerateCommitMessage(projectContext string, changes changes.Changes) (string, error) {
	resp, err := p.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: p.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: SystemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Project Context:\n\n%s\n\nChanges:\n\n%s", projectContext, changes.ToString()),
				},
			},
			Temperature: 0.7,
			MaxTokens:   50,
		} ,
	)
	if err != nil {
		return "", fmt.Errorf("error calling OpenAI API: %v", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI API")
	}

	return resp.Choices[0].Message.Content, nil
}
