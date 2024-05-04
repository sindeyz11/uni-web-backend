package persistence

import (
	"database/sql"
	"uni-web/internal/domain/repository"
)

type Repositories struct {
	Form     repository.FormRepository
	Language repository.LanguageRepository
	Db       *sql.DB
}

func NewRepositories(conn *sql.DB) (*Repositories, error) {
	return &Repositories{
		Form:     NewFormRepository(conn),
		Language: NewLanguageRepository(conn),
		Db:       conn,
	}, nil
}

func (r *Repositories) Close() error {
	return r.Db.Close()
}
