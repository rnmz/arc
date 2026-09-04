package common

import (
	"fmt"
	"time"
)

type CryptHashTextError struct {
	Time time.Time
}

func (err CryptHashTextError) Error() string {
	return fmt.Sprintf("%s", err.Time)
}

type CryptGenerateKeyError struct {
	Time time.Time
}

func (err CryptGenerateKeyError) Error() string {
	return fmt.Sprintf("%s", err.Time)
}
