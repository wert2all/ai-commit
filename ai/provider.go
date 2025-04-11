package ai

import (
	"github.com/wert2all/ai-commit/project"
)

// Provider defines the interface for AI providers
type Provider interface {
	GenerateCommitMessage(projectContext project.ProjectContext) (string, error)
}
