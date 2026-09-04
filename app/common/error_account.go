package common

import (
	"fmt"
	"time"
)

type AccountSameLoginError struct {
	Login string
	Time  time.Time
}

func (err AccountSameLoginError) Error() string {
	return fmt.Sprintf("%s", err.Login)
}

type AccountPasswordTooShortError struct {
	Length int
	Time   time.Time
}

func (err AccountPasswordTooShortError) Error() string {
	return fmt.Sprint("")
}

type AccountPasswordTooWeakError struct {
	Length int
	Time   time.Time
}

func (err AccountPasswordTooWeakError) Error() string {
	return fmt.Sprint("")
}

type AccountEmailNotFoundError struct {
	Email string
	Time  time.Time
}

func (err AccountEmailNotFoundError) Error() string {
	return fmt.Sprintf("%s", err.Email)
}

type AccountInvalidEmailError struct {
	Email string
	Time  time.Time
}

func (err AccountInvalidEmailError) Error() string {
	return fmt.Sprintf("%s", err.Email)
}

type AccountInvalidPasswordError struct {
	Time time.Time
}

func (err AccountInvalidPasswordError) Error() string {
	return fmt.Sprintf("%s", err.Time)
}

type AccountUsernameTooShortError struct {
	Name string
	Time time.Time
}

func (err AccountUsernameTooShortError) Error() string {
	return fmt.Sprint("")
}

type AccountUsernameTooLongError struct {
	Name string
	Time time.Time
}

func (err AccountUsernameTooLongError) Error() string {
	return fmt.Sprint("")
}

type AccountInvalidUsernameError struct {
	Name string
	Time time.Time
}

func (err AccountInvalidUsernameError) Error() string {
	return fmt.Sprintf("%s", err.Time)
}

type AccountNotCreatedError struct {
	Time time.Time
}

func (err AccountNotCreatedError) Error() string {
	return fmt.Sprintf("%s", err.Time)
}
