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

func (f *Form) getHandler(w http.ResponseWriter, r *http.Request, formLanguages []FormLanguage) (data Task5Data, err error) {
	session, _ := store.Get(r, "cookie-name")
	if userId, _ := session.Values["user_id"]; userId != nil {
		var form *entity.Form
		form, _ = f.formApp.GetFormByUserId(userId.(int))

		for index, language := range formLanguages {
			if slices.Contains(form.Languages, language.Id) {
				formLanguages[index].Selected = true
			}
		}
		data = Task5Data{
			Languages: formLanguages,
			Form:      *form,
		}
	} else {
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
	}

	return data, nil
}

func (f *Form) postHandler(
	w http.ResponseWriter, r *http.Request, formLanguages []FormLanguage,
) (data Task5Data, err error) {
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
		session, _ := store.Get(r, "cookie-name")
		if userId, _ := session.Values["user_id"]; userId != nil {
			form.UserId = userId.(int)
			// todo оптимизировать две строки снизу =)
			userForm, _ := f.formApp.GetFormByUserId(userId.(int))
			form.Id = userForm.Id

			_, err = f.formApp.UpdateForm(&form)
			if err != nil {
				return data, err
			}
			data = Task5Data{
				SuccessMessage: "Форма была обновлена.",
			}
			return data, nil
		}
		user, _ := f.userApp.CreateNewUser()
		userId := user.Id
		f.Login(r, w, userId)
		if _, err = f.formApp.SaveFormWithUser(&form, userId); err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return data, err
		}
		saveFormInCookies(w, form)
		data = Task5Data{
			SuccessMessage: fmt.Sprintf(
				"Форма была отправлена, спасибо! Ваш логин: %s Ваш пароль: %s",
				user.Login,
				user.Password,
			),
		}
	}

	return data, nil
}

func (f *Form) Task5(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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

	// tired
	if r.Method == http.MethodGet {
		data, err = f.getHandler(w, r, formLanguages)
	} else {
		data, err = f.postHandler(w, r, formLanguages)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
