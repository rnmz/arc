package dto

import (
	"github.com/google/uuid"
)

type FileDTO struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Extension  string    `json:"extension"`
	Size       int       `json:"size"`
	Encrypt    bool      `json:"encrypt"`
	OwnerID    uuid.UUID `json:"ownerId"`
	ViewStatus string    `json:"viewStatus"`
	FolderID   uuid.UUID `json:"folderId"`
}
type FileListDTO struct {
	Files []FileDTO `json:"files"`
	Page  int       `json:"page"`
}
