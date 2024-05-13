package interfaces

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
	Errors    []string
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

func (f *Form) Task3(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/templates/task3.html"))
	languages, err := f.languageApp.GetAllLanguages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	if r.Method == "GET" {
		data := Task3Data{
			Languages: languages,
		}
		tmpl.Execute(w, data)
		return
	} else if r.Method == "POST" {
		form := entity.Form{
			Fio:       r.PostFormValue("fio"),
			Phone:     r.PostFormValue("phone"),
			Email:     r.PostFormValue("email"),
			Birthday:  r.PostFormValue("birthday"),
			Gender:    r.PostFormValue("gender"),
			Biography: r.PostFormValue("biography"),
		}
		formErrors := form.Validate([]int{1})

		var data Task3Data

		if len(formErrors) > 0 {
			var errorMessages []string

			for _, v := range formErrors {
				errorMessages = append(errorMessages, v)
			}

			data = Task3Data{
				Languages: languages,
				Errors:    errorMessages,
				Message:   "Ошибка. Форма содержала неверные данные: ",
			}
		} else {
			f.formApp.SaveForm(&form)
			data = Task3Data{
				Message: "Форма была отправлена, спасибо!",
			}
		}
		// todo сохранение языков
		tmpl.Execute(w, data)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
