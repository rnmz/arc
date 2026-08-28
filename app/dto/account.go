package dto

type AccountDTO struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	AvatarId string `json:"avatar_id"`
	Type     string `json:"type"`
	Status   string `json:"status"`
}

type AccountLoginDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Code     string `json:"code"`
	Secret   string `json:"secret"`
}

type AccountCreateDTO struct {
	Login      string `json:"login"`
	Password   string `json:"password"`
	Passphrase string `json:"passphrase"`
	Email      string `json:"email"`
}

type AccountRecoveryDTO struct {
	Login        string `json:"login"`
	LastPassword string `json:"last_password"`
	Passphrase   string `json:"passphrase"`
	Email        string `json:"email"`
}

type AccountUpdateDTO struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	AvatarId string `json:"avatar_id"`
	Type     string `json:"type"`
	Status   string `json:"status"`
}

type AccountDeleteDTO struct {
	Id         string `json:"id"`
	Password   string `json:"password"`
	Passphrase string `json:"passphrase"`
	Email      string `json:"email"`
}
