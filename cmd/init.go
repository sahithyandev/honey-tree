package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/sahithyandev/honey-tree/helpers"
	"github.com/sahithyandev/honey-tree/helpers/gitmanager"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

const (
	_HONEY_TREE_CACHE_DIR = ".honey-tree"
)

var (
	// TODO Test saveLocally feature
	saveLocally bool
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init BOILERPLATE_GIT_REPO TARGET_DIR",
	Short: "Initializes a project from a honey-tree-boilerplate",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var boilerplateDirectory, targetLocation = args[0], args[1]

		if helpers.DoesExist(targetLocation) {
			fmt.Println(targetLocation, "already exists")
			os.Exit(2)
		}

		var HOME_DIR, err = homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println("HOME", HOME_DIR)

		var cacheLocation = path.Join(HOME_DIR, _HONEY_TREE_CACHE_DIR, boilerplateDirectory)

		if saveLocally {
			// TODO safely abort if boilerplateDirectory is local.
			// cache it in the $HOME/.honey-tree/$BOILERPLATE_NAME
			// if it already exists inside .honey-tree, delete it.
			if helpers.DoesExist(cacheLocation) {
				os.RemoveAll(cacheLocation)
			}
			gitmanager.CloneRepo(boilerplateDirectory, cacheLocation)
		}

		// check if boilerplate exists inside $HOME/.honey-tree/
		// if it does, use it. show a warning.
		if helpers.DoesExist(cacheLocation) {
			fmt.Printf("%v is available locally (at %v). Local version will be used.\n", boilerplateDirectory, cacheLocation)
			boilerplateDirectory = cacheLocation
		}

		gitmanager.CloneRepo(boilerplateDirectory, targetLocation)
		gitmanager.ResetGitRepo(targetLocation)

		fmt.Printf("New project created at %v", targetLocation)
	},
}

func init() {
	initCmd.Flags().BoolVarP(&saveLocally, "save-locally", "s", false, "save the boilerplate locally (for future offline use)")

	rootCmd.AddCommand(initCmd)
}
