package persistence

import (
	"database/sql"
	"uni-web/internal/domain/repository"
	"uni-web/internal/infrastructure/persistence/impl"
)

type Repositories struct {
	Form     repository.FormRepository
	Language repository.LanguageRepository
	Db       *sql.DB
}

func NewRepositories(conn *sql.DB) (*Repositories, error) {
	return &Repositories{
		Form:     impl.NewFormRepository(conn),
		Language: impl.NewLanguageRepository(conn),
		Db:       conn,
	}, nil
}

func (r *Repositories) Close() error {
	return r.Db.Close()
}
