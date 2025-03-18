package utils

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func RandBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("Error: could not create random bytes: %v", err)
	}
	return b, nil
}

func CreateNewHash(n []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(n, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}
