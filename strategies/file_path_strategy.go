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
	if nil == fileEntity.Id {
		err = common.ServerError("File.Id is empty")
		return
	}

	if nil == fileEntity.Created {
		err = common.ServerError("File.Created is empty")
		return
	}

	fileExtension, err2 := mime.ExtensionsByType(fileEntity.Type)

	if nil != err2 {
		err = common.NewServerError(err2)
		return
	}

	filePath = fmt.Sprintf(
		"uploads/%d/%d/%s%s",
		fileEntity.Created.Year(),
		int(fileEntity.Created.Month()),
		fileEntity.Id,
		fileExtension[0],
	)
	return
}



