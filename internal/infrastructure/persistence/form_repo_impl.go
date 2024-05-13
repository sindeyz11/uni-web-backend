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
	dbErr := make(map[string]string)

	query := `INSERT INTO forms (fio, phone, email, birthday, gender, biography)
				VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.Conn.Exec(query, form.Fio, form.Phone, form.Email, form.Birthday, form.Gender, form.Biography)
	if err != nil {
		dbErr["Error saving a form"] = err.Error()
		return nil, dbErr
	}

	return form, nil
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
