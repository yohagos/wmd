package sessions

import (
	"github.com/gorilla/sessions"
)

const key = "sLdWTRr2wyPDSIXIg7VrbMaa"

// Store var - Session Key
var Store = sessions.NewCookieStore([]byte(key))

// SessionInit func
func SessionInit() {
	Store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		MaxAge:   60 * 120}
}
