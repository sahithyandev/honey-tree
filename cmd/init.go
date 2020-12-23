package cmd

import (
	"fmt"
	"os"

	"github.com/sahithyandev/honey-tree/helpers"
	"github.com/sahithyandev/honey-tree/helpers/gitmanager"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init BOILERPLATE_GIT_REPO TARGET_DIR",
	Short: "Initializes a project from a honey-tree-boilerplate",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var boilerplateDirectory, targetLocation = args[0], args[1]

		if helpers.IsFolderExists(targetLocation) {
			fmt.Println(targetLocation, "already exists")
			os.Exit(2)
		}

		gitmanager.CloneRepo(boilerplateDirectory, targetLocation)
		gitmanager.ResetGitRepo(targetLocation)

		fmt.Printf("New project created at %v", targetLocation)
	},
}

func init() {

	rootCmd.AddCommand(initCmd)
}
