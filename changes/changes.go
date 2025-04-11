package changes

import (
	"fmt"
	"os/exec"
	"strings"
)

type (
	Changes interface {
		Diff() []byte
	}
	changesImpl struct {
		changed []byte
	}
)

// Diff implements Changes.
func (c *changesImpl) Diff() []byte { return c.changed }

func NewChanges() (Changes, error) {
	changedCmd := exec.Command("git", "diff", "--cached", "--diff-algorithm=minimal")
	changes, err := changedCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error getting staged changes: %v", err)
	}

	// If no changes, return error
	if strings.TrimSpace(string(changes[:])) == "" {
		return nil, fmt.Errorf("no changes detected in the repository")
	}

	return &changesImpl{changed: changes}, nil
}
