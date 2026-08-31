package data

import (
	"arc/app/common"
	"arc/app/entity"

	"github.com/google/uuid"
)

func newFile(fileId uuid.UUID) {
}

func getFile(fileId uuid.UUID) (entity.FileEntity, error) {

}

func deleteFile(fileId uuid.UUID) error {

}

func changeViewStatus(fileId uuid.UUID, viewStatus common.ViewStatus) (entity.FileEntity, error) {

}

func newFolder(name string, ownerId uuid.UUID) (entity.FolderEntity, error) {

}

func getFolder(fileId uuid.UUID) (entity.FolderEntity, error) {

}

func deleteFolder(fileId uuid.UUID) error {

}
