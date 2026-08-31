package crypt

import (
	"crypto/sha256"

	"golang.org/x/crypto/bcrypt"
)

func HashText(text string) (string, error) {
	sha := sha256.Sum256([]byte(text))
	bytes, err := bcrypt.GenerateFromPassword(sha[:], bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckHashText(text, hash string) bool {
	sha := sha256.Sum256([]byte(text))
	err := bcrypt.CompareHashAndPassword([]byte(hash), sha[:])
	return err == nil
}
