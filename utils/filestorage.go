package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// CreateNewDirectory func
func CreateNewDirectory(user, item string) string {
	absPath, err := filepath.Abs(StoragePath)
	if err != nil {
		LoggingErrorFile("Error reading given path : ")
		LoggingErrorFile(err.Error())
	}
	userDir := user + "\\" + item

	var path []string
	path = strings.Split(userDir, "\\")
	for _, v := range path {
		absPath = absPath + "\\" + v
		if _, err := os.Stat(absPath); os.IsNotExist(err) {
			LoggingInfoFile("Creating new directory : " + v)
			os.Mkdir(absPath, 0755)
		} else {
			LoggingInfoFile("Directory already exists : " + absPath)
		}
	}

	newPath, err := filepath.Abs(StoragePath)
	newPath = newPath + "\\" + userDir
	return newPath
}
