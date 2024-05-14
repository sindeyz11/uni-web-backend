package entity

import "regexp"

type Form struct {
	Id        int
	Fio       string
	Phone     string
	Email     string
	Birthday  string
	Gender    string
	Biography string
}

func (form *Form) Validate(languages []int) map[string]string {
	var errors = make(map[string]string)

	reg := regexp.MustCompile("[A-zА-яёЁ\\s]{1,150}")
	if match := reg.Match([]byte(form.Fio)); !match {
		errors["Fio"] = "ФИО должно содержать только буквы и пробелы и быть не длиннее 150 символов"
	}

	reg = regexp.MustCompile("\\+\\d{11}")
	if match := reg.Match([]byte(form.Phone)); !match {
		errors["Phone"] = "Телефон должен быть указан в формате +79189999999"
	}

	reg = regexp.MustCompile("[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+")
	if match := reg.Match([]byte(form.Email)); !match {
		errors["Email"] = "Введите действующий email"
	}

	if form.Gender != "f" && form.Gender != "m" {
		errors["Gender"] = "Указан несуществующий пол"
	}

	for _, num := range languages {
		if num > 11 || num < 0 {
			errors["Languages"] = "Указан несуществующий язык программирования"
			break
		}
	}

	return errors
}
