package entity

import (
	"net/http"
	"regexp"
)

type Form struct {
	Id        int    `json:"id"`
	Fio       string `json:"fio"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Birthday  string `json:"birthday"`
	Gender    string `json:"gender"`
	Biography string `json:"biography"`
	Languages []int  `json:"languages"`
}

func (form *Form) Validate(languages []int) map[string]string {
	var errors = make(map[string]string)

	reg := regexp.MustCompile("[A-zА-яёЁ\\s]{1,150}")
	if match := reg.Match([]byte(form.Fio)); !match {
		errors["fio"] = "ФИО должно содержать только буквы и пробелы и быть не длиннее 150 символов"
	}

	reg = regexp.MustCompile("\\+\\d{11}")
	if match := reg.Match([]byte(form.Phone)); !match {
		errors["phone"] = "Телефон должен быть указан в формате +79189999999"
	}

	reg = regexp.MustCompile("[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+")
	if match := reg.Match([]byte(form.Email)); !match {
		errors["email"] = "Введите действующий email"
	}

	if form.Gender != "f" && form.Gender != "m" {
		errors["gender"] = "Указан несуществующий пол"
	}

	if len(languages) == 0 {
		errors["languages"] = "Список языков пуст"
	}

	for _, num := range languages {
		if num > 11 || num < 0 {
			errors["languages"] = "Указан несуществующий язык программирования"
			break
		}
	}

	return errors
}

func GetFormFromRequest(r *http.Request) Form {
	return Form{
		Fio:       r.PostFormValue("fio"),
		Phone:     r.PostFormValue("phone"),
		Email:     r.PostFormValue("email"),
		Birthday:  r.PostFormValue("birthday"),
		Gender:    r.PostFormValue("gender"),
		Biography: r.PostFormValue("biography"),
	}
}
