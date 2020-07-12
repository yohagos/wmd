package utils

import (
	"log"
	"os"
	"path/filepath"
)

var (
	envString     = "GOWMD"
	logString     = "\\logs"
	storageString = "storage"
	logFileName   = "\\general.log"
	// EnvironmentPath var
	EnvironmentPath string
	// StoragePath var
	StoragePath = EnvironmentPath + storageString
	// LogPath var
	LogPath = EnvironmentPath + logString
	// LogFilePath var
	LogFilePath string
)

// Init function - Checking environment variable
func Init() {
	logString = logString[1:]
	CheckEnvironmentVariable()
	LoggingMapToFile()
}

// CheckEnvironmentVariable func
func CheckEnvironmentVariable() {
	name, ok := os.LookupEnv(envString)
	if ok {
		EnvironmentPath = name
		StoragePath = name + "\\" + storageString
		LogPath = name + "\\" + logString
		LogFilePath = LogPath + logFileName

		CreateLogDirectory()
		CreateStorageDirectory()

		AddLoggingToMap([]byte(LogInfo), []byte("Environment path : "+EnvironmentPath))
		log.Println(LogInfo + LogTabulator + "Environment path : " + EnvironmentPath)

		AddLoggingToMap([]byte(LogInfo), []byte("Storage path : "+StoragePath))
		log.Println(LogInfo + LogTabulator + "Storage path : " + StoragePath)

		AddLoggingToMap([]byte(LogInfo), []byte("Log path : "+LogPath))
		log.Println(LogInfo + LogTabulator + "Log path : " + LogPath)
	} else {
		log.Println(LogError + LogTabulator + "Cannot find Environment Variable 'GOWMD'")
		os.Exit(3)
	}
}

// CreateStorageDirectory func
func CreateStorageDirectory() {
	absPath, err := filepath.Abs(EnvironmentPath)
	if err != nil {
		LoggingErrorFile("Error reading given path : ")
		LoggingErrorFile(err.Error())
	}
	StoragePath = absPath + "\\" + storageString
	if _, err := os.Stat(StoragePath); os.IsNotExist(err) {
		AddLoggingToMap([]byte(LogInfo), []byte("Creating Storage directory.."))
		os.MkdirAll(StoragePath, 0755)
	}
	LoggingWarningFile("Storage Directory already exists")
}

// CreateLogDirectory func
func CreateLogDirectory() {
	absPath, err := filepath.Abs(EnvironmentPath)
	if err != nil {
		LoggingErrorFile("Error reading given path : ")
		LoggingErrorFile(err.Error())
		return
	}
	LogPath = absPath + "\\" + logString
	if _, err := os.Stat(LogPath); os.IsNotExist(err) {
		AddLoggingToMap([]byte(LogInfo), []byte("Creating Log directory.."))
		os.MkdirAll(LogPath, 0755)
		CreateFile()
		return
	}
	LoggingWarningFile("Log Directory already exists")
}
