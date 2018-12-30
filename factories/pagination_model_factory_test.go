package factories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaginationModelFactory(t *testing.T) {
	t.Run("NewPaginationModelFactory", func(t *testing.T) {
		_, isPaginationModelFactory := NewPaginationModelFactory().(*paginationModelFactoryImpl)

		assert.True(t, isPaginationModelFactory)
	})

	t.Run("CreatePaginationQuery", func(t *testing.T) {
		paginationModelFactory := new(paginationModelFactoryImpl)
		paginationQuery := paginationModelFactory.CreatePaginationQuery()

		assert.Equal(t, paginationQueryDefaultLimit, paginationQuery.Limit)
		assert.Equal(t, 0, paginationQuery.Start)
		assert.Equal(t, 1, paginationQuery.Page)
	})

	t.Run("CreatePaginationResult", func(t *testing.T) {
		paginationModelFactory := new(paginationModelFactoryImpl)
		assert.NotNil(t, paginationModelFactory.CreatePaginationResult())
	})
}
