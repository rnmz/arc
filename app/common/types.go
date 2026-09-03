package common

type SignatureStatus int64

const (
	SignatureStatusVerified    SignatureStatus = 1
	SignatureStatusNotVerified SignatureStatus = 32
)

type ViewStatus int64

const (
	ViewStatusPublic  ViewStatus = 1
	ViewStatusPrivate ViewStatus = 32
	ViewStatusLinked  ViewStatus = 64
)

type AccountStatus string

const (
	AccountStatusActive  AccountStatus = "active"
	AccountStatusBanned  AccountStatus = "banned"
	AccountStatusDeleted AccountStatus = "deleted"
)

type AccountRole string

const (
	AccountRoleUser  AccountRole = "user"
	AccountRoleAdmin AccountRole = "admin"
)
