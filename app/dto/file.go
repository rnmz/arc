package dto

import "time"

type FileInfoDTO struct {
	Id          string     `json:"id"`
	Crypt       bool       `json:"crypt"`
	Name        string     `json:"name"`
	Size        int64      `json:"size"` // MB
	FolderId    string     `json:"folder_id"`
	FolderName  string     `json:"folder_name"`
	AccessType  string     `json:"access"`
	AccessLink  string     `json:"access_link"`
	AccessUsers []string   `json:"access_users"`
	UploadDate  *time.Time `json:"upload_date"`
}

type FileListDTO struct {
	Files []FileInfoDTO `json:"files"`
	Page  int64         `json:"page"`
}
