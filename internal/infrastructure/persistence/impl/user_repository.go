package impl

import (
	"database/sql"
	"errors"
	"uni-web/internal/domain/entity"
	"uni-web/internal/domain/repository"
	"uni-web/internal/infrastructure/security"
)

type UserRepo struct {
	Conn *sql.DB
}

func NewUserRepository(conn *sql.DB) *UserRepo {
	return &UserRepo{Conn: conn}
}

var _ repository.UserRepository = &UserRepo{}

func (r *UserRepo) CreateNewUser() (*entity.User, error) {
	password := entity.GenerateRandomPassword()
	user := entity.User{
		Login:    entity.GenerateRandomLogin(),
		Password: password,
	}
	hashedPassword, err := security.HashPassword(password)

	query := `INSERT INTO "user" (login, password)
				VALUES ($1, $2);`
	_, err = r.Conn.Exec(query,
		user.Login, hashedPassword,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUserByLogin(login string) (*entity.User, error) {
	var resUser entity.User
	query := `SELECT * FROM "user" WHERE login=$1;`
	row := r.Conn.QueryRow(query, login)
	err := row.Scan(&resUser.Id, &resUser.Login, &resUser.Password)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &resUser, nil
}
