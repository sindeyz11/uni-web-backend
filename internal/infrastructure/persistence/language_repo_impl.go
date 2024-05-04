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

func (r *LanguageRepo) GetAllLanguages() []entity.Language {
	//TODO implement me
	panic("implement me")
}
