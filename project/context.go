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
const systemPrompt = `You are an expert commit message generator. Generate a concise, descriptive, and semantically meaningful commit message strictly following the Conventional Commits specification (https://www.conventionalcommits.org/).

FORMAT: <type>[optional scope]: <description>

TYPES:
- feat: A new feature or significant enhancement
- fix: A bug fix
- docs: Documentation changes only
- style: Code style changes (formatting, missing semi-colons, etc; no code change)
- refactor: Code changes that neither fix bugs nor add features
- perf: Performance improvements
- test: Adding or correcting tests
- build: Changes to build system or external dependencies
- ci: Changes to CI configuration files and scripts
- chore: Other changes that don't modify src or test files
- revert: Reverts a previous commit

SCOPE:
- Optional parenthesized noun describing the section of the codebase affected
- Use lowercase with hyphens for multi-word scopes
- Common examples: core, ui, api, auth, data, utils, testing

DESCRIPTION:
- Use imperative, present tense (e.g., "change" not "changed" or "changes")
- Don't capitalize first letter
- No period at the end
- Maximum 72 characters
- Be specific and clear about what changed and why

BREAKING CHANGES:
- Add "!" after type/scope to indicate breaking changes (e.g., feat(api)!: remove user endpoints)
- Include BREAKING CHANGE: in the body for details

ANALYSIS PROCESS:
1. Examine file paths and extensions to identify affected components
2. Review code changes to determine type of change
3. Analyze diff content to understand what functionality was modified
4. Consider project structure to determine appropriate scope
5. Check branch name for additional context

Return ONLY the commit message without any explanations, markdown, or additional text.`

type (
	ProjectContext struct {
		Context      string
		SystemPrompt string
	}

	ContextBuilder interface {
		AddChanges()
		AddLanguages()
		AddGitBranch()
		AddChangedFilesContent()

		Build() (*ProjectContext, error)
	}
	contextBuilderImpl struct {
		files               []string
		errors              []error
		changes             changes.Changes
		changedFilesContent map[string]string
		languages           []string
		branch              *string
	}
)

// AddGitBranch implements ContextBuilder.
func (c *contextBuilderImpl) AddGitBranch() {
	// Get git branch info
	branchCmd := exec.Command("git", "branch", "--show-current")
	branchOut, err := branchCmd.Output()
	if err != nil {
		c.errors = append(c.errors, err)
	}
	branchString := string(branchOut[:])
	c.branch = &branchString
}

// AddLanguages implements ContextBuilder.
func (c *contextBuilderImpl) AddLanguages() {
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
}

// AddChanges implements ContextBuilder.
func (c *contextBuilderImpl) AddChanges() {
	changes, err := changes.NewChanges()
	if err != nil {
		c.errors = append(c.errors, err)
	}
	c.changes = changes
}

// Build implements ContextBuilder.
func (c *contextBuilderImpl) Build() (*ProjectContext, error) {
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
		context.Write(c.changes.Diff())

	}

	if len(c.changedFilesContent) > 0 {
		context.WriteString("\n=== Changed files content ===\n")
		for filename, content := range c.changedFilesContent {
			fileHeader := fmt.Sprintf("\n== Filename: %s ==\n", filename)
			context.WriteString(fileHeader)
			context.WriteString(content)
		}
	}

	return &ProjectContext{
		Context:      context.String(),
		SystemPrompt: systemPrompt,
	}, nil
}

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

func (c *contextBuilderImpl) AddChangedFilesContent() {
	changedFiles := c.changes.ChangedFiles()
	c.changedFilesContent = make(map[string]string, 0)
	for _, file := range changedFiles {
		c.changedFilesContent[file] = readFileContent(file)
	}
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
	for file := range strings.SplitSeq(string(filesOut), "\n") {
		if file = strings.TrimSpace(file); file != "" {
			files = append(files, file)
		}
	}

	return &contextBuilderImpl{
		errors:              make([]error, 0),
		files:               files,
		languages:           make([]string, 0),
		branch:              nil,
		changes:             nil,
		changedFilesContent: map[string]string{},
	}, nil
}
