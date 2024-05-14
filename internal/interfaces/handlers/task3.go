package handlers

import (
	"html/template"
	"net/http"
	"uni-web/internal/domain/entity"
)

func (f *Form) Task3(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/templates/task3.html"))
	languages, err := f.languageApp.GetAllLanguages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
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
			data = Task3Data{
				Languages: languages,
				Errors:    formErrors,
				Message:   "Ошибка. Форма содержала неверные данные: ",
			}
		} else {
			f.formApp.SaveForm(&form, []int{1, 2, 3})
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
