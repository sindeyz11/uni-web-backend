package application

import (
	"uni-web/internal/domain/entity"
	"uni-web/internal/domain/repository"
)

type userApp struct {
	ur repository.UserRepository
}

var _ UserAppInterface = &userApp{}

type UserAppInterface interface {
	CreateNewUser() (*entity.User, error)
	GetUserByLogin(login string) (*entity.User, error)
}

func (u *userApp) CreateNewUser() (*entity.User, error) {
	return u.ur.CreateNewUser()
}

func (u *userApp) GetUserByLogin(login string) (*entity.User, error) {
	return u.ur.GetUserByLogin(login)
}
