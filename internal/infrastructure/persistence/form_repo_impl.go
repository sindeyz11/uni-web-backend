package persistence

import (
	"database/sql"
	"uni-web/internal/domain/entity"
	"uni-web/internal/domain/repository"
)

type FormRepo struct {
	Conn *sql.DB
}

func NewFormRepository(conn *sql.DB) *FormRepo {
	return &FormRepo{Conn: conn}
}

var _ repository.FormRepository = &FormRepo{}

func (r *FormRepo) SaveForm(form *entity.Form) (*entity.Form, map[string]string) {
	//TODO implement me
	panic("implement me")
}

func (r *FormRepo) GetForm(formId int) (*entity.Form, error) {
	//TODO implement me
	panic("implement me")
}

func (r *FormRepo) GetAllForms() ([]entity.Form, error) {
	//TODO implement me
	panic("implement me")
}

func (r *FormRepo) UpdateForm(form *entity.Form) (*entity.Form, map[string]string) {
	//TODO implement me
	panic("implement me")
}

func (r *FormRepo) DeleteForm(formId int) (int, error) {
	//TODO implement me
	panic("implement me")
}
