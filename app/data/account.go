package data

import (
	"arc/app/crypt"
	"context"

	"github.com/jmoiron/sqlx"
)

func CreateNewAccount(db *sqlx.DB, ctx context.Context, name, email, password, publicKey string) error {
	hashedPassword, hashErr := crypt.HashText(password)
	if hashErr != nil {
		return hashErr
	}

	tx, txErr := db.BeginTxx(ctx, nil)
	if txErr != nil {
		return txErr
	}
	tx.Exec("INSERT INTO accounts (name, email, password, hashed_password, hashed_password_encrypted) VALUES ($1, $2, $3, $4)", name, email, hashedPassword, hashedPassword)
	return tx.Commit()
}

func Login(login, password string) (string, error) {
	return "", nil
}

func AccountRecovery(email, login, lastPassword, publicKey, otpCode string) (string, error) {
	return "", nil
}
