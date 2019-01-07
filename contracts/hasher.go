package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	Hasher interface {
		Make(password string) (string, errors.Error)
		Check(hashedPassword, password string) errors.Error
	}
)
