package repository

import "uni-web/internal/domain/entity"

type FormRepository interface {
	SaveForm(form *entity.Form) (*entity.Form, map[string]string)
	GetForm(formId int) (*entity.Form, error)
	GetAllForms() ([]entity.Form, error)
	UpdateForm(form *entity.Form) (*entity.Form, map[string]string)
	DeleteForm(formId int) (int, error)
}
