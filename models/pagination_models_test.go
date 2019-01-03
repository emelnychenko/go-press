package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaginationModels(t *testing.T) {
	t.Run("Offset", func(t *testing.T) {
		paginationLimit := 39
		paginationStart := 21
		paginationPage := 82

		paginationQuery := &PaginationQuery{Limit: paginationLimit, Start: paginationStart, Page: paginationPage}
		assert.Equal(t, (paginationPage-1)*paginationLimit+paginationStart, paginationQuery.Offset())
	})
}
