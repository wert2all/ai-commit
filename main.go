package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/sashabaranov/go-openai"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is not set")
	}

	client := openai.NewClient(apiKey)
	commitMsg, err := generateCommitMessage(client)
	if err != nil {
		log.Fatal("Error generating commit message:", err)
	}

	fmt.Println("Generated commit message:")
	fmt.Println(commitMsg)
}

type ProjectInfo struct {
	Languages    []string
	Dependencies map[string]string
	Config       map[string]string
}

func detectProjectType(files []string) []string {
	languages := make(map[string]bool)

	for _, file := range files {
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

	result := make([]string, 0, len(languages))
	for lang := range languages {
		result = append(result, lang)
	}
	return result
}

func readFileContent(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(content)
}

func getProjectDependencies(repoRoot string, files []string) map[string]string {
	deps := make(map[string]string)

	// Check for various dependency files
	for _, file := range files {
		fullPath := filepath.Join(repoRoot, file)
		switch filepath.Base(file) {
		case "go.mod":
			deps["Go Dependencies"] = readFileContent(fullPath)
		case "package.json":
			deps["Node.js Dependencies"] = readFileContent(fullPath)
		case "requirements.txt":
			deps["Python Dependencies"] = readFileContent(fullPath)
		case "composer.json":
			deps["PHP Dependencies"] = readFileContent(fullPath)
		case "pom.xml":
			deps["Maven Dependencies"] = readFileContent(fullPath)
		case "build.gradle":
			deps["Gradle Dependencies"] = readFileContent(fullPath)
		case "Gemfile":
			deps["Ruby Dependencies"] = readFileContent(fullPath)
		case "Cargo.toml":
			deps["Rust Dependencies"] = readFileContent(fullPath)
		}
	}

	return deps
}

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

func getProjectContext() (string, error) {
	// Get repository root
	rootCmd := exec.Command("git", "rev-parse", "--show-toplevel")
	rootOut, err := rootCmd.Output()
	if err != nil {
		return "", fmt.Errorf("error getting repository root: %v", err)
	}
	repoRoot := strings.TrimSpace(string(rootOut))

	// Get project structure using git ls-files
	filesCmd := exec.Command("git", "ls-files")
	filesCmd.Dir = repoRoot
	filesOut, err := filesCmd.Output()
	if err != nil {
		return "", fmt.Errorf("error getting git files: %v", err)
	}

	// Split output into lines and filter empty lines
	files := make([]string, 0)
	for _, file := range strings.Split(string(filesOut), "\n") {
		if file = strings.TrimSpace(file); file != "" {
			files = append(files, file)
		}
	}

	// Detect project languages
	languages := detectProjectType(files)

	// Get dependencies
	deps := getProjectDependencies(repoRoot, files)

	// Get configuration
	config := getProjectConfig(repoRoot, files)

	// Get git branch info
	branchCmd := exec.Command("git", "branch", "--show-current")
	branchOut, err := branchCmd.Output()
	if err != nil {
		log.Printf("Warning: error getting current branch: %v", err)
	}

	// Combine project context
	var context bytes.Buffer

	// Add languages
	context.WriteString("=== Project Languages ===\n")
	for _, lang := range languages {
		context.WriteString(lang + "\n")
	}

	// Add dependencies
	context.WriteString("\n=== Dependencies ===\n")
	for depType, depContent := range deps {
		if depContent != "" {
			context.WriteString(fmt.Sprintf("--- %s ---\n%s\n", depType, depContent))
		}
	}

	// Add configuration
	context.WriteString("\n=== Configuration ===\n")
	for configType, configContent := range config {
		if configContent != "" {
			context.WriteString(fmt.Sprintf("--- %s ---\n%s\n", configType, configContent))
		}
	}

	// Add project structure
	context.WriteString("\n=== Project Structure ===\n")
	for _, file := range files {
		context.WriteString(file + "\n")
	}

	// Add git branch
	context.WriteString("\n=== Current Branch ===\n")
	context.Write(branchOut)

	return context.String(), nil
}

func getGitChanges() (string, error) {
	// Get staged changes
	stagedCmd := exec.Command("git", "diff", "--cached", "--diff-algorithm=minimal")
	stagedOut, err := stagedCmd.Output()
	if err != nil {
		return "", fmt.Errorf("error getting staged changes: %v", err)
	}

	// Get unstaged changes
	unstagedCmd := exec.Command("git", "diff", "--diff-algorithm=minimal")
	unstagedOut, err := unstagedCmd.Output()
	if err != nil {
		return "", fmt.Errorf("error getting unstaged changes: %v", err)
	}

	// Get untracked files
	untrackedCmd := exec.Command("git", "ls-files", "--others", "--exclude-standard")
	untrackedOut, err := untrackedCmd.Output()
	if err != nil {
		return "", fmt.Errorf("error getting untracked files: %v", err)
	}

	// Combine all changes
	var changes bytes.Buffer
	changes.WriteString("=== Staged Changes ===\n")
	changes.Write(stagedOut)
	changes.WriteString("\n=== Unstaged Changes ===\n")
	changes.Write(unstagedOut)
	changes.WriteString("\n=== Untracked Files ===\n")
	changes.Write(untrackedOut)

	return changes.String(), nil
}

func generateCommitMessage(client *openai.Client) (string, error) {
	// Get project context
	projectContext, err := getProjectContext()
	if err != nil {
		log.Printf("Warning: error getting project context: %v", err)
	}

	// Get git changes
	changes, err := getGitChanges()
	if err != nil {
		return "", fmt.Errorf("error getting git changes: %v", err)
	}

	// If no changes, return error
	if strings.TrimSpace(changes) == "" {
		return "", fmt.Errorf("no changes detected in the repository")
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
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
