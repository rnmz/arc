package data

import (
	"arc/app/common"
	"arc/app/crypt"
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

func CreateNewAccount(db *sqlx.DB, ctx context.Context, name, email, password, publicKey string) error {
	hashedPassword, hashErr := crypt.HashText(password)
	if hashErr != nil {
		return hashErr
	}

	tx, txErr := db.BeginTxx(ctx)
	if txErr != nil {
		return common.TxxError{
			TxFunctionName: "CreateNewAccount",
			Time:           time.Now(),
		}
	}

	return nil
}

func Login(login, password string) (string, error) {
	return "", nil
}

func AccountRecovery(email, login, lastPassword, publicKey, otpCode string) (string, error) {
	return "", nil
}
