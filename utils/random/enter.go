package random

import "math/rand"

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandStringByCode(rg string, length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = rg[rand.Intn(len(rg))]
	}
	return string(b)
}
