package handlers

import (
	"html/template"
	"net/http"
	"uni-web/internal/domain/entity"
)

type Task3Data struct {
	Languages []entity.Language
	Errors    map[string]string
	Message   string
}

func (f *Form) Task3(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/templates/task3.html"))
	allLanguages, err := f.languageApp.GetAllLanguages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		data := Task3Data{
			Languages: allLanguages,
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
		languages := entity.LanguagesParseForm(r.Form["languages"])

		formErrors := form.Validate(languages)
		var data Task3Data
		if len(formErrors) > 0 {
			data = Task3Data{
				Languages: allLanguages,
				Errors:    formErrors,
				Message:   "Ошибка. Форма содержала неверные данные: ",
			}
		} else {
			if _, err = f.formApp.SaveForm(&form, languages); err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			data = Task3Data{
				Message: "Форма была отправлена, спасибо!",
			}
		}
		tmpl.Execute(w, data)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
