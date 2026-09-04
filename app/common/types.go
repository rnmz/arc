package common

type SignatureStatus bool

const (
	SignatureStatusVerified    SignatureStatus = true
	SignatureStatusNotVerified SignatureStatus = false
)

type ViewStatus string

const (
	ViewStatusPublic  ViewStatus = "PUBLIC"
	ViewStatusPrivate ViewStatus = "PRIVATE"
	ViewStatusLink    ViewStatus = "PUBLIC_LINK"
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
