package entity

import (
	"time"
)

type Form struct {
	Id       int
	Fio      string
	Phone    string
	Email    string
	Birthday time.Time
}

func (form *Form) Validate() map[string]string {
	var errorMessages = make(map[string]string)

	// Todo

	return errorMessages
}
