package service

import (
	"arc/app/crypt"
	"arc/app/data"
	"context"

	"github.com/jmoiron/sqlx"
)

var accountData = data.AccountData{}

func Create(db *sqlx.DB, ctx context.Context, name, email, login, password string) {
	hashedPass, hashErr := crypt.HashText(password)
	accountData.CreateNewAccount()
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
