package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

type (
	fileEntityFactoryImpl struct {
	}
)

func NewFileEntityFactory() contracts.FileEntityFactory {
	return new(fileEntityFactoryImpl)
}

func (*fileEntityFactoryImpl) CreateFileEntity() *entities.FileEntity {
	return entities.NewFileEntity()
}
