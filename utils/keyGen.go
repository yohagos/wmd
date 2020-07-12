package utils

import (
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
var length = 10

// RandomKeyMail func
func RandomKeyMail() string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[seededRand.Intn(len(letters))]
	}
	return string(b)
}

// Invoking
/*
	randomStringBytes(24)
	--> sLdWTRr2wyPDSIXIg7VrbMaa
*/

// Used to create the Session Key -> length 24 characters
/*func randomStringSessionKey(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}*/
