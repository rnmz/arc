package crypt

import (
	"crypto/hmac"
	"crypto/rand"
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

func GenerateAccountKey() ([]byte, error) {
	raw := make([]byte, 32)
	if _, err := rand.Read(raw); err != nil {
		return nil, err
	}
	return raw, nil
}

func HashAccountKey(masterKey, rawKey []byte) (string, error) {
	sum := peppered(masterKey, rawKey)
	bytesHash, err := bcrypt.GenerateFromPassword(sum, bcrypt.DefaultCost)
	return string(bytesHash), err
}

func CheckAccountKey(masterKey, rawKey []byte, hash string) bool {
	sum := peppered(masterKey, rawKey)
	return bcrypt.CompareHashAndPassword([]byte(hash), sum) == nil
}

func peppered(masterKey, rawKey []byte) []byte {
	h := hmac.New(sha256.New, masterKey)
	h.Write(rawKey)
	return h.Sum(nil)
}
