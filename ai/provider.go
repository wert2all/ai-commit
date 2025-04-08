package ai

// Provider defines the interface for AI providers
type Provider interface {
	GenerateCommitMessage(projectContext, changes string) (string, error)
}
