package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/oleiade/reflections"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"
	"uni-web/internal/domain/entity"
)

type FormLanguage struct {
	Id       int
	Title    string
	Selected bool
}

func convertLanguagesToFormLanguages(allLanguages []entity.Language) []FormLanguage {
	var formLanguages []FormLanguage
	for _, language := range allLanguages {
		fl := FormLanguage{
			Id:    language.Id,
			Title: language.Title,
		}
		formLanguages = append(formLanguages, fl)
	}
	return formLanguages
}

func saveErrorsInCookies(w http.ResponseWriter, r *http.Request, formErrors map[string]string) {
	// Delete fixed forms errors from cookies
	cookies := r.Cookies()
	for _, cookie := range cookies {
		reg := regexp.MustCompile("(.+)-formError")
		if match := reg.FindStringSubmatch(cookie.Name); match != nil {
			if _, ok := formErrors[match[1]]; !ok {
				cookie := http.Cookie{
					Name:    cookie.Name,
					Value:   "idc",
					Expires: time.Now().Add(-time.Hour),
				}
				http.SetCookie(w, &cookie)
			}
		}
	}

	// Add from formErrors to cookies
	for key, value := range formErrors {
		cookieName := fmt.Sprintf("%s-formError", strings.ToLower(key))
		encodedValue := url.QueryEscape(value)

		cookie := http.Cookie{
			Name:    cookieName,
			Value:   encodedValue,
			Expires: time.Now().Add(365 * 24 * time.Hour),
		}
		http.SetCookie(w, &cookie)
	}
}

func getFormErrorsFromCookies(r *http.Request) map[string]string {
	formErrors := make(map[string]string)
	cookies := r.Cookies()

	for _, cookie := range cookies {
		reg := regexp.MustCompile("(.+)-formError")
		if match := reg.FindStringSubmatch(cookie.Name); match != nil {
			if decodedValue, err := url.QueryUnescape(cookie.Value); err == nil {
				formErrors[match[1]] = decodedValue
			}
		}
	}
	return formErrors
}

func saveFormInCookies(w http.ResponseWriter, form entity.Form) {
	expires := time.Now().Add(365 * 24 * time.Hour)
	val := reflect.ValueOf(form)
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		//fmt.Printf("Field Name: %s, Field Value: %v\n", fieldType.Name, field.Interface())

		if fieldType.Name == "Id" || fieldType.Name == "UserId" {
			continue
		}

		var cookie http.Cookie
		if fieldType.Name == "Languages" {
			data := field.Interface().([]int)
			s, _ := json.Marshal(data)
			cookie = http.Cookie{
				Name:    strings.ToLower(fieldType.Name),
				Value:   string(s),
				Expires: expires,
			}
		} else {
			cookie = http.Cookie{
				Name:    strings.ToLower(fieldType.Name),
				Value:   field.Interface().(string),
				Expires: expires,
			}
		}
		http.SetCookie(w, &cookie)

	}
}

func getFormFromCookies(r *http.Request) (form entity.Form, err error) {
	val := reflect.ValueOf(form)
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		//field := val.Field(i)
		fieldType := typ.Field(i)
		//fmt.Printf("Field Name: %s, Field Value: %v\n", fieldType.Name, field.Interface())

		if fieldType.Name == "Id" || fieldType.Name == "UserId" {
			continue
		}

		if fieldType.Name == "Languages" {
			cookie, err := r.Cookie(strings.ToLower(fieldType.Name))
			if err != nil {
				return form, err
			}
			languages, err := entity.LanguagesParseString(cookie.Value)
			if err != nil {
				return form, err
			}
			form.Languages = languages
			continue
		}

		cookie, err := r.Cookie(strings.ToLower(fieldType.Name))
		if err != nil {
			return form, err
		}
		if reflections.SetField(&form, fieldType.Name, cookie.Value) != nil {
			return form, err
		}
	}
	return form, nil
}
