package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/wert2all/ai-commit/ai"
	"github.com/wert2all/ai-commit/commit"
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

	if commit.AskUser() {
		commit.Commit(commitMsg, absProjectDir)
		fmt.Println("Successfully committed changes with the generated message!")
	} else {
		fmt.Println("Commit cancelled.")
	}
}
