package changes

import (
	"fmt"
	"os/exec"
	"strings"
)

type (
	Changes interface {
		Diff() []byte
		ChangedFiles() []string
	}
	changesImpl struct {
		changed      []byte
		changedFiles []string
	}
)

// Diff implements Changes.
func (c *changesImpl) Diff() []byte { return c.changed }

func (c *changesImpl) ChangedFiles() []string { return c.changedFiles }

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

	// Extract changed files from the diff output
	changedFiles := extractChangedFilesFromDiff(changes)

	return &changesImpl{
		changed:      changes,
		changedFiles: changedFiles,
	}, nil
}

// extractChangedFilesFromDiff parses the diff output and returns a list of unique changed filenames.
func extractChangedFilesFromDiff(diff []byte) []string {
	lines := strings.Split(string(diff), "\n")
	filesMap := make(map[string]struct{})
	for _, line := range lines {
		if strings.HasPrefix(line, "diff --git ") {
			parts := strings.Split(line, " ")
			if len(parts) >= 4 {
				// The format is: diff --git a/<file> b/<file>
				bPath := parts[3]
				// Remove the "b/" prefix
				if strings.HasPrefix(bPath, "b/") {
					filename := bPath[2:]
					filesMap[filename] = struct{}{}
				}
			}
		}
	}
	files := make([]string, 0, len(filesMap))
	for f := range filesMap {
		files = append(files, f)
	}
	return files
}
