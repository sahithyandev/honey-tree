package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/sahithyandev/honey-tree/helpers"
	"github.com/sahithyandev/honey-tree/helpers/gitmanager"
	"gopkg.in/validator.v2"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run inside a directory to check if it is a honey-tree-boilerplate",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var projectDir = ""
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

		// check configFile
		var configFileError = testConfigFile(projectDir)
		if configFileError != nil {
			fmt.Println(configFileError)
			os.Exit(2)
		}

		fmt.Printf("%v is a honey-tree-boilerplate.\n", projectDir)
	},
}

type configFormat struct {
	Name        string `validate:"nonzero,nonnil"`
	Description string
	Language    string
}

// using named returns
// be careful
func readConfigFile(configFilePath string) (config configFormat, err error) {
	config = configFormat{}
	var fileData []byte

	fileData, err = ioutil.ReadFile(configFilePath)
	if err != nil {
		return
	}

	err = json.Unmarshal(fileData, &config)
	if err != nil {
		return
	}

	err = validator.Validate(config)
	if err != nil {
		return
	}

	return
}

func testConfigFile(projectDir string) error {
	const configFileName = "honey-tree.config.json"
	var configFilePath = path.Join(projectDir, configFileName)

	// Check if it exists
	if !helpers.DoesExist(configFilePath) {
		return fmt.Errorf("config file (%v) not found", configFilePath)
	}

	var config, err = readConfigFile(configFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(config)

	// check if

	validator.Validate(config)

	return nil
}

func init() {
	rootCmd.AddCommand(testCmd)
}
