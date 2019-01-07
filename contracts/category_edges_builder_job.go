package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	CategoryEdgesBuilderJob interface {
		BuildCategoriesEdges() (err errors.Error)
	}
)
