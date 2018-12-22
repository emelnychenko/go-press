package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/enums"
)

type (
	SubjectResolver interface {
		ResolveSubject(subjectId *common.ModelId, subjectType enums.SubjectType) (subject interface{}, err common.Error)
	}
)
