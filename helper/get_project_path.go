package helper

import (
	"os"
	"path"
)

func GetProjectPath(dir string, filename string) (string, error) {
	wd, getWdErr := os.Getwd()
	if getWdErr != nil {
		return "", getWdErr
	}
	dirPath := path.Clean(wd + dir)
	filePath := path.Join(dirPath, filename)

	return filePath, nil
}
