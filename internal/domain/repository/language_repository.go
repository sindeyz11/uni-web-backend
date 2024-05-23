package repository

import "uni-web/internal/domain/entity"

type LanguageRepository interface {
	GetAllLanguages() ([]entity.Language, error)
	// operating with []int of languages ids
	//GetLanguagesByFormId(formId int) ([]int, error)
	//CreateLanguagesByFormId(formId int, languages []int) (int, error)
	//DeleteLanguagesByFormId(formId int) (int, error)
}
