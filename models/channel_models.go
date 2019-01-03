package models

import (
	"github.com/emelnychenko/go-press/common"
	"time"
)

type (
	ChannelId = common.ModelId

	Channel struct {
		Id      *ChannelId `json:"id" xml:"id"`
		Name    string     `json:"name" xml:"name"`
		Created *time.Time `json:"created" xml:"created"`
		Updated *time.Time `json:"updated" xml:"updated"`
	}

	ChannelCreate struct {
		Name string `json:"name" xml:"name"`
	}

	ChannelUpdate struct {
		Name string `json:"name" xml:"name"`
	}

	ChannelPaginationQuery struct {
		*PaginationQuery
	}
)
