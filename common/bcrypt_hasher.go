package common

import (
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

func NewBCryptHasher() *BCryptHasher {
	return &BCryptHasher{cost: BCryptPasswordCost}
}

func (h *BCryptHasher) Make(password string) (string, Error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), h.cost);

	if nil != err {
		return "", NewServerError(err)
	}

	return string(hash), nil
}

func (*BCryptHasher) Check(hashedPassword, password string) Error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if nil != err {
		return NewServerError(err)
	}

	return nil
}

