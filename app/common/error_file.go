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
