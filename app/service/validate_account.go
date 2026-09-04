package service

import (
	"arc/app/common"
	"regexp"
	"time"
	"unicode"
)

var (
	onlyEnglishAndDigitsRegex = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	hasDigitRegex             = regexp.MustCompile(`[0-9]`)

	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

func validatePassword(password string) error {
	if len(password) < 8 {
		return common.AccountPasswordTooShortError{
			Length: len(password),
			Time:   time.Now(),
		}
	}

	if !onlyEnglishAndDigitsRegex.MatchString(password) {
		return common.AccountInvalidPasswordError{
			Time: time.Now(),
		}
	}

	if !hasDigitRegex.MatchString(password) {
		return common.AccountPasswordTooWeakError{
			Length: len(password),
			Time:   time.Now(),
		}
	}

	return nil
}

func validateEmail(email string) error {
	if !emailRegex.MatchString(email) {
		return common.AccountInvalidEmailError{
			Email: email,
			Time:  time.Now(),
		}
	}
	return nil
}

func validateUsername(username string) error {
	if username == "" {
		return common.AccountInvalidUsernameError{
			Name: username,
			Time: time.Now(),
		}
	}

	if len(username) < 2 {
		return common.AccountUsernameTooShortError{
			Name: username,
			Time: time.Now(),
		}
	}
	if len(username) > 32 {
		return common.AccountUsernameTooLongError{
			Name: username,
			Time: time.Now(),
		}
	}

	for _, r := range username {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != ' ' {
			return common.AccountInvalidUsernameError{
				Name: username,
				Time: time.Now(),
			}
		}
	}
	return nil
}
