package models

import (
	"github.com/emelnychenko/go-press/enums"
	"time"
)

type (
	Post struct {
		Id          *PostId           `json:"id" xml:"id"`
		Author      interface{}       `json:"author" xml:"author"`
		Title       string            `json:"title" xml:"title"`
		Description string            `json:"description" xml:"description"`
		Content     string            `json:"content" xml:"content"`
		Status      enums.PostStatus  `json:"status" xml:"status"`
		Privacy     enums.PostPrivacy `json:"privacy" xml:"privacy"`
		Picture     *File             `json:"picture" xml:"picture"`
		Video       *File             `json:"video" xml:"video"`
		Views       int               `json:"views" xml:"views"`
		Published   *time.Time        `json:"published" xml:"published"`
		Created     *time.Time        `json:"created" xml:"created"`
		Updated     *time.Time        `json:"updated" xml:"updated"`
	}

	PostCreate struct {
		Title       string            `json:"title" xml:"title"`
		Description string            `json:"description" xml:"description"`
		Content     string            `json:"content" xml:"content"`
		Status      enums.PostStatus  `json:"status" xml:"status"`
		Privacy     enums.PostPrivacy `json:"privacy" xml:"privacy"`
		Views       int               `json:"views" xml:"views"`
		Published   *time.Time        `json:"published" xml:"published"`
	}

	PostUpdate struct {
		Title       string            `json:"title" xml:"title"`
		Description string            `json:"description" xml:"description"`
		Content     string            `json:"content" xml:"content"`
		Status      enums.PostStatus  `json:"status" xml:"status"`
		Privacy     enums.PostPrivacy `json:"privacy" xml:"privacy"`
		Views       int               `json:"views" xml:"views"`
		Published   *time.Time        `json:"published" xml:"published"`
	}
)
