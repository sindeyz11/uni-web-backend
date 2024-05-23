package impl

import (
	"database/sql"
	"errors"
	"strings"
	"uni-web/internal/domain/entity"
	"uni-web/internal/domain/repository"
	"uni-web/internal/utils"
)

// TODO add transactions?

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
				VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`
	formId := -734
	err := r.Conn.QueryRow(query,
		form.Fio, form.Phone, form.Email, form.Birthday, form.Gender, form.Biography,
	).Scan(&formId)
	if err != nil {
		return nil, err
	}
	// saving languages
	_, err = r.CreateLanguagesByFormId(formId, languages)
	if err != nil {
		return nil, err
	}
	return form, nil
}

func (r *FormRepo) SaveFormWithUser(form *entity.Form, userId int) (*entity.Form, error) {
	// saving form
	query := `INSERT INTO form (fio, phone, email, birthday, gender, biography, id_user)
				VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`
	formId := -734
	err := r.Conn.QueryRow(query,
		form.Fio, form.Phone, form.Email, form.Birthday, form.Gender, form.Biography, userId,
	).Scan(&formId)
	if err != nil {
		return nil, err
	}
	form.Id = formId

	// saving languages
	_, err = r.CreateLanguagesByFormId(formId, form.Languages)
	if err != nil {
		return nil, err
	}
	return form, nil
}

func (r *FormRepo) GetForm(formId int) (*entity.Form, error) {
	panic("implement me")
}

func (r *FormRepo) GetAllForms() ([]entity.Form, error) {
	panic("implement me")
}

func (r *FormRepo) UpdateForm(form *entity.Form) (*entity.Form, error) {
	query := `UPDATE form
				SET fio = $1,
				    phone = $2,
				    email = $3,
				    birthday = $4,
				    gender = $5,
				    biography = $6
				WHERE id_user = $7;`
	_, err := r.Conn.Exec(query,
		form.Fio, form.Phone, form.Email, form.Birthday, form.Gender, form.Biography, form.UserId,
	)
	if err != nil {
		return nil, errors.New("an error occurred while updating form")
	}

	if _, err := r.DeleteLanguagesByFormId(form.Id); err != nil {
		return nil, errors.New("an error occurred while deleting languages")
	}

	if _, err := r.CreateLanguagesByFormId(form.Id, form.Languages); err != nil {
		return nil, errors.New("an error occurred while creating languages")
	}

	return form, nil
}

func (r *FormRepo) DeleteForm(formId int) (int, error) {
	panic("implement me")
}

func (r *FormRepo) GetLanguagesByFormId(formId int) ([]int, error) {
	query := `SELECT id_language FROM forms_languages WHERE id_form=$1;`
	rows, err := r.Conn.Query(query, formId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var languages []int

	for rows.Next() {
		var langId int
		if err := rows.Scan(&langId); err != nil {
			return languages, err
		}
		languages = append(languages, langId)
	}
	if err = rows.Err(); err != nil {
		return languages, err
	}
	return languages, nil
}

func (r *FormRepo) CreateLanguagesByFormId(formId int, languages []int) (int, error) {
	var vals []interface{}
	for _, language := range languages {
		vals = append(vals, formId, language)
	}
	sqlStr := `INSERT INTO forms_languages (id_form, id_language) VALUES %s`
	sqlStr = utils.ReplaceSQL(sqlStr, "(?, ?)", len(languages))
	stmt, _ := r.Conn.Prepare(sqlStr)
	if _, err := stmt.Exec(vals...); err != nil {
		return -1, err
	}
	return formId, nil
}

func (r *FormRepo) DeleteLanguagesByFormId(formId int) (int, error) {
	query := `DELETE FROM forms_languages WHERE id_form=$1;`
	_, err := r.Conn.Exec(query, formId)
	if err != nil {
		return -734, errors.New("form/languages not found")
	}
	return formId, nil
}

func (r *FormRepo) GetFormByUserId(userId int) (*entity.Form, error) {
	var resForm entity.Form
	query := `SELECT id, fio, phone, email, birthday, gender, biography FROM form WHERE id_user=$1;`
	row := r.Conn.QueryRow(query, userId)
	err := row.Scan(
		&resForm.Id, &resForm.Fio, &resForm.Phone, &resForm.Email,
		&resForm.Birthday, &resForm.Gender, &resForm.Biography,
	)
	if err != nil {
		return nil, errors.New("form not found")
	}
	resForm.Birthday = strings.Split(resForm.Birthday, "T")[0]

	languages, err := r.GetLanguagesByFormId(resForm.Id)
	if err != nil {
		return nil, errors.New("form not found")
	}
	resForm.Languages = languages
	return &resForm, nil
}
