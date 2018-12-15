package models

type (
	Post struct {
		Id *PostId `json:"id" xml:"id"`
		Content string `json:"content" xml:"content"`
	}
)
