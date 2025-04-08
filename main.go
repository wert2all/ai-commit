package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/wert2all/ai-commit/ai"
	"github.com/wert2all/ai-commit/project"
)

func main() {
	providerName := flag.String("provider", "openai", "AI provider to use (openai, claude, mistral, gemini, local)")
	model := flag.String("model", "", "Model to use (e.g., gpt-3.5-turbo, claude-2, mistral-medium, gemini-pro)")
	projectDir := flag.String("dir", ".", "Project directory path")
	flag.Parse()

	// Convert relative path to absolute
	absProjectDir, err := filepath.Abs(*projectDir)
	if err != nil {
		log.Fatalf("Error resolving project directory path: %v", err)
	}

	provider, err := ai.NewProvider(*providerName, *model)
	if err != nil {
		log.Fatal("Error creating AI provider:", err)
	}

	contextBuilder, err := project.NewBuilder(absProjectDir)
	if err != nil {
		log.Fatal(err)
	}

	projectContext, err := contextBuilder.
		AddLanguages().
		AddChanges().
		AddGitBranch().
		Build()
	if err != nil {
		log.Fatal(err)
	}

	commitMsg, err := provider.GenerateCommitMessage(*projectContext)
	if err != nil {
		log.Fatal("Error generating commit message:", err)
	}

	fmt.Println("Generated commit message:")
	fmt.Println(commitMsg)

	// Ask user for confirmation
	fmt.Print("Do you want to commit with this message? (yes/no): ")
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading user input:", err)
	}

	response = strings.TrimSpace(strings.ToLower(response))
	if response == "yes" || response == "y" {
		// Execute git commit with proper escaping of special characters
		cmd := exec.Command("git", "commit", "-m", commitMsg)
		cmd.Dir = absProjectDir
		// Set the environment to ensure proper handling of special characters
		cmd.Env = append(os.Environ(), "LANG=en_US.UTF-8")
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal("Error executing git commit:", string(output))
		}
		fmt.Println("Successfully committed changes with the generated message!")
	} else {
		fmt.Println("Commit cancelled.")
	}
}
