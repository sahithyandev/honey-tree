package helpers

import (
	"fmt"
	"os"
)

// IsFolderExists checks if a folder exists
func IsFolderExists(directory string) bool {
	var output, err = os.Stat(directory)

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
