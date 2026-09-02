package data

import (
	"arc/app/common"
	"arc/app/entity"
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type FolderData struct{}

func (_ FolderData) NewFolder(db *sqlx.DB, ctx context.Context, name string, path string, ownerId uuid.UUID, viewStatus common.ViewStatus) error {
	tx, txErr := db.BeginTxx(ctx, nil)
	if txErr != nil {
		return txErr
	}
	_, _ = tx.Exec(`INSERT INTO folders (name, path, owner_id, view_status) VALUES ($1, $2, $3, $4)`, name, path, ownerId, viewStatus)
	return tx.Commit()
}

func (_ FolderData) GetFolder(db *sqlx.DB, ctx context.Context, fileId uuid.UUID) (entity.FolderEntity, error) {
	var folder entity.FolderEntity
	err := db.GetContext(ctx, &folder, `SELECT * FROM folders WHERE id = $1`, fileId)
	return folder, err
}

func (_ FolderData) DeleteFolder(db *sqlx.DB, ctx context.Context, fileId uuid.UUID) error {
	tx, txErr := db.BeginTxx(ctx, nil)
	if txErr != nil {
		return txErr
	}
	_, _ = tx.Exec(`DELETE FROM folders WHERE id = $1`, fileId)
	return tx.Commit()
}
