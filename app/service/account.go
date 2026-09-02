package service

import (
	"arc/app/crypt"
	"arc/app/data"
	"context"

	"github.com/jmoiron/sqlx"
)

var accountData = data.AccountData{}

func Create(db *sqlx.DB, ctx context.Context, planId int, name, email, login, password string) error {
	hashedPass, hashErr := crypt.HashText(password)
	if hashErr != nil {
		return hashErr
	}
	accountData.CreateNewAccount(db, ctx, planId, name, email, login, hashedPass)
}

func Login() {

}

func Logout() {

}

func Recover() {

}

func ChangeUserRole() {

}

func GetAccountInfoById() {

}

func GetAllAccounts() {

}
