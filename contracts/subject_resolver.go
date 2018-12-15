package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/enums"
	"github.com/emelnychenko/go-press/models"
)

type (
	SubjectResolver interface {
		ResolveSubject(subjectId *models.ModelId, subjectType enums.SubjectType) (subject interface{}, err common.Error)
	}
)
