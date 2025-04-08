package ai

import "fmt"

// SystemPrompt is the standard prompt for all AI providers
const SystemPrompt = `You are a commit message generator. Generate a concise and descriptive commit message 
following the Conventional Commits specification (https://www.conventionalcommits.org/).
The message should be in the format: type(scope): description
where type is one of: feat, fix, docs, style, refactor, test, or chore.
Analyze both the project context and git changes provided to generate an appropriate commit message.
Consider the project structure, dependencies, and current branch when determining the scope.
Return only the commit message, nothing else.`

// Provider defines the interface for AI providers
type Provider interface {
	GenerateCommitMessage(projectContext, changes string) (string, error)
}

// GenerateCommitMessagePrompt creates a standardized prompt for commit message generation
func GenerateCommitMessagePrompt(projectContext, changes string) string {
	return fmt.Sprintf(`%s

Project Context:
%s

Git Changes:
%s`, SystemPrompt, projectContext, changes)
}
