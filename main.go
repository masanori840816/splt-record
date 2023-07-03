package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	templ    *template.Template
	settings *AppSettings
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handleMainPageRequest(w, r, t)
}
func main() {
	settings, err := NewAppSettings()
	if err != nil {
		log.Panicln(err.Error())
	}
	http.Handle("/js/", http.FileServer(http.Dir("templates")))
	http.Handle("/css/", http.FileServer(http.Dir("templates")))
	http.Handle("/", &templateHandler{settings: &settings})
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", settings.Host, settings.Port), nil))
}
func handleMainPageRequest(w http.ResponseWriter, r *http.Request, t *templateHandler) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", "register.html")))
	})
	t.templ.Execute(w, t.settings.BaseURL)
}
