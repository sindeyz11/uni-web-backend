package entity

import (
	"regexp"
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

func LanguagesParseString(languages string) ([]int, error) {
	re := regexp.MustCompile("\\d+")
	matches := re.FindAllString(languages, -1)

	var numbers []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}
