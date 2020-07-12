package routes

import (
	"github.com/gorilla/mux"
	"github.com/wmd/mails"
	"github.com/wmd/middleware"
	"github.com/wmd/models"
	"github.com/wmd/sessions"
	"github.com/wmd/utils"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const maxUploadSize = 25 * 1024

var router = mux.NewRouter()

var (
	path string
)

// NewRouter func
func NewRouter() *mux.Router {
	utils.LoggingInfoFile("Router is starting...")
	router.HandleFunc("/", middleware.AuthRequired(indexGETHandler)).Methods("GET")

	router.HandleFunc("/login", loginGETHandler).Methods("GET")
	router.HandleFunc("/login", loginPOSTHandler).Methods("POST")

	router.HandleFunc("/newUpload", middleware.AuthRequired(newUploadGETHandler)).Methods("GET")
	router.HandleFunc("/newUpload", middleware.AuthRequired(newUploadPOSTHandler)).Methods("POST")

	router.HandleFunc("/logout", logoutGETHandler).Methods("GET")

	router.HandleFunc("/register", registerGETHandler).Methods("GET")
	router.HandleFunc("/register", registerPOSTHandler).Methods("POST")

	router.HandleFunc("/verif", verifyGETHandler).Methods("GET")
	router.HandleFunc("/verif", verifyPOSTHandler).Methods("POST")

	router.HandleFunc("/info", infoGETHandler).Methods("GET")
	router.HandleFunc("/contact", contactGETHandler).Methods("GET")

	router.HandleFunc("/{itemname}", middleware.AuthRequired(itemGETHandler)).Methods("GET")

	utils.LoggingInfoFile("File server will start loading..")

	fs := http.FileServer(http.Dir("static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	router.NotFoundHandler = router.NewRoute().HandlerFunc(notFoundGETHandler).GetHandler()
	utils.LoggingInfoFile("Web Server will start running... anticipating Requests")
	return router
}

func indexGETHandler(w http.ResponseWriter, r *http.Request) {
	uploads, err := models.GetAllUploads()
	if err != nil {
		utils.InternalServerError(w)
		return
	}

	session, _ := sessions.Store.Get(r, "session")
	untypeduser_id := session.Values["user_id"]
	currentuser_id, _ := untypeduser_id.(int64)

	var display bool
	display = models.CheckProf(currentuser_id)
	utils.ExecuteTemplate(w, "index.html", struct {
		Title   string
		Uploads []*models.Upload
		Display bool
	}{
		Title:   "All Uploads",
		Uploads: uploads,
		Display: display,
	})
}

func loginGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func loginPOSTHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	user, err := models.AuthentificateUser(username, password)
	if err != nil {
		switch err {
		case utils.ErrUserNotFound:
			utils.ExecuteTemplate(w, "login.html", "unknown user")
		case utils.ErrInvalidLogin:
			utils.ExecuteTemplate(w, "login.html", "invalid login")
		default:
			utils.InternalServerError(w)
		}
		return
	}
	user_id, err := user.GetID()
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	session, _ := sessions.Store.Get(r, "session")
	session.Values["user_id"] = user_id
	err = session.Save(r, w)
	utils.CheckError(err)
	http.Redirect(w, r, "/", 302)
}

func newUploadGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "newUpload.html", struct {
		Title string
	}{
		Title: "New Upload",
	})
}

func newUploadPOSTHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	untypeduser_id := session.Values["user_id"]
	currentuser_id, ok := untypeduser_id.(int64)
	if !ok {
		utils.InternalServerError(w)
		return
	}
	if r.Method == http.MethodPost {
		r.ParseMultipartForm(10 * 1024)
		itemname := utils.RemoveBlanks(r.PostForm.Get("item"))
		description := r.PostForm.Get("desc")
		scientificType := utils.RemoveBlanks(r.PostForm.Get("typ"))
		cat := utils.RemoveBlanks(r.PostForm.Get("cate"))

		err := r.ParseMultipartForm(maxUploadSize)
		utils.CheckError(err)
		var fileExt string
		y := r.MultipartForm.File

		file, handler, err := r.FormFile("myFile")
		if err != nil {
			utils.InternalServerError(w)
			utils.LoggingErrorFile(err.Error())
			return
		}
		defer file.Close()
		for _, v := range y {
			for _, x := range v {
				if x.Header.Get("Content-Type") == "application/pdf" && strings.Contains(handler.Filename, ".pdf") {
					fileExt = ".pdf"
				}
			}
		}

		username, _ := models.ReturnUsername(currentuser_id)
		temp := utils.CreateNewDirectory(username, itemname)
		if temp != "" {
			data, err2 := ioutil.ReadAll(file)
			utils.CheckError(err2)

			path = temp + "\\" + itemname + fileExt
			err2 = ioutil.WriteFile(path, data, 0777)

			utils.CheckError(err2)
			_, err21 := models.NewUpload(currentuser_id, itemname, description, scientificType, cat, path)
			if err21 != nil {
				utils.InternalServerError(w)
				utils.LoggingErrorFile(err21.Error())
				return
			}
			utils.LoggingInfoFile("New Upload : " + itemname)
		} else {
			utils.ExecuteTemplate(w, "newUpload.html", utils.ErrFileUpload)
		}
	}
	session.Values["user_id"] = currentuser_id
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}

func logoutGETHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	delete(session.Values, "user_id")
	session.Save(r, w)
	http.Redirect(w, r, "/login", 302)
}

func registerGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}

func registerPOSTHandler(w http.ResponseWriter, r *http.Request) {
	rand := utils.RandomKeyMail()
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	mail := r.PostForm.Get("mail")

	models.NewVerification(mail, username, password, []byte(rand))
	mails.SendMail(username, mail, rand)

	http.Redirect(w, r, "/login", 302)
	return
}

func verifyGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "verify.html", nil)
}

func verifyPOSTHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	verifCode := r.PostForm.Get("verif")

	var b = false
	verSlice, err := models.GetAllVerif()
	utils.CheckError(err)

	u := ""
	p := ""
	m := ""
	prof := ""

	for _, v := range verSlice {
		rand, _ := v.GetRand()
		if strings.EqualFold(verifCode, string(rand)) {
			b = true
			u, err = v.GetUsername()
			p, err = v.GetPassword()
			m, err = v.GetMail()
			v.DelVerif()
		}
	}
	if !b {
		utils.InternalServerError(w)
		return
	}

	if strings.Contains(m, "yosef") {
		prof = "Professor"
	}

	err = models.RegisterUser(prof, m, u, p)
	if err == utils.ErrUsernameTaken {
		utils.ExecuteTemplate(w, "login.html", "user has not registered himself. please check your emails")
		return
	} else if err != nil {
		utils.InternalServerError(w)
		return
	}
	utils.ExecuteTemplate(w, "verify.html", struct {
		Title string
	}{
		Title: u,
	})
	http.Redirect(w, r, "/login", 302)
}

func itemGETHandler(w http.ResponseWriter, r *http.Request) {
	i := r.URL.RequestURI()[1:]
	if strings.EqualFold("favicon.ico", i) {
		return
	}
	path, err := models.UploadsPath(i)
	utils.CheckError(err)
	if path == "" {
		log.Printf("404: %s", r.URL.String())
		http.NotFound(w, r)
		return
	}

	f, err := os.Open(path)
	utils.CheckError(err)
	defer f.Close()

	w.Header().Set("Content-type", "application/pdf")
	if _, err := io.Copy(w, f); err != nil {
		utils.LoggingErrorFile(err.Error())
		w.WriteHeader(500)
	}

	utils.ExecuteTemplateWithContentType(w, "outputPDF.html", struct {
		Title string
	}{
		Title: i,
	})
}

func infoGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "info.html", nil)
}

func contactGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "contact.html", nil)
}

func notFoundGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "error.html", struct {
		Error error
	}{
		Error: utils.ErrPageNotFound,
	})
}
