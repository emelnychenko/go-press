package hashers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
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

func (h *bCryptHasherImpl) Make(password string) (hashedPassword string, err errors.Error) {
	hash, bCryptErr := bcrypt.GenerateFromPassword([]byte(password), h.cost)

	if bCryptErr != err {
		err = errors.NewSystemErrorFromBuiltin(bCryptErr)
	}

	hashedPassword = string(hash)
	return
}

func (*bCryptHasherImpl) Check(hashedPassword, password string) (err errors.Error) {
	bCryptErr := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if nil != bCryptErr {
		err = errors.NewSystemErrorFromBuiltin(bCryptErr)
	}

	return
}
