package main

import (
	"fmt"
	"os"

	"github.com/wert2all/ai-commit/ai"
	"github.com/wert2all/ai-commit/commit"
	"github.com/wert2all/ai-commit/project"
	"github.com/wert2all/ai-commit/ui"
)

var cardWidth = 60

func main() {
	config, err := ai.ReadConfig()
	if err != nil {
		fmt.Println(ui.NewError(err.Error(), cardWidth))
		os.Exit(1)
	}

	provider, err := ai.NewProvider(*config)
	if err != nil {
		fmt.Println(ui.NewError(err.Error(), cardWidth))
		os.Exit(1)
	}

	contextBuilder, err := project.NewBuilder(config.Directory)
	if err != nil {
		fmt.Println(ui.NewError(err.Error(), cardWidth))
		os.Exit(1)
	}

	projectContext, err := contextBuilder.
		AddLanguages().
		AddChanges().
		AddGitBranch().
		Build()
	if err != nil {
		fmt.Println(ui.NewError(err.Error(), cardWidth))
		os.Exit(1)
	}

	commitMsg, err := provider.GenerateCommitMessage(*projectContext)
	if err != nil {
		fmt.Println(ui.NewError(err.Error(), cardWidth))
		os.Exit(1)
	}

	fmt.Println(ui.NewCard("Commit message", commitMsg, cardWidth))

	if config.Options.WithCommit {
		if shouldCommit := commit.AskUser(); shouldCommit {
			commit.Commit(commitMsg, config.Directory)
			fmt.Println("Successfully committed changes with the generated message!")
		} else {
			fmt.Println("Commit cancelled.")
		}
	}
}
