package models

import (
	"time"
)

type (
	BannerId = ModelId

	Banner struct {
		Id      *BannerId  `json:"id" xml:"id"`
		Title   string     `json:"title" xml:"title"`
		Key     string     `json:"key" xml:"key"`
		Created *time.Time `json:"created" xml:"created"`
		Updated *time.Time `json:"updated" xml:"updated"`
	}

	BannerCreate struct {
		Title string `json:"title" xml:"title"`
		Key   string `json:"key" xml:"key"`
	}

	BannerUpdate struct {
		Title string `json:"title" xml:"title"`
		Key   string `json:"key" xml:"key"`
	}

	BannerPaginationQuery struct {
		*PaginationQuery
	}
)
