package repository

import "uni-web/internal/domain/entity"

type FormRepository interface {
	SaveForm(form *entity.Form, languages []int) (*entity.Form, error)
	GetForm(formId int) (*entity.Form, error)
	GetAllForms() ([]entity.Form, error)
	UpdateForm(form *entity.Form) (*entity.Form, error)
	DeleteForm(formId int) (int, error)
	//GetFormByUserId(userId int) (*entity.Form, error)
}
