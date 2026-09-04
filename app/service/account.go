package service

import (
	"arc/app/common"
	"arc/app/config"
	"arc/app/crypt"
	"arc/app/data"
	"arc/app/dto"
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

var accountData = data.AccountData{}

func Create(
	db *sqlx.DB, ctx context.Context,
	accountInfo dto.AccountInfoDTO,
	cfg config.Config,
) (string, error) {
	accountUsernameErr := validateUsername(accountInfo.Name)
	if accountUsernameErr != nil {
		return "", accountUsernameErr
	}
	passHash, passHashErr := crypt.HashText(accountInfo.Password)
	if passHashErr != nil {
		return "", passHashErr
	}

	pubKey, pubKeyErr := crypt.GenerateKey()
	if pubKeyErr != nil {
		return "", common.CryptGenerateKeyError{Time: time.Now()}
	}
	hashedKey, hashedKeyErr := crypt.HashKey([]byte(cfg.MasterKey), pubKey)
	if hashedKeyErr != nil {
		return "", common.CryptHashTextError{Time: time.Now()}
	}

	err := accountData.CreateNewAccount(
		db, ctx,
		accountInfo.Name, accountInfo.Login, accountInfo.Email, false,
		passHash, time.Now(), common.AccountRoleUser,
		common.AccountStatusActive, 1, hashedKey,
	)
	if err != nil {
		return "", common.AccountNotCreatedError{Time: time.Now()}
	}

	return string(pubKey), nil
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
