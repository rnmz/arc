package entity

import "github.com/google/uuid"

type FolderEntity struct {
	Id         uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	Path       string    `db:"path"`
	OwnerID    uuid.UUID `db:"owner_id"`
	ViewStatus string    `db:"view_status"`
}
