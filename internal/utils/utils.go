package utils

import "math/rand"

func GenerateRandomSequence(chars string, length int) string {
	result := make([]byte, length)

	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}

	return string(result)
}
