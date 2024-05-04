package interfaces

import (
	"html/template"
	"net/http"
)

func Task1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/templates/task1.html"))
	tmpl.Execute(w, nil)
}

func Task2(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/templates/task2.html"))
	tmpl.Execute(w, nil)
}

func Task3(w http.ResponseWriter, r *http.Request) {
	// Todo
}
