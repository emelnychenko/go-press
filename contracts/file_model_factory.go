package contracts

import (
	"github.com/emelnychenko/go-press/models"
)

type (
	FileModelFactory interface {
		CreateFilePaginationQuery() *models.FilePaginationQuery
		CreateFile() *models.File
		CreateFileUpload() *models.FileUpload
		CreateFileUpdate() *models.FileUpdate
	}
)
