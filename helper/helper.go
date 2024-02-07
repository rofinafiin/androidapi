package helper

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())

	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(result)
}
