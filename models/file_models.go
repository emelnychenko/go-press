package models

import (
	"time"
)

type (
	FileId = ModelId

	File struct {
		Id      *FileId    `json:"id" xml:"id"`
		Name    string     `json:"name" xml:"name"`
		Size    int64      `json:"size" xml:"size"`
		Type    string     `json:"type" xml:"type"`
		Created *time.Time `json:"created" xml:"created"`
		Updated *time.Time `json:"updated" xml:"updated"`
	}

	FileUpload struct {
		Name string
		Size int64
		Type string
	}

	FileUpdate struct {
		Name string `json:"name" xml:"name" validate:"required"`
	}

	FilePaginationQuery struct {
		*PaginationQuery
	}
)
