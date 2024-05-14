package handlers

import (
	"html/template"
	"net/http"
	"uni-web/internal/application"
	"uni-web/internal/domain/entity"
)

type Form struct {
	formApp     application.FormAppInterface
	languageApp application.LanguageAppInterface
}

func NewForm(fApp application.FormAppInterface, lApp application.LanguageAppInterface) *Form {
	return &Form{
		formApp:     fApp,
		languageApp: lApp,
	}
}

type Task3Data struct {
	Languages []entity.Language
	Errors    map[string]string
	Message   string
}

func Task1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/templates/task1.html"))
	tmpl.Execute(w, nil)
}

func Task2(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/templates/task2.html"))
	tmpl.Execute(w, nil)
}
