package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	SubjectResolver interface {
		ResolveSubject(subjectId *models.SubjectId, subjectType models.SubjectType) (interface{}, errors.Error)
	}
)
