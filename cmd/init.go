package cmd

import (
	"fmt"
	"os"

	"github.com/sahithyandev/honey-tree/gitmanager"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init BOILERPLATE_GIT_REPO TARGET_DIR",
	Short: "Initializes a project from a honey-tree-boilerplate",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var boilerplateDirectory, targetLocation = args[0], args[1]
		
		if isFolderExists(targetLocation) {
			fmt.Println(targetLocation, "already exists")
			os.Exit(2)
		}
		
		gitmanager.CloneRepo(boilerplateDirectory, targetLocation)	
		gitmanager.ResetGitRepo(targetLocation)
	},
}

func isFolderExists(directory string) bool {
	var output,err = os.Stat(directory)
	
	if os.IsNotExist(err) {
		return false
	} else if err == nil {
		return true
	}
	
	fmt.Println(output)
	fmt.Println(err)
	os.Exit(2)
	return false
}

func init() {
	
	rootCmd.AddCommand(initCmd)
}
