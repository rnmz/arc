package service

import (
	"arc/app/data"
	"arc/app/dto"
	"context"

	"github.com/jmoiron/sqlx"
)

var accountData = data.AccountData{}

func Create(db *sqlx.DB, ctx context.Context) error {
}

func Login() (string, error) {

}

func Logout() error {
	return nil
}

func Recover() (string, error) {

}

func ChangeUserRole() (dto.AccountDTO, error) {

}

func GetAccountInfoById() (dto.AccountDTO, error) {

}

func GetAllAccounts() ([]dto.AccountDTO, error) {

}
