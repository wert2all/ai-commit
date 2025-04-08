package ai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
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

func (p *OpenAIProvider) GenerateCommitMessage(projectContext, changes string) (string, error) {
	resp, err := p.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: p.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: `You are a commit message generator. Generate a concise and descriptive commit message 
					following the Conventional Commits specification (https://www.conventionalcommits.org/).
					The message should be in the format: type(scope): description
					where type is one of: feat, fix, docs, style, refactor, test, or chore.
					Analyze both the project context and git changes provided to generate an appropriate commit message.
					Consider the project structure, dependencies, and current branch when determining the scope.
					Return only the commit message, nothing else.`,
				},
				{
					Role: openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Project Context:\n\n%s\n\nChanges:\n\n%s", projectContext, changes),
				},
			},
			Temperature: 0.7,
			MaxTokens:  50,
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
