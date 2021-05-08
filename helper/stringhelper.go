package helper

import (
	"crypto/rand"
)

func GenerateSecret(len int) []byte {
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		panic("could not generate secret: " + err.Error())
	}

	return b
}
