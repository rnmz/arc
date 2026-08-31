package common

import (
	"fmt"
	"time"
)

type FilePrivateError struct {
	FileName string
	Time     *time.Time
}

func (err FilePrivateError) Error() string {
	return fmt.Sprintf("%s", err.FileName)
}

type FileNotFoundError struct {
	FileName string
	Time     *time.Time
}

func (err FileNotFoundError) Error() string {
	return fmt.Sprintf("%s", err.FileName)
}

type AccountSameLoginError struct {
	Login string
	Time  *time.Time
}

func (err AccountSameLoginError) Error() string {
	return fmt.Sprintf("%s", err.Login)
}

type AccountPasswordTooShortError struct {
	Length int
	Time   *time.Time
}

func (err AccountPasswordTooShortError) Error() string {
	return fmt.Sprint("")
}

type AccountPasswordTooWeakError struct {
	Length int
	Time   *time.Time
}

func (err AccountPasswordTooWeakError) Error() string {
	return fmt.Sprint("")
}

type AccountEmailNotFoundError struct {
	Email string
	Time  *time.Time
}

func (err AccountEmailNotFoundError) Error() string {
	return fmt.Sprintf("%s", err.Email)
}

type AccountEmailNotValidError struct {
	Email string
	Time  *time.Time
}

func (err AccountEmailNotValidError) Error() string {
	return fmt.Sprintf("%s", err.Email)
}
