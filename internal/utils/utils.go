package utils

import "math/rand"

func GenerateRandomStrSequence(chars string, length int) string {
	result := make([]byte, length)

	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}

	return string(result)
}
