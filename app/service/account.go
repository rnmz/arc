package service

import (
	"arc/app/data"
	"arc/app/dto"
	"context"

	"github.com/jmoiron/sqlx"
)

var accountData = data.AccountData{}

func Create(
	db *sqlx.DB, ctx context.Context,
	accountInfo dto.AccountInfoDTO,
) error {
	accountUsernameErr := validateUsername(accountInfo.Name)
}

func Login() (string, error) {

}

func Logout() error {
	return nil
}

func Recover() (string, error) {

}

func ChangeUserRole() (dto.AccountInfoDTO, error) {

}

func GetAccountInfoById() (dto.AccountInfoDTO, error) {

}

func GetAllAccounts() ([]dto.AccountInfoDTO, error) {

}
