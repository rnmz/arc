package entity

import "github.com/google/uuid"

type FolderEntity struct {
	Id         uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	Path       string    `db:"path"`
	OwnerID    uuid.UUID `db:"owner_id"`
	ViewStatus int64     `db:"view_status"`
}
