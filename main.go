package main

import (
	"fmt"
	"os"

	"github.com/wert2all/ai-commit/ai"
	"github.com/wert2all/ai-commit/commit"
	"github.com/wert2all/ai-commit/project"
	"github.com/wert2all/ai-commit/ui"
)

var (
	cardWidth = 60
	Version   = "dev"
)

func main() {
	config, err := ai.ReadConfig()
	if err != nil {
		handleError(err)
	}

	if config.Options.ShowVersion {
		fmt.Printf("Version: %s\n", Version)
		os.Exit(0)
	}

	provider, err := ai.NewProvider(*config)
	if err != nil {
		handleError(err)
	}

	contextBuilder, err := project.NewBuilder(config.Directory)
	if err != nil {
		handleError(err)
	}

	contextBuilder.AddLanguages()
	contextBuilder.AddGitBranch()
	contextBuilder.AddChanges()

	if config.Options.WithChangedFilesContent {
		contextBuilder.AddChangedFilesContent()
	}

	projectContext, err := contextBuilder.Build()
	if err != nil {
		handleError(err)
	}

	commitMsg, err := provider.GenerateCommitMessage(*projectContext)
	if err != nil {
		handleError(err)
	}

	fmt.Println(ui.NewProviderInfo(provider.GetProviderInfo()))
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

func handleError(err error) {
	fmt.Println(ui.NewError(err.Error(), cardWidth))
	os.Exit(1)
}
