package models

import (
	"time"
)

type (
	PollId = ModelId

	Poll struct {
		Id      *PollId    `json:"id" xml:"id"`
		Title   string     `json:"title" xml:"title"`
		Created *time.Time `json:"created" xml:"created"`
		Updated *time.Time `json:"updated" xml:"updated"`
	}

	PollCreate struct {
		Title string `json:"title" xml:"title"`
	}

	PollUpdate struct {
		Title string `json:"title" xml:"title"`
	}

	PollPaginationQuery struct {
		*PaginationQuery
	}
)
