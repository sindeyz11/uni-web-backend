package application

import (
	"uni-web/internal/domain/entity"
	"uni-web/internal/domain/repository"
)

type formApp struct {
	fr repository.FormRepository
}

var _ FormAppInterface = &formApp{}

type FormAppInterface interface {
	SaveForm(*entity.Form) (*entity.Form, map[string]string)
	GetForm(int) *entity.Form
	GetAllForms() []entity.Form
	UpdateForm(*entity.Form) (*entity.Form, map[string]string)
	DeleteForm(int)
}

func (f *formApp) SaveForm(form *entity.Form) (*entity.Form, map[string]string) {
	return f.fr.SaveForm(form)
}

func (f *formApp) GetForm(formId int) *entity.Form {
	return f.fr.GetForm(formId)
}

func (f *formApp) GetAllForms() []entity.Form {
	return f.fr.GetAllForms()
}

func (f *formApp) UpdateForm(form *entity.Form) (*entity.Form, map[string]string) {
	return f.fr.UpdateForm(form)
}

func (f *formApp) DeleteForm(formId int) {
	f.fr.DeleteForm(formId)
}
