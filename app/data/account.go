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

func (_ AccountData) CreateNewAccount(
	db *sqlx.DB, ctx context.Context,
	name, login, email string, emailVerified bool,
	passwordHash string, registrationDate time.Time, role common.AccountRole,
	status common.AccountStatus, planId int, publicKey string,
) error {
	tx, txErr := db.BeginTxx(ctx, nil)
	if txErr != nil {
		return txErr
	}
	_, _ = tx.Exec(`INSERT INTO accounts (
  		name, login, email, email_verified,
	  	password_hash, registration_date, role,
	  	status, plan_id, public_key) VALUES (
		$1, $2, $3, $4,
	 	$5, $6, $7,
	 	$8, $9, $10
	  	)`, name, login, email, emailVerified, passwordHash, registrationDate, role, status, planId, publicKey,
	)
	return tx.Commit()
}

func (_ AccountData) Login(db *sqlx.DB, ctx context.Context, login, password string) (entity.Account, error) {
	var userData entity.Account
	resErr := db.SelectContext(ctx, &userData, `SELECT 1 FROM accounts WHERE login = $1 and password_hash = $2`, login, password)
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

func (_ AccountData) ChangePlan(db *sqlx.DB, ctx context.Context, id uuid.UUID, newPlanId int) error {
	tx, txErr := db.BeginTxx(ctx, nil)
	if txErr != nil {
		return txErr
	}
	_, _ = tx.Exec(`UPDATE accounts SET plan_id = $1 WHERE id = $2`, newPlanId, id)
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

func (_ AccountData) GetById(db *sqlx.DB, ctx context.Context, id uuid.UUID) (entity.Account, error) {
	var userData entity.Account
	err := db.GetContext(ctx, &userData, `SELECT 1 FROM accounts WHERE id = $1`, id)
	return userData, err
}

func (_ AccountData) GetAccounts(db *sqlx.DB, ctx context.Context, page, itemsPerPage int) ([]entity.Account, error) {
	var accounts []entity.Account
	err := db.GetContext(ctx, &accounts, `SELECT * FROM accounts LIMIT $1 OFFSET $2`, page, itemsPerPage)
	return accounts, err
}

func (_ AccountData) GetAccountsCount(db *sqlx.DB, ctx context.Context) (int, error) {
	var count int
	err := db.GetContext(ctx, &count, `SELECT COUNT(*) FROM accounts`)
	return count, err
}
