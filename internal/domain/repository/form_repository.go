package repository

import "uni-web/internal/domain/entity"

type FormRepository interface {
	SaveForm(*entity.Form) (*entity.Form, map[string]string)
	GetForm(int) *entity.Form
	GetAllForms() []entity.Form
	UpdateForm(*entity.Form) (*entity.Form, map[string]string)
	DeleteForm(int)
}
