package crypt

type SignatureStatus int64

const (
	SignatureValid   SignatureStatus = 1
	SignatureInvalid SignatureStatus = 64
)

func CheckSignature(key string, payload any) SignatureStatus {
	return SignatureInvalid
}
