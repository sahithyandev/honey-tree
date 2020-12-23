package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/sahithyandev/honey-tree/helpers"
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
		fmt.Printf("Checking %v...\n", projectDir)

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

func readConfigFile(configFilePath string) configFormat {
	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		fmt.Print(err)
	}

	var config configFormat

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("error:", err)
	}

	var validationErr = validator.Validate(config)
	if validationErr != nil {
		fmt.Println("ValidationError", validationErr)
	}
	return config
}

func testConfigFile(projectDir string) error {
	const configFileName = "honey-tree.config.json"
	var configFilePath = path.Join(projectDir, configFileName)

	// Check if it exists
	if !helpers.DoesExist(configFilePath) {
		return fmt.Errorf("config file (%v) not found", configFilePath)
	}

	// TODO Check the config file format
	// var config = readConfigFile(configFilePath)

	// TODO Validate the json

	return nil
}

func init() {
	rootCmd.AddCommand(testCmd)
}
