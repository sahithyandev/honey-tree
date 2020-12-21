package cmd

import (
	"github.com/sahithyandev/honey-tree/gitmanager"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a project from a honey-tree-boilerplate",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var boilerplateDirectory, targetLocation = args[0], args[1]
		
		if gitmanager.IsGitRepo(boilerplateDirectory) {
			// gitmanager.CloneRepo(boilerplateDirectory, targetLocation)	
			gitmanager.ResetGitRepo(targetLocation)
		}
	},
}


func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
