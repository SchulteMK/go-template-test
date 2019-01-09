package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"path"
)

type server struct {
	mux          *mux.Router
	templatePath string
}

type common struct {
	Title string
	Name  string
}

func New(baseUrl, templatePath string) (*server, error) {
	s := &server{
		mux:          mux.NewRouter(),
		templatePath: templatePath,
	}

	commonTemplates, err := template.ParseFiles(
		path.Join(s.templatePath, "layout.html"),
		path.Join(s.templatePath, "footer.html"),
		path.Join(s.templatePath, "header.html"))
	if err != nil {
		log.Fatal(err)
	}

	router := s.mux.PathPrefix(baseUrl).Subrouter()

	router.HandleFunc("/", s.createIndex(commonTemplates)).Methods("GET")

	return s, nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	s.mux.ServeHTTP(w, r)
}
