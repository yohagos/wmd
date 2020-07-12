package utils

import (
	"errors"
	"net/http"
)

var (
	// ErrNothingFound error declaration
	ErrNothingFound = errors.New("a page was tried to reach, which does not exists")
	// ErrPageNotFound error declaration
	ErrPageNotFound = errors.New("page not found. Error 404")
	// ErrUserNotFound error declaration
	ErrUserNotFound = errors.New("user not found")
	// ErrInvalidLogin error declaration
	ErrInvalidLogin = errors.New("invalid login")
	// ErrUsernameTaken error declaration
	ErrUsernameTaken = errors.New("username taken")
	// ErrFileUpload error declaration
	ErrFileUpload = errors.New("error while file upload")
	// ErrDirectoryCannotBeCreated will be thrown if the directories cannot be created
	ErrDirectoryCannotBeCreated = errors.New("You may not have the rights to create directories")
)

// InternalServerError func
func InternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error"))
}
