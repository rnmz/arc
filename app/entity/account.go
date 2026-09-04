package entity

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id               uuid.UUID `db:"id"`
	PfpFile          string    `db:"pfp_file"`
	Name             string    `db:"name"`
	Login            string    `db:"login"`
	Email            string    `db:"email"`
	EmailVerified    bool      `db:"email_verified"`
	PasswordHash     string    `db:"password_hash"`
	RegistrationDate time.Time `db:"registration_date"`
	TwoFactor        bool      `db:"two_factor"`
	Role             string    `db:"role"`
	Status           string    `db:"status"`
	PlanId           int       `db:"plan_id"`
	PublicKey        string    `db:"public_key"`
}
