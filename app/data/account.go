package data

import (
	"arc/app/common"
	"arc/app/entity"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AccountData struct{}

func (_ AccountData) CreateNewAccount(db *sqlx.DB, ctx context.Context, defaultPlanId int, name, email, login, password, publicKey string) error {
	tx, txErr := db.BeginTxx(ctx, nil)
	if txErr != nil {
		return txErr
	}
	_, _ = tx.Exec(`INSERT INTO accounts (
  		name, email, login,
	  	password, registration_date,
	  	status, public_key) VALUES ($1, $2, $3, $4, $5, $6, $7)`, name, email, login, password, time.Now(), common.AccountStatusUser, publicKey)
	return tx.Commit()
}

func (_ AccountData) Login(db *sqlx.DB, ctx context.Context, login, password string) (entity.Account, error) {
	var userData entity.Account
	resErr := db.SelectContext(ctx, &userData, `SELECT 1 FROM accounts WHERE login = $1 and password = $2`, login, password)
	return userData, resErr
}

func (_ AccountData) AccountRecovery(db *sqlx.DB, ctx context.Context, email, login, publicKey string) (entity.Account, error) {
	var userData entity.Account
	resErr := db.SelectContext(ctx, &userData, `SELECT 1 FROM accounts WHERE email = $1 and login = $2 and public_key = $3`, email, login, publicKey)
	return userData, resErr
}

func (_ AccountData) ChangeStatus(db *sqlx.DB, ctx context.Context, id uuid.UUID, status common.AccountStatus) error {
	tx, txErr := db.BeginTxx(ctx, nil)
	if txErr != nil {
		return txErr
	}
	_, _ = tx.Exec(`UPDATE accounts SET status = $1 WHERE id = $2`, status, id)
	return tx.Commit()
}

func (_ AccountData) ChangePlan(db *sqlx.DB, ctx context.Context, id uuid.UUID, newPlan string) error {
	tx, txErr := db.BeginTxx(ctx, nil)
	if txErr != nil {
		return txErr
	}
	_, _ = tx.Exec(`UPDATE accounts SET plan = $1 WHERE id = $2`, newPlan, id)
	return tx.Commit()
}

func (_ AccountData) ChangeEmail(db *sqlx.DB, ctx context.Context, id uuid.UUID, newEmail string) error {
	tx, txErr := db.BeginTxx(ctx, nil)
	if txErr != nil {
		return txErr
	}
	_, _ = tx.Exec(`UPDATE accounts SET email = $1 WHERE id = $2`, newEmail, id)
	return tx.Commit()
}

func (_ AccountData) GetById(db *sqlx.DB, ctx context.Context, id string) (entity.Account, error) {
	var userData entity.Account
	err := db.GetContext(ctx, &userData, `SELECT 1 FROM accounts WHERE id = $1`, id)
	return userData, err
}

func (_ AccountData) GetAccounts(db *sqlx.DB, ctx context.Context, page, itemsPerPage int) ([]entity.Account, error) {
	var accounts []entity.Account
	err := db.GetContext(ctx, &accounts, `SELECT * FROM accounts`)
	return accounts, err
}

func (_ AccountData) GetAccountsCount(db *sqlx.DB, ctx context.Context, id string) (int, error) {
	var count int
	err := db.GetContext(ctx, &count, `SELECT COUNT(*) FROM accounts`)
	return count, err
}
