package nserver

import (
	"fmt"
	"net/http"
	"text/template"
)

type Nserver struct {
	port string
}

var templates = template.Must(template.ParseFiles("static/index.html"))

func (serv Nserver) StartServer(port string) {
	serv.port = ":" + port
	fs := http.FileServer(http.Dir("static/"))
	http.HandleFunc("/", serv.indexHandler)
	http.HandleFunc("/api", serv.apiHandler)
	http.HandleFunc("/about", serv.aboutHandler)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/registration", serv.registrationHandler)
	fmt.Println("Server is running on port: ", serv.port)
	http.ListenAndServe(serv.port, nil)
}

func (serv Nserver) indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (serv Nserver) apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the API Page!")
}

func (serv Nserver) aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the About Page!")
}

func (serv Nserver) registrationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Registration Page!")
}
