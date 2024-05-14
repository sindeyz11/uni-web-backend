package impl

import (
	"database/sql"
	"uni-web/internal/domain/entity"
	"uni-web/internal/domain/repository"
	"uni-web/internal/utils"
)

type FormRepo struct {
	Conn *sql.DB
}

func NewFormRepository(conn *sql.DB) *FormRepo {
	return &FormRepo{Conn: conn}
}

var _ repository.FormRepository = &FormRepo{}

func (r *FormRepo) SaveForm(form *entity.Form, languages []int) (*entity.Form, error) {
	// saving form
	query := `INSERT INTO form (fio, phone, email, birthday, gender, biography)
				VALUES ($1, $2, $3, $4, $5, $6);`
	_, err := r.Conn.Exec(query, form.Fio, form.Phone, form.Email, form.Birthday, form.Gender, form.Biography)
	if err != nil {
		return nil, err
	}

	// saving languages todo form id, languages
	var vals []interface{}
	for _, language := range languages {
		vals = append(vals, 7, language)
	}
	sqlStr := `INSERT INTO forms_languages (id_form, id_language) VALUES %s`
	sqlStr = utils.ReplaceSQL(sqlStr, "(?, ?)", len(languages))
	stmt, _ := r.Conn.Prepare(sqlStr)
	if _, err = stmt.Exec(vals...); err != nil {
		return nil, err

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

func (r *FormRepo) UpdateForm(form *entity.Form) (*entity.Form, error) {
	//TODO implement me
	panic("implement me")
}

func (r *FormRepo) DeleteForm(formId int) (int, error) {
	//TODO implement me
	panic("implement me")
}
