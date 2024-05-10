package application

import (
	"uni-web/internal/domain/entity"
	"uni-web/internal/domain/repository"
)

type languageApp struct {
	lr repository.LanguageRepository
}

var _ LanguageAppInterface = &languageApp{}

type LanguageAppInterface interface {
	GetAllLanguages() ([]entity.Language, error)
}

func (l *languageApp) GetAllLanguages() ([]entity.Language, error) {
	return l.lr.GetAllLanguages()
}
