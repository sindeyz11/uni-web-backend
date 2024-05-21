package handlers

import (
	"html/template"
	"net/http"
	"uni-web/internal/application"
)

type Form struct {
	formApp     application.FormAppInterface
	languageApp application.LanguageAppInterface
	userApp     application.UserAppInterface
}

func NewForm(
	fApp application.FormAppInterface,
	lApp application.LanguageAppInterface,
	uApp application.UserAppInterface,
) *Form {
	return &Form{
		formApp:     fApp,
		languageApp: lApp,
		userApp:     uApp,
	}
}

func Task1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/templates/task1.html"))
	tmpl.Execute(w, nil)
}

func Task2(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/templates/task2.html"))
	tmpl.Execute(w, nil)
}
