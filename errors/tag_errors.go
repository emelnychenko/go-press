package errors

import (
	"fmt"
)

func NewTagNotFoundError(request string) Error {
	message := fmt.Sprintf("The tag was not found on request: %s", request)
	return NewNotFoundError(message)
}

func NewTagByIdNotFoundError(tagId fmt.Stringer) Error {
	request := fmt.Sprintf("id = %s", tagId)
	return NewTagNotFoundError(request)
}

//NewTagXrefNotFoundError
func NewTagXrefNotFoundError(request string) Error {
	message := fmt.Sprintf("The tag reference was not found on request: %s", request)
	return NewNotFoundError(message)
}

//NewTagXrefNotFoundByReferenceError
func NewTagXrefNotFoundByReferenceError(tagId fmt.Stringer, objectType string, objectId fmt.Stringer) Error {
	request := fmt.Sprintf("tagId = %s, objectType = %s, objectId = %s", tagId, objectType, objectId)
	return NewTagXrefNotFoundError(request)
}
