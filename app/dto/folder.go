package dto

import "github.com/google/uuid"

type FolderInfoDTO struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	OwnerID    uuid.UUID `json:"ownerId"`
	ViewStatus string    `json:"viewStatus"`
}

type FolderListDTO struct {
	Folders []FolderInfoDTO `json:"folders"`
	Page    int             `json:"page"`
}

type FolderFilesDTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Files []FileDTO `json:"files"`
	Page  int       `json:"page"`
}
