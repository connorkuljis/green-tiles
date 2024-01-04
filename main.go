package main

import (
	"embed"
	"github.com/gorilla/sessions"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"
)

var server = true

const SessionName = "session"

// Server encapsulates all dependencies for the web server.
// HTTP handlers access information via receiver types.
type Server struct {
	Router       *http.ServeMux
	Port         string
	TemplatesDir string // location of html templates, makes template parsing less verbose.
	StaticDir    string // location of static assets
	FileSystem   fs.FS  // in-memory or disk
	Sessions     *sessions.CookieStore
}

//go:embed templates/* static/*
var inMemoryFS embed.FS

func main() {
	if server {
		port := "8080"

		store := sessions.NewCookieStore([]byte("special_key"))

		log.Println("[ 💿 Spinning up server on http://localhost:" + port + " ]")

		s := Server{
			Router:       http.NewServeMux(),
			Port:         port,
			TemplatesDir: "templates",
			StaticDir:    "static",
			FileSystem:   inMemoryFS,
			Sessions:     store,
		}

		s.routes()

		err := http.ListenAndServe(":"+s.Port, s.Router)
		if err != nil {
			panic(err)
		}
	} else {
		TakeScreenshot("connorkuljis", Double, 0)
	}
}

func compileTemplates(s *Server, files []string) *template.Template {
	for i := range files {
		files[i] = filepath.Join(s.TemplatesDir, files[i])
	}

	tmpl, err := template.ParseFS(s.FileSystem, files...)
	if err != nil {
		panic(err)
	}

	return tmpl
}

func (s *Server) routes() {
	// check http handler documentation for the appropriate pattern for static content.
	// eg: chi -> "/static/*"
	s.Router.Handle("/static/", http.FileServer(http.FS(s.FileSystem)))
	s.Router.Handle("/screenshots/", http.StripPrefix("/screenshots/", http.FileServer(http.Dir("screenshots"))))
	s.Router.HandleFunc("/", s.handleIndex())
	s.Router.HandleFunc("/generate", s.handleGenerate())
}

func (s *Server) handleIndex() http.HandlerFunc {
	type PageData struct {
		Username string
		Option   string
		Offset   string
	}

	indexHTML := []string{
		"root.html",
		"head.html",
		"layout.html",
		"components/hero.html",
		"components/footer.html",
	}

	tmpl := compileTemplates(s, indexHTML)

	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := s.Sessions.Get(r, SessionName)

		var pageData PageData

		switch val := session.Values["username"].(type) {
		case string:
			pageData.Username = val
		}

		switch val := session.Values["option"].(type) {
		case string:
			pageData.Option = val
		}

		switch val := session.Values["offset"].(type) {
		case string:
			pageData.Offset = val
		}

		tmpl.ExecuteTemplate(w, "root", pageData)
	}
}

func (s *Server) handleGenerate() http.HandlerFunc {
	type PageData struct {
		Username string
		Filename string
	}

	type FormData struct {
		username string
		option   int
		offset   int
	}

	generateHTMLPartial := []string{
		"partials/contribution.html",
	}

	tmpl := compileTemplates(s, generateHTMLPartial)

	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := s.Sessions.Get(r, SessionName)

		r.ParseForm()

		username := r.Form.Get("username")
		optionStr := r.Form.Get("option")
		offsetStr := r.Form.Get("offset")

		if username == "" {
			log.Println("empty username")
		}

		option, err := strconv.Atoi(optionStr)
		if err != nil {
			log.Printf("Unable to convert option: %s to int\n", optionStr)
		}

		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			log.Printf("Unable to convert offset: %s to int\n", offsetStr)
		}

		session.Values["username"] = username
		session.Values["option"] = optionStr
		session.Values["offset"] = offsetStr
		session.Save(r, w)

		formData := FormData{username: username, option: option, offset: offset}

		log.Println("FormData: ", formData)

		imgName, err := TakeScreenshot(formData.username, formData.option, formData.offset)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		data := PageData{Username: username, Filename: imgName}

		tmpl.ExecuteTemplate(w, "contribution", data)
	}
}
