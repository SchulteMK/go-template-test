package server

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

type index struct {
	common
	Age string
}

func (s *server) createIndex(t *template.Template) func(http.ResponseWriter, *http.Request) {
	site, err := t.ParseFiles(path.Join(s.templatePath, "index.html"))
	if err != nil {
		log.Fatal(err)
	}
	indexCommon := common{
		Title: "Main Page",
		Name:  "Marcel",
	}

	return func(w http.ResponseWriter, r *http.Request) {
		data := &index{common: indexCommon, Age: "456"}
		err := site.ExecuteTemplate(w, "layout", data)
		if err != nil {
			log.Fatal(err)
		}
	}
}
