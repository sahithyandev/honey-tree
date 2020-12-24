package cmd

import (
	"fmt"
	"os"

	"github.com/sahithyandev/honey-tree/helpers"
	"github.com/sahithyandev/honey-tree/helpers/gitmanager"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run inside a directory to check if it is a honey-tree-boilerplate",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var projectDir = "."
		if len(args) > 0 {
			projectDir = args[0]
		}
		fmt.Printf("Checking %v ...\n", projectDir)

		// check if the projectDir exists
		if !helpers.DoesExist(projectDir) {
			fmt.Printf("%v directory not found\n", projectDir)
			os.Exit(2)
		}

		// check if it is a git repo
		if !gitmanager.IsGitRepo(projectDir) {
			fmt.Printf("%v must be a git repository\n", projectDir)
			os.Exit(2)
		}

		readConfig(projectDir)

		fmt.Printf("%v is a honey-tree-boilerplate.\n", projectDir)
	},
}

func readConfig(projectDir string) {
	viper.SetConfigName(HoneyTreeConfigFileName)
	viper.SetConfigType(HoneyTreeConfigFileType)

	viper.AddConfigPath(projectDir)

	var err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No config file")
		} else {
			// Config file was found but another error was produced
			fmt.Printf("Error occured while reading %v", HoneyTreeConfigFile())
		}
		os.Exit(2)
	}
}

func init() {
	rootCmd.AddCommand(testCmd)
}
