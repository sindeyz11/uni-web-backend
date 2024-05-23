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
	SaveForm(*entity.Form, []int) (*entity.Form, error)
	SaveFormWithUser(*entity.Form, int) (*entity.Form, error)
	GetForm(formId int) (*entity.Form, error)
	GetAllForms() ([]entity.Form, error)
	UpdateForm(*entity.Form) (*entity.Form, error)
	DeleteForm(formId int) (int, error)
	GetLanguagesByFormId(formId int) ([]int, error)
	CreateLanguagesByFormId(formId int, languages []int) (int, error)
	DeleteLanguagesByFormId(formId int) (int, error)
	GetFormByUserId(userId int) (*entity.Form, error)
}

func (f *formApp) SaveForm(form *entity.Form, languages []int) (*entity.Form, error) {
	return f.fr.SaveForm(form, languages)
}

func (f *formApp) SaveFormWithUser(form *entity.Form, userId int) (*entity.Form, error) {
	return f.fr.SaveFormWithUser(form, userId)
}

func (f *formApp) GetForm(formId int) (*entity.Form, error) {
	return f.fr.GetForm(formId)
}

func (f *formApp) GetAllForms() ([]entity.Form, error) {
	return f.fr.GetAllForms()
}

func (f *formApp) UpdateForm(form *entity.Form) (*entity.Form, error) {
	return f.fr.UpdateForm(form)
}

func (f *formApp) DeleteForm(formId int) (int, error) {
	return f.fr.DeleteForm(formId)
}

func (f *formApp) GetLanguagesByFormId(formId int) ([]int, error) {
	return f.fr.GetLanguagesByFormId(formId)
}

func (f *formApp) CreateLanguagesByFormId(formId int, languages []int) (int, error) {
	return f.fr.CreateLanguagesByFormId(formId, languages)
}

func (f *formApp) DeleteLanguagesByFormId(formId int) (int, error) {
	return f.fr.DeleteLanguagesByFormId(formId)
}

func (f *formApp) GetFormByUserId(userId int) (*entity.Form, error) {
	return f.fr.GetFormByUserId(userId)
}
