package contracts

import (
	"github.com/emelnychenko/go-press/models"
)

type (
	FileModelFactory interface {
		CreateFile() *models.File
		CreateFileUpload() *models.FileUpload
		CreateFileUpdate() *models.FileUpdate
	}
)
