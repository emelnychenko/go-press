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

func (h *bCryptHasherImpl) Make(password string) (hashedPassword string, err common.Error) {
	hash, bCryptErr := bcrypt.GenerateFromPassword([]byte(password), h.cost)

	if bCryptErr != err {
		err = common.NewSystemErrorFromBuiltin(bCryptErr)
	}

	hashedPassword = string(hash)
	return
}

func (*bCryptHasherImpl) Check(hashedPassword, password string) (err common.Error) {
	bCryptErr := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if nil != bCryptErr {
		err = common.NewSystemErrorFromBuiltin(bCryptErr)
	}

	return
}
