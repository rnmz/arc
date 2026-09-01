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
	AccountStatusUser  = "ROLE_USER"
	AccountStatusModer = "ROLE_MODER"
	AccountStatusAdmin = "ROLE_ADMIN"

	AccountStatusTrial  = "ROLE_TRIAL"
	AccountStatusBanned = "ROLE_BANNED"
)
