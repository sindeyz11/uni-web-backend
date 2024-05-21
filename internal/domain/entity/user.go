package entity

import (
	"math/rand"
	"uni-web/internal/utils"
)

type User struct {
	Id       int
	Login    string
	Password string
}

func GenerateRandomLogin() string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	length := 8
	return utils.GenerateRandomSequence(chars, length)
}

func GenerateRandomPassword() string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+-=[]{}|;:,.<>?`~"
	length := rand.Intn(5) + 8
	return utils.GenerateRandomSequence(chars, length)
}
