package contracts

import "github.com/emelnychenko/go-press/models"

type (
	PaginationModelFactory interface {
		CreatePaginationQuery() *models.PaginationQuery
		CreatePaginationResult() *models.PaginationResult
	}
)
