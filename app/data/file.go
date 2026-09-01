package data

import (
	"arc/app/common"
	"arc/app/entity"
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func addFile(db *sqlx.DB, ctx context.Context,
	name string, size int,
	path string, encrypt bool,
	ownerId uuid.UUID, viewStatus common.ViewStatus,
	folderId uuid.UUID) error {
	tx, txErr := db.BeginTxx(ctx, nil)
	if txErr != nil {
		return txErr
	}
	_, _ = tx.Exec(`INSERT INTO files (name, size, path,
   		encrypt, owner_id, view_status,
	   	folder_id) VALUES ($1, $2, $3, $4, $5, $6, $7);`,
		name, size, path, encrypt, ownerId, viewStatus, folderId)
	return tx.Commit()
}

func getFile(db *sqlx.DB, ctx context.Context, id uuid.UUID) (entity.FileEntity, error) {
	var file entity.FileEntity
	err := db.GetContext(ctx, &file, `SELECT * FROM files WHERE id = $1`, id)
	return file, err
}

func deleteFile(db *sqlx.DB, ctx context.Context, fileId uuid.UUID) error {
	tx, txErr := db.BeginTxx(ctx, nil)
	if txErr != nil {
		return txErr
	}
	_, resErr := tx.Exec(`DELETE FROM files WHERE id = $1`, fileId)
	if resErr != nil {
		return resErr
	}
	return tx.Commit()
}

func changeViewStatus(db *sqlx.DB, ctx context.Context, fileId uuid.UUID, viewStatus common.ViewStatus) (entity.FileEntity, error) {
	tx, txErr := db.BeginTxx(ctx, nil)
	if txErr != nil {
		return entity.FileEntity{}, txErr
	}
	_, resErr := tx.Exec(`UPDATE files SET view_status = $1 WHERE id = $2`, viewStatus, fileId)
	if resErr != nil {
		return entity.FileEntity{}, resErr
	}
	return getFile(db, ctx, fileId)
}

func newFolder(db *sqlx.DB, ctx context.Context, name string, path string, ownerId uuid.UUID, viewStatus common.ViewStatus) error {
	tx, txErr := db.BeginTxx(ctx, nil)
	if txErr != nil {
		return txErr
	}
	_, _ = tx.Exec(`INSERT INTO folders (name, path, owner_id, view_status) VALUES ($1, $2, $3, $4)`, name, path, ownerId, viewStatus)
	return tx.Commit()
}

func getFolder(db *sqlx.DB, ctx context.Context, fileId uuid.UUID) (entity.FolderEntity, error) {
	var folder entity.FolderEntity
	err := db.GetContext(ctx, &folder, `SELECT * FROM folders WHERE id = $1`, fileId)
	return folder, err
}

func deleteFolder(db *sqlx.DB, ctx context.Context, fileId uuid.UUID) error {
	tx, txErr := db.BeginTxx(ctx, nil)
	if txErr != nil {
		return txErr
	}
	_, resErr := tx.Exec(`DELETE FROM folders WHERE id = $1`, fileId)
	if resErr != nil {
		return resErr
	}
	return tx.Commit()
}
