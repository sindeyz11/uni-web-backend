package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"slices"
	"uni-web/internal/domain/entity"
)

type Task5Data struct {
	Languages      []FormLanguage
	Errors         map[string]string
	Message        string
	SuccessMessage string
	Form           entity.Form
}

func (f *Form) Task5(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	session, _ := store.Get(r, "cookie-name")
	if auth, _ := session.Values["authenticated"].(bool); auth {
		http.Error(w, "authorized=)", http.StatusForbidden)
		return
	}

	allLanguages, err := f.languageApp.GetAllLanguages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("assets/templates/task5.html"))
	formLanguages := convertLanguagesToFormLanguages(allLanguages)
	var data Task5Data

	if r.Method == "GET" {
		formErrors := getFormErrorsFromCookies(r)

		if form, err := getFormFromCookies(r); err == nil {
			for index, language := range formLanguages {
				if slices.Contains(form.Languages, language.Id) {
					formLanguages[index].Selected = true
				}
			}
			data = Task5Data{
				Languages: formLanguages,
				Errors:    formErrors,
				Message:   "Ошибка. Форма содержала неверные данные: ",
				Form:      form,
			}
		} else {
			data = Task5Data{
				Languages: formLanguages,
				Message:   "Ошибка. Форма содержала неверные данные: ",
				Errors:    formErrors,
			}
		}
	} else {
		form := entity.GetFormFromRequest(r)
		languages := entity.LanguagesParseForm(r.Form["languages"])
		form.Languages = languages
		formErrors := form.Validate(languages)
		saveErrorsInCookies(w, r, formErrors)

		if len(formErrors) > 0 {
			for index, language := range formLanguages {
				if slices.Contains(languages, language.Id) {
					formLanguages[index].Selected = true
				}
			}
			data = Task5Data{
				Languages: formLanguages,
				Errors:    formErrors,
				Message:   "Ошибка. Форма содержала неверные данные: ",
				Form:      form,
			}
		} else {
			if _, err = f.formApp.SaveForm(&form, languages); err != nil {
				http.Error(w, err.Error(), http.StatusBadGateway)
				return
			}
			user, _ := f.userApp.CreateNewUser()
			saveFormInCookies(w, form)
			data = Task5Data{
				SuccessMessage: fmt.Sprintf(
					"Форма была отправлена, спасибо! Ваш логин: %s Ваш пароль: %s",
					user.Login,
					user.Password,
				),
			}
		}
	}
	tmpl.Execute(w, data)
}
