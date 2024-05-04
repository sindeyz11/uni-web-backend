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
	GetAllLanguages() []entity.Language
}

func (l *languageApp) GetAllLanguages() []entity.Language {
	return l.lr.GetAllLanguages()
}
