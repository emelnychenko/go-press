package strategies

import (
	"fmt"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"mime"
)

type (
	filePathStrategyImpl struct {
	}
)

func NewFilePathStrategy() contracts.FilePathStrategy {
	return &filePathStrategyImpl{}
}

func (*filePathStrategyImpl) BuildPath(fileEntity *entities.FileEntity) (filePath string, err common.Error) {
	fileExtensions, mimeErr := mime.ExtensionsByType(fileEntity.Type)

	if nil != mimeErr {
		err = common.NewSystemErrorFromBuiltin(mimeErr)
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
