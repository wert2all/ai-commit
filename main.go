package main

import (
	"fmt"
	"log"

	"github.com/wert2all/ai-commit/ai"
	"github.com/wert2all/ai-commit/commit"
	"github.com/wert2all/ai-commit/project"
)

func main() {
	config, err := ai.ReadConfig()
	if err != nil {
		log.Fatalf("Error read config: %v", err)
	}

	provider, err := ai.NewProvider(*config)
	if err != nil {
		log.Fatal("Error creating AI provider:", err)
	}

	contextBuilder, err := project.NewBuilder(config.Directory)
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

	if config.Options.WithCommit {
		if shouldCommit := commit.AskUser(); shouldCommit {
			commit.Commit(commitMsg, config.Directory)
			fmt.Println("Successfully committed changes with the generated message!")
		} else {
			fmt.Println("Commit cancelled.")
		}
	}
}
