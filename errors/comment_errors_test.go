package errors

import (
	"fmt"
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCommentErrors(t *testing.T) {
	t.Run("CommentNotFoundError", func(t *testing.T) {
		requestString := "request"
		err := NewCommentNotFoundError(requestString)
		assert.Equal(t, fmt.Sprintf("The comment was not found on request: %s", requestString), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})

	t.Run("CommentNotFoundByIdError", func(t *testing.T) {
		commentId := new(models.CommentId)
		err := NewCommentByIdNotFoundError(commentId)
		assert.Equal(t, fmt.Sprintf("The comment was not found on request: id = %s", commentId), err.Error())
		assert.Equal(t, http.StatusNotFound, err.Code())
	})
}
