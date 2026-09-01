package common

import (
	"fmt"
	"time"
)

type TxxError struct {
	TxFunctionName string
	Time           time.Time
}

func (err TxxError) Error() string {
	return fmt.Sprintf("%s", err.TxFunctionName)
}
