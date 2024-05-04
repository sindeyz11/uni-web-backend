package repository

import "uni-web/internal/domain/entity"

type LanguageRepository interface {
	GetAllLanguages() []entity.Language
}
