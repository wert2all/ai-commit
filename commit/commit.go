package commit

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func AskUser() bool {
	// Ask user for confirmation
	fmt.Print("Do you want to commit with this message? (yes/no): ")
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading user input:", err)
	}

	response = strings.TrimSpace(strings.ToLower(response))
	return response == "yes" || response == "y"
}

func Commit(commitMsg string, absProjectDir string) {
	// Execute git commit with proper escaping of special characters
	cmd := exec.Command("git", "commit", "-m", commitMsg)
	cmd.Dir = absProjectDir
	// Set the environment to ensure proper handling of special characters
	cmd.Env = append(os.Environ(), "LANG=en_US.UTF-8")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("Error executing git commit:", string(output))
	}

}
