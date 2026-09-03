package entity

import "github.com/google/uuid"

type FileEntity struct {
	Id         uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	Extension  string    `db:"extension"`
	Size       int64     `db:"size"`
	Path       string    `db:"path"`
	Encrypt    bool      `db:"encrypt"`
	OwnerID    uuid.UUID `db:"owner_id"`
	ViewStatus int64     `db:"view_status"`
	FolderId   uuid.UUID `db:"folder_id"`
}
