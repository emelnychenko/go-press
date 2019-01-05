package models

import (
	"github.com/emelnychenko/go-press/common"
	"time"
)

type (
	CategoryId = common.ModelId

	Category struct {
		Id      *CategoryId `json:"id" xml:"id"`
		Name    string      `json:"name" xml:"name"`
		Created *time.Time  `json:"created" xml:"created"`
		Updated *time.Time  `json:"updated" xml:"updated"`
	}

	CategoryTree struct {
		*Category
		Categories []*CategoryTree `json:"categories" xml:"categories"`
	}

	CategoryCreate struct {
		Name string `json:"name" xml:"name"`
	}

	CategoryUpdate struct {
		Name string `json:"name" xml:"name"`
	}

	CategoryPaginationQuery struct {
		*PaginationQuery
	}
)
