package gitmanager

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecuteCommand(commandName string, args ...string) string {
	var output, err = exec.Command(commandName, args...).Output()

	if err != nil {
		fmt.Println(fmt.Errorf("Error occured while running %v. %v", commandName, err).Error())
		os.Exit(2)
	}
	return string(output)
}

func InitGitRepo(directory string) string {
	return ExecuteCommand("git", "init", directory)
}

func IsGitRepo(directory string) bool {
	var _, err = os.Stat(directory + "/.git")

	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CloneRepo(repoLink string, targetLocation string) string {
	return ExecuteCommand("git", "clone", repoLink, targetLocation)
}

func ResetGitRepo(directory string) string {
	RemoveDir(directory + "/.git/")
	return InitGitRepo(directory)
}

func RemoveDir(directory string) {
	// Open the directory and read all its files.
	dirRead, _ := os.Open(directory)
	dirFiles, _ := dirRead.Readdir(0)

	// Loop over the directory's files.
	for index := range dirFiles {
		fileHere := dirFiles[index]

		// Get name of file and its full path.
		nameHere := fileHere.Name()
		fullPath := directory + nameHere

		// Remove the file.
		os.Remove(fullPath)
		fmt.Println("Removed file:", fullPath)
		
		// Remove the folder finally
		os.Remove(directory)
	}
}

func main() {

}
