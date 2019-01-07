package strategies

import (
	"fmt"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"mime"
)

type (
	filePathStrategyImpl struct {
	}
)

func NewFilePathStrategy() contracts.FilePathStrategy {
	return &filePathStrategyImpl{}
}

func (*filePathStrategyImpl) BuildPath(fileEntity *entities.FileEntity) (filePath string, err errors.Error) {
	fileExtensions, mimeErr := mime.ExtensionsByType(fileEntity.Type)

	if nil != mimeErr {
		err = errors.NewSystemErrorFromBuiltin(mimeErr)
		return
	}

	filePath = fmt.Sprintf(
		"uploads/%d/%d/%s%s",
		fileEntity.Created.Year(),
		int(fileEntity.Created.Month()),
		fileEntity.Id,
		fileExtensions[0],
	)
	return
}
