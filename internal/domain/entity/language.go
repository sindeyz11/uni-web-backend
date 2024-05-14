package entity

import (
	"strconv"
)

type Language struct {
	Id    int
	Title string
}

func LanguagesParseForm(formLanguages []string) []int {
	var result []int
	for _, languageIdStr := range formLanguages {
		num, err := strconv.Atoi(languageIdStr)
		if err == nil {
			result = append(result, num)
		}
	}
	return result
}
