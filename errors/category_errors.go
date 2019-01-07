package errors

import (
	"fmt"
)

//NewCategoryNotFoundError
func NewCategoryNotFoundError(request string) Error {
	message := fmt.Sprintf("The category was not found on request: %s", request)
	return NewNotFoundError(message)
}

//NewCategoryByIdNotFoundError
func NewCategoryByIdNotFoundError(categoryId fmt.Stringer) Error {
	request := fmt.Sprintf("id = %s", categoryId)
	return NewCategoryNotFoundError(request)
}

//NewCategoryXrefNotFoundError
func NewCategoryXrefNotFoundError(request string) Error {
	message := fmt.Sprintf("The category reference was not found on request: %s", request)
	return NewNotFoundError(message)
}

//NewCategoryXrefNotFoundByReferenceError
func NewCategoryXrefNotFoundByReferenceError(categoryId fmt.Stringer, objectType string, objectId fmt.Stringer) Error {
	request := fmt.Sprintf("categoryId = %s, objectType = %s, objectId = %s", categoryId, objectType, objectId)
	return NewCategoryXrefNotFoundError(request)
}
