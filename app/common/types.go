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

type AccountStatus int64

const (
	AccountStatusUser   AccountStatus = 1
	AccountStatusBanned AccountStatus = 32
	AccountStatusAdmin  AccountStatus = 64
)
