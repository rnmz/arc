package data

import (
	"arc/app/entity"
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

func CreateNewAccount(db *sqlx.DB, ctx context.Context, name, login, email, password, publicKey string) error {
	tx, txErr := db.BeginTxx(ctx, nil)
	if txErr != nil {
		return txErr
	}
	_, _ = tx.Exec(`INSERT INTO accounts (
  		name, login, email,
	  	password, registration_date,
	  	last_login_date, public_key) VALUES ($1, $2, $3, $4, $5, $6, $7)`, name, login, email, password, time.Now(), time.Now(), publicKey)
	return tx.Commit()
}

func Login(db *sqlx.DB, ctx context.Context, login, password string) (entity.Account, error) {
	var userData entity.Account
	resErr := db.SelectContext(ctx, &userData, `SELECT 1 FROM accounts WHERE login = $1 and password = $2`, login, password)
	if resErr != nil {
		return entity.Account{}, resErr
	}
	return userData, nil
}

func AccountRecovery(db *sqlx.DB, ctx context.Context, email, login, publicKey string) (entity.Account, error) {
	var userData entity.Account
	resErr := db.SelectContext(ctx, &userData, `SELECT 1 FROM accounts WHERE email = $1 and login = $2 and public_key = $3`, email, login, publicKey)
	if resErr != nil {
		return entity.Account{}, resErr
	}
	return userData, nil
}
