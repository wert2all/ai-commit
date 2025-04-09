package project

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
		AddGitBranch() ContextBuilder

		Build() (*ProjectContext, error)
	}
	contextBuilderImpl struct {
		files     []string
		errors    []error
		changes   changes.Changes
		languages []string
		branch    *string
	}
)

// AddGitBranch implements ContextBuilder.
func (c contextBuilderImpl) AddGitBranch() ContextBuilder {
	// Get git branch info
	branchCmd := exec.Command("git", "branch", "--show-current")
	branchOut, err := branchCmd.Output()
	if err != nil {
		c.errors = append(c.errors, err)
	}
	branchString := string(branchOut[:])
	c.branch = &branchString
	return c
}

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

	context.WriteString("\n=== Project Structure ===\n")
	for _, file := range c.files {
		context.WriteString(file + "\n")
	}

	if len(c.languages) > 0 {
		context.WriteString("\n=== Project Languages ===\n")
		for _, lang := range c.languages {
			context.WriteString(lang + "\n")
		}
	}

	if c.branch != nil {
		context.WriteString("\n=== Git branch ===\n")
		context.WriteString(*c.branch)
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

// nolint
func readFileContent(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(content)
}

// nolint
func getProjectConfig(repoRoot string, files []string) map[string]string {
	config := make(map[string]string)

	// Check for various config files
	for _, file := range files {
		fullPath := filepath.Join(repoRoot, file)
		switch filepath.Base(file) {
		case ".env":
			config["Environment Config"] = readFileContent(fullPath)
		case "docker-compose.yml", "docker-compose.yaml":
			config["Docker Compose Config"] = readFileContent(fullPath)
		case "Dockerfile":
			config["Dockerfile"] = readFileContent(fullPath)
		case ".gitlab-ci.yml", ".github/workflows/":
			config["CI/CD Config"] = readFileContent(fullPath)
		case "nginx.conf":
			config["Nginx Config"] = readFileContent(fullPath)
		case "webpack.config.js":
			config["Webpack Config"] = readFileContent(fullPath)
		}
	}

	return config
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
		branch:    nil,
	}, nil
}
