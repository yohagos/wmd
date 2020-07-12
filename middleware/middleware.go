package middleware

import (
	"net/http"

	"github.com/wmd/sessions"
)

// HandleFunc func
type HandleFunc func(http.ResponseWriter, *http.Request)

// AuthRequired handler
func AuthRequired(handler HandleFunc) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := sessions.Store.Get(r, "session")
		_, ok := session.Values["user_id"]
		if !ok {
			http.Redirect(w, r, "/login", 302)
			return
		}
		handler(w, r)
	}
}

// ProfRequired handler
/*func ProfRequired(handler HandleFunc) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := sessions.Store.Get(r, "session")
		untypeduser_id := session.Values["user_id"]
		currentuser_id, _ := untypeduser_id.(int64)
		if !models.CheckProf(currentuser_id) {
			s, _ := models.ReturnUsername(currentuser_id)
			utils.LoggingWarningFile("User " + s + " does not have the authorization to upload a document.")
			http.Redirect(w, r, "/", 302)
			return
		}
		handler(w, r)
	}
}*/

/*func TypeHandler(handler HandleFunc) HandleFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		/*session, _ := sessions.Store.Get(r, "session")
		untypeduser_id := session.Values["user_id"]
		currentuser_id, _ := untypeduser_id.(int64)
		if !models.CheckProf(currentuser_id) {
			s, _ := models.ReturnUsername(currentuser_id)
			utils.LoggingWarningFile("User " + s + " does not have the authorization to upload a document.")
			http.Redirect(w, r, "/", 302)
			return
		}
		x := r.URL.RequestURI()[1:]

		handler(w, r)
	}
}*/
