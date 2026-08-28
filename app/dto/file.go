package dto

import "time"

type FileInfoDTO struct {
	Id         string     `json:"id"`
	Crypt      bool       `json:"crypt"`
	Name       string     `json:"name"`
	Size       int64      `json:"size"` // MB
	FolderId   string     `json:"folder_id"`
	FolderName string     `json:"folder_name"`
	UploadDate *time.Time `json:"upload_date"`
}

type FileListDTO struct {
	Files []FileInfoDTO `json:"files"`
	Page  int64         `json:"page"`
}
