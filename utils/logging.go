package utils

import "log"

var (
	// LogInfo const
	LogInfo = "INFO"
	// LogWarning const
	LogWarning = "WARNING"
	// LogError const
	LogError = "ERROR"
	// LogTabulator const
	LogTabulator = " : "
	// LogMap var
	LogMap [][]byte
	// LevelMap var
	LevelMap [][]byte
)

// AddLoggingToMap func
func AddLoggingToMap(level []byte, lmsg []byte) {
	LogMap = append(LogMap, lmsg)
	LevelMap = append(LevelMap, level)
}

// LoggingMapToFile func
func LoggingMapToFile() {
	err := CreateFile()
	if IsError(err) {
		log.Println(LogError + LogTabulator + "Cannot find File..")
		return
	}
	slvl := LevelMap
	slog := LogMap
	for i := 0; i < len(slvl) && i < len(slog); i++ {
		WriteFile(string(slvl[i]), string(slog[i]))
	}
}

// LoggingInfoFile func
func LoggingInfoFile(s string) {
	log.Println(LogInfo + LogTabulator + s)
	WriteFile(LogInfo, s)
}

// LoggingErrorFile func
func LoggingErrorFile(s string) {
	log.Println(LogError + LogTabulator + s)
	WriteFile(LogError, s)
}

// LoggingWarningFile func
func LoggingWarningFile(s string) {
	log.Println(LogWarning + LogTabulator + s)
	WriteFile(LogWarning, s)
}
