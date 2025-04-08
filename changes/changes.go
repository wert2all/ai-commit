package changes

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type (
	Changes interface {
		ToString() string
	}
	changesImpl struct {
		changed []byte
	}
)

// ToString implements Changes.
func (c changesImpl) ToString() string {
	var changes bytes.Buffer
	changes.WriteString("=== Changes ===\n")
	changes.Write(c.changed)
	return changes.String()
}

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
