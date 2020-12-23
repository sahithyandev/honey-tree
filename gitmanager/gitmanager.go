package gitmanager

import (
	"fmt"
	"os"
	"os/exec"
)

// ExecuteCommand executes a command in the shell
func ExecuteCommand(commandName string, args ...string) string {
	var output, err = exec.Command(commandName, args...).Output()

	if err != nil {
		fmt.Println(fmt.Errorf("Error occured while running %v. %v", commandName, err).Error())
		os.Exit(2)
	}
	return string(output)
}

// InitGitRepo initializes a git repository in the given directory
func InitGitRepo(directory string) string {
	return ExecuteCommand("git", "init", directory)
}

// IsGitRepo checks if a directory is a repository
// By checking if it has .git folder
func IsGitRepo(directory string) bool {
	var _, err = os.Stat(directory + "/.git")

	if os.IsNotExist(err) {
		return false
	}
	return true
}

// CloneRepo clones a git repository to the target location
func CloneRepo(repoLink string, targetLocation string) string {
	return ExecuteCommand("git", "clone", repoLink, targetLocation)
}

// ResetGitRepo reset a git repository
// A new repository will be initiated in the directory
func ResetGitRepo(directory string) string {
	// RemoveDir(directory + "/.git/")
	var err = os.RemoveAll(directory + "/.git")
	if err != nil {
		fmt.Println(fmt.Errorf("%v", err).Error())
		os.Exit(1)
	}

	return InitGitRepo(directory)
}
