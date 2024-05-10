package interfaces

import (
	"fmt"
	"html/template"
	"net/http"
	"uni-web/internal/application"
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

func Task1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/templates/task1.html"))
	tmpl.Execute(w, nil)
}

func Task2(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/templates/task2.html"))
	tmpl.Execute(w, nil)
}

func (form *Form) Task3(w http.ResponseWriter, r *http.Request) {
	languages, err := form.languageApp.GetAllLanguages()
	if err != nil {
		return
	}
	fmt.Println(languages)
	tmpl := template.Must(template.ParseFiles("assets/templates/task2.html"))
	tmpl.Execute(w, nil)
}
