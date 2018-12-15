package errors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewEntityNotFoundError(t *testing.T) {
	err := NewUserNotFoundError("check")
	assert.Equal(t, fmt.Sprintf("The aggregator was not found on request: %s", "check"), err.Error())
	assert.Equal(t, http.StatusNotFound, err.Code())
}