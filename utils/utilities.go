package utils

import (
	"strings"
)

// RemoveBlanks func
func RemoveBlanks(str string) string {
	str = strings.Replace(str, " ", "_", -1)
	return str
}

// CheckError func
func CheckError(err error) error {
	if err != nil {
		LoggingErrorFile(err.Error())
		return err
	}
	return nil
}
