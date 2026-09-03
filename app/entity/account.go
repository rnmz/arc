package entity

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id               uuid.UUID `db:"id"`
	Pfp              string    `db:"pfp"`
	Name             string    `db:"name"`
	Login            string    `db:"login"`
	Email            string    `db:"email"`
	EmailVerified    bool      `db:"email_verified"`
	PasswordHash     string    `db:"password_hash"`
	RegistrationDate time.Time `db:"registration_date"`
	TwoFactor        bool      `db:"two_factor"`
	Status           string    `db:"status"`
	PlanId           string    `db:"plan_id"`
	PublicKey        string    `db:"public_key"`
}

type Pfp struct {
	Id   int64  `db:"id"`
	Path string `db:"path"`
}
