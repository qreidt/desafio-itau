package util

import "math/rand"

const l = 62
const numbers = "0123456789"
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandStringBytes(n uint8) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = letters[rand.Intn(l)]
	}

	return string(b)
}

func RandNumbers(n uint8) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = numbers[rand.Intn(10)]
	}

	return string(b)
}
