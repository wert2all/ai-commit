package project

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/wert2all/ai-commit/changes"
)

// SystemPrompt is the standard prompt for all AI providers
const systemPrompt = `You are a commit message generator. Generate a concise and descriptive commit message 
following the Conventional Commits specification (https://www.conventionalcommits.org/).
The message should be in the format: type(scope): description
where type is one of: feat, fix, docs, style, refactor, test, or chore.
Analyze both the project context and git changes provided to generate an appropriate commit message.
Consider the project structure, dependencies, and current branch when determining the scope.
Return only the commit message, nothing else.`

type (
	ProjectContext struct {
		Context      string
		SystemPrompt string
	}

	ContextBuilder interface {
		AddChanges() ContextBuilder
		AddLanguages() ContextBuilder
		Build() (*ProjectContext, error)
	}
	contextBuilderImpl struct {
		files     []string
		errors    []error
		changes   changes.Changes
		languages []string
	}
)

// AddLanguages implements ContextBuilder.
func (c contextBuilderImpl) AddLanguages() ContextBuilder {
	languages := make(map[string]bool)

	for _, file := range c.files {
		switch {
		case strings.HasSuffix(file, ".go"):
			languages["Go"] = true
		case strings.HasSuffix(file, ".js") || strings.HasSuffix(file, ".ts"):
			languages["JavaScript/TypeScript"] = true
		case strings.HasSuffix(file, ".py"):
			languages["Python"] = true
		case strings.HasSuffix(file, ".php"):
			languages["PHP"] = true
		case strings.HasSuffix(file, ".java"):
			languages["Java"] = true
		case strings.HasSuffix(file, ".rb"):
			languages["Ruby"] = true
		case strings.HasSuffix(file, ".rs"):
			languages["Rust"] = true
		}
	}

	for lang := range languages {
		c.languages = append(c.languages, lang)
	}

	return c
}

// AddChanges implements ContextBuilder.
func (c contextBuilderImpl) AddChanges() ContextBuilder {
	changes, err := changes.NewChanges()
	if err != nil {
		c.errors = append(c.errors, err)
	}
	c.changes = changes
	return c
}

// Build implements ContextBuilder.
func (c contextBuilderImpl) Build() (*ProjectContext, error) {
	if len(c.errors) != 0 {
		return nil, c.errors[0]
	}
	var context bytes.Buffer

	if len(c.languages) > 0 {
		context.WriteString("\n=== Project Languages ===\n")
		for _, lang := range c.languages {
			context.WriteString(lang + "\n")
		}
	}

	if c.changes != nil {
		context.WriteString("\n=== Changes ===\n")
		context.Write(c.changes.Value())
	}

	return &ProjectContext{
		Context:      context.String(),
		SystemPrompt: systemPrompt,
	}, nil
}

func NewBuilder(projectDir string) (ContextBuilder, error) {
	// Use provided project directory as repository root
	repoRoot := projectDir

	// Get project structure using git ls-files
	filesCmd := exec.Command("git", "ls-files")
	filesCmd.Dir = repoRoot
	filesOut, err := filesCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error getting git files: %v", err)
	}

	// Split output into lines and filter empty lines
	files := make([]string, 0)
	for _, file := range strings.Split(string(filesOut), "\n") {
		if file = strings.TrimSpace(file); file != "" {
			files = append(files, file)
		}
	}

	return contextBuilderImpl{
		errors:    make([]error, 0),
		files:     files,
		languages: make([]string, 0),
	}, nil
}
