package lib

import "math/rand"

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"

func GenerateUID () string {
	b := make([]byte, 8)
	for i := 0; i < len(b); i++ {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}