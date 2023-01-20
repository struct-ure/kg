package util

import (
	"errors"
	"os"
)

// DirectoryExists return true if the path exists.
func DirectoryExists(path string) bool {
	if fileinfo, err := os.Stat(path); err == nil {
		return fileinfo.IsDir()
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		panic(err)
	}
}
