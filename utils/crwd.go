package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Quelle https://gist.github.com/novalagung/13c5c8f4d30e0c4bff27

// CreateFile func
func CreateFile() error {
	file, err := os.Create(LogFilePath) //, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 755
	if IsError(err) {
		return err
	}
	defer file.Close()
	return nil
}

// WriteFile func
func WriteFile(level, message string) {
	// open file using READ & WRITE permission
	file, err := os.OpenFile(LogFilePath, os.O_WRONLY|os.O_APPEND, 0644)
	if IsError(err) {
		return
	}
	defer file.Close()

	log.SetOutput(file)

	switch level {
	case "INFO":
		log.Println(strings.ToUpper(level) + " - " + message)
	case "WARNING":
		log.Println(strings.ToUpper(level) + " - " + message)
	case "ERROR":
		log.Println(strings.ToUpper(level) + " - " + message)
	default:
		fmt.Fprintf(os.Stderr, "Invalid level value %q, allowed values are: debug, info, notice, warning, error, critical, alert, emergency and none\n", level)
		os.Exit(2)
	}

	log.SetOutput(os.Stdout)

	if IsError(err) {
		return
	}

	// save changes
	err = file.Sync()
	if IsError(err) {
		return
	}
}

// IsError func
func IsError(err error) bool {
	if err != nil {
		log.Println(LogError + LogTabulator + "Error occurred")
		log.Println(err)
	}
	return (err != nil)
}
