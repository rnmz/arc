package common

import (
	"fmt"
	"time"
)

type RoleCantPromoteError struct {
	UserId int64
	RoleId int64
	Time   *time.Time
}

func (err RoleCantPromoteError) Error() string {
	return fmt.Sprintf("")
}
