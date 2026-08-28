package dto

type FolderInfoDTO struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Size        int64    `json:"size"`
	Crypt       string   `json:"crypt"`
	AccessType  string   `json:"access"`
	AccessLink  string   `json:"access_link"`
	AccessUsers []string `json:"access_users"`
}
