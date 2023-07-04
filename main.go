package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"

	db "github.com/splt-record/db"
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

	dbCtx := db.NewSpltRecordContext(settings.ConnectionString)

	http.HandleFunc("/records", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			log.Println("Get")
		case "POST":
			log.Println("POST")
		case "PUT":
			log.Println("PUT")
		case "DELETE":
			log.Println("DELETE")
		}
		w.Write([]byte("Hello"))
	})
	http.HandleFunc("/weapons/all", func(w http.ResponseWriter, r *http.Request) {
		res, err := GetAllWeapons(dbCtx)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(502)
		}
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
	})
	http.Handle("/", &templateHandler{settings: &settings})
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", settings.Host, settings.Port), nil))
}
func handleMainPageRequest(w http.ResponseWriter, r *http.Request, t *templateHandler) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", "register.html")))
	})
	t.templ.Execute(w, t.settings.BaseURL)
}
