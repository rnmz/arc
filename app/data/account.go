package data

import (
	"arc/app/common"
	"arc/app/entity"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Account struct{}

func (account Account) CreateNewAccount(db *sqlx.DB, ctx context.Context, name, login, email, password, publicKey string) error {
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

func (account Account) Login(db *sqlx.DB, ctx context.Context, login, password string) (entity.Account, error) {
	var userData entity.Account
	resErr := db.SelectContext(ctx, &userData, `SELECT 1 FROM accounts WHERE login = $1 and password = $2`, login, password)
	if resErr != nil {
		return entity.Account{}, resErr
	}
	return userData, nil
}

func (account Account) AccountRecovery(db *sqlx.DB, ctx context.Context, email, login, publicKey string) (entity.Account, error) {
	var userData entity.Account
	resErr := db.SelectContext(ctx, &userData, `SELECT 1 FROM accounts WHERE email = $1 and login = $2 and public_key = $3`, email, login, publicKey)
	if resErr != nil {
		return entity.Account{}, resErr
	}
	return userData, nil
}

func (account Account) ChangeRole(db *sqlx.DB, ctx context.Context, id uuid.UUID, status common.AccountStatus) error {
	tx, txErr := db.BeginTxx(ctx, nil)
	if txErr != nil {
		return txErr
	}
	_, _ = tx.Exec(`UPDATE accounts SET status = $1 WHERE id = $2`, status, id)
	return tx.Commit()
}

func (account Account) GetById(db *sqlx.DB, ctx context.Context, id string) (entity.Account, error) {
	var userData entity.Account
	err := db.GetContext(ctx, &userData, `SELECT 1 FROM accounts WHERE id = $1`, id)
	return userData, err
}

func (account Account) GetAccounts(db *sqlx.DB, ctx context.Context, page, itemsPerPage int) ([]entity.Account, error) {
	var accounts []entity.Account
	err := db.GetContext(ctx, &accounts, `SELECT * FROM accounts`)
	return accounts, err
}

func (account Account) GetAccountsCount(db *sqlx.DB, ctx context.Context, id string) (int, error) {
	var count int
	err := db.GetContext(ctx, &count, `SELECT COUNT(*) FROM accounts`)
	return count, err
}
