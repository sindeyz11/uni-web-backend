package repository

import "uni-web/internal/domain/entity"

type UserRepository interface {
	CreateNewUser() (*entity.User, error)
	GetUserByLogin(login string) (*entity.User, error)
}
