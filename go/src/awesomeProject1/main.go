package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	TemplatesList := []string{"templates/blank.html"}
	HandlersData := NewHandlers()
	HandlersData.InitFromTmpls(TemplatesList)

	http.HandleFunc("/", HandlersData.IndexPageHandler)
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/blank.html")
	})
	http.HandleFunc("/support", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/Support.html")
	})
	http.HandleFunc("/dogovor", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/dogovor.html")
	})
	http.HandleFunc("/done", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/done.html")

	})
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/download.html")
	})
	http.HandleFunc("/feature", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/Feature.html")
	})
	http.HandleFunc("/price", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/Price.html")
	})
	http.HandleFunc("/privacy-policy", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/privacy-policy.html")
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/AboutUs.html")
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

type Handlers struct {
	Templates map[string]string
}

func (h Handlers) IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Извините, данная страница не найдена !", 404)
		return
	}
	IndexPageObj := new(IndexPage)
	IndexPageObj.Header.Title = "Главная"
	IndexPageObj.Header.H1MainTitle = "Добро Пожаловать !"
	tmpl := template.Must(template.ParseFiles("templates/Support.html"))
	tmpl.Execute(w, IndexPageObj)
}

func NewHandlers() *Handlers {
	handlers := new(Handlers)
	templatesMap := make(map[string]string)
	handlers.Templates = templatesMap
	return handlers
}

func (h *Handlers) InitFromTmpls(tmpls []string) {
	for _, tmpl := range tmpls {
		tmpBuf, err := ioutil.ReadFile(tmpl)
		if err != nil {
			panic(err)
		}
		h.Templates[tmpl] = string(tmpBuf)
	}
}

type IndexPage struct {
	Header PageHead
}

type PageHead struct {
	Title       string
	H1MainTitle string
}