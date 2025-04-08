package changes

import (
	"bytes"
	"fmt"
	"os/exec"
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
	changedOut, err := changedCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error getting staged changes: %v", err)
	}

	return &changesImpl{changed: changedOut}, nil
}
