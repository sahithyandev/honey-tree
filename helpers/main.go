package helpers

import (
	"fmt"
	"os"
)

// DoesExist checks if a folder exists
func DoesExist(path string) bool {
	var output, err = os.Stat(path)

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
