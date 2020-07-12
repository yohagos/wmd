package main

import (
	"github.com/wmd/sessions"
	"net/http"

	"github.com/wmd/models"
	"github.com/wmd/routes"
	"github.com/wmd/utils"
)

func main() {
	utils.AddLoggingToMap([]byte("INFO"), []byte("Server is starting now.."))
	utils.Init()
	utils.LoadTemplates("templates/*.html")
	models.Init()
	sessions.SessionInit()

	r := routes.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
