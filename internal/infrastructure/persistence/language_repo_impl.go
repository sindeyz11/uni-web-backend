package persistence

import (
	"database/sql"
	"uni-web/internal/domain/entity"
	"uni-web/internal/domain/repository"
)

type LanguageRepo struct {
	Conn *sql.DB
}

func NewLanguageRepository(conn *sql.DB) *LanguageRepo {
	return &LanguageRepo{Conn: conn}
}

var _ repository.LanguageRepository = &LanguageRepo{}

func (r *LanguageRepo) GetAllLanguages() ([]entity.Language, error) {
	// TODO ADD CACHE FOR 1 WEEK

	query := `SELECT * FROM languages;`
	rows, err := r.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var languages []entity.Language

	for rows.Next() {
		var lang entity.Language
		if err := rows.Scan(&lang.Id, &lang.Title); err != nil {
			return languages, err
		}
		languages = append(languages, lang)
	}
	if err = rows.Err(); err != nil {
		return languages, err
	}
	return languages, nil
}
