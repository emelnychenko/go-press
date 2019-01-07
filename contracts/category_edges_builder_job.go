package contracts

import (
	"github.com/emelnychenko/go-press/common"
)

type (
	CategoryEdgesBuilderJob interface {
		BuildCategoriesEdges() (err common.Error)
	}
)
