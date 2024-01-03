package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var server = true

// Server encapsulates all dependencies for the web server.
// HTTP handlers access information via receiver types.
type Server struct {
	Router       *http.ServeMux
	Port         string
	TemplatesDir string // location of html templates, makes template parsing less verbose.
	StaticDir    string // location of static assets
	FileSystem   fs.FS  // in-memory or disk
}

//go:embed templates/* static/*
var inMemoryFS embed.FS

func main() {
	if server {
		port := "8080"

		log.Println("[ 💿 Spinning up server on http://localhost:" + port + " ]")

		s := Server{
			Router:       http.NewServeMux(),
			Port:         port,
			TemplatesDir: "templates",
			StaticDir:    "static",
			FileSystem:   inMemoryFS,
		}

		s.routes()

		err := http.ListenAndServe(":"+s.Port, s.Router)
		if err != nil {
			panic(err)
		}
	} else {
		takeScreenshot("connorkuljis")
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
	indexHTML := []string{
		"root.html",
		"head.html",
		"layout.html",
		"components/hero.html",
		"components/footer.html",
	}

	tmpl := compileTemplates(s, indexHTML)

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "root", nil)
	}
}

func (s *Server) handleGenerate() http.HandlerFunc {
	type PageData struct {
		Username string
		Filename string
	}
	generateHTMLPartial := []string{
		"partials/contribution.html",
	}

	tmpl := compileTemplates(s, generateHTMLPartial)

	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		username := r.Form.Get("username")

		if username == "" {
			log.Fatal("empty username")
		}

		fmt.Println(username)

		filename := takeScreenshot(username)

		data := PageData{Username: username, Filename: filename}

		tmpl.ExecuteTemplate(w, "contribution", data)
	}
}
