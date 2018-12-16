package hasher

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"golang.org/x/crypto/bcrypt"
)

type (
	BCryptHasher struct {
		cost int
	}
)

const (
	BCryptPasswordCost = 12
)

func NewBCryptHasher() contracts.Hasher {
	return &BCryptHasher{cost: BCryptPasswordCost}
}

func (h *BCryptHasher) Make(password string) (string, common.Error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)

	if nil != err {
		return "", common.NewServerError(err)
	}

	return string(hash), nil
}

func (*BCryptHasher) Check(hashedPassword, password string) common.Error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if nil != err {
		return common.NewServerError(err)
	}

	return nil
}
