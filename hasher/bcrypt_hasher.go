package hasher

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"golang.org/x/crypto/bcrypt"
)

type (
	bCryptHasherImpl struct {
		cost int
	}
)

const (
	BCryptPasswordCost = 12
)

func NewBCryptHasher() contracts.Hasher {
	return &bCryptHasherImpl{cost: BCryptPasswordCost}
}

func (h *bCryptHasherImpl) Make(password string) (string, common.Error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)

	if nil != err {
		return "", common.NewServerError(err)
	}

	return string(hash), nil
}

func (*bCryptHasherImpl) Check(hashedPassword, password string) common.Error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if nil != err {
		return common.NewServerError(err)
	}

	return nil
}
